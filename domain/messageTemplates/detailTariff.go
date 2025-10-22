package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DetailTariffMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, tariffId uint) external.TextMessage {

	s := tariff.Storage{DB: conn}
	currentTariff, _ := s.GetById(tariffId)
	text := "<b>Тариф:</b> " + currentTariff.Name +
		"\n<b>Стоимость:</b> " + strconv.Itoa(int(currentTariff.Price)) + " " + currentTariff.Currency +
		"\n<b>Кол-во доменов: </b>" + strconv.Itoa(int(currentTariff.DomainsCount)) +
		"\n<b>Кол-во сокращений для одного домена: </b>" + strconv.Itoa(int(currentTariff.LinksForDomainCount))

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подключить тариф", "tariffConnect|"+strconv.FormatUint(uint64(tariffId), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "tariffSettings"),
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
