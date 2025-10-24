package domain

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

func (r *Storage) GetDomainsByClientID(tgUserId int64) ([]Domains, error) {

	var domains []Domains
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("domain/get-for-client")

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Get(url, headers)
	err := json.Unmarshal(result.Data, &dataMap)

	if tariffsData, exists := dataMap["domains"]; exists {
		err = json.Unmarshal(tariffsData, &domains)
		if err != nil {
		}
	}

	return domains, nil
}

func (r *Storage) GetDomainByID(tgUserId int64, id uint) (*Domains, error) {

	var domain Domains
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("domain/show/" + strconv.FormatUint(uint64(id), 10))

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Get(url, headers)
	err := json.Unmarshal(result.Data, &dataMap)

	if tariffsData, exists := dataMap["domain"]; exists {
		err = json.Unmarshal(tariffsData, &domain)
		if err != nil {
			return nil, err
		}
	}

	return &domain, nil
}

func (r *Storage) CreateDomain(tgUserId int64, domainName string) (*Domains, error) {

	var domain Domains
	var errorMes string
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("domain")

	params := map[string]interface{}{
		"name": domainName,
	}

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Post(url, params, headers)

	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
	}

	if domainData, exists := dataMap["domain"]; exists {
		err = json.Unmarshal(domainData, &domain)
		if err != nil {
		}
	}

	if errorMessage, exists := dataMap["error"]; exists {
		_ = json.Unmarshal(errorMessage, &errorMes)
		return nil, fmt.Errorf(errorMes)
	}

	return &domain, nil
}

func (r *Storage) DeleteDomainByID(tgUserId int64, id uint) bool {
	var dataMap map[string]json.RawMessage
	var resultDeleted bool
	url := fmt.Sprintf("domain/" + strconv.FormatUint(uint64(id), 10))

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Delete(url, headers)
	err := json.Unmarshal(result.Data, &dataMap)

	if tariffsData, exists := dataMap["domain_deleted"]; exists {
		err = json.Unmarshal(tariffsData, &resultDeleted)
		if err != nil {
			return false
		}
	}

	return true
}
