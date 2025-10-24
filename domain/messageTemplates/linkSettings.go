package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func LinkSettingsMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
	text := "Вот все твои ебаные ссылки, которые ты сокращал:"

	s := url.Storage{DB: conn}
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	urls, _ := s.GetUrlsByClientID(client.TgUserId)
	if len(urls) == 0 {
		text = "Нету созданных ссылок"
	}
	for i := 0; i < len(urls); i++ {
		desc := "Нет описания"

		btnText := urls[i].From + "|" + desc

		if urls[i].Active == true {
			btnText = "✅ " + btnText
		} else {
			btnText = "❌ " + btnText
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(btnText, "detailLink|"+strconv.FormatUint(uint64(urls[i].ID), 10)),
		))
	}

	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
	))

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: keyboard,
	}
	return mess
}
