package middleware

import (
	"im-golang/dao"
	"im-golang/model"
	"im-golang/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := util.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		nickname := claims.UserId
		DB := dao.DBConn()
		var user model.User
		DB.Where("nickname = ?", nickname).First(&user)

		if len(user.Mobile) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
