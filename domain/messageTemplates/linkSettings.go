package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

func LinkSettingsMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Вот все твои ебаные ссылки, которые ты сокращал:"
	s := url.Storage{DB: conn}
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "cabinet"
	urls, _ := s.GetUrlsByClientID(client.ID)

	if len(urls) == 0 {
		text = "Нету созданных ссылок"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(urls)+1)
	for i := 0; i < len(urls); i++ {
		callbackData := "detailLink|" + strconv.FormatUint(uint64(urls[i].ID), 10)
		btnText := "https://" + urls[i].Domain.Domain + "/" + urls[i].From + "|" + urls[i].Description
		if urls[i].Active == true {
			btnText = "✅ " + btnText
		} else {
			btnText = "❌ " + btnText
		}
		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}

	rows[len(urls)] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
		Text:         "В кабинет",
		CallbackData: &backBtnCB,
	})
	buttons = rows

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: keyboard,
	}
	return mess
}
