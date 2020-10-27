package twitter

import (
	"errors"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mrjones/oauth"
)

const (
	requestTokenURL = "https://api.twitter.com/oauth/request_token"
	authenticateURL = "https://api.twitter.com/oauth/authenticate"
	accessTokenURL  = "https://api.twitter.com/oauth/access_token"
)

var (
	callbackURL    string
	consumerKey    string
	consumerSecret string
)

func init() {
	callbackURL = os.Getenv("TWITTER_CALLBACK_URL")
	consumerKey = os.Getenv("TWITTER_API_KEY")
	consumerSecret = os.Getenv("TWITTER_API_SECRET")
}

func NewClient() *oauth.Consumer {
	return oauth.NewConsumer(consumerKey, consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authenticateURL,
			AccessTokenUrl:    accessTokenURL,
		})
}

func OAuth(c *gin.Context) (string, error) {
	client := NewClient()
	// リクエストトークンを取得
	rToken, loginURL, err := client.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		return "", nil
	}

	session := sessions.Default(c)
	session.Set("twitter_request_token", rToken.Token)
	session.Set("twitter_request_secret", rToken.Secret)
	session.Save()

	// 認可画面へリダイレクト
	return loginURL, nil
}

func Callback(c *gin.Context) (string, error) {
	// 認可証明書を取り出す
	var token, verificationCode string
	token = c.DefaultQuery("oauth_token", "")
	verificationCode = c.DefaultQuery("oauth_verifier", "")
	if token == "" || verificationCode == "" {
		return "/", errors.New("cannot get oauth_token/oauth_verifier")
	}

	// リクエストトークンの照合
	session := sessions.Default(c)
	rt := session.Get("twitter_request_token").(string)
	if rt == "" {
		return "/", errors.New("cannot get request token")
	}
	if token != rt {
		return "/", errors.New("request token is not correct")
	}
	// リクエストトークンの準備
	rs := session.Get("twitter_request_secret").(string)
	if rs == "" {
		return "/", errors.New("cannot get request secret")
	}
	rToken := oauth.RequestToken{Token: rt, Secret: rs}

	client := NewClient()
	// アクセストークンを取得
	aToken, err := client.AuthorizeToken(&rToken, verificationCode)
	if err != nil {
		return "/", err
	}
	session.Set("twitter_access_token", aToken.Token)
	session.Set("twitter_access_secret", aToken.Secret)
	session.Save()

	// リダイレクト
	return "/twitter", nil
}

func GetAccessToken(c *gin.Context) *oauth.AccessToken {
	session := sessions.Default(c)
	vat := session.Get("twitter_access_token")
	vas := session.Get("twitter_access_secret")
	if vat == nil || vas == nil {
		return nil
	}
	at, as := vat.(string), vas.(string)
	if at == "" || as == "" {
		return nil
	}
	aToken := oauth.AccessToken{Token: at, Secret: as}
	return &aToken
}

func UnOAuth(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete("twitter_access_token")
	session.Delete("twitter_access_secret")
	session.Save()
	return nil
}
