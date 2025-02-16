package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

func TariffsSettingsMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Плоти деньги чорт блять"

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "cabinet"
	s := tariff.Storage{DB: conn}

	tariffs, _ := s.Get()

	if len(tariffs) == 0 {
		text = "Нет актиывных тарифов"
	}
	//rows := make([][]tgbotapi.InlineKeyboardButton, 1)
	rows := make([][]tgbotapi.InlineKeyboardButton, len(tariffs)+1)

	for i := 0; i < len(tariffs); i++ {
		callbackData := "detailTariff|" + strconv.FormatUint(uint64(tariffs[i].ID), 10)
		btnText := tariffs[i].Name
		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(tariffs)] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
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
