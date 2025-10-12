package useCase

import (
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CallBackQuery(callBackQuery *tgbotapi.CallbackQuery, botAPI tgbotapi.BotAPI, conn *gorm.DB) {

	clientProfile := AuthClient(conn, callBackQuery.From.ID, callBackQuery.From.UserName)

	messageTemplate := findTemplate(callBackQuery.Data, *clientProfile, conn, &tgbotapi.Message{})

	deleteMessage(clientProfile.LastTgMessageId, clientProfile.TgUserId, botAPI)
	result := send(&messageTemplate, botAPI)
	s := telegramProfile.Storage{DB: conn}

	s.UpdateLastMessageClient(result.MessageID, clientProfile.ID)
}
