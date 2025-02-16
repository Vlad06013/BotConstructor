package domain

type Domains struct {
	ID        uint   `json:"id" gorm:"primary_key;column:id"`
	TgUserId  uint   `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	Domain    string `json:"domain" gorm:"column:domain;default:null"`
	Active    bool   `json:"active" gorm:"column:active;default:false"`
	Status    int64  `json:"status" gorm:"column:status;default:0"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}
