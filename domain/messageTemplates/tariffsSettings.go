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
	s := tariff.Storage{DB: conn}

	tariffs, _ := s.Get(client.TgUserId)

	for i := 0; i < len(tariffs); i++ {
		if tariffs[i].IsMy {
			text = "<b>Текущий тариф:</b> " + tariffs[i].Name +
				"\n<b>Стоимость:</b> " + strconv.Itoa(int(tariffs[i].Price)) + " " + tariffs[i].Currency +
				"\n<b>Кол-во доменов: </b>" + strconv.Itoa(int(tariffs[i].DomainsCount)) +
				"\n<b>Кол-во сокращений: </b>" + strconv.Itoa(int(tariffs[i].LinksLimit)) +
				"\n<b>Действителен до : </b>" + tariffs[i].End

		} else {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(tariffs[i].Name, "detailTariff|"+strconv.FormatUint(uint64(tariffs[i].ID), 10)),
			))
		}
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
