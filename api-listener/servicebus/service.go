package servicebus

import (
	"context"
	"encoding/json"
	"errors"
	"fingw-listener-req/api"
	"fingw-listener-req/domain/model"
	"fingw-listener-req/pkg/env"
	incoming_event "fingw-listener-req/pkg/incomimg_event.go"
	"fingw-listener-req/pkg/utils"
	"fmt"
	"runtime/debug"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/samber/lo"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"
)

type serviceBus struct {
	ai       appinsightsx.Appinsightsx
	topic    *topic.Topic
	finGwApi api.FinGwAPI
}

type ServiceBus interface {
	SubscriptionSuccess()
}

func NewServiceBus(ai appinsightsx.Appinsightsx, finGwApi api.FinGwAPI) ServiceBus {
	t := topic.New()
	ai.Info(appinsightsx.LoggerRequest{
		Tags:    utils.TAGS_SERVICE_NEW_SERVICE,
		Process: utils.PROCESS_START_PROJECT,
	})
	return serviceBus{ai, t, finGwApi}
}

func (s serviceBus) SubscriptionSuccess() {
	s.topic.SetTopic(env.Env().SB_TOPIC).SetSubscription(env.Env().SB_SUBSCRIPTION).Subscribe(s.Callback)
}

func (s serviceBus) CallProcess(message topic.MessageResponse) (als model.AppinsightLogStruct, errProp model.ErrorProps) {
	eventType, err := utils.EventType(message)

	if err != nil {
		errProp = utils.ErrorData(err)
		return
	}

	switch {
	case utils.ContainsEventType(eventType,
		lo.Map(incoming_event.Registered_APIs, func(event incoming_event.APIConfig, _ int) string {
			return event.EventType
		})):
		req := model.RequestData[model.PayloadData]{}
		err = json.Unmarshal(message.Data, &req)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		//TODO: api package
		als = model.AppinsightLogStruct{
			TxID:      req.TxID,
			Event:     "Request",
			EventType: eventType,
		}

		request, found := lo.Find(incoming_event.Registered_APIs, func(event incoming_event.APIConfig) bool {
			return event.EventType == eventType
		})
		if !found {
			s.ai.Error(
				appinsightsx.LoggerRequest{})
		}

		err = s.finGwApi.ToDoImplement2(request, req)
		if err != nil {
			s.ai.Error(
				appinsightsx.LoggerRequest{
					Error: err,
				})
		}

	default:
		errProp = utils.ErrorData(errors.New("call process event type not found"))
		return
	}
	return als, errProp
}

func (s serviceBus) Callback(message topic.MessageResponse, receiver *azservicebus.Receiver, msg *azservicebus.ReceivedMessage) error {
	defer s.HandlePanic(message)

	als, errProp := s.CallProcess(message)

	var err error
	if errProp.Error != nil {
		err = s.AppinsightMonitorLog(als, message, errProp)
		deadLetterOptions := &azservicebus.DeadLetterOptions{
			ErrorDescription: to.Ptr(errProp.Error.Error()),
		}
		receiver.DeadLetterMessage(context.TODO(), msg, deadLetterOptions)
	} else {
		err = s.AppinsightMonitorLog(als, message, model.ErrorProps{})
	}

	if err != nil {
		s.ai.Error(appinsightsx.LoggerRequest{
			Tags:    utils.TAGS_SERVICE_CALLBACK,
			Process: utils.PROCESS_PANIC_LOG,
			Error:   fmt.Sprintf("Error AppinsightMonitorLog: %v", err),
		})
	}

	err = receiver.CompleteMessage(context.Background(), msg, &azservicebus.CompleteMessageOptions{})
	if err != nil {
		s.ai.Error(appinsightsx.LoggerRequest{
			Tags:    utils.TAGS_SERVICE_CALLBACK,
			Process: utils.PROCESS_PANIC_LOG,
			Error:   fmt.Sprintf("[Subscription] Failed to complete message: %v", err),
		})
	}

	return nil
}

func (s serviceBus) HandlePanic(message topic.MessageResponse) {
	var msg map[string]interface{}
	err := json.Unmarshal(message.Data, &msg)
	if err != nil {
		s.ai.Error(appinsightsx.LoggerRequest{
			Tags:    utils.TAGS_PANIC,
			Process: utils.PROCESS_PANIC_LOG,
			Error:   fmt.Sprintf("Error HandlePanic: %v", err),
		})
	}

	if r := recover(); r != nil {
		msgData := struct {
			Event   string
			Message interface{} `json:"payload"`
		}{
			Event:   "Panic",
			Message: msg,
		}

		panicData := struct {
			Recover    interface{}
			DebugStack string
		}{
			Recover:    r,
			DebugStack: string(debug.Stack()),
		}

		s.ai.Error(appinsightsx.LoggerRequest{
			Tags:    utils.TAGS_PANIC,
			Process: utils.PROCESS_PANIC_LOG,
			Param:   msgData,
			Panic:   panicData,
		})
	}
}

func (s serviceBus) AppinsightMonitorLog(als model.AppinsightLogStruct, message topic.MessageResponse, errProp model.ErrorProps) error {
	var msg map[string]interface{}
	err := json.Unmarshal(message.Data, &msg)
	if err != nil {
		return err
	}

	if errProp.Error != nil {
		als.Event = "Exception"
		msgData := utils.GetLogMessage(als, msg)

		errData := model.LogErrorStruct{
			MessageError: errProp.Error.Error(),
			FileError:    errProp.FileError,
			LineError:    errProp.LineError,
		}

		s.ai.Error(appinsightsx.LoggerRequest{
			TxID:    als.TxID,
			Tags:    utils.TAGS_LOG,
			Process: utils.PROCESS_APP_LOG,
			Error:   errData,
			Param:   msgData,
		})

	} else {

		s.ai.Info(appinsightsx.LoggerRequest{
			TxID:    als.TxID,
			Tags:    utils.TAGS_LOG,
			Process: utils.PROCESS_APP_LOG,
			Param:   utils.GetLogMessage(als, msg),
		})
	}

	return nil
}
