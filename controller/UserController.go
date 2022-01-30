package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"oceanLearn/common"
	"oceanLearn/dto"
	"oceanLearn/model"
	"oceanLearn/util"
)

/**
  用户注册
*/
func Register(c *gin.Context) {
	//使用map获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(c.Request.Body).Decode(&requestMap)
	
	//获取参数
	//var requestUser = model.User{}
	//c.ShouldBindJSON(&requestUser)
	
	// 获取参数
	DB := common.GetDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	// 数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//如果没有传name,给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)
	// 判断手机号是否存在
	if isPhoneExist(DB, phone) {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	//创建用户
	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "加密错误!"})
		return
	}

	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hassedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

/**
用户登录
*/
func Login(c *gin.Context) {
	//获取参数
	DB := common.GetDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	//数据校验
	if phone == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "请输入手机号"})
		return
	}
	if len(phone) != 11 {
		c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "请输入密码"})
		return
	}

	// 判断手机号是否存在
	//if isPhoneExist(DB, phone) {
	//	c.JSON(http.StatusOK, gin.H{"code": 422, "msg": "用户已存在"})
	//	return
	//}
	//判断用户是否存在
	userInfo := getUserByName(DB, name)
	log.Println(userInfo)
	if userInfo.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "该用户不存在"})
		return
	}

	//判断密码是否正确
	bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "密码错误"})
		return
	}
	//发放token
	token,err := common.ReleaseToken(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登陆成功", "data": gin.H{"token": token}})
	return

}

/**
  根据姓名查询用户信息
*/
func getUserByName(db *gorm.DB, name string) model.User {
	var userInfo model.User
	db.Where("name = ?", name).First(&userInfo)
	return userInfo
}

/**
验证手机号
*/
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ? ", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

/**

 */
func Info(c *gin.Context)  {
	user,_ := c.Get("user")
	c.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
	
}