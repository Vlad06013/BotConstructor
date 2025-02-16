package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CabinetMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Добро пожаловать в ваш личный кабинет!\n " +
		"Тут можно управлять всеми вашими сокращенными ссылками."

	//s := domain.Storage{DB: conn}
	//s.CreateDomain(client.ID, textMessage)

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("Сократить ссылку", "shotLink"),
			tgbotapi.NewInlineKeyboardButtonData("Сократить ссылку", "chooseDomainToShotLink"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Управление ссылками", "linkSettings"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Управление доменами", "domainSettings"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ТАРИФЫ", "tariffSettings"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
