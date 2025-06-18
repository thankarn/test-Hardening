package customer

import (
	"go-starter-api/domain/model"
)

type reciveService struct {
	reciveRepository ReciveRepository
}

type ReciveService interface {
	// GetCustomerAll() ([]entity.CustomerModel, error)
	InsertRecive(req model.ReciveInsertRequest) (bool, error)
}

func NewReciveService(reciveRepository ReciveRepository) ReciveService {
	return reciveService{reciveRepository}
}

// func (c reciveService) GetCustomerAll() ([]entity.CustomerModel, error) {
// 	model := []entity.CustomerModel{}
// 	res, err := c.reciveRepository.GetCustomerAll(model)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (c reciveService) InsertRecive(req model.ReciveInsertRequest) (bool, error) {
	res, err := c.reciveRepository.InsertRecive(req)
	if err != nil {
		return false, err
	}
	return res, nil
}
