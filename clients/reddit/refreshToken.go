package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// RefreshToken обновляет access_token, если он истек или при 403 или 401
func (c *RedditClient) RefreshToken() error {
	// Формируем параметры запроса для обновления токена
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", c.refreshToken)

	url := fmt.Sprintf("%s/api/v1/access_token", c.cfg.URL)
	// Создаем POST-запрос к Reddit API для обновления токена
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Устанавливаем базовую аутентификацию и заголовки
	req.SetBasicAuth(c.cfg.AuthData.ClientID, c.cfg.AuthData.ClientSecret)
	req.Header.Set("User-Agent", c.cfg.AuthData.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Выполняем запрос
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем статус ответа, чтобы убедиться, что запрос успешен
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to refresh token, status: %d", resp.StatusCode)
	}

	// Декодируем ответ и обновляем токен
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	// Обновляем токен, если он присутствует в ответе
	if newToken, ok := result["access_token"].(string); ok {
		c.token = newToken
		fmt.Println("🔄 Access token refreshed successfully!")
		return nil
	}

	return fmt.Errorf("no access token in response")
}
