package linksTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ChangeLinkCommentMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, urlId uint) external.TextMessage {

	s := url.Storage{DB: conn}
	url, _ := s.GetUrlByID(urlId)
	if url == nil {

	}
	text := "Окей нахуй, введи сюда комментарий для ссылки \n" + url.From
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)
	client.NextMessage = "save_comment_link|" + strconv.FormatUint(uint64(url.ID), 10)
	c := telegramProfile.Storage{DB: conn}
	c.UpdateClient(client)

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
