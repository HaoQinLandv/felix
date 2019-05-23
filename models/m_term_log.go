package models

import (
	"time"
)

type TermLog struct {
	Id        uint      `gorm:"primary_key"`
	MachineId uint      `json:"machine_id" `
	UserId    uint      `json:"user_id" `
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Log       string    `json:"log" gorm:"text"`
}

//Create
func (m *TermLog) Create() (err error) {

	return db.Create(m).Error
}
