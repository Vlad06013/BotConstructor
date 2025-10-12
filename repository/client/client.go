package client

import "github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"

type Client struct {
	ID        uint   `json:"id" gorm:"primary_key;column:id"`
	Name      string `json:"name" gorm:"column:name;default:noname"`
	Email     string `json:"email" gorm:"column:email;default:null"`
	Phone     string `json:"phone" gorm:"column:phone;default:null"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`

	TelegramProfile telegramProfile.TelegramProfile `gorm:"foreignKey:client_id"`
}
