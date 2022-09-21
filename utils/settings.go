package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	SigningKey string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("读取配置文件错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8080")
	SigningKey = file.Section("server").Key("SigningKey").MustString("xblog2022")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("xblog")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("xblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("xblog")
	DbName = file.Section("database").Key("DbName").MustString("xblog")
}
