package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// PostComment отправляет комментарий с текстом на пост, идентифицируемый по thingID.
// Для thingID используйте формат "t3_<postID>" (например, "t3_abcdef").
// Функция выполняет POST-запрос к Reddit API.
func (c *RedditClient) PostComment(thingID, text string) error {
	// Формируем параметры запроса
	data := url.Values{}
	data.Set("thing_id", thingID)
	data.Set("text", text)
	data.Set("api_type", "json")

	// Создаем POST-запрос к Reddit API для отправки комментария
	req, err := http.NewRequest("POST", "https://oauth.reddit.com/api/comment", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Устанавливаем заголовки: авторизация с помощью bearer-токена и User-Agent
	req.Header.Set("Authorization", "bearer "+c.token)
	req.Header.Set("User-Agent", c.cfg.AuthData.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Выполняем запрос
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		// Если токен истек, обновляем его и повторяем запрос
		if err := c.RefreshToken(); err != nil {
			return err
		}

		// Повторяем запрос с обновленным токеном
		req.Header.Set("Authorization", "bearer "+c.token)
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post comment, status: %d", resp.StatusCode)
	}

	// Декодируем ответ в структуру RedditCommentResponse
	var result RedditCommentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	// Проверяем наличие ошибок в ответе Reddit
	if len(result.JSON.Errors) > 0 {
		return fmt.Errorf("failed to post comment, errors: %v", result.JSON.Errors)
	}

	fmt.Println("💬 Комментарий успешно отправлен!")
	return nil
}
