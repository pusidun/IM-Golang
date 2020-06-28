package controller

import (
	"im-golang/dao"
	"im-golang/model"
	"im-golang/serializer"
	"im-golang/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Ping(c *gin.Context) {
	var res = serializer.Response{
		Status_Code: http.StatusOK,
		Msg:         "PONG",
	}
	res.CtxSetResponse(c)
}

func UserRegister(c *gin.Context) {
	DB := dao.DBConn()
	var res serializer.Response

	//Get params
	name := c.PostForm("name")
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	// Valid
	if len(mobile) != 11 {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "手机号必须为11位",
		}
		res.CtxSetResponse(c)
		return
	}
	if len(password) < 6 {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "密码不得少于6位",
		}
		res.CtxSetResponse(c)
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(DB, mobile) {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "手机已注册",
		}
		res.CtxSetResponse(c)
		return
	}

	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		res = serializer.Response{
			Status_Code: http.StatusInternalServerError,
			Msg:         "密码加密错误",
		}
		res.CtxSetResponse(c)
		return
	}
	newUser := model.User{
		Nickname: name,
		Mobile:   mobile,
		Password: string(cryptPassword),
	}
	DB.Create(&newUser)

	res = serializer.Response{
		Status_Code: http.StatusOK,
		Msg:         "注册成功",
	}
	res.CtxSetResponse(c)
}

func isTelephoneExist(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func UserLogin(c *gin.Context) {
	DB := dao.DBConn()
	var res serializer.Response

	//Get params
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	// Valid
	if len(mobile) != 11 {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "手机号必须为11位",
		}
		res.CtxSetResponse(c)
		return
	}
	if len(password) < 6 {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "密码不得少于6位",
		}
		res.CtxSetResponse(c)
		return
	}
	var user model.User
	DB.Where("mobile = ?", mobile).First(&user)
	if len(user.Nickname) == 0 {
		res = serializer.Response{
			Status_Code: http.StatusUnprocessableEntity,
			Msg:         "用户不存在",
		}
		res.CtxSetResponse(c)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		res = serializer.Response{
			Status_Code: http.StatusBadRequest,
			Msg:         "密码错误",
		}
		res.CtxSetResponse(c)
		return
	}

	// token
	token, err := util.GenerateToken(user)
	if err != nil {
		res = serializer.Response{
			Status_Code: http.StatusInternalServerError,
			Msg:         "token generate err",
		}
		res.CtxSetResponse(c)
		return
	}

	res = serializer.Response{
		Status_Code: http.StatusOK,
		Msg:         "登陆成功",
		Data:        gin.H{"token": token},
	}
	res.CtxSetResponse(c)
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	res := serializer.Response{
		Status_Code: http.StatusOK,
		Data:        gin.H{"user": user},
	}
	res.CtxSetResponse(c)
}
