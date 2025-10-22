package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TariffsSettingsMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {

	text := "у вас нет подключенного тарифа"
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "cabinet"
	s := tariff.Storage{DB: conn}

	tariffs, _ := s.Get()
	myActiveTariffs, _ := s.GetMy(client.TgUserId)
	if myActiveTariffs != nil {
		text = "<b>Текущий тариф:</b> " + myActiveTariffs.Name +
			"\n<b>Стоимость:</b> " + strconv.Itoa(int(myActiveTariffs.Price)) + " " + myActiveTariffs.Currency +
			"\n<b>Кол-во доменов: </b>" + strconv.Itoa(int(myActiveTariffs.DomainsCount)) +
			"\n<b>Кол-во сокращений для одного домена: </b>" + strconv.Itoa(int(myActiveTariffs.LinksForDomainCount)) +
			"\n<b>Действителен до : </b>" + myActiveTariffs.End

	}

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

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: keyboard,
	}
	return mess
}
