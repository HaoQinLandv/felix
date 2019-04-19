package miniApp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var videoDir = ""

func Run(dir string) {
	videoDir = dir
	router := gin.Default()
	router.StaticFS("/static", http.Dir(videoDir))

	// Listen and serve on 0.0.0.0:8080
	router.Run(":4444")
}
