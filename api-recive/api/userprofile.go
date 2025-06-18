package api

import (
	"encoding/json"
	"fmt"
	"go-starter-api/domain/model"

	"go-starter-api/pkg/env"
	"net/http"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"
)

type userprofileApi struct {
	http http_base.HttpBase
	ai   appinsightsx.Appinsightsx
}

type UserprofileApi interface {
	GetUserProfileByEmail(email string) (*model.UserprofileResponse, error)
}

func NewUserprofileApi(http http_base.HttpBase, ai appinsightsx.Appinsightsx) UserprofileApi {
	return userprofileApi{http, ai}
}

func (u userprofileApi) GetUserProfileByEmail(email string) (*model.UserprofileResponse, error) {
	url := fmt.Sprintf("%s/employees/search/email?email=%s", env.Env().USER_PROFILE_BASE_URL, email)

	header := http.Header{
		"Accept":       []string{"application/json;odata=verbose"},
		"Content-Type": []string{"application/json"},
	}

	res, bodyBuff, err := u.http.Get(url, header)
	if err != nil {
		return nil, fmt.Errorf("%s [Request]: %s", "Error GetUserProfileByEmail", err.Error())
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s [StatusCode:%d]: %s", "Error GetUserProfileByEmail", res.StatusCode, string(bodyBuff))
	}

	var result model.UserprofileResponse
	err = json.Unmarshal(bodyBuff, &result)
	if err != nil {
		return nil, fmt.Errorf("%s [Unmarshal]: %s", "Error GetUserProfileByEmail", err.Error())
	}

	if len(result.Error) > 0 {
		return &result, fmt.Errorf("%s [Result Error]: %s", "Error GetUserProfileByEmail", result.Error)
	}

	if len(result.Data.(map[string]interface{})) == 0 {
		return &result, fmt.Errorf("%s [Check Length]: %s", "Error GetUserProfileByEmail", "email not found data")
	}

	u.ai.Dependency(appinsightsx.LoggerRequest{
		DependencyLog: &appinsightsx.DependencyLog{
			Name:           "Endpoint employees/search",
			DependencyType: "HTTP",
			Target:         url,
			Success:        true,
		},
	})
	return &result, nil
}
