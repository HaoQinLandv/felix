package ssh2ws

import (
	"fmt"
	"github.com/dejavuzhou/felix/ssh2ws/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RunSsh2ws(bindAddress, user, password string, isDev bool) {
	r := gin.Default()
	r.MaxMultipartMemory = 32 << 20

	r.Use(static.Serve("/", static.LocalFile("dist", true)))
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
	api := r.Group("api")
	{
		api.GET("ssh", controllers.SshAll)
		api.POST("ssh", controllers.SshCreate)
		api.GET("ssh/:id", controllers.SshOne)
		api.PATCH("ssh/:id", controllers.SshUpdate)
		api.DELETE("ssh/:id", controllers.SshDelete)

		api.GET("sftp/:id", controllers.SftpLs)
		api.GET("sftp/:id/dl", controllers.SftpDl)
		api.GET("sftp/:id/cat", controllers.SftpCat)
		api.GET("sftp/:id/rm", controllers.SftpRm)
		api.GET("sftp/:id/rename", controllers.SftpRename)
		api.GET("sftp/:id/mkdir", controllers.SftpMkdir)
		api.POST("sftp/:id/up", controllers.SftpUp)
		//api.GET("sftp/:id/rm", controllers.SftpLs)
		//api.GET("sftp/:id/cat", controllers.SftpLs)
	}

	//http.HandleFunc("/ws/sftp", controllers.AuthPreChecker(files.SftpEstablish{}))
	fmt.Println("listening on port", bindAddress)
	fmt.Printf("auth user:%s,password:%s", user, password)
	// listen http
	if err := r.Run(bindAddress); err != nil {
		log.Println(err)
	}
}
