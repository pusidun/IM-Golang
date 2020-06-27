package controller

import (
	"fmt"
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
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "PONG",
	})
}

func UserRegister(c *gin.Context) {
	DB := dao.DBConn()

	//Get params
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	fmt.Println(name, telephone, password)
	// Valid
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "密码不得少于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机已注册"})
		return
	}

	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密错误"})
		return
	}
	newUser := model.User{
		Nickname: name,
		Mobile:   telephone,
		Password: string(cryptPassword),
	}
	DB.Create(&newUser)

	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("mobile = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func UserLogin(c *gin.Context) {
	DB := dao.DBConn()

	//Get params
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// Valid
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "密码不得少于6位"})
		return
	}
	var user model.User
	DB.Where("mobile = ?", telephone).First(&user)
	if len(user.Nickname) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// token
	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token generate err"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
