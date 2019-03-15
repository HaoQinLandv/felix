package models

import "github.com/jinzhu/gorm"

type Config struct {
	gorm.Model
	User     string    `json:"user"`
	Password string    `json:"password"`
	Pkey     string    `json:"pkey"`
	Hosts    []Machine `json:"hosts"`
}
