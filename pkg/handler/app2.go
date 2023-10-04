package handler

import (
	"encoding/base64"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NewApp2Server() *gin.Engine {
	sessionMap = map[string]loginUser{}
	engine := gin.New()
	engine.Use(func(ctx *gin.Context) {
		slog.Info("Request.", "Path", ctx.Request.URL.Path)
	})
	dir := GetDirectory()
	engine.LoadHTMLGlob(filepath.Join(dir, "statics/app2/html/*"))
	engine.GET("/app", app2LoginPageHandler)
	engine.GET("/api/app", app2APIAppHandler)
	return engine
}

func app2LoginPageHandler(c *gin.Context) {
	token := c.GetHeader("X-Access-Token")
	decodeToken, _ := base64.StdEncoding.DecodeString(token)
	c.HTML(http.StatusOK, "test.html", gin.H{"Text": string(decodeToken)})
}

func app2APIAppHandler(c *gin.Context) {
	c.SetCookie(sessionKeyName, "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
