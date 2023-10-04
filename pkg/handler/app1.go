package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

type loginUser struct {
	Name string `json:"name"`
}

var (
	sessionMap     map[string]loginUser
	sessionKeyName = "sample_login"
)

func GetDirectory() string {
	_, filePath, _, _ := runtime.Caller(0)
	fixtureDir := filepath.Dir(filePath)
	_, err := os.Stat(fixtureDir)
	if err != nil {
		panic(err)
	}
	return fixtureDir
}

func NewApp1Server() *gin.Engine {
	sessionMap = map[string]loginUser{}
	engine := gin.New()
	engine.Use(func(ctx *gin.Context) {
		slog.Info("Request.", "Path", ctx.Request.URL.Path)
	})
	dir := GetDirectory()
	engine.LoadHTMLGlob(filepath.Join(dir, "statics/app1/html/*"))
	engine.GET("/login", app1LoginPageHandler)
	engine.POST("/login", app1LoginHandler)
	engine.GET("/auth", app1AuthHandler)
	engine.GET("/app", app1AppPage)
	return engine
}

func app1AuthHandler(c *gin.Context) {
	slog.Info("request.")
	v, err := c.Cookie(sessionKeyName)
	if err != nil {
		slog.ErrorContext(c.Request.Context(), err.Error())
		c.Header("Location", "/login")
		c.Status(http.StatusUnauthorized)
		return
	}
	user, ok := sessionMap[v]
	if !ok {
		c.Header("Location", "/login")
		c.Status(http.StatusUnauthorized)
		return
	}
	buf, _ := json.Marshal(user)
	c.Header("X-Access-Token", base64.StdEncoding.EncodeToString(buf))
	c.Status(http.StatusOK)
}

func app1LoginPageHandler(c *gin.Context) {
	c.SetCookie(sessionKeyName, "", -1, "/", "localhost", false, true)
	c.HTML(http.StatusOK, "login.html", nil)
}

func app1LoginHandler(c *gin.Context) {
	var loginForm = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BindJSON(&loginForm); err != nil {
		slog.ErrorContext(c.Request.Context(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "ログインエラー"})
	}
	if loginForm.Email == "user@example.com" && loginForm.Password == "password" {
		b := make([]byte, 64)
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			slog.ErrorContext(c.Request.Context(), "ランダムな文字作成時にエラーが発生しました。")
		}
		newSessionKey := base64.URLEncoding.EncodeToString(b)
		loginUser := loginUser{
			Name: "テストユーザー",
		}
		sessionMap[newSessionKey] = loginUser
		c.SetCookie(sessionKeyName, newSessionKey, 0, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ログインユーザーが違います。"})
	}
}

func app1AppPage(c *gin.Context) {
	c.HTML(http.StatusOK, "app.html", nil)
}
