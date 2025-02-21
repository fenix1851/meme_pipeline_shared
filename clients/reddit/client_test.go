package reddit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/fenix1851/meme_pipeline_shared/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RedditClientTestSuite struct {
	suite.Suite
	client *RedditClient
}

func (suite *RedditClientTestSuite) SetupTest() {

}

func (suite *RedditClientTestSuite) TestRefreshToken() {
	configureSuite(suite, "Parse")
	err := suite.client.RefreshToken()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "new_token", suite.client.token)
}

func (suite *RedditClientTestSuite) TestGetTopPosts() {
	configureSuite(suite, "Parse")
	posts, err := suite.client.GetTopPostsFromSubreddit("all", "100", "day")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), posts, 100)
}

func (suite *RedditClientTestSuite) TestPostComment() {
	configureSuite(suite, "Submit")
	err := suite.client.PostComment("t3_1iuim4y", "yea at this point ypu can do literelly anything in the game")
	assert.NoError(suite.T(), err)
}

func TestRedditClientTestSuite(t *testing.T) {
	suite.Run(t, new(RedditClientTestSuite))
}

func configureSuite(suite *RedditClientTestSuite, configType string) {
	config, err := config.LoadConfig(fmt.Sprintf("./configFor%sTests.yaml", configType))
	assert.NoError(suite.T(), err)
	httpClient := &http.Client{}
	suite.client = NewRedditClient(httpClient, *config)
	suite.client.cfg = config.ClientReddit
}
