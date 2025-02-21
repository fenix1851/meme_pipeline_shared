package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fenix1851/meme_pipeline_shared/models"
	"gorm.io/datatypes"
)

// GetTopPosts retrieves the top posts from Reddit based on the specified limit and period.
func (c *RedditClient) GetTopPostsFromSubreddit(subreddit string, limit string, period string) (RedditPostResponses, error) {
	requestURL := fmt.Sprintf("%s/r/%s/top?t=%s&limit=%s", c.cfg.URL, subreddit, period, limit)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", c.cfg.AuthData.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Refresh token if unauthorized
	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("⚠️ 403/401 received - Refreshing token...")
		if err := c.RefreshToken(); err != nil {
			return nil, err
		}
		fmt.Printf("\n\n\ntoken: %s\n\n\n", c.token)
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("User-Agent", c.cfg.AuthData.UserAgent)
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get posts, status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode into our typed struct instead of using maps.
	var apiResp RedditAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	// Convert the API response into our models.RedditPost
	var posts RedditPostResponses
	for _, child := range apiResp.Data.Children {
		postData := child.Data

		// Extract image URLs if available.
		picUrls := extractImageURLs(postData.Preview)

		picUrlsJSON, err := json.Marshal(picUrls)
		if err != nil {
			return nil, err
		}

		post := &models.RedditPost{
			PostLink:  fmt.Sprintf("https://www.reddit.com%s", postData.Permalink),
			PostTitle: postData.Title,
			PostText:  postData.Selftext,
			Upvotes:   postData.Ups,
			Comments:  postData.NumComments,
			PicUrls:   datatypes.JSON(picUrlsJSON),
		}

		posts = append(posts, post)
	}

	// Optionally, print the posts JSON for debugging.
	postsJSON, err := json.MarshalIndent(posts, "", "  ")
	if err == nil {
		fmt.Println(string(postsJSON))
	}

	return posts, nil
}
