package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"
	"github.com/rs/zerolog"
)

func ContainsEventType(e string, eventArr []string) bool {
	for _, event := range eventArr {
		if event == e {
			return true
		}
	}
	return false
}

func TerminalLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	return zerolog.New(output).With().Timestamp().Logger()
}

func EventType(message topic.MessageResponse) (string, error) {
	customJson, err := json.Marshal(message.Custom)
	if err != nil {
		return "", err
	}

	var customFormat struct {
		EventType string `json:"eventType"`
	}

	err = json.Unmarshal(customJson, &customFormat)
	if err != nil {
		return "", err
	}

	return customFormat.EventType, nil
}
