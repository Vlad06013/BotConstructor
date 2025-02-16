package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

func ChooseDomainToShotLinkMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Выберит домен для которого сокращаем ссылку"
	s := domain.Storage{DB: conn}
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "cabinet"
	domains, _ := s.GetDomainsByClientID(client.ID)

	if len(domains) == 0 {
		text = "Нет подключенных доменов"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(domains)+1)
	for i := 0; i < len(domains); i++ {
		callbackData := "shotLink|" + strconv.FormatUint(uint64(domains[i].ID), 10)
		btnText := domains[i].Domain
		if domains[i].Active == true {
			btnText = "✅ " + btnText
		} else {
			btnText = "❌ " + btnText
		}
		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}

	rows[len(domains)] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
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
