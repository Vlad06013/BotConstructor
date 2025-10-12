package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func SaveLinkDestinationMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, message *tgbotapi.Message, urlId uint) external.TextMessage {

	s := url.Storage{DB: conn}
	host, _ := parseUrl(message.Text)

	if host == "" {
		text := "Неверная ссылка:  <b>" + message.Text + " </b> \n"

		mess := external.TextMessage{
			Text:   text,
			ChatId: client.TgUserId,
		}
		return mess
	}
	d := url.Storage{DB: conn}
	urlFound, _ := d.GetUrlByID(urlId)

	s.UpdateUrlDestination(message.Text, urlFound.ID)
	client.NextMessage = ""
	c := telegramProfile.Storage{DB: conn}
	c.UpdateClient(client)

	text := "Заебись! Мы поменяли всё что надо не заёбывай \n" +
		"Теперь она ведет на: \n <b>" + message.Text + " </b> \n "

	buttons := tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)
	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
