package messageTemplates

import (
	"github.com/Vlad06013/BotConstructor.git/domain/module/api"
	"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/url"
	"strconv"

	//"github.com/Vlad06013/BotConstructor.git/repository/domain"
	"github.com/Vlad06013/BotConstructor.git/repository/tgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"math/rand"
	urlParsing "net/url"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func parseUrl(s string) (string, string) {
	u, err := urlParsing.Parse(s)

	if err != nil {
		panic(err)
	}

	return u.Host, u.Path
}

func SaveInputShotLinkMessage(client tgUser.Clients, conn *gorm.DB, message *tgbotapi.Message, domainId uint) api.TextMessage {

	s := url.Storage{DB: conn}
	host, _ := parseUrl(message.Text)

	if host == "" {
		text := "Неверная ссылка:  <b>" + message.Text + " </b> \n"

		mess := api.TextMessage{
			Text:   text,
			ChatId: client.TgUserId,
		}
		return mess
	}
	d := domain.Storage{DB: conn}
	domainFound, _ := d.GetDomainByID(domainId)

	//domainFound, _ := d.GetByNameAndClientId(host, client.ID)
	//
	if domainFound == nil {
		text := "Не найден подключенный домен:  <b>" + host + " </b> \n"

		mess := api.TextMessage{
			Text:   text,
			ChatId: client.TgUserId,
		}
		return mess
	}
	shotLink := randSeq(6)

	newUrl := url.Urls{
		DomainId:    domainId,
		ClientId:    client.ID,
		From:        shotLink,
		To:          message.Text,
		Description: "",
		Active:      false,
	}
	s.CreateUrl(newUrl)

	client.NextMessage = ""
	c := tgUser.Storage{DB: conn}
	c.UpdateClient(client)

	text := "Вот твоя ссылка пидарас:  <b>https://" + domainFound.Domain + "/" + shotLink + " </b> \n" +
		"Теперь она ведет на: \n <b>" + message.Text + " </b> \n "

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сократить еще одну ссылку", "shotLink|"+strconv.FormatUint(uint64(domainFound.ID), 10)),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("Оставить комментарий к этой ссылке", "wait_input_comment_shot_link"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В кабинет", "cabinet"),
		),
	)
	mess := api.TextMessage{
		Text:    text,
		ChatId:  client.TgUserId,
		Buttons: buttons,
	}
	return mess
}
