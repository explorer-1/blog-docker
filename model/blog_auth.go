package model

import (
	"log"

	"gorm.io/gorm"
)

type BlogAuth struct {
	gorm.Model

	Username string
	Password string
}

func (b *BlogAuth) TableName() string {
	return "blog_auth"
}

func GetUserCountByUsername(name string) (int, bool) {
	var count int64
	err := DB.Where("username = ?", name).Find(&BlogAuth{}).Count(&count).Error
	if err != nil {
		return 0, false
	}
	return int(count), true
}

func GetUserPasswordByUsername(name string) (string, bool) {
	var password string
	err := DB.Model(&BlogAuth{}).Where("username = ?", name).Select("password").Scan(&password).Error
	if err != nil {
		return "", false
	}

	return password, true
}

func GetUserIdByUsername(name string) (int, bool) {
	var id int
	err := DB.Model(&BlogAuth{}).Where("username = ?", name).Select("id").Scan(&id).Error
	if err != nil {
		log.Printf("get user id by name error : %v\n", err)
		return -1, false
	}

	return id, true
}
