package useCase

import (
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TextMessage(message *tgbotapi.Message, botAPI tgbotapi.BotAPI, conn *gorm.DB) {

	clientProfile := AuthClient(conn, message.From.ID, message.From.UserName)
	messageTemplate := getMessageTemplate(*clientProfile, conn, message)

	go deleteMessage(clientProfile.LastTgMessageId, clientProfile.TgUserId, botAPI)
	result := send(&messageTemplate, botAPI)
	s := telegramProfile.Storage{DB: conn}
	s.UpdateLastMessageClient(result.MessageID, clientProfile.ID)
}
