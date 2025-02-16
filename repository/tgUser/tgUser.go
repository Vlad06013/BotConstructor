package tgUser

type Clients struct {
	ID              uint   `json:"id" gorm:"primary_key;column:id"`
	Name            string `json:"name" gorm:"column:name;default:noname"`
	Email           string `json:"email" gorm:"column:email;default:null"`
	Phone           string `json:"phone" gorm:"column:phone;default:null"`
	TgUserId        int64  `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	NextMessage     string `json:"next_message" gorm:"column:next_message"`
	LastTgMessageId int    `json:"last_tg_message_id" gorm:"column:last_tg_message_id;default:null"`
	TgUserName      string `json:"tg_user_name" gorm:"column:tg_user_name"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       string `json:"updated_at" gorm:"column:updated_at"`
}
