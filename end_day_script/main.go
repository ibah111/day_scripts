package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const webhookURL = "https://chat.nbkfinance.ru/rest/4018/e28d2pdnglnkmdms/timeman.close"

type BitrixRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

func main() {
	// Создаем запрос для окончания рабочего дня
	request := BitrixRequest{
		Method: "timeman.entry.finish",
		Params: map[string]interface{}{
			"USER_ID": 1, // Замените на нужный ID пользователя
		},
	}

	// Преобразуем запрос в JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Ошибка при создании JSON: %v\n", err)
		return
	}

	// Отправляем POST запрос
	resp, err := http.Post(webhookURL+"/timeman.close", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %v\n", err)
		return
	}

	fmt.Printf("Время окончания рабочего дня: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("Ответ от Битрикс24: %s\n", string(body))
}
