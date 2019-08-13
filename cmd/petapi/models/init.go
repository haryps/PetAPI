package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	conn, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}

	defer db.Close()

	db = conn
	db.Debug().AutoMigrate(&Pet{})
}

//GetDB db connection
func GetDB() *gorm.DB {
	return db
}
