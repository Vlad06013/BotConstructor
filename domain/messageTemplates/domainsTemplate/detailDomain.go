package domainsTemplate

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	//"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DetailDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, domainId uint) external.TextMessage {
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	urlLength := "0"
	s := domain.Storage{DB: conn}
	domainCurrent, _ := s.GetDomainByID(client.TgUserId, domainId)
	urls := domainCurrent.Urls
	urlLength = strconv.FormatUint(uint64(len(urls)), 10)

	text := "Сейчас на домене " + domainCurrent.Domain + " привязано ссылок:" + urlLength + ". Вот они:"

	if len(urls) == 0 {
		text = "Нет подключенных ссылок"
	}
	for i := 0; i < len(urls); i++ {
		btnText := urls[i].From + "|" + urls[i].Description
		if urls[i].Active == true {
			btnText = "✅ " + btnText
		} else {
			btnText = "❌ " + btnText
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(btnText, "detailLink|"+strconv.FormatUint(uint64(urls[i].ID), 10)),
		))
	}

	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить нахуй этот домен и все ссылки",
			"deleteDomain|"+strconv.FormatUint(uint64(domainId), 10)),
	))

	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "domainSettings"),
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
