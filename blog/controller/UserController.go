package controller

import (
	"blog/Dao"
	"blog/model"
	"blog/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	db := Dao.GetDB()
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11为",
		})
		return
	}
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isHasTelephone(db, telephone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "用户已经存在",
		})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)
	log.Println(name, telephone, password)
	context.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func isHasTelephone(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
