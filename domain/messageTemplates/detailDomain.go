package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DetailDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, domainId uint) external.TextMessage {
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	s := domain.Storage{DB: conn}
	u := url.Storage{DB: conn}
	backBtnCB := "domainSettings"
	cabinetBtnCB := "cabinet"
	deleteDomainBtnCB := "deleteDomain|" + strconv.FormatUint(uint64(domainId), 10)
	domain, _ := s.GetDomainByID(domainId)
	urls, _ := u.GetUrlByDomainID(domain.ID)
	urlLength := strconv.FormatUint(uint64(len(urls)), 10)

	text := "Сейчас на домене " + domain.Domain + " привязано ссылок:" + urlLength + ". Вот они:"

	if len(urls) == 0 {
		text = "Нет подключенных ссылок"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(urls)+2)
	for i := 0; i < len(urls); i++ {
		callbackData := "detailLink|" + strconv.FormatUint(uint64(urls[i].ID), 10)
		btnText := "https://" + domain.Domain + "/" + urls[i].From + "|" + urls[i].Description
		if urls[i].Active == true {
			btnText = "✅ " + btnText
		} else {
			btnText = "❌ " + btnText
		}
		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(urls)] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
		Text:         "Удалить нахуй этот домен и все ссылки",
		CallbackData: &deleteDomainBtnCB,
	})
	rows[len(urls)+1] = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: &backBtnCB,
		},
		tgbotapi.InlineKeyboardButton{
			Text:         "В кабинет",
			CallbackData: &cabinetBtnCB,
		},
	)
	//rows[len(urls)] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
	//	Text:         "Удалить нахуй этот домен и все ссылки",
	//	CallbackData: &deleteDomainBtnCB,
	//})
	//rows[len(urls)+1] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
	//	Text:         "Назад",
	//	CallbackData: &backBtnCB,
	//})
	buttons = rows

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: keyboard,
	}
	return mess
}
