package model

type ReciveInsertRequest struct {
	Name         string  `json:"name" gorm:"column:names"`
	Credit       float64 `json:"credit" gorm:"column:credits"`
	ErrorMassage string  `json:"error_massage" gorm:"column:error_massage"`
	Transaction  string  `json:"transaction" gorm:"column:transactions"`
	TxID         string  `json:"TxID" gorm:"column:TxID"`
}
