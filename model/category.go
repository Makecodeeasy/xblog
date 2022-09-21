package model

import (
	"xblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(32);not null" json:"name"`
}

// TableName overrides the table name used by categories to `category`
func (Category) TableName() string {
	return "category"
}

// CheckCategory 查询分类
func CheckCategory(username string) (code int) {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateCategory(id int, data *Category) int {
	var categoryMap = make(map[string]interface{})
	categoryMap["name"] = data.Name
	err := db.Model(&data).Where("id = ?", id).Updates(&categoryMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategoryList() ([]Category, int) {
	var categories []Category
	var code int
	err := db.Find(&categories).Error
	if err != nil {
		code = errmsg.ERROR
	} else {
		code = errmsg.SUCCESS
	}
	return categories, code
}
