package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DetailLinkMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, urlId uint) external.TextMessage {

	s := url.Storage{DB: conn}
	var text string
	currentUrl, _ := s.GetUrlByID(urlId)
	if currentUrl != nil {
		text = "Ссылка: <b>https://" + currentUrl.From + "</b>\n\nВедёт: <b>" + currentUrl.To + "</b>\n\n" + "Комментарий: \n" + currentUrl.Description
	} else {
		text = "Ошибка"
	}

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Изменить конечную точку", "changeLinkDestination|"+strconv.FormatUint(uint64(currentUrl.ID), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить эту ссылку", "deleteLink|"+strconv.FormatUint(uint64(currentUrl.ID), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Изменить комментарий", "changeComment|"+strconv.FormatUint(uint64(currentUrl.ID), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "linkSettings"),
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
