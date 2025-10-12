package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Storage struct {
	*gorm.DB
}

func (r *Storage) UpdateDomain(domain Domains) *Domains {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	r.Save(&Domains{
		ID:        domain.ID,
		ClientID:  domain.ClientID,
		Domain:    domain.Domain,
		Active:    domain.Active,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: dateTime,
	})
	return &domain
}

func (r *Storage) GetDomainsByClientID(tgID uint) ([]Domains, error) {
	var domains []Domains
	if err := r.Find(&domains, "client_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	return domains, nil
}

func (r *Storage) GetDomainByID(id uint) (*Domains, error) {
	var domain Domains
	if err := r.First(&domain, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &domain, nil
}

func (r *Storage) CreateDomain(id uint, domainName string) *Domains {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	domain := Domains{
		ClientID:  id,
		Domain:    domainName,
		Active:    false,
		Status:    "wait_connection",
		CreatedAt: dateTime,
		UpdatedAt: dateTime,
	}
	r.Create(&domain)
	return &domain
}

func (r *Storage) GetByNameAndClientId(domainName string, clientID uint) (*Domains, error) {
	var domain Domains
	if err := r.Where("domain = ? AND tg_user_id = ?", domainName, clientID).First(&domain).Error; err != nil {
		return nil, err
	}
	return &domain, nil
}

func (r *Storage) DeleteDomainByID(id uint) {
	r.Delete(&Domains{}, id)
}
