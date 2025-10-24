package tariffsTemplate

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/payment"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func BuyTariff(client telegramProfile.TelegramProfile, conn *gorm.DB, tariffId uint) external.TextMessage {
	s := payment.Storage{DB: conn}
	createdPayment, _ := s.TariffPayment(client.TgUserId, tariffId)

	var text string
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	if createdPayment.PaymentUrl == "" && createdPayment.Status == "confirmed" {
		text = "<b>Тарифный план успешно подключен</b>\n\n"
	} else {
		text = "Ваша ссылка на оплату"

		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Оплатить", createdPayment.PaymentUrl),
		))
	}
	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
	))
	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	//text := "Ваша ссылка на оплату"
	//
	//buttons := tgbotapi.NewInlineKeyboardMarkup(
	//	tgbotapi.NewInlineKeyboardRow(
	//		tgbotapi.NewInlineKeyboardButtonURL("Оплатить", createdPayment.PaymentUrl),
	//	),
	//	tgbotapi.NewInlineKeyboardRow(
	//		tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
	//	),
	//)

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: keyboard,
	}
	return mess
}
