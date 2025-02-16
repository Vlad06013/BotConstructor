package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func NewDomainMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Для начала работы вам нужно прикрепить к боту своё доменное имя "

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Припарковать свой домен", "connect_domain"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
