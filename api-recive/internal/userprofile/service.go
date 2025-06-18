package userprofile

import (
	"go-starter-api/api"
	"go-starter-api/domain/model"
)

type userprofileService struct {
	userprofileApi api.UserprofileApi
}

type UserprofileService interface {
	GetEmail(email string) (*model.UserprofileResponse, error)
}

func NewCustomerService(userprofileApi api.UserprofileApi) UserprofileService {
	return userprofileService{userprofileApi}
}

func (u userprofileService) GetEmail(email string) (*model.UserprofileResponse, error) {
	res, err := u.userprofileApi.GetUserProfileByEmail(email)
	if err != nil {
		return nil, err
	}
	return res, nil
}
