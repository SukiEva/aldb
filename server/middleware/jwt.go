package middleware

import (
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/SukiEva/aldb/server/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims *jwt.Claims
		code := e.CODE.Success
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			code = e.CODE.AuthEmpty
		} else {
			parts := strings.Split(authHeader, ".") // header、payload、signature
			if len(parts) != 3 {
				code = e.CODE.AuthFormatError
			} else {
				var err error
				claims, err = jwt.ParseToken(authHeader)
				if err != nil {
					code = e.CODE.AuthTokenError
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.CODE.AuthTokenTimeOut
				}
			}
		}
		if code != e.CODE.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.ParseCode(code),
			})
			c.Abort()
			return
		}
		// 将当前请求信息保存到请求的上下文
		c.Set("UserEmail", claims.UserEmail)
		c.Set("UserPwd", claims.UserPwd)
		c.Next() //后续处理函数 c.Get("UserName") 获取当前请求的用户信息
	}
}
