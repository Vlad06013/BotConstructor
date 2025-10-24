package domainsTemplate

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ConnectDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
	text := "Для парковки вашего домена, откройте личный кабинет на сайте регистраторе и пропишите a-name запись вашего домена на наш IP адрес - 88.888.888.888"

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Я сделяль!", "wait_input_domain"),
		),
	)

	mess := external.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
