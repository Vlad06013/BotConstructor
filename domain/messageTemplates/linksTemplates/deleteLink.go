package linksTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DeleteLinkMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, urlId uint) external.TextMessage {

	var text string
	s := url.Storage{DB: conn}
	currentUrl, _ := s.GetUrlByID(urlId)
	if currentUrl == nil {
		text = "Ошибка"

	} else {
		text = "Удалить ссылку " + currentUrl.From + "\n Которая ведет сюда: \n" + currentUrl.To
	}
	//text := "Удалить ссылку https://" + url.Domain.Domain + "/" + currentUrl.From + "\n Которая ведет сюда: \n" + currentUrl.To

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "confirmDeleteLink|"+strconv.FormatUint(uint64(currentUrl.ID), 10)),
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
