package tariffsTemplate

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TariffsSettingsMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
	var text string
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	s := tariff.Storage{DB: conn}

	tariffs, _ := s.Get(client.TgUserId)
	myTariff, nonActiveTariffError := s.GetMy(client.TgUserId)

	if myTariff != nil {
		text = "<b>Текущий тариф:</b> " + myTariff.Name +
			"\n<b>Стоимость:</b> " + strconv.Itoa(int(myTariff.Price)) + " " + myTariff.Currency +
			"\n<b>Кол-во доменов: </b>" + strconv.Itoa(int(myTariff.DomainsCount)) +
			"\n<b>Кол-во сокращений: </b>" + strconv.Itoa(int(myTariff.LinksLimit)) +
			"\n<b>Действителен до : </b>" + myTariff.End
	}
	if nonActiveTariffError != nil {
		text = nonActiveTariffError.Error()
	}

	for i := 0; i < len(tariffs); i++ {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(tariffs[i].Name, "detailTariff|"+strconv.FormatUint(uint64(tariffs[i].ID), 10)),
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
