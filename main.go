package main

import (
	"xblog/model"
	"xblog/routes"
)

func main() {
	// 引用数据库
	model.InitDB()
	routes.InitRouter()
}
