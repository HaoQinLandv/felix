package controllers

import (
	"errors"
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func jsonError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": msg})
}
func jsonAuthError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"ok": false, "msg": msg})
}

func jsonData(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": true, "data": data})
}

//func jsonPagination(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
//	c.AbortWithStatusJSON(200, gin.H{"ok": true, "data": list, "total": total, "offset": query.Offset, "limit": query.Size})
//}
func jsonSuccess(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, gin.H{"ok": true, "msg": "success"})
}
func jsonPagination(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
	c.JSON(200, gin.H{"ok": true, "data": list, "total": total, "page": query.Page, "size": query.Size})
}
func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func parseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}
