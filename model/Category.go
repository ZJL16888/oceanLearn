package model

type Category struct {
	ID        int    `json:"id" gorm:"primary_key" `
	//gorm.Model        // gorm.Model 是一个包含一些基本字段的结构体, 包含的字段有 ID，CreatedAt， UpdatedAt， DeletedAt
	Name string `json:"name" gorm:"type:varchar(50); not null;unique"`
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
	
}
