package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "xblog/api/v1"
	"xblog/middleware"
	"xblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	v1g := r.Group("api/v1")
	v1g.Use(middleware.JwtToken())
	{
		// 用户模块的路由
		v1g.POST("user/add", v1.AddUser)

		v1g.PUT("user/:id", v1.UpdateUser)
		//v1g.GET("user/:id", v1.GetUser)
		v1g.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由
		v1g.POST("category/add", v1.AddCategory)

		v1g.PUT("category/:id", v1.UpdateCategory)
		v1g.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由
		v1g.POST("article/add", v1.AddArticle)

		v1g.PUT("article/:id", v1.UpdateArticle)
		v1g.DELETE("article/:id", v1.DeleteArticle)

	}

	publicApiGroup := r.Group("api/v1")
	{
		// 测试接口
		publicApiGroup.GET("hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "ok",
			})
		})
		publicApiGroup.GET("users", v1.GetUserList)
		publicApiGroup.GET("article/:id", v1.GetArticle)
		publicApiGroup.GET("category/articles", v1.GetCategoryArticleList)
		publicApiGroup.GET("articles", v1.GetArticleList)
		publicApiGroup.GET("categories", v1.GetCategoryList)
		publicApiGroup.POST("login", v1.Login)
	}

	serverAddress := fmt.Sprintf(":%s", utils.HttpPort)
	r.Run(serverAddress)
}
