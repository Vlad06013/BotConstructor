package payment

type Payment struct {
	ID          uint   `json:"id" gorm:"primary_key;column:id"`
	PaymentId   string `json:"payment_id" gorm:"column:payment_id"`
	PaymentUrl  string `json:"payment_url" gorm:"column:payment_url"`
	Amount      uint   `json:"amount" gorm:"column:amount"`
	Status      string `json:"status" gorm:"column:status"`
	OrderNumber string `json:"order_number" gorm:"column:order_number;"`
}
