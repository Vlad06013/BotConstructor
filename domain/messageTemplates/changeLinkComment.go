package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

func ChangeLinkCommentMessage(client tgUser.Clients, conn *gorm.DB, urlId uint) api.TextMessage {

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
	c := tgUser.Storage{DB: conn}
	c.UpdateClient(client)

	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
