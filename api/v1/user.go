package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xblog/model"
	"xblog/utils/errmsg"
)

func UserExists(ctx *gin.Context) {

}

func AddUser(ctx *gin.Context) {
	// 添加用户
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	fmt.Printf("data: %v", data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetUserList(ctx *gin.Context) {
	// 查询用户列表
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code := model.GetUserList(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUser(ctx *gin.Context) {
	// 查询用户
}

func UpdateUser(ctx *gin.Context) {
	// 更新用户
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &data)
	}
	if code == errmsg.USERNAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteUser(ctx *gin.Context) {
	// 删除用户
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
