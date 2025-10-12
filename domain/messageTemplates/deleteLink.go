package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DeleteLinkMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, urlId uint) external.TextMessage {

	s := url.Storage{DB: conn}
	url, _ := s.GetUrlByID(urlId)
	if url == nil {
	}
	text := "Удалить ссылку https://" + url.Domain.Domain + "/" + url.From + "\n Которая ведет сюда: \n" + url.To

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "confirmDeleteLink|"+strconv.FormatUint(uint64(url.ID), 10)),
		),
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
