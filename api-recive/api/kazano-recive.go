package api

import (
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"
)

type apiRecive struct {
	http http_base.HttpBase
	ai   appinsightsx.Appinsightsx
}

type ApiRecive interface {
	Getrecive(email string) (bool, error)
}

func NewReciveApi(http http_base.HttpBase, ai appinsightsx.Appinsightsx) ApiRecive {
	return apiRecive{http, ai}
}

func (u apiRecive) Getrecive(email string) (bool, error) {

	//insert to  DB
	return true, nil
}
