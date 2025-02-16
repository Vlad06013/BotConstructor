package useCase

import (
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CallBackQuery(callBackQuery *tgbotapi.CallbackQuery, botAPI tgbotapi.BotAPI, conn *gorm.DB) {

	s := tgUser.Storage{DB: conn}
	client := s.InitClient(callBackQuery.From.ID, callBackQuery.From.UserName)
	messageTemplate := findTemplate(callBackQuery.Data, *client, conn, &tgbotapi.Message{})

	deleteMessage(client.LastTgMessageId, client.TgUserId, botAPI)
	result := send(&messageTemplate, botAPI)
	s.UpdateLastMessageClient(result.MessageID, client.ID)
}
