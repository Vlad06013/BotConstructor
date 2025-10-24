package url

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Vlad06013/BotConstructor.git/domain/infrastructure/external/ApiClientBackend"
	"github.com/jinzhu/gorm"
)

type Storage struct {
	*gorm.DB
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

func (r *Storage) CreateUrl(tgUserId int64, domainId uint, to string) (*Urls, error) {

	var urlCreated Urls
	var errorMes string
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("url")

	params := map[string]interface{}{
		"domain_id": strconv.FormatUint(uint64(domainId), 10),
		"to":        to,
	}

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Post(url, params, headers)

	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
	}

	if urlData, exists := dataMap["url"]; exists {
		err = json.Unmarshal(urlData, &urlCreated)
		if err != nil {
		}
	}

	if errorMessage, exists := dataMap["error"]; exists {
		_ = json.Unmarshal(errorMessage, &errorMes)
		return nil, fmt.Errorf(errorMes)
	}

	return &urlCreated, nil
}

func (r *Storage) UpdateUrlDestination(to string, urlId uint) {
	r.Model(&Urls{}).Where("id =?", urlId).Update("to", to)
}

func (r *Storage) UpdateUrlComment(description string, urlId uint) {
	r.Model(&Urls{}).Where("id =?", urlId).Update("description", description)
}

func (r *Storage) GetUrlsByClientID(tgUserId int64) ([]Urls, error) {

	var urls []Urls
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("url/get-for-client")

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Get(url, headers)
	err := json.Unmarshal(result.Data, &dataMap)

	if urlsData, exists := dataMap["urls"]; exists {
		err = json.Unmarshal(urlsData, &urls)
		if err != nil {
		}
	}

	return urls, nil

}
