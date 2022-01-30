package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oceanLearn/common"
	"oceanLearn/model"
	"oceanLearn/util"
	"oceanLearn/vo"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB

}

func (c CategoryController) PageList(ctx *gin.Context) {
	panic("implement me")
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	
	return CategoryController{DB:db}
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定 body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	
	// 数据验证
	if requestCategory.Name == "" {
		util.Fail(ctx,nil,"分类名称必填")
		return
	}
	
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
		
		util.Fail(ctx,nil,"找不到该分类")
		return
	}
	//更新分类
	c.DB.Model(&updateCategory).Update("name",requestCategory.Name)
	util.Success(ctx,gin.H{"category":updateCategory},"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category,categoryId).RecordNotFound(){
		util.Fail(ctx,nil,"找不到该分类")
		return
	}
	
	//返回分类
	util.Success(ctx,gin.H{"category":category},"获取成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{},categoryId).Error; err != nil {
		//fmt.Println("ss")
		//panic(err)
		
		util.Fail(ctx,nil,"删除失败")
		return
	}
	
	util.Success(ctx,nil,"删除成功")
}

func (c CategoryController) Create (ctx *gin.Context)  {
	var requestCategory vo.CreateCategory
	if error := ctx.ShouldBind(&requestCategory); error != nil{
		util.Fail(ctx,nil,"分类名称必填")
		return
	}
	
	//ctx.Bind(&requestCategory)
	//// 数据验证
	//if requestCategory.Name == "" {
	//	util.Fail(ctx,nil,"分类名称必填")
	//	return
	//}
	category := model.Category{Name: requestCategory.Name}
	
	c.DB.Create(&category)
	util.Success(ctx,gin.H{"category":category},"")
	
}
func (c CategoryController) Patch (ctx *gin.Context)  {

}