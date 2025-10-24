package tariffsTemplate

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/tariff"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TariffConnectMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, tariffId uint) external.TextMessage {
	s := tariff.Storage{DB: conn}
	currentTariff, _ := s.GetById(client.TgUserId, tariffId)

	text := "Будет подключен тариф " + currentTariff.Name + ". К оплате: " + strconv.Itoa(int(currentTariff.Price)) + " " + currentTariff.Currency

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подключаем тариф", "buyTariff|"+strconv.FormatUint(uint64(tariffId), 10)),
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
