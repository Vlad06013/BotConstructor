package telegramProfile

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Storage struct {
	*gorm.DB
}

func (r *Storage) UpdateClient(client TelegramProfile) *TelegramProfile {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	r.Save(&TelegramProfile{
		ID:              client.ID,
		ClientID:        client.ClientID,
		TgUserId:        client.TgUserId,
		NextMessage:     client.NextMessage,
		LastTgMessageId: client.LastTgMessageId,
		TgUserName:      client.TgUserName,
		CreatedAt:       client.CreatedAt,
		UpdatedAt:       dateTime,
	})
	return &client
}

func (r *Storage) UpdateLastMessageClient(lastTgMessageId int, clientId uint) {
	r.Model(&TelegramProfile{}).Where("id =?", clientId).Update("last_tg_message_id", lastTgMessageId)
}

func (r *Storage) GetClientByTGID(tgID int64) (*TelegramProfile, error) {
	var telegramProfile TelegramProfile
	if err := r.First(&telegramProfile, "tg_user_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	return &telegramProfile, nil
}

func (r *Storage) CreateClientTelegramProfile(clientID uint, tgID int64, name string) *TelegramProfile {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	telegramProfile := TelegramProfile{
		ClientID:   clientID,
		TgUserId:   tgID,
		TgUserName: name,
		CreatedAt:  dateTime,
		UpdatedAt:  dateTime,
	}
	r.Create(&telegramProfile)
	return &telegramProfile
}
