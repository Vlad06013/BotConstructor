package linksTemplates

import (
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/jinzhu/gorm"
)

func WaitInputShotLinkMessage(client telegramProfile.TelegramProfile, conn *gorm.DB, domainId uint) external.TextMessage {
	text := "Пришли боту ссылку, которую ты пидор хочешь сократить"
	//client.NextMessage = "save_input_link|"+domainId
	client.NextMessage = "save_input_link|" + strconv.FormatUint(uint64(domainId), 10)

	s := telegramProfile.Storage{DB: conn}
	s.UpdateClient(client)
	mess := external.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
