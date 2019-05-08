package utils

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

const INDEX = "index.html"

// Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs http.FileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		urlPath := strings.TrimSpace(c.Request.URL.Path)
		if urlPath == urlPrefix {
			urlPath = path.Join(urlPrefix, INDEX)
		}
		f, err := fs.Open(urlPath)
		if err != nil {
			return
		}
		fi, err := f.Stat()
		if err != nil || !fi.IsDir() {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}

	}
}
