package postgres

import (
	"fmt"
	config "github.com/Vlad06013/BotConstructor.git"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection(dbc config.ConfigDBPostgres) *gorm.DB {

	db, err := gorm.Open("postgres", "host="+dbc.DB_HOST+" user="+dbc.DB_USERNAME+" password="+dbc.DB_PASSWORD+" dbname="+dbc.DB_NAME+" port="+dbc.DB_PORT+" sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("Не удалось подключиться к базе данных ")
	}
	//db.AutoMigrate(&Bot{})
	fmt.Println("DB connected")
	return db

}
