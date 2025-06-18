package servicebus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-stater-listener/domain/model"
	"go-stater-listener/pkg/env"
	"go-stater-listener/pkg/utils"
	"runtime/debug"

	bpLogCenter "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/log_center/logx"
	bpLogCenterModel "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/log_center/model"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type serviceBus struct {
	logM  bpLogCenter.LogCenter
	topic *topic.Topic
}

type ServiceBus interface {
	SubscriptionSuccess()
}

func NewServiceBus(logM bpLogCenter.LogCenter) ServiceBus {
	t := topic.New()
	logM.Info(bpLogCenterModel.LoggerRequest{
		Tags:    utils.TAGS_SERVICE_NEW_SERVICE,
		Process: utils.PROCESS_START_PROJECT,
	})
	return serviceBus{logM, t}
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
	case utils.ContainsEventType(eventType, utils.EventTypeRequestStruct):
		req := model.RequestData{}
		err = json.Unmarshal(message.Data, &req)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		als = model.AppinsightLogStruct{
			TxID:      req.TxID,
			Event:     "Request",
			EventType: eventType,
			Source:    req.Source,
		}

		err := s.RequestProcess(eventType, req)
		if err.Error != nil {
			errProp = err
			return
		}

	case utils.ContainsEventType(eventType, utils.EventTypeResponseStruct):
		res := model.ResponseData{}
		err = json.Unmarshal(message.Data, &res)
		if err != nil {
			errProp = utils.ErrorData(err)
			return
		}

		als = model.AppinsightLogStruct{
			TxID:      res.TxID,
			Event:     "Response",
			EventType: eventType,
			Source:    res.Source,
		}

		err := s.ResponseProcess(eventType, message)
		if err.Error != nil {
			errProp = err
			return
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
	if errProp.Error != nil {
		deadLetterOptions := &azservicebus.DeadLetterOptions{
			ErrorDescription: to.Ptr(errProp.Error.Error()),
		}
		receiver.DeadLetterMessage(context.TODO(), msg, deadLetterOptions)

		err := s.AppinsightMonitorLog(als, message, errProp)
		if err != nil {
			s.logM.Error(bpLogCenterModel.LoggerRequest{
				Tags:    utils.TAGS_SERVICE_CALLBACK,
				Process: utils.PROCESS_PANIC_LOG,
				Error:   fmt.Sprintf("Error AppinsightMonitorLog: %v", err),
			})
		}
	} else {
		err := s.AppinsightMonitorLog(als, message, model.ErrorProps{})
		if err != nil {
			s.logM.Error(bpLogCenterModel.LoggerRequest{
				Tags:    utils.TAGS_SERVICE_CALLBACK,
				Process: utils.PROCESS_PANIC_LOG,
				Error:   fmt.Sprintf("Error AppinsightMonitorLog: %v", err),
			})
		}
	}

	err := receiver.CompleteMessage(context.Background(), msg, &azservicebus.CompleteMessageOptions{})
	if err != nil {
		s.logM.Error(bpLogCenterModel.LoggerRequest{
			Tags:    utils.TAGS_SERVICE_CALLBACK,
			Process: utils.PROCESS_COMPLETE_ERROR,
			Error:   fmt.Sprintf("[Subscription] Failed to complete message: %v", err),
		})
	}
	return nil
}

func (s serviceBus) HandlePanic(message topic.MessageResponse) {
	var msg map[string]interface{}
	err := json.Unmarshal(message.Data, &msg)
	if err != nil {
		s.logM.Error(bpLogCenterModel.LoggerRequest{
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

		s.logM.Error(bpLogCenterModel.LoggerRequest{
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
		msgData := model.LogFormatStruct{
			Event:     "Exception",
			EventType: als.EventType,
			Source:    als.Source,
			Message:   msg,
		}

		errData := model.LogErrorStruct{
			MessageError: errProp.Error.Error(),
			FileError:    errProp.FileError,
			LineError:    errProp.LineError,
		}

		s.logM.Error(bpLogCenterModel.LoggerRequest{
			TxID:    als.TxID,
			Tags:    utils.TAGS_LOG,
			Process: utils.PROCESS_APP_LOG,
			Error:   errData,
			Param:   msgData,
		})

	} else {
		msgData := model.LogFormatStruct{
			Event:     als.Event,
			EventType: als.EventType,
			Source:    als.Source,
			Message:   msg,
		}

		switch {
		case utils.ContainsEventType(als.EventType, utils.EventTypeRequestStruct), utils.ContainsEventType(als.EventType, utils.EventTypeResponseStruct):
			s.logM.Info(bpLogCenterModel.LoggerRequest{
				TxID:    als.TxID,
				Tags:    utils.TAGS_LOG,
				Process: utils.PROCESS_APP_LOG,
				Param:   msgData,
			})

		default:
			return errors.New("appinsight monitor log event type not found")
		}
	}
	return nil
}
