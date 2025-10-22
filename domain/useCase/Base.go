package useCase

import (
	"strconv"
	"strings"

	"github.com/Vlad06013/BotConstructor.git/domain/messageTemplates"
	"github.com/Vlad06013/BotConstructor.git/domain/module/external"
	"github.com/Vlad06013/BotConstructor.git/repository/telegramProfile"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func send(messageTemplate *external.TextMessage, botApi tgbotapi.BotAPI) *tgbotapi.Message {
	messageTemplate.Bot = botApi
	var output external.OutputMessage = messageTemplate
	return output.Send()
}

func deleteMessage(messageId int, chatID int64, botApi tgbotapi.BotAPI) {
	msg := external.DeleteMessage{
		MessageID: messageId,
		Bot:       botApi,
		ChatId:    chatID,
	}
	msg.DeleteMessage()
}

func getMessageTemplate(client telegramProfile.TelegramProfile, conn *gorm.DB, message *tgbotapi.Message) external.TextMessage {

	if client.NextMessage != "" {
		return findTemplate(client.NextMessage, client, conn, message)
	}
	if message.Text == "cabinet" {
		return findTemplate(message.Text, client, conn, message)
	} else {
		return messageTemplates.StartMessage(client, conn)
	}
}

func findTemplate(callBack string, client telegramProfile.TelegramProfile, conn *gorm.DB, message *tgbotapi.Message) external.TextMessage {

	if callBack == "start" {
		return messageTemplates.StartMessage(client, conn)
	}

	if callBack == "new_domain" {
		return messageTemplates.NewDomainMessage(client, conn)
	}

	if callBack == "connect_domain" {
		return messageTemplates.ConnectDomainMessage(client, conn)
	}

	if callBack == "wait_input_domain" {
		return messageTemplates.WaitInputDomainMessage(client, conn)
	}

	if callBack == "save_input_domain" {
		return messageTemplates.SaveInputDomainMessage(client, conn, message)
	}

	if callBack == "cabinet" {
		return messageTemplates.CabinetMessage(client, conn)
	}

	if callBack == "chooseDomainToShotLink" {
		return messageTemplates.ChooseDomainToShotLinkMessage(client, conn)
	}

	if strings.Contains(callBack, "shotLink") {
		res := strings.Split(callBack, "|")
		u64, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.WaitInputShotLinkMessage(client, conn, uint(u64))
	}

	if strings.Contains(callBack, "save_input_link") {
		res := strings.Split(callBack, "|")
		domainId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.SaveInputShotLinkMessage(client, conn, message, uint(domainId))
	}
	if callBack == "linkSettings" {
		return messageTemplates.LinkSettingsMessage(client, conn)
	}

	if callBack == "domainSettings" {
		return messageTemplates.DomainSettingsMessage(client, conn)
	}

	if callBack == "tariffSettings" {
		return messageTemplates.TariffsSettingsMessage(client, conn)
	}

	//if callBack == "tariffConnect" {
	//	return messageTemplates.TariffConnectMessage(client, conn)
	//}

	if strings.Contains(callBack, "detailLink") {
		res := strings.Split(callBack, "|")
		u64, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.DetailLinkMessage(client, conn, uint(u64))
	}

	if strings.Contains(callBack, "detailDomain") {
		res := strings.Split(callBack, "|")
		u64, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.DetailDomainMessage(client, conn, uint(u64))
	}
	if strings.Contains(callBack, "detailTariff") {
		res := strings.Split(callBack, "|")
		u64, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.DetailTariffMessage(client, conn, uint(u64))
	}

	if strings.Contains(callBack, "changeLinkDestination") {
		res := strings.Split(callBack, "|")
		u64, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.ChangeLinkDestinationMessage(client, conn, uint(u64))
	}

	if strings.Contains(callBack, "save_destination_link") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.SaveLinkDestinationMessage(client, conn, message, uint(urlId))
	}

	if strings.Contains(callBack, "deleteLink") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.DeleteLinkMessage(client, conn, uint(urlId))
	}

	if strings.Contains(callBack, "confirmDeleteLink") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.ConfirmDeleteLinkMessage(client, conn, uint(urlId))
	}

	if strings.Contains(callBack, "confirmDeleteDomain") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.ConfirmDeleteDomainMessage(client, conn, uint(urlId))
	}

	if strings.Contains(callBack, "changeComment") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.ChangeLinkCommentMessage(client, conn, uint(urlId))
	}

	if strings.Contains(callBack, "save_comment_link") {
		res := strings.Split(callBack, "|")
		urlId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.SaveInputCommentShotLinkMessage(client, conn, message, uint(urlId))
	}

	if strings.Contains(callBack, "deleteDomain") {
		res := strings.Split(callBack, "|")
		domainId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.DeleteDomainMessage(client, conn, uint(domainId))
	}

	if strings.Contains(callBack, "tariffConnect") {
		res := strings.Split(callBack, "|")
		tariffId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.TariffConnectMessage(client, conn, uint(tariffId))
	}

	if strings.Contains(callBack, "createOrder") {
		res := strings.Split(callBack, "|")
		tariffId, _ := strconv.ParseUint(res[1], 10, 32)
		return messageTemplates.CreateOrder(client, conn, uint(tariffId))
	}

	return messageTemplates.StartMessage(client, conn)
}
