package entity

type Recive struct {
	ID           int     `json:"id" gorm:"column:ID"`
	Name         string  `json:"name" gorm:"column:names"`
	Credit       float64 `json:"credit" gorm:"column:credits"`
	ErrorMassage string  `json:"error_massage" gorm:"column:error_massage"`
	Transaction  string  `json:"transaction" gorm:"column:transactions"`
}

func (t Recive) TableName() string {
	return "RECIVE"
}
