package internal

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
)

func UserAll(c *gin.Context) {
	mdl := models.User{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	var mcs []models.User
	for _, vv := range *list {
		vv.Password = ""
		mcs = append(mcs, vv)
	}
	jsonPagination(c, mcs, total, query)
}

func UserCreate(c *gin.Context) {
	var mdl models.User
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

func UserUpdate(c *gin.Context) {
	var mdl models.User
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

func UserDelete(c *gin.Context) {
	var mdl models.User
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
