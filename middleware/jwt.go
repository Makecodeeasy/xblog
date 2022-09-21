package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
	"xblog/utils"
	"xblog/utils/errmsg"
)

type XClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) string {
	claims := XClaim{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "xblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(utils.SigningKey))
	if err != nil {
		// todo 错误处理
		fmt.Println("err: ", err)
	}
	return ss
}

func ValidateToken(tokenString string) (*XClaim, int) {

	token, _ := jwt.ParseWithClaims(tokenString, &XClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SigningKey), nil
	})

	if claim, ok := token.Claims.(*XClaim); ok && token.Valid {
		// 验证成功
		fmt.Println("token验证成功了！！！")
		return claim, errmsg.SUCCESS
	} else {
		fmt.Println("token验证失败")
		return claim, errmsg.ERROR
	}

}

func JwtToken() gin.HandlerFunc {
	var code int
	return func(context *gin.Context) {
		requestToken := context.Request.Header.Get("Authorization")
		if requestToken == "" {
			code = errmsg.TOKEN_NOT_EXIST
			context.Abort()
		}
		requestTokenSlice := strings.SplitN(requestToken, " ", 2)
		fmt.Println("slice: ", requestTokenSlice)
		if len(requestTokenSlice) != 2 || requestTokenSlice[0] != "Bearer" {
			code = errmsg.TOKEN_FORMAT_WRONG
			fmt.Println("token格式错误")
			context.Abort()
		}

		claim, validateCode := ValidateToken(requestTokenSlice[1])
		if validateCode == errmsg.SUCCESS {
			code = errmsg.SUCCESS
		} else {
			if time.Now().Unix() > claim.ExpiresAt.Unix() {
				code = errmsg.TOKEN_EXPIRED
				fmt.Println("token已过期了啊")
				context.Abort()
			} else {
				code = errmsg.TOKEN_WRONG
				context.Abort()
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		context.Set("username", claim.Username)
		context.Next()
	}
}
