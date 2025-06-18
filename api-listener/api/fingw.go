package api

import (
	"bytes"
	"encoding/json"
	"fingw-listener-req/domain/model"
	incoming_event "fingw-listener-req/pkg/incomimg_event.go"
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsight"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base/get_token"
	httpClient "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_client"
)

type (
	finGwApi struct {
		http      httpClient.HttpClient
		http_base http_base.HttpBase
		ai        appinsight.Appinsight
	}
	FinGwAPI interface {
		ToDoImplement() error
		ToDoImplement2(data incoming_event.APIConfig, payload model.RequestData[model.PayloadData]) error
	}
)

func NewFinGwApi(htt3p httpClient.HttpClient) FinGwAPI {
	acquireToken := get_token.NewBanpuAcquireToken()
	http := http_base.NewHttpClient(acquireToken)
	ai := appinsight.NewAppinsights()
	return finGwApi{htt3p, http, ai}
}

func (api finGwApi) ToDoImplement() error {
	return nil
}

func (api finGwApi) ToDoImplement2(data incoming_event.APIConfig, payload model.RequestData[model.PayloadData]) error {
	log.Info().Msg("Initial Calling")
	url := data.APIEndpoint

	// Can do something like modify payload here if needed
	if data.HTTPMethod == incoming_event.HTTP_METHODS.GET {
		return api.Get(url, payload)
	} else if data.HTTPMethod == incoming_event.HTTP_METHODS.POST {
		return api.Post(url, payload)
	} else if data.HTTPMethod == incoming_event.HTTP_METHODS.PUT {
		return api.Put(url, payload)
	} else {
		return api.Delete(url, payload)
	}
}

func (api finGwApi) Post(url string, payload any) error {
	var res any
	bufReq := new(bytes.Buffer)
	if err := json.NewEncoder(bufReq).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	if err := api.http.Post(url, bufReq, &res); err != nil {
		return fmt.Errorf("HTTP POST request failed: %w", err)
	}

	return nil
}

func (api finGwApi) Get(url string, res any) error {
	if err := api.http.Get(url, res); err != nil {
		return fmt.Errorf("HTTP GET request failed: %w", err)
	}

	return nil
}

func (api finGwApi) Put(url string, payload any) error {
	var res any
	bufReq := new(bytes.Buffer)
	if err := json.NewEncoder(bufReq).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	if err := api.http.Put(url, bufReq, &res); err != nil {
		return fmt.Errorf("HTTP PUT request failed: %w", err)
	}

	return nil
}

func (api finGwApi) Delete(url string, payload any) error {
	var res any
	bufReq := new(bytes.Buffer)
	if err := json.NewEncoder(bufReq).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	if err := api.http.Delete(url, bufReq, &res); err != nil {
		return fmt.Errorf("HTTP DELETE request failed: %w", err)
	}

	return nil
}
