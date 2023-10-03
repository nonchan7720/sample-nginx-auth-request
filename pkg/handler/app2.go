package handler

import (
	"encoding/base64"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NewApp2Server() *gin.Engine {
	sessionMap = map[string]loginUser{}
	engine := gin.New()
	dir := GetDirectory()
	engine.LoadHTMLGlob(filepath.Join(dir, "statics/app2/html/*"))
	engine.GET("/app", app2LoginPageHandler)
	return engine
}

func app2LoginPageHandler(c *gin.Context) {
	token := c.GetHeader("X-Access-Token")
	decodeToken, _ := base64.StdEncoding.DecodeString(token)
	c.HTML(http.StatusOK, "test.html", gin.H{"Text": string(decodeToken)})
}
