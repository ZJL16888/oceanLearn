package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"oceanLearn/common"
	"oceanLearn/model"
	"oceanLearn/util"
)

func Register(c *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	// 数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "手机号必须为11位"})
	}
	if len(password) < 6 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "密码不能少于6位"})
	}
	//如果没有传name,给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)
	// 判断手机号是否存在
	if isPhoneExist(DB, phone) {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "用户已存在"})
	}

	//创建用户
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

/**
验证手机号
*/

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user  model.User
	db.Where("phone = ? ", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
