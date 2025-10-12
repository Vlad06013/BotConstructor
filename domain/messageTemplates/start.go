package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func StartMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
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

	mess := external.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
