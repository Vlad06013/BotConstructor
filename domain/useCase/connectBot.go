package useCase

import (
	"fmt"
	"log"

	config "github.com/Vlad06013/BotConstructor.git"
	"github.com/Vlad06013/BotConstructor.git/postgres"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
)

var Bot tgbotapi.BotAPI

func Listen() {

	Bot, err := tgbotapi.NewBotAPI("8424114449:AAHxP8pajRHHC-C4OpaN4yl-NvMNlzJ-dzQ")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	conn := ConnectDB()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			fmt.Println(update.Message.Text)
			//var output api.OutputMessage = api.TextMessage{
			//	Text:    saa,
			//	Bot:     *botApi,
			//	Buttons: nil,
			//	ChatId:  update.Message.Chat.ID,
			//}
			//output.Send()
			TextMessage(update.Message, *Bot, conn)
		}
		if update.CallbackQuery != nil {
			CallBackQuery(update.CallbackQuery, *Bot, conn)
			//bu.CallbackQueryMessageHandler(bot, update.CallbackQuery)
		}
		//if update.MyChatMember != nil {
		//	ReadMyChatMember(db,)
		//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
		//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		//}
	}
}

func GetBot() tgbotapi.BotAPI {
	return Bot
}

func ConnectDB() *gorm.DB {

	var cfg config.ConfigDBPostgres
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}
	var conn = postgres.NewConnection(cfg)
	return conn
}
