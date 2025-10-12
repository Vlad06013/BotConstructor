package domain

type Domains struct {
	ID        uint   `json:"id" gorm:"primary_key;column:id"`
	ClientID  uint   `json:"client_id" gorm:"column:client_id"`
	Domain    string `json:"domain" gorm:"column:domain;default:null"`
	Active    bool   `json:"active" gorm:"column:active;default:false"`
	Status    string `json:"status" gorm:"column:status"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}
