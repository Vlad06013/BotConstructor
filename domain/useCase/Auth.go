package useCase

import (
	"github.com/Vlad06013/BotConstructor.git/repository/client"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/jinzhu/gorm"
)

func AuthClient(conn *gorm.DB, tgID int64, name string) *telegramProfile.TelegramProfile {
	s := telegramProfile.Storage{DB: conn}

	clientProfile, err := s.GetClientByTGID(tgID)

	if err != nil {

		r := client.Storage{DB: conn}
		clientCreated := r.CreateClient()

		clientProfile = s.CreateClientTelegramProfile(clientCreated.ID, tgID, name)
	}

	return clientProfile

}
