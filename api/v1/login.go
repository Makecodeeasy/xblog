package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xblog/middleware"
	"xblog/model"
	"xblog/utils/errmsg"
)

func Login(context *gin.Context) {
	var data model.User
	var token string
	context.BindJSON(&data)
	code := model.UserLogin(data.Username, data.Password)
	fmt.Println("登录结果code= ", code)
	if code == errmsg.SUCCESS {
		token = middleware.GenerateToken(data.Username)
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})

}
