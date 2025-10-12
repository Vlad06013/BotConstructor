package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/jinzhu/gorm"
)

func WaitInputDomainMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
	text := "Теперь пришлите домен, который вы привязали к своему личному кабинету в ответ на это сообщение. \n Дoмен пришлите в виде vladgey.com"
	client.NextMessage = "save_input_domain"
	s := telegramProfile.Storage{DB: conn}
	s.UpdateClient(client)
	mess := external.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
