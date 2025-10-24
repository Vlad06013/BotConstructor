package linksTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	"github.com/jinzhu/gorm"
)

func WaitInputCommentShotLinkMessage(client telegramProfile.TelegramProfile, conn *gorm.DB) external.TextMessage {
	text := "Тут ты можешь оставить коментарий, чтобы не забыть куда эта всратая ссылка ведет."
	client.NextMessage = "save_input_comment_shot_link"
	s := telegramProfile.Storage{DB: conn}
	s.UpdateClient(client)
	mess := external.TextMessage{
		Text:   text,
		ChatId: client.TgUserId,
	}
	return mess
}
