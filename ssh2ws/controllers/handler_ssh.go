package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SshAll(c *gin.Context) {
	mdl := models.Machine{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	var mcs []models.Machine
	for _, vv := range *list {
		vv.Password = ""
		vv.Key = ""
		mcs = append(mcs, vv)
	}
	jsonPagination(c, mcs, total, query)
}
func SshOne(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl := models.Machine{Model: gorm.Model{ID: id}}
	mac, err := mdl.One()
	if handleError(c, err) {
		return
	}
	info, err := flx.FetchHardwareInfo(mac)
	if handleError(c, err) {
		return
	}
	//data := gin.H{"mac":mac,"info":info}
	jsonData(c, info)
}
func SshCreate(c *gin.Context) {
	var mdl models.Machine
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

func SshUpdate(c *gin.Context) {
	var mdl models.Machine
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c, "")
}

func SshDelete(c *gin.Context) {
	var mdl models.Machine
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.ID = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c, "")
}
