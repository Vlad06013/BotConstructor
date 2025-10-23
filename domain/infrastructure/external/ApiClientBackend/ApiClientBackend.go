package ApiClientBackend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Message string `json:"message"`
	//Data    []map[string]interface{} `json:"data"`
	//Data map[string]interface{} `json:"data"`
	//Data interface{} `json:"data"`
	Data json.RawMessage `json:"data"`
}

var url = "http://127.0.0.1/api/telegram-bot/v1/"

func send(uri string, method string, body interface{}, headers interface{}) (response *http.Response, err error) {
	fmt.Println("отправка запроса:", method, " ", url+uri)
	var resp *http.Response

	jsonData, err := json.Marshal(body)

	req, err := http.NewRequest(method, url+uri, bytes.NewBuffer(jsonData))
	fmt.Println(jsonData)

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	headersMap, _ := headers.(map[string]interface{})
	value, exists := headersMap["auth-telegram-id"]

	if exists {
		req.Header.Set("auth-telegram-id", value.(string)) // замените на нужное значение
	}

	// Отправляем запрос
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		// обработка ошибки
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Получен неожиданный статус:", resp.Status)
		defer resp.Body.Close()

		return
	}
	return resp, err
}

func decode(resp *http.Response) (response ApiResponse, err error) {

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var result ApiResponse
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	return result, err
}

func Get(uri string, headers interface{}) ApiResponse {
	resp, _ := send(uri, http.MethodGet, nil, headers)
	result, _ := decode(resp)

	return result
}

func Post(uri string, body interface{}, headers interface{}) ApiResponse {
	resp, _ := send(uri, http.MethodPost, body, headers)
	result, _ := decode(resp)

	return result
}
