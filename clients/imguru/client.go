package imguru

import (
	"net/http"
)

type ImgurClient struct {
	clientID   string
	httpClient *http.Client
}

func NewClient(clientID string) *ImgurClient {
	return &ImgurClient{
		clientID:   clientID,
		httpClient: &http.Client{},
	}
}
