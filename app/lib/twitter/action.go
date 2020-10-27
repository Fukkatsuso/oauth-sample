package twitter

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mrjones/oauth"
)

const (
	accountVerifyCredsURL = "https://api.twitter.com/1.1/account/verify_credentials.json"
	tweetURL              = "https://api.twitter.com/1.1/statuses/update.json"
	userTimelineURL       = "https://api.twitter.com/1.1/statuses/user_timeline.json"
)

func GetUser(c *gin.Context, token *oauth.AccessToken, user *User) error {
	client := NewClient()
	params := map[string]string{}
	resp, err := client.Get(accountVerifyCredsURL, params, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("twitter is unavailable")
	}
	if resp.StatusCode >= 400 {
		return errors.New("twitter request is invalid")
	}

	err = json.NewDecoder(resp.Body).Decode(user)
	return err
}

func Tweet(c *gin.Context, token *oauth.AccessToken, post *NewPost) error {
	if len(post.Status) == 0 || len(post.Status) > 140 {
		return errors.New("status must be 0~140 chars")
	}
	client := NewClient()
	params := map[string]string{
		"status": post.Status,
	}
	resp, err := client.Post(tweetURL, params, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("twitter is unavailable")
	}
	if resp.StatusCode >= 400 {
		return errors.New("twitter request is invalid")
	}

	return nil
}

func GetUserTimeline(c *gin.Context, token *oauth.AccessToken, id string, tl *UserTimeline) error {
	client := NewClient()
	params := map[string]string{
		"user_id": id,
	}
	resp, err := client.Get(userTimelineURL, params, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("twitter is unavailable")
	}
	if resp.StatusCode >= 400 {
		return errors.New("twitter request is invalid")
	}

	err = json.NewDecoder(resp.Body).Decode(&tl.Posts)
	return err
}
