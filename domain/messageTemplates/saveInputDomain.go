package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func SaveInputDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, message *tgbotapi.Message) external.TextMessage {
	s := domain.Storage{DB: conn}
	createdDomain, err := s.CreateDomain(client.TgUserId, message.Text)

	var text string
	if err != nil {
		text = err.Error()
	} else {
		text = "Заебись братишка, ты сохранили твой домен <b>" + createdDomain.Domain + " </b> теперь можешь пойти расслабиться поесть пряников или подрочить на хентай, или чем вы там маркетологи ебаные занимаетесь.\n " +
			"Когда домен привяжется, тебе петуху придет уведомление от бота.Что бы эта хуйня заработа подключи тариф в кабинете"
	}

	client.NextMessage = ""
	c := telegramProfile.Storage{DB: conn}
	c.UpdateClient(client)

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В личный кабинет", "cabinet"),
		),
	)
	mess := external.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
