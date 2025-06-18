package servicebus

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-stater-listener/domain/model"
	"go-stater-listener/pkg/env"
	"go-stater-listener/pkg/utils"
	"time"
	_ "time/tzdata"

	m "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/mail"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"
	"github.com/google/uuid"
)

func (s serviceBus) ResponseProcess(eventType string, message topic.MessageResponse) model.ErrorProps {
	res := model.ResponseData{}
	err := json.Unmarshal(message.Data, &res)
	if err != nil {
		return utils.ErrorData(err)
	}

	switch eventType {
	case "STARTER_EVENT_RESPONSE":
		err := s.SendEmailMessage(res)
		if err.Error != nil {
			return err
		}

	default:
		return utils.ErrorData(errors.New("response process event type not found"))
	}
	return model.ErrorProps{}
}

func (s serviceBus) SendEmailMessage(res model.ResponseData) model.ErrorProps {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return utils.ErrorData(err)
	}

	mail := m.NewEmail()
	mail.Process(m.EmailProcess{
		TxId:        uuid.New().String(),
		From:        env.Env().EMAIL_FROM,
		To:          []string{res.Payload.Email},
		CC:          nil,
		Subject:     fmt.Sprintf("go starter listener: %s", time.Now().In(loc).Format("02/01/2006 15:04:05")),
		Body:        fmt.Sprintf("payload message: %s", res.Payload.Message),
		Attachments: []string{},
		ContentType: "TEXT",
		Container:   env.Env().AZURE_MAIL_BLOB_CONTAINER,
	})

	return model.ErrorProps{}
}
