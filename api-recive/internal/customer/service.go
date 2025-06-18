package customer

import (
	"go-starter-api/domain/entity"
	"go-starter-api/domain/model"
)

type customerService struct {
	customerRepository CustomerRepository
}

type CustomerService interface {
	GetCustomerAll() ([]entity.CustomerModel, error)
	InsertCustomer(req model.CustomerInsertRequest) (entity.CustomerModel, error)
}

func NewCustomerService(customerRepository CustomerRepository) CustomerService {
	return customerService{customerRepository}
}

func (c customerService) GetCustomerAll() ([]entity.CustomerModel, error) {
	model := []entity.CustomerModel{}
	res, err := c.customerRepository.GetCustomerAll(model)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c customerService) InsertCustomer(req model.CustomerInsertRequest) (entity.CustomerModel, error) {
	res, err := c.customerRepository.InsertCustomer(req)
	if err != nil {
		return entity.CustomerModel{}, err
	}
	return res, nil
}
