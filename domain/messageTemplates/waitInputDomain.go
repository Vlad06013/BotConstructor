package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	"github.com/jinzhu/gorm"
)

func WaitInputDomainMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Теперь пришлите домен, который вы привязали к своему личному кабинету в ответ на это сообщение. \n Дoмен пришлите в виде vladgey.com"
	client.NextMessage = "save_input_domain"
	s := tgUser.Storage{DB: conn}
	s.UpdateClient(client)
	mess := api.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
