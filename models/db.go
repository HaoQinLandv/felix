package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
)

var db *gorm.DB
var dbPath string

var DefaultUser, DefaultPassword string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal("get home dir failed:", err)
	}
	dbPath = path.Join(dir, ".felix.db")
}
func CreateSqliteDB(verbose bool) {
	sqlite, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		logrus.WithError(err).Fatalf("master fail to open its sqlite db in %s. please install master first.", dbPath)
		return
	}

	db = sqlite
	//TODO::optimize
	db.AutoMigrate(Config{}, Machine{}, Task{}, User{}, Ginbro{}, TermLog{})
	db.LogMode(verbose)

	if DefaultPassword == "" || DefaultUser == "" {
		logrus.Error("wrong default user and password")
		return
	}

	userObj := User{Username: DefaultUser, Password: DefaultPassword, Email: "admin@felix.mojotv.cn", Avatar: "https://tech.mojotv.cn/assets/image/logo01.png"}
	if err := userObj.CreateInitUser(); err != nil {
		logrus.WithError(err).Error("create init user failed")
	}

}

func FlushSqliteDb() error {
	db.Close()
	return os.RemoveAll(dbPath)
}
