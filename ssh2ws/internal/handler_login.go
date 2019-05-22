package internal

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type jwtUser struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form "password"`
}

func GetLoginHandler(user, password string, expire time.Duration, secretBytes []byte) gin.HandlerFunc {

	return func(c *gin.Context) {

		var mdl jwtUser
		err := c.ShouldBind(&mdl)
		if handleError(c, err) {
			return
		}
		if mdl.User != user || mdl.Password != password {
			jsonError(c, "user authentication is failed")
			return
		}
		ip := c.ClientIP()
		obj, err := jwtGenerateToken(ip, expire, secretBytes)
		if handleError(c, err) {
			return
		}
		jsonData(c, obj)
	}
}

const bearerLength = len("Bearer ")

func JwtAuthMiddleware(secretBytes []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.GetQuery("_t")
		if !ok {
			hToken := c.GetHeader("Authorization")
			if len(hToken) < bearerLength {
				jsonAuthError(c, "header Authorization has not Bearer token")
				return
			}
			token = strings.TrimSpace(hToken[bearerLength:])
		}
		user, err := jwtParseUser(token, secretBytes)
		if err != nil {
			jsonAuthError(c, err.Error())
			return
		}
		//store the user Model in the context
		c.Set("user", user)
		c.Next()
		// after request
	}
}
