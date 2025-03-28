package messages

type RabitUrlMessage struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

const YoutubeUrlType = "youtube_url"

type RabbitChannelMessage struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

const YoutubeChannelType = "youtube_channel"
