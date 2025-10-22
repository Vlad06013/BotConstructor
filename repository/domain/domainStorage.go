package domain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Vlad06013/BotConstructor.git/domain/infrastructure/external/ApiClientBackend"
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

func (r *Storage) CreateDomain(tgUserId int64, domainName string) *Domains {

	url := fmt.Sprintf("domain")

	params := map[string]interface{}{
		"name": domainName,
	}

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	//
	result := ApiClientBackend.Post(url, params, headers)
	var domain Domains

	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// Обработка ошибочного случая
	}

	if domainData, exists := dataMap["domain"]; exists {
		err = json.Unmarshal(domainData, &domain)
		if err != nil {
			// обработка ошибки
		}
	}

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
