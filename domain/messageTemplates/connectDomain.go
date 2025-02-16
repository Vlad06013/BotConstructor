package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ConnectDomainMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Для парковки вашего домена, откройте личный кабинет на сайте регистраторе и пропишите a-name запись вашего домена на наш IP адрес - 88.888.888.888"

	//buttons := []tgbotapi.InlineKeyboardButton{
	//	tgbotapi.NewInlineKeyboardButtonData("Я сделяль!", "wait_input_domain"),
	//	tgbotapi.NewInlineKeyboardButtonData("Назад", "new_domain"),
	//}
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Я сделяль!", "wait_input_domain"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
