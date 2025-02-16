package tgUser

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Storage struct {
	*gorm.DB
}

func (r *Storage) InitClient(tgID int64, name string) *Clients {

	client, err := r.GetClientByTGID(tgID)
	if err != nil {
		client = r.CreateClient(tgID, name)
	}

	return client
}

func (r *Storage) UpdateClient(client Clients) *Clients {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	r.Save(&Clients{
		ID:              client.ID,
		Name:            client.Name,
		Email:           client.Email,
		Phone:           client.Phone,
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
	r.Model(&Clients{}).Where("id =?", clientId).Update("last_tg_message_id", lastTgMessageId)
}

func (r *Storage) GetClientByTGID(tgID int64) (*Clients, error) {
	var client Clients
	if err := r.First(&client, "tg_user_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *Storage) CreateClient(tgID int64, name string) *Clients {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	client := Clients{
		TgUserId:   tgID,
		TgUserName: name,
		CreatedAt:  dateTime,
		UpdatedAt:  dateTime,
	}
	r.Create(&client)
	return &client
}
