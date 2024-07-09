package model

import (
	"log"

	"gorm.io/gorm"
)

type BlogTag struct {
	gorm.Model

	Name      string `json:"name"`
	State     int    `json:"state"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func (b *BlogTag) TableName() string {
	return "blog_tag"
}

func GetTagCountByName(name string) (int, bool) {
	var count int64
	err := DB.Where("name = ?", name).Find(&BlogTag{}).Count(&count).Error
	if err != nil {
		return 0, false
	}
	return int(count), true
}

func AddTag(name string, state int, createdBy string) bool {
	err := DB.Create(&BlogTag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error

	if err != nil {
		log.Printf("add tag err : %v\n", err)
		return false
	}

	return true
}

func DeleteTagByTagName(name string) bool {
	err := DB.Where("name = ?", name).Delete(&BlogTag{}).Error
	return err == nil
}
