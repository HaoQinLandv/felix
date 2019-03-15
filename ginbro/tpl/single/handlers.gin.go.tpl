package handlers

import (
	"fmt"
	"{{.ProjectPackage}}/models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"{{.ProjectPackage}}/config"
	"log"
	"path"
)

var r = gin.Default()
var groupApi *gin.RouterGroup

//in the same package init executes in file'name alphabet order
func init() {
	if config.GetBool("app.enable_cors") {
		enableCorsMiddleware()
	}
	if sp := config.GetString("app.static_path"); sp != "" {
		r.Use(static.Serve("/", static.LocalFile(sp, true)))
		if config.GetBool("app.enable_not_found") {
			r.NoRoute(func(c *gin.Context) {
				file := path.Join(sp, "index.html")
				c.File(file)
			})
		}
	}

	if config.GetBool("app.enable_swagger") && config.GetString("app.env") != "prod" {
		//add edit your own swagger.doc.yml file in ./swagger/doc.yml
		//generateSwaggerDocJson()
		r.Static("swagger", "./swagger")
	}
	prefix := config.GetString("app.api_prefix")
	api := "api"
	if prefix != "" {
		api = fmt.Sprintf("%s/%s", api, prefix)
	}
	groupApi = r.Group(api)

	if config.GetString("app.env") != "prod" {
		r.GET("/app/info", func(c *gin.Context) {
			c.JSON(200, config.GetStringMapString("app"))
		})
	}

}

//ServerRun start the gin server
func ServerRun() {

	addr := config.GetString("app.addr")
	if config.GetBool("app.enable_https") {
		log.Fatal(autotls.Run(r, addr))
	} else {
		log.Printf("visit http://%s/swagger for RESTful APIs Document", addr)
		log.Printf("visit http://%s/ for front-end static html files", addr)
		log.Printf("visit http://%s/app/info for app info only on not-prod mode", addr)
		r.Run(addr)
	}
}

//Close gin app
func Close() {
	models.Close()
}
