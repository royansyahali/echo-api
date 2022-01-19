package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ...
func ConnectDB() *gorm.DB {
	dsn := "root@/myapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	// See "Important settings" section.
	sqlDB.SetConnMaxLifetime(time.Minute * 60)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)

	return db
}
