package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"oceanLearn/model"
)


var DB *gorm.DB
/**
连接数据库
*/
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	fmt.Println("driverName")
	fmt.Println(driverName)
	//host := "localhost"
	//port := "3306"
	//database := "ginessntial"
	//username := "root"
	//password := "root"
	//charset := "utf8"

	//args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%parseTime=true",
	//	username,
	//	password,
	//	host,
	//	port,
	//	database,
	//	charset)
	//db, err := gorm.Open(driverName, "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:@/ginessntial?charset=utf8&parseTime=True&loc=Local")

	//db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}