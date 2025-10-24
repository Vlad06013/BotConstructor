package domainsTemplate

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ConfirmDeleteDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, domainId uint) external.TextMessage {

	d := domain.Storage{DB: conn}
	deleted := d.DeleteDomainByID(client.TgUserId, domainId)
	text := "Окей мы всё нахуй удалили. Иди нахуй"

	if !deleted {
		text = "Ошибка удаления. Иди нахуй"
	}

	buttons := tgbotapi.NewInlineKeyboardMarkup(
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
