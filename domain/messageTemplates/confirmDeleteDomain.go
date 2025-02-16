package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ConfirmDeleteDomainMessage(client tgUser.Clients, conn *gorm.DB, urlId uint) api.TextMessage {

	//s := url.Storage{DB: conn}
	d := domain.Storage{DB: conn}
	d.DeleteDomainByID(urlId)
	text := "Окей мы всё нахуй удалили. Иди нахуй"

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)

	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
