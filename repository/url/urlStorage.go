package url

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Storage struct {
	*gorm.DB
}

//
//func (r *Storage) UpdateDomain(domain Urls) *Urls {
//
//	location, _ := time.LoadLocation("Europe/Moscow")
//	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
//
//	r.Save(&Urls{
//		ID:        domain.ID,
//		TgUserId:  domain.TgUserId,
//		Domain:    domain.Domain,
//		Active:    domain.Active,
//		Status:    domain.Status,
//		CreatedAt: domain.CreatedAt,
//		UpdatedAt: dateTime,
//	})
//	return &domain
//}

func (r *Storage) GetUrlsByClientID(tgID uint) ([]Urls, error) {
	var urls []Urls
	if err := r.Find(&urls, "tg_user_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	r.Preload("Domain").Find(&urls)
	return urls, nil
}

func (r *Storage) GetUrlByID(id uint) (*Urls, error) {
	var urls Urls
	if err := r.First(&urls, "id = ?", id).Error; err != nil {
		return nil, err
	}
	r.Preload("Domain").First(&urls)
	return &urls, nil
}

func (r *Storage) GetUrlByDomainID(domainId uint) ([]Urls, error) {
	var urls []Urls
	if err := r.Find(&urls, "domain_id = ?", domainId).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (r *Storage) DeleteUrlByID(id uint) {
	r.Delete(&Urls{}, id)
}

func (r *Storage) CreateUrl(url Urls) *Urls {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	domain := Urls{
		DomainId:    url.DomainId,
		ClientId:    url.ClientId,
		From:        url.From,
		To:          url.To,
		Description: url.Description,
		Active:      url.Active,
		CreatedAt:   dateTime,
		UpdatedAt:   dateTime,
	}
	r.Create(&domain)
	return &domain
}

func (r *Storage) UpdateUrlDestination(to string, urlId uint) {
	r.Model(&Urls{}).Where("id =?", urlId).Update("to", to)
}

func (r *Storage) UpdateUrlComment(description string, urlId uint) {
	r.Model(&Urls{}).Where("id =?", urlId).Update("description", description)
}
