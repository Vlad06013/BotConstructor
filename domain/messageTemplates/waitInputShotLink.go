package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	"github.com/jinzhu/gorm"
	"strconv"
)

func WaitInputShotLinkMessage(client tgUser.Clients, conn *gorm.DB, domainId uint) api.TextMessage {
	text := "Пришли боту ссылку, которую ты пидор хочешь сократить"
	//client.NextMessage = "save_input_link|"+domainId
	client.NextMessage = "save_input_link|" + strconv.FormatUint(uint64(domainId), 10)

	s := tgUser.Storage{DB: conn}
	s.UpdateClient(client)
	mess := api.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
