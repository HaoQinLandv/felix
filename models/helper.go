package models

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

func parseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}

//PaginationQuery gin handler query binding struct
type PaginationQuery struct {
	Where  string `form:"where"`  // todo deprecated
	Fields string `form:"fields"` // todo deprecated
	Order  string `form:"order"`  // todo deprecated
	Size   uint   `form:"size"`
	Page   uint   `form:"page"`
}

//String to string
func (pq *PaginationQuery) String() string {
	return fmt.Sprintf("w=%v_f=%s_o=%s_op=%d_l=%d", pq.Where, pq.Fields, pq.Order, pq.Page, pq.Size)
}

func crudAll(m interface{}, q *PaginationQuery, list interface{}) (total uint, err error) {
	var tx *gorm.DB
	total, tx = getResourceCount(m, q)
	if q.Fields != "" {
		columns := strings.Split(q.Fields, ",")
		if len(columns) > 0 {
			tx = tx.Select(q.Fields)
		}
	}
	if q.Order != "" {
		tx = tx.Order(q.Order)
	}
	if q.Size <= 0 {
		q.Size = 15
	}
	if q.Page < 1 {
		q.Page = 1
	}
	offset := (q.Page - 1) * q.Size
	err = tx.Offset(offset).Limit(q.Size).Find(list).Error
	return
}

func crudAllV2(countQuery *gorm.DB, query *gorm.DB, q *PaginationQuery, list interface{}) (total uint, err error) {
	err = countQuery.Count(&total).Error
	if err != nil {
		return 0, err
	}
	if q.Size <= 0 {
		q.Size = 15
	}
	if q.Page < 1 {
		q.Page = 1
	}
	offset := (q.Page - 1) * q.Size
	err = query.Offset(offset).Limit(q.Size).Find(list).Error
	return
}

func crudOne(m interface{}, one interface{}) (err error) {
	if db.Where(m).First(one).RecordNotFound() {
		return errors.New("resource is not found")
	}
	return nil
}

func crudUpdate(m interface{}, where interface{}) (err error) {
	db := db.Model(where).Updates(m)
	if err = db.Error; err != nil {
		return
	}
	if db.RowsAffected != 1 {
		return errors.New("id is invalid and resource is not found")
	}
	return nil
}

func crudDelete(m interface{}) (err error) {
	//WARNING When delete a record, you need to ensure it’s primary field has value, and GORM will use the primary key to delete the record, if primary field’s blank, GORM will delete all records for the model
	//primary key must be not zero value
	db := db.Delete(m)
	if err = db.Error; err != nil {
		return
	}
	if db.RowsAffected != 1 {
		return errors.New("resource is not found to destroy")
	}
	return nil
}
func getResourceCount(m interface{}, q *PaginationQuery) (uint, *gorm.DB) {
	var tx = db.Model(m)
	conditions := strings.Split(q.Where, ",")
	for _, val := range conditions {
		w := strings.SplitN(val, ":", 2)
		if len(w) == 2 {
			bindKey, bindValue := w[0], w[1]
			if intV, err := strconv.ParseInt(bindValue, 10, 64); err == nil {
				// bind value is int
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, intV)
			} else if fV, err := strconv.ParseFloat(bindValue, 64); err == nil {
				// bind value is float
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, fV)
			} else if bindValue != "" {
				// bind value is string
				field := fmt.Sprintf("`%s` LIKE ?", bindKey)
				sV := fmt.Sprintf("%%%s%%", bindValue)
				tx = tx.Where(field, sV)
			}
		}
	}
	var count uint
	tx.Count(&count)
	return count, tx
}
