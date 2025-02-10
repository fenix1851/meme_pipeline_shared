package messages

// TextBlock представляет структуру для хранения данных текстового блока.
type TextBlock struct {
	X           int    `json:"x"`           // Координата X верхнего левого угла
	Y           int    `json:"y"`           // Координата Y верхнего левого угла
	Width       int    `json:"width"`       // Ширина текстового блока
	Height      int    `json:"height"`      // Высота текстового блока
	Description string `json:"description"` // Описание текстового блока
}

// MemeGenerationRequest - структура запроса на генерацию мема
type MemeGenerationRequest struct {
	Type       string    `json:"type"`        // Тип сообщения
	TemplateID int       `json:"template_id"` // ID шаблона
	Text       TextBlock `json:"text_block"`  // Текстовый блок
}

// Константа для типа сообщения
const MemeGenerationRequestType = "meme_generation"
