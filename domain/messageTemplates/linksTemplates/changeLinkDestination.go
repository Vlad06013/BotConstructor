package linksTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ChangeLinkDestinationMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, urlId uint) external.TextMessage {
	var text string
	s := url.Storage{DB: conn}

	currentUrl, _ := s.GetUrlByID(urlId)
	if currentUrl != nil {
		text = "Окей нахуй, введи сюда новую конечную точнку для ссылки " + currentUrl.From
	} else {
		text = "Ошибка"
	}
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)
	client.NextMessage = "save_destination_link|" + strconv.FormatUint(uint64(currentUrl.ID), 10)
	c := telegramProfile.Storage{DB: conn}
	c.UpdateClient(client)

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
