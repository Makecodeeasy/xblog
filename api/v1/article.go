package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xblog/model"
	"xblog/utils/errmsg"
)

// AddArticle 增加文章
func AddArticle(ctx *gin.Context) {
	var data model.Article
	_ = ctx.ShouldBindJSON(&data)
	fmt.Printf("data: %v", data)
	code := model.CreateArticle(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetArticleList(ctx *gin.Context) {
	// 查询文章列表
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code := model.GetArticleList(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticle(ctx *gin.Context) {
	// 查询文章
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetCategoryArticleList(ctx *gin.Context) {
	// 查询分类下的文章
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	categoryID, _ := strconv.Atoi(ctx.Query("cid"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code := model.GetCategoryArticle(categoryID, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func UpdateArticle(ctx *gin.Context) {
	// 更新文章
	var data model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)

	code := model.UpdateArticle(id, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteArticle(ctx *gin.Context) {
	// 删除用户
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
