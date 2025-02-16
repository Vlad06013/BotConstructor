package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TextMessage struct {
	Text string
	Bot  tgbotapi.BotAPI
	//Buttons []tgbotapi.InlineKeyboardButton
	Buttons tgbotapi.InlineKeyboardMarkup
	ChatId  int64
}
type DeleteMessage struct {
	MessageID int
	Bot       tgbotapi.BotAPI
	ChatId    int64
}

type OutputMessage interface {
	Send() *tgbotapi.Message
}

func (o TextMessage) Send() *tgbotapi.Message {

	msg := tgbotapi.NewMessage(o.ChatId, o.Text)
	msg.ParseMode = "HTML"
	buttons := o.Buttons

	if len(buttons.InlineKeyboard) != 0 {
		msg.ReplyMarkup = buttons
	}

	res, err := o.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}

	return &res
}
func (o DeleteMessage) DeleteMessage() {
	msg := tgbotapi.NewDeleteMessage(o.ChatId, o.MessageID)
	o.Bot.Send(msg)
}
