package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"xblog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int;not null" json:"role"`
	Email    string `gorm:"type:varchar(32)" json:"email"`
	Active   int    `gorm:"type:int;not null;default 1;" json:"active"`
}

// TableName overrides the table name used by users to `user`
func (User) TableName() string {
	return "user"
}

// CheckUser 查询用户
func CheckUser(username string) (code int) {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	// data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeSave  是gorm框架钩子方法
func (u *User) BeforeSave() {
	// 通过钩子方法将密码加密, 该钩子方法gorm自动调用，类似的钩子还有AfterSave。文档：https://gorm.io/zh_CN/docs/hooks.html
	u.Password = ScryptPassword(u.Password)
}

// ScryptPassword 字符串加密
func ScryptPassword(password string) string {
	const keyLen = 12
	salt := make([]byte, 8)
	salt = []byte{16, 32, 9, 47, 25, 96, 33, 2}
	hashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatalln(err)
	}
	return base64.StdEncoding.EncodeToString(hashPw)
}

func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateUser(id int, data *User) int {
	var user User
	var userMap = make(map[string]interface{})
	userMap["username"] = data.Username
	userMap["role"] = data.Role
	userMap["email"] = data.Email
	err := db.Model(&user).Where("id = ?", id).Updates(&userMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUserList(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var code int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		code = errmsg.ERROR
	} else {
		code = errmsg.SUCCESS
	}
	return users, code
}

func UserLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.USER_NOT_EXIST
	}
	if user.Password != ScryptPassword(password) {
		return errmsg.PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.USER_NO_PERMISSION
	}
	if user.ID >= 1 {
		return errmsg.SUCCESS
	}

	return errmsg.ERROR
}
