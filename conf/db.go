package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB(path string) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	DB = db
}
