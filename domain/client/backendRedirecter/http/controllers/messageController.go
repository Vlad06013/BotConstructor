package controllers

import (
	"fmt"
	"net/http"

	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/domain/useCase"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Hello, world!"}`)

	mess := external.TextMessage{
		Text:   "Hello, world!",
		ChatId: 878108763,
		Bot:    useCase.GetBot(),
	}
	fmt.Print(mess.Text, mess.ChatId, mess.Bot)
	var output external.OutputMessage = mess
	output.Send()

}
