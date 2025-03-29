package messages

type RabbitUrlMessage struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

const YoutubeUrlType = "youtube_url"

type RabbitChannelMessage struct {
	Type           string `json:"type"`
	Url            string `json:"url"`
	NumberOfVideos int    `json:"number_of_videos"`
}

const YoutubeChannelType = "youtube_channel"
