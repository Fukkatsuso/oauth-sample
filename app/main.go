package main

import (
	"net/http"
	"os"

	"github.com/Fukkatsuso/oauth-sample/app/lib/twitter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// mysqlsession "github.com/kimiazhu/ginweb-contrib/sessions"
)

func main() {
	endpoint := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@" + os.Getenv("MYSQL_PROTOCOL") + "/" + os.Getenv("MYSQL_DATABASE")
	endpoint += "?parseTime=true&loc=Local"
	// store, err := mysqlsession.NewMySQLStore(endpoint, "sessions", []byte("secret"))
	// if err != nil {
	// 	panic("failed to create session store")
	// }
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24,
	})

	r := gin.Default()
	r.LoadHTMLGlob("views/html/*")
	// sessions.Sessions(name string, store Store)
	// type Store interface {
	// 	sessions.Store: "github.com/gorilla/sessions"
	// 	Options(Options)
	// }
	// sessions.Store interface {
	// 	Get(r *http.Request, name string) (*Session, error)
	// 	New(r *http.Request, name string) (*Session, error)
	// 	Save(r *http.Request, w http.ResponseWriter, s *Session) error
	// }
	r.Use(sessions.Sessions("session", store))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "OAuth sample app!",
		})
	})

	r.GET("/twitter", func(c *gin.Context) {
		aToken := twitter.GetAccessToken(c)
		if aToken == nil {
			c.Redirect(http.StatusSeeOther, "/twitter/oauth")
			return
		}
		// プロフィール取得
		user := twitter.User{}
		err := twitter.GetUser(c, aToken, &user)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/twitter/oauth")
			return
		}
		// タイムライン取得
		tl := twitter.UserTimeline{}
		err = twitter.GetUserTimeline(c, aToken, user.ID, &tl)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/twitter/oauth")
			return
		}
		// ユーザーページ表示
		c.HTML(http.StatusOK, "twitter.html", gin.H{
			"title":    "user page",
			"user":     user,
			"timeline": tl,
		})
	})
	r.GET("/twitter/oauth", func(c *gin.Context) {
		loginURL, err := twitter.OAuth(c)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		c.Redirect(http.StatusFound, loginURL)
	})
	r.GET("/twitter/callback", func(c *gin.Context) {
		redirectURL, err := twitter.Callback(c)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/twitter/oauth")
			return
		}
		c.Redirect(http.StatusFound, redirectURL)
	})
	r.POST("/twitter/unoauth", func(c *gin.Context) {
		err := twitter.UnOAuth(c)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "twitter unauthorize successed",
		})
	})
	r.POST("/twitter/post", func(c *gin.Context) {
		aToken := twitter.GetAccessToken(c)
		if aToken == nil {
			c.Redirect(http.StatusSeeOther, "/twitter/oauth")
			return
		}
		post := twitter.NewPost{
			Status: c.PostForm("content"),
		}
		err := twitter.Tweet(c, aToken, &post)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/twitter")
			return
		}
		c.Redirect(http.StatusFound, "/twitter")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
