package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func StartFreePeriodMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {

	s := tariff.Storage{DB: conn}
	freeTariff, _ := s.StartFreeTariff(client.TgUserId)

	text := "Активирован<b> Тариф:</b> " + freeTariff.Name +
		"\n<b>Кол-во доменов: </b>" + strconv.Itoa(int(freeTariff.DomainsCount)) +
		"\n<b>Кол-во сокращений для одного домена: </b>" + strconv.Itoa(int(freeTariff.LinksLimit))

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Давайте начнем работу", "new_domain"),
		),
	)

	mess := external.TextMessage{
		Text:    text,
		Buttons: buttons,
		ChatId:  client.TgUserId,
	}
	return mess
}
