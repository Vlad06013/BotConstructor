package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func StartMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Привет <b>" + client.TgUserName + "</b>!" + "\n" +
		"Я бот для создания и управления короткими ссылками" + "\n" +
		"Я могу делать короткие ссылки, управлять ими и менять финальное содержимое любой ссылки!"

	//buttons := []tgbotapi.InlineKeyboardButton{
	//	tgbotapi.NewInlineKeyboardButtonData("Давайте начнем работу", "new_domain"),
	//}

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Давайте начнем работу", "new_domain"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
