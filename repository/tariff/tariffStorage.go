package tariff

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

func (r *Storage) Get(tgUserId int64) ([]Tariffs, error) {
	var tariffs []Tariffs

	url := fmt.Sprintf("tariff")

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Get(url, headers)

	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// В `data` может быть не объект, а другой формат
		// Обработка ошибочного случая
	}
	// Проверка наличия ключа "tariffs"
	if tariffsData, exists := dataMap["tariffs"]; exists {
		err = json.Unmarshal(tariffsData, &tariffs)
		if err != nil {
			// обработка ошибки
		}

		//for _, item := range tariffs {
		//fmt.Printf("Тариф ID: %d, Название: %s\n", item.ID, item.Name)
		//}
	}

	return tariffs, nil

}

func (r *Storage) GetById(tgUserId int64, tariffId uint) (*Tariffs, error) {

	var tariff Tariffs
	url := fmt.Sprintf("tariff/show/" + strconv.FormatUint(uint64(tariffId), 10))

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Get(url, headers)

	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// Обработка ошибочного случая
	}

	if tariffData, exists := dataMap["tariff"]; exists {
		err = json.Unmarshal(tariffData, &tariff)
		if err != nil {
			return nil, nil
		}
		return &tariff, nil

	}
	return nil, nil
}

func (r *Storage) GetMy(tgUserId int64) (*MyTariff, error) {
	var myTariff MyTariff
	url := fmt.Sprintf("tariff/get-my")
	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Get(url, headers)

	fmt.Print(result)
	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// Обработка ошибочного случая
	}

	if tariffData, exists := dataMap["tariff"]; exists {
		err = json.Unmarshal(tariffData, &myTariff)
		if err != nil {
			return nil, nil
		}
		return &myTariff, nil

	}
	return nil, nil

}

func (r *Storage) GetFreeTariff(tgUserId int64) (*Tariffs, error) {

	var tariff Tariffs
	url := fmt.Sprintf("tariff/free-tariff")

	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Get(url, headers)

	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// Обработка ошибочного случая
	}

	if tariffData, exists := dataMap["tariff"]; exists {
		err = json.Unmarshal(tariffData, &tariff)
		if err != nil {
			return nil, nil
		}
		return &tariff, nil

	}
	return nil, nil
}

func (r *Storage) StartFreeTariff(tgUserId int64) (*Tariffs, error) {
	var connectedTariff Tariffs
	url := fmt.Sprintf("tariff/start-free-tariff")
	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}

	result := ApiClientBackend.Post(url, nil, headers)

	fmt.Print(result)
	var dataMap map[string]json.RawMessage
	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
		// Обработка ошибочного случая
	}

	if tariffData, exists := dataMap["tariff"]; exists {
		err = json.Unmarshal(tariffData, &connectedTariff)
		if err != nil {
			return nil, nil
		}
		return &connectedTariff, nil

	}
	return nil, nil

}
