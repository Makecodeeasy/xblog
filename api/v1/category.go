package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xblog/model"
	"xblog/utils/errmsg"
)

func AddCategory(ctx *gin.Context) {
	// 添加分类
	var data model.Category
	_ = ctx.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetCategoryList(ctx *gin.Context) {
	// 查询分类列表
	data, code := model.GetCategoryList()
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategory(ctx *gin.Context) {
	// 查询用户
}

func UpdateCategory(ctx *gin.Context) {
	// 更新分类
	var data model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code := model.UpdateCategory(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteCategory(ctx *gin.Context) {
	// 删除用户
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
