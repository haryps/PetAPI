package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {

	conn, err := gorm.Open("mysql", "root:MyNewPass@tcp(127.0.0.1:3306)/petapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("Connection established")

	Db = conn
	Db.Debug().AutoMigrate(&Pet{})
}

//GetDB db connection
func GetDB() *gorm.DB {
	return Db
}
