package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	"github.com/jinzhu/gorm"
)

func WaitInputCommentShotLinkMessage(client tgUser.Clients, conn *gorm.DB) api.TextMessage {
	text := "Тут ты можешь оставить коментарий, чтобы не забыть куда эта всратая ссылка ведет."
	client.NextMessage = "save_input_comment_shot_link"
	s := tgUser.Storage{DB: conn}
	s.UpdateClient(client)
	mess := api.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
