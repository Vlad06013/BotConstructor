package useCase

import (
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TextMessage(message *tgbotapi.Message, botAPI tgbotapi.BotAPI, conn *gorm.DB) {

	s := tgUser.Storage{DB: conn}
	client := s.InitClient(message.From.ID, message.From.UserName)
	messageTemplate := getMessageTemplate(*client, conn, message)

	deleteMessage(client.LastTgMessageId, client.TgUserId, botAPI)
	result := send(&messageTemplate, botAPI)
	s.UpdateLastMessageClient(result.MessageID, client.ID)
}
