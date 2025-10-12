package url

import "github.com/Vlad06013/BotConstructor.git/repository/domain"

type Urls struct {
	ID          uint            `json:"id" gorm:"primary_key;column:id"`
	DomainId    uint            `json:"domain_id" gorm:"column:domain_id"`
	From        string          `json:"from" gorm:"column:from"`
	To          string          `json:"to" gorm:"column:to;unique"`
	Description string          `json:"description" gorm:"column:description"`
	Active      bool            `json:"active" gorm:"column:active;default:false"`
	CreatedAt   string          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string          `json:"updated_at" gorm:"column:updated_at"`
	Domain      *domain.Domains `gorm:"foreignKey:DomainId"`
}
