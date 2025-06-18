package customer

import (
	"go-starter-api/domain/entity"
	"go-starter-api/domain/model"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

type CustomerRepository interface {
	GetCustomerAll(model []entity.CustomerModel) ([]entity.CustomerModel, error)
	InsertCustomer(model model.CustomerInsertRequest) (entity.CustomerModel, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return customerRepository{db}
}

type CustomerDto struct {
	FirstName string `json:"firstName" gorm:"column:FIRST_NAME"`
}

func (t CustomerDto) TableName() string {
	return "CUSTOMER"
}

func (c customerRepository) GetCustomerAll(model []entity.CustomerModel) ([]entity.CustomerModel, error) {
	tx := c.db.Find(&model).Limit(100)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (c customerRepository) InsertCustomer(req model.CustomerInsertRequest) (entity.CustomerModel, error) {
	res := entity.CustomerModel{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	tx := c.db.Omit("CreatedDate", "CreatedBy", "UpdatedDate", "UpdatedBy").Create(&res)
	if tx.Error != nil {
		return entity.CustomerModel{}, tx.Error
	}
	return res, nil
}
