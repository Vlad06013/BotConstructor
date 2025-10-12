package ApiClientBackend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiResponse struct {
	Message string `json:"message"`
	//Data    []map[string]interface{} `json:"data"`
	Data map[string]interface{} `json:"data"`
}

var url = "http://127.0.0.1/api/telegram-bot/v1/"

func send(uri string, method string, body interface{}) (response *http.Response, err error) {
	fmt.Println("отправка запроса:", url+uri)
	var resp *http.Response

	if method == "POST" {
		// Кодируем параметры в JSON
		jsonData, err := json.Marshal(body)
		if err != nil {
			log.Fatal(err)
		}
		resp, err = http.Post(url+uri, "application/json", bytes.NewBuffer(jsonData))
	}
	if method == "GET" {
		resp, err = http.Get(url + uri)
		if err != nil {
			fmt.Println("Ошибка при отправке запроса:", err)
			return
		}
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

func Get(uri string) ApiResponse {
	resp, _ := send(uri, http.MethodGet, nil)
	result, _ := decode(resp)

	return result
}

func Post(uri string, body interface{}) ApiResponse {
	resp, _ := send(uri, http.MethodPost, body)
	result, _ := decode(resp)

	return result
}
