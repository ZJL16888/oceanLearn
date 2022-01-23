package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model        // gorm.Model 是一个包含一些基本字段的结构体, 包含的字段有 ID，CreatedAt， UpdatedAt， DeletedAt
	Name       string `gorm:"type:varchar(20);not null"`
	Phone      string `gorm:"type:varchar(11);not null;unique"`
	Password   string `gorm:"size:255;not null"`
}
