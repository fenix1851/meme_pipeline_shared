package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *OpenAIClient) SendChatGPTRequest(messages []OpenAIMessage, model string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	requestDTO := ChatGPTRequest{
		Model:       model,
		Messages:    messages,
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
	}

	// Сериализация DTO в JSON
	requestBody, err := json.Marshal(requestDTO)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации JSON запроса: %v", err)
	}

	// Создание HTTP-запроса
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка от API OpenAI: %s", body)
	}

	// Декодирование ответа от OpenAI
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("OpenAI не вернул ответ")
	}

	return response.Choices[0].Message.Content, nil
}
