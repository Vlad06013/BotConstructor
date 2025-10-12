package client

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Storage struct {
	*gorm.DB
}

func (r *Storage) CreateClient() *Client {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	client := Client{
		CreatedAt: dateTime,
		UpdatedAt: dateTime,
	}
	r.Create(&client)
	return &client
}
