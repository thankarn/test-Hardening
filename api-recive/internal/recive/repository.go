package customer

import (
	"go-starter-api/domain/entity"
	"go-starter-api/domain/model"

	"gorm.io/gorm"
)

type reciveRepository struct {
	db *gorm.DB
}

type ReciveRepository interface {
	// GetCustomerAll(model []entity.NewReciveService) ([]entity.NewReciveService, error)
	InsertRecive(model model.ReciveInsertRequest) (bool, error)
}

func NewReciveRepository(db *gorm.DB) ReciveRepository {
	return reciveRepository{db}
}

type ReciveDto struct {
	FirstName string `json:"firstName" gorm:"column:FIRST_NAME"`
}

func (t ReciveDto) TableName() string {
	return "RECIVE"
}

// func (c reciveRepository) GetCustomerAll(model []entity.NewReciveService) ([]entity.NewReciveService, error) {
// 	tx := c.db.Find(&model).Limit(100)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return model, nil
// }

func (c reciveRepository) InsertRecive(req model.ReciveInsertRequest) (bool, error) {
	res := entity.Recive{
		Name:         "",
		Credit:       0,
		ErrorMassage: "",
		Transaction:  "",
	}

	tx := c.db.Omit("CreatedDate", "CreatedBy", "UpdatedDate", "UpdatedBy").Create(&res)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
