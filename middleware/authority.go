package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"imooc.com/datasource"
	"net/http"
)
const (
	// 可自定义盐值
	TokenSalt = "default_salt"
)

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		return
		token,err := c.Request.Cookie("token")
		if err == nil{
			datasource.GetRedis().Get(token.Value)
			c.Next()
		}else{
			c.Abort()
			c.HTML(http.StatusUnauthorized, "401.html", nil)
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
		}
	}
}
