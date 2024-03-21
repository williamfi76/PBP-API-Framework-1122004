package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "gorm.io/gorm"
)

// func connect() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/pbp_fiber?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

func connectGorm() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/pbp_fiber")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
