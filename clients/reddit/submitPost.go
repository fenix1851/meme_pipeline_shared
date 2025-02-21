package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SubmitPost отправляет пост в Reddit. Если post.URL не пустой, то отправляется link-пост,
// иначе отправляется текстовый (self) пост.
func (c *RedditClient) SubmitPost(post PostData) (subredditLink string, err error) {
	// Формируем параметры запроса
	data := url.Values{}
	data.Set("sr", post.Subreddit)
	data.Set("title", post.Title)
	data.Set("api_type", "json")
	data.Set("kind", "link")
	data.Set("url", post.Url)

	// Создаем POST-запрос к Reddit API для отправки поста
	req, err := http.NewRequest("POST", "https://oauth.reddit.com/api/submit", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "bearer "+c.token)
	req.Header.Set("User-Agent", c.cfg.AuthData.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Выполняем запрос
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Если токен устарел (403), пробуем обновить и повторить запрос
	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("🚫 Доступ запрещен, пробую обновить токен...")
		if err := c.RefreshToken(); err != nil {
			return "", err
		}

		// Обновляем заголовок авторизации и повторяем запрос
		req.Header.Set("Authorization", "bearer "+c.token)
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
	}

	// Если слишком много запросов (429), ждем 5 секунд и повторяем запрос
	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("⚠️ Слишком много запросов, жду 5 секунд...")
		time.Sleep(7 * time.Second)
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
	}

	// Если ответ не 200, читаем тело и возвращаем ошибку
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("не удалось отправить пост, статус: %d, ответ: %s", resp.StatusCode, string(bodyBytes))
	}

	SubmitPostResponse := SubmitPostResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&SubmitPostResponse); err != nil {
		return "", err
	}

	fmt.Println("✅ Пост успешно отправлен!")
	fmt.Printf("Ответ сервера: %+v\n", SubmitPostResponse)

	// Возвращаем ссылку на пост
	return SubmitPostResponse.JSON.Data.URL, nil
}
