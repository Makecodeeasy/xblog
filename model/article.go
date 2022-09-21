package model

import (
	"gorm.io/gorm"
	"xblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title    string `gorm:"varchar(100);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`
	Describe string `gorm:"type:varchar(200)" json:"describe"`
	Content  string `gorm:"type:longtext" json:"content"`
	Img      string `gorm:"type:varchar(100)" json:"img"`
}

// TableName overrides the table name used by articles to `article`
func (Article) TableName() string {
	return "article"
}

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticleList(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return nil, errmsg.ERROR
	} else {
		return articles, errmsg.SUCCESS
	}

}

func UpdateArticle(id int, data *Article) int {
	var article Article
	var articleMap = make(map[string]interface{})
	articleMap["title"] = data.Title
	articleMap["cid"] = data.Cid
	articleMap["describe"] = data.Describe
	articleMap["content"] = data.Content
	articleMap["img"] = data.Img
	err := db.Model(&article).Where("id = ? ", id).Updates(articleMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ARTICAL_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

func GetCategoryArticle(categoryID int, pageSize int, pageNum int) ([]Article, int) {
	var categoryArticleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", categoryID).Find(&categoryArticleList).Error
	if err != nil {
		return nil, errmsg.CATEGORY_NOT_EXIST
	}
	return categoryArticleList, errmsg.SUCCESS
}
