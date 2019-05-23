package internal

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Login(c *gin.Context) {
	var mdl models.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	ip := c.ClientIP()
	data, err := mdl.Login(ip)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func JwtMiddleware(c *gin.Context) {
	token, ok := c.GetQuery("_t")
	if !ok {
		hToken := c.GetHeader("Authorization")
		if len(hToken) < bearerLength {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "header Authorization has not Bearer token"})
			return
		}
		token = strings.TrimSpace(hToken[bearerLength:])
	}
	user, err := models.JwtParseUser(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}
	//store the user Model in the context
	c.Set("user", user)
	c.Next()
	// after request
}

const bearerLength = len("Bearer ")
