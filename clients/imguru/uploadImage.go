package imguru

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// UploadMedia принимает имя файла, получает бинарные данные через getBinary,
// затем запрашивает у Reddit создание медиа-ассета, загружает файл и возвращает asset_id,
// который можно использовать в /api/submit для загрузки картинки.
func (client *ImgurClient) UploadImage(imageData []byte) (string, error) {
	// Конвертируем бинарные данные в Base64, так как Imgur API принимает изображения в этом формате.
	encoded := base64.StdEncoding.EncodeToString(imageData)

	// Подготавливаем JSON payload для Imgur, чтобы отправить изображение в нужном формате.
	payload := map[string]string{
		"image": encoded,
		"type":  "base64",
		"name":  uuid.NewString()[:5], // можно добавить title/description при необходимости
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("не удалось создать JSON payload: %w", err)
	}

	// Создаем запрос к Imgur API для загрузки изображения.
	imgurURL := "https://api.imgur.com/3/upload"
	req, err := http.NewRequest("POST", imgurURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("не удалось создать запрос: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", client.clientID))

	// Выполняем запрос к Imgur API.
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса к Imgur: %w", err)
	}
	defer resp.Body.Close()

	// Если слишком много запросов, ждем 20 секунд и повторяем запрос.
	if resp.StatusCode == http.StatusTooManyRequests {
		time.Sleep(time.Second * 20)
		resp, err = client.httpClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("ошибка при выполнении запроса к Imgur: %w", err)
		}
		defer resp.Body.Close()
	}

	// Проверяем статус ответа, чтобы убедиться, что запрос успешен.
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("imgur вернул ошибку, статус: %d, ответ: %s", resp.StatusCode, string(bodyBytes))
	}

	// Декодируем JSON-ответ от Imgur.
	var imgurResp struct {
		Data struct {
			Link string `json:"link"`
		} `json:"data"`
		Success bool `json:"success"`
		Status  int  `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&imgurResp); err != nil {
		return "", fmt.Errorf("ошибка при декодировании ответа от Imgur: %w", err)
	}

	// Проверяем, что загрузка на Imgur прошла успешно.
	if !imgurResp.Success {
		return "", fmt.Errorf("загрузка на Imgur не удалась, статус: %d", imgurResp.Status)
	}

	// Возвращаем публичную ссылку на изображение, чтобы использовать её в Reddit посте.
	return imgurResp.Data.Link, nil
}
