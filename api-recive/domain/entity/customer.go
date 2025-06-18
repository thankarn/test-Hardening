package entity

import "time"

type CustomerModel struct {
	ID          int       `json:"id" gorm:"column:ID"`
	FirstName   string    `json:"firstName" gorm:"column:FIRST_NAME"`
	LastName    string    `json:"lastName" gorm:"column:LAST_NAME"`
	Email       *string    `json:"email" gorm:"column:EMAIL"`
	Phone       *string    `json:"phone" gorm:"column:PHONE"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:CREATED_DATE"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:CREATED_BY"`
	UpdatedDate *time.Time `json:"updatedDate" gorm:"column:UPDATED_DATE"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:UPDATED_BY"`
}

func (t CustomerModel) TableName() string {
	return "CUSTOMER"
}
