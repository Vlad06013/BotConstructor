package domain

import "github.com/Vlad06013/BotConstructor.git/repository/url"

type Domains struct {
	ID     uint       `json:"id" gorm:"primary_key;column:id"`
	Domain string     `json:"domain" gorm:"column:domain;default:null"`
	Active bool       `json:"active" gorm:"column:active;default:false"`
	Urls   []url.Urls `json:"urls"`
}
