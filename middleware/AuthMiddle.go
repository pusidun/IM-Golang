package middleware

import (
	"im-golang/dao"
	"im-golang/model"
	"im-golang/serializer"
	"im-golang/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			res := serializer.Response{
				Status_Code: http.StatusUnauthorized,
				Msg:         "权限不足",
			}
			res.CtxSetResponse(c)
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := util.ParseToken(tokenString)
		if err != nil || !token.Valid {
			res := serializer.Response{
				Status_Code: http.StatusUnauthorized,
				Msg:         "权限不足",
			}
			res.CtxSetResponse(c)
			c.Abort()
			return
		}

		mobile := claims.UserId
		DB := dao.DBConn()
		var user model.User
		DB.Where("mobile = ?", mobile).First(&user)

		if len(user.Mobile) == 0 {
			res := serializer.Response{
				Status_Code: http.StatusUnauthorized,
				Msg:         "用户不存在",
			}
			res.CtxSetResponse(c)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
