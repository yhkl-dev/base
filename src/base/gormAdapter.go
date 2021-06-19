package base

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	db, err := gorm.Open("mysql", "root:123456@tcp(49.232.252.187:3306)/jstock?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	return &GormAdapter{DB: db}
}
