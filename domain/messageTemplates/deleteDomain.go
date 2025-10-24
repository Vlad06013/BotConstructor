package messageTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func DeleteDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, domainId uint) external.TextMessage {

	s := domain.Storage{DB: conn}
	currentDomain, _ := s.GetDomainByID(client.TgUserId, domainId)
	if currentDomain == nil {
	}
	text := "Братишка ты уверен? Это полностью удалит информацию об этом редиректе из нашего бота и все ссылки которые были прикриплены к этому домену"

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "confirmDeleteDomain|"+strconv.FormatUint(uint64(currentDomain.ID), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "domainSettings"),
		),
	)

	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
