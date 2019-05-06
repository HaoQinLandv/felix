package ssh2ws

import (
	"fmt"
	"github.com/dejavuzhou/felix/ssh2ws/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RunSsh2ws(bindAddress, user, password string, isDev bool) {
	r := gin.Default()
	//r.Use(static.Serve("/", static.LocalFile("dist", true)))
	if isDev {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"}, //https://foo.com
			AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false, //enable cookie
			AllowOriginFunc: func(origin string) bool {
				return true
				//return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour, //cache options result decrease request lag
		}))
	}

	//http.Handle("/", http.FileServer(http.Dir("dist")))
	// api
	//http.HandleFunc("/api/sftp/upload", controllers.SftpUp)
	//http.HandleFunc("/api/sftp/ls", controllers.SftpLs)
	//http.HandleFunc("/api/sftp/dl", controllers.SftpDl)
	//http.HandleFunc("/api/ssh/view", controllers.SshAll)
	r.GET("/ws/ssh/:id", controllers.WsSsh)

	r.GET("api/ssh", controllers.SshAll)
	r.POST("api/ssh", controllers.SshCreate)
	r.GET("api/ssh/:id", controllers.SshOne)
	r.PATCH("api/ssh/:id", controllers.SshUpdate)
	r.DELETE("api/ssh/:id", controllers.SshDelete)

	//http.HandleFunc("/ws/sftp", controllers.AuthPreChecker(files.SftpEstablish{}))
	fmt.Printf("listening on port %s", bindAddress)
	fmt.Printf("auth user:%s,password:%s", user, password)
	// listen http
	if err := r.Run(bindAddress); err != nil {
		log.Println(err)
	}
}
