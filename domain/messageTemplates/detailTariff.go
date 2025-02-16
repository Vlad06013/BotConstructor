package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DetailTariffMessage(client tgUser.Clients, conn *gorm.DB, tariffId uint) api.TextMessage {

	s := tariff.Storage{DB: conn}
	tariff, _ := s.GetById(tariffId)
	text := "<b>Тариф:</b>\n" + tariff.Name + "\n\n<b>Подробнее:</b>\n" + tariff.Description
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "tariffSettings"),
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)
	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
