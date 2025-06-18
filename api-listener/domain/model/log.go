package model

import "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"

type AppinsightLogStruct struct {
	TxID        string
	Event       string
	EventType   string
	Source      string
	MessageData topic.MessageResponse
}

type LogFormatStruct struct {
	Event     string
	EventType string
	Source    string
	Message   map[string]interface{} `json:"payload"`
}

type LogErrorStruct struct {
	MessageError string `json:"messageError"`
	FileError    string `json:"fileError"`
	LineError    int    `json:"lineError"`
}
