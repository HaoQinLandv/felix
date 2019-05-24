package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TermLog struct {
	gorm.Model
	// this a bug in GORM of HasOne in SQLite
	//Machine   Machine   `gorm:"association_autoupdate:false;association_autocreate:true"`
	//User      User      `gorm:"association_autoupdate:false;association_autocreate:true"`
	UserName    string    `json:"user_name"`
	MachineName string    `json:"machine_name"`
	MachineIp   string    `json:"machine_ip"`
	UserId      uint      `json:"user_id" gorm:"index"`
	MachineId   uint      `json:"machine_id" gorm:"index"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Log         string    `json:"log" gorm:"size:10240"`
}

func (m *TermLog) One() (one *Machine, err error) {
	one = &Machine{}
	err = crudOne(m, one)
	return
}

//All get all for pagination
func (m *TermLog) All(q *PaginationQuery) (list *[]TermLog, total uint, err error) {
	list = &[]TermLog{}
	total, err = crudAll(m, q, list)
	return
}

//Create insert a row
func (m *TermLog) Create() (err error) {
	//solve gorm sqlite \r failed
	//m.Log = strings.Replace(m.Log,"\r","<br>",-1)
	return db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", true).Create(m).Error
}

//Delete destroy a row
func (m *TermLog) Delete() (err error) {
	if m.ID == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
