package internal

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
)

func TermLogAll(c *gin.Context) {
	mdl := models.TermLog{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}
func TermLogOne(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl := &models.TermLog{}
	mdl.ID = id
	tLog, err := mdl.One()
	if handleError(c, err) {
		return
	}

	jsonData(c, tLog)
}

func TermLogDelete(c *gin.Context) {
	var mdl models.TermLog
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.ID = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c, "item has been removed")
}
func TermLogUpdate(c *gin.Context) {
	var mdl models.TermLog
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
