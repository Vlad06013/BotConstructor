package domain

type Domains struct {
	ID     uint   `json:"id" gorm:"primary_key;column:id"`
	Domain string `json:"domain" gorm:"column:domain;default:null"`
	Active bool   `json:"active" gorm:"column:active;default:false"`
}
