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

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		logrus.WithError(err).Fatal(" get home dir failed")
	}
	dbPath = path.Join(dir, ".felix.db")
}
func CreateSqliteDB(verbose bool) {
	sqlite, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("master fail to open its sqlite db in %s. please install master first. error:%s", dbPath, err)
	} else {
		db = sqlite
		db.AutoMigrate(Config{}, Machine{})
		db.LogMode(verbose)
	}
}

func FlushSqliteDb() error {
	db.Close()
	return os.RemoveAll(dbPath)
}
