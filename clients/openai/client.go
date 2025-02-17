package openai

import (
	"github.com/fenix1851/meme_pipeline_shared/config"
)

// NewOpenAIClient создаёт новый клиент OpenAI
func NewOpenAIClient(cfg *config.Config) *OpenAIClient {
	return &OpenAIClient{
		APIKey:      cfg.OpenAI.APIKey,
		Model:       cfg.OpenAI.Model,
		MaxTokens:   cfg.OpenAI.MaxTokens,
		Temperature: cfg.OpenAI.Temperature,
	}
}

// OpenAIClient структура для работы с OpenAI API
type OpenAIClient struct {
	APIKey      string
	Model       string
	MaxTokens   int
	Temperature float64
}

// Формирование DTO для JSON-запроса
type ChatGPTRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	MaxTokens   int             `json:"max_tokens"`
	Temperature float64         `json:"temperature"`
}

// ImageURLContent содержит ссылку на изображение в виде объекта
type ImageURLContent struct {
	URL string `json:"url"`
}

// MessagePart представляет часть сообщения: текст или изображение.
type MessagePart struct {
	Type     string           `json:"type"`
	Text     string           `json:"text,omitempty"`
	ImageURL *ImageURLContent `json:"image_url,omitempty"`
}

// OpenAIMessage структура для передачи сообщений в OpenAI API
type OpenAIMessage struct {
	Role    string        `json:"role"`
	Content []MessagePart `json:"content"`
}
