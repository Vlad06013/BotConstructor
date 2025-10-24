package payment

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

func (r *Storage) TariffPayment(tgUserId int64, tariffId uint) (Payment, error) {
	var payment Payment
	var dataMap map[string]json.RawMessage

	url := fmt.Sprintf("order/buy-tariff")
	body := map[string]interface{}{
		"tariff_id": tariffId,
	}
	headers := map[string]interface{}{
		"auth-telegram-id": strconv.FormatUint(uint64(tgUserId), 10),
	}
	result := ApiClientBackend.Post(url, body, headers)

	err := json.Unmarshal(result.Data, &dataMap)
	if err != nil {
	}

	if paymentData, exists := dataMap["payment"]; exists {
		err = json.Unmarshal(paymentData, &payment)
		if err != nil {
		}
	}

	return payment, nil
}
