package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

func DeleteDomainMessage(client tgUser.Clients, conn *gorm.DB, domainId uint) api.TextMessage {

	s := domain.Storage{DB: conn}
	domain, _ := s.GetDomainByID(domainId)
	if domain == nil {
	}
	text := "Братишка ты уверен? Это полностью удалит информацию об этом редиректе из нашего бота и все ссылки которые были прикриплены к этому домену"

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "confirmDeleteDomain|"+strconv.FormatUint(uint64(domain.ID), 10)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "domainSettings"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
