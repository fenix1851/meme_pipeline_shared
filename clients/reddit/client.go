package reddit

import (
	"net/http"
	"strings"

	"github.com/fenix1851/meme_pipeline_shared/config"
	"github.com/fenix1851/meme_pipeline_shared/models"
)

type RedditClient struct {
	httpClient   *http.Client
	cfg          *config.ClientRedditConfig
	token        string
	refreshToken string
}

func NewRedditClient(httpClient *http.Client, cfg config.Config) *RedditClient {
	return &RedditClient{
		httpClient:   httpClient,
		cfg:          cfg.ClientReddit,
		token:        cfg.ClientReddit.AccessToken,
		refreshToken: cfg.ClientReddit.RefreshToken,
	}
}

type RedditAPIResponse struct {
	Data RedditData `json:"data"`
}

// RedditData contains the children posts.
type RedditData struct {
	Children []RedditChild `json:"children"`
}

// RedditChild wraps individual post data.
type RedditChild struct {
	Data RedditPostData `json:"data"`
}

// RedditPostData contains the fields we're interested in.
type RedditPostData struct {
	Permalink   string       `json:"permalink"`
	Title       string       `json:"title"`
	Selftext    string       `json:"selftext"`
	Ups         int          `json:"ups"`
	NumComments int          `json:"num_comments"`
	Preview     *PreviewData `json:"preview,omitempty"`
}

// PreviewData contains images information.
type PreviewData struct {
	Images []PreviewImage `json:"images"`
}

// PreviewImage contains the source image.
type PreviewImage struct {
	Source ImageSource `json:"source"`
}

// ImageSource contains the URL of the image.
type ImageSource struct {
	URL string `json:"url"`
}

// RedditCommentResponse описывает структуру JSON-ответа Reddit при публикации комментария.
type RedditCommentResponse struct {
	JSON struct {
		Errors [][]interface{} `json:"errors"`
	} `json:"json"`
}

type RedditPostResponses []*models.RedditPost

type PostData struct {
	Title     string `json:"title"`
	Subreddit string `json:"subreddit"`
	Url       string `json:"url"`
}

type SubmitPostResponse struct {
	JSON struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
		Errors []interface{} `json:"errors"`
	} `json:"json"`
}

// extractThreadLink извлекает ссылку на тред из ссылки поста
func ExtractThreadLink(postLink string) string {
	parts := strings.Split(postLink, "/")
	if len(parts) < 6 {
		return postLink // если формат не соответствует, вернем оригинал
	}
	return strings.Join(parts[:6], "/") // https://www.reddit.com/r/golang/comments/threadID
}

// extractThreadName извлекает имя сабреддита
func ExtractThreadName(threadLink string) string {
	parts := strings.Split(threadLink, "/")
	if len(parts) < 6 {
		return "Unknown Thread"
	}
	return parts[4] // /r/golang/comments/... -> golang
}

// extractImageURLs extracts and cleans image URLs from the PreviewData.
func extractImageURLs(preview *PreviewData) []string {
	var picUrls []string
	if preview == nil {
		return picUrls
	}

	for _, img := range preview.Images {
		url := img.Source.URL
		// Replace HTML-encoded ampersands with actual ampersands.
		url = strings.ReplaceAll(url, "&amp;", "&")
		picUrls = append(picUrls, url)
	}
	return picUrls
}
