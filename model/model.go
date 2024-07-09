package model

import (
	"blog/common/confsetting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	sec, err := confsetting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("[model] cfg getSection database err : %v\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sec.Key("USERNAME").MustString("root"), sec.Key("PASSWORD").MustString("2138108Aa!"),
		sec.Key("HOST").MustString("127.0.0.1:3306"), sec.Key("NAME").MustString("db6"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[model] gorm open err : %v\n", err)
	}
}
