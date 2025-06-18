package utils

import "fingw-listener-req/domain/model"

func GetLogMessage(als model.AppinsightLogStruct, msg map[string]interface{}) model.LogFormatStruct {
	return model.LogFormatStruct{
		Event:     als.Event,
		EventType: als.EventType,
		Source:    als.Source,
		Message:   msg,
	}
}
