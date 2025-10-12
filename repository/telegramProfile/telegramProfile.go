package telegramProfile

type TelegramProfile struct {
	ID              uint   `json:"id" gorm:"primary_key;column:id"`
	ClientID        uint   `json:"client_id" gorm:"column:client_id"`
	TgUserId        int64  `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	NextMessage     string `json:"next_message" gorm:"column:next_message"`
	LastTgMessageId int    `json:"last_tg_message_id" gorm:"column:last_tg_message_id;default:null"`
	TgUserName      string `json:"tg_user_name" gorm:"column:tg_user_name"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       string `json:"updated_at" gorm:"column:updated_at"`
}
