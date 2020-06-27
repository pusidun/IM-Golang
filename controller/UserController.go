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
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "PONG",
	})
}

func UserLogin(c *gin.Context) {
	name := c.Param("nickname")
	fmt.Println("Login %s", name)
	return
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
		fmt.Println("随机生成name ", name)
	}

	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机已注册"})
		return
	}
	fmt.Println(name, telephone, password)
	newUser := model.User{
		Nickname: name,
		Mobile:   telephone,
		Password: password,
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
