package servicebus

import (
	"encoding/json"
	"errors"
	"go-stater-listener/domain/model"
	"go-stater-listener/pkg/utils"
)

func (s serviceBus) RequestProcess(eventType string, req model.RequestData) model.ErrorProps {
	switch eventType {
	case "STARTER_EVENT_REQUEST":
		data, err := json.Marshal(req)
		if err != nil {
			return utils.ErrorData(err)
		}

		props := map[string]interface{}{
			"eventType": "STARTER_EVENT_RESPONSE",
		}

		s.topic.Emit(data, props)

	default:
		return utils.ErrorData(errors.New("request process event type not found"))
	}

	return model.ErrorProps{}
}
