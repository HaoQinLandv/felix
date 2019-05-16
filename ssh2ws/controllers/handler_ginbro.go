package controllers

import (
	"github.com/dejavuzhou/felix/ginbro"
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
)

func GinbroGen(c *gin.Context) {

	var m models.Ginbro
	err := c.ShouldBind(&m)
	if handleError(c, err) {
		return
	}
	app, err := ginbro.Run(m)
	if handleError(c, err) {
		return
	}
	err = app.ListAppFileTree()
	if handleError(c, err) {
		return
	}
	jsonData(c, app)
}
func GinbroDb(c *gin.Context) {
	var gb models.Ginbro
	err := c.ShouldBind(&gb)
	if handleError(c, err) {
		return
	}
	data, err := ginbro.FetchDbColumn(gb)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)

}
