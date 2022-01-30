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

type IPostController interface {
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("id desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)
	//if err != nil {
	//	panic(err)
	//}
	//总条数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	util.Success(ctx, gin.H{"data": posts, "total": total}, "成功")

}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePost
	if error := ctx.ShouldBind(&requestPost); error != nil {
		util.Fail(ctx, nil, "文章名称必填")
		return
	}

	//ctx.Bind(&requestCategory)
	//// 数据验证
	//if requestCategory.Name == "" {
	//	util.Fail(ctx,nil,"分类名称必填")
	//	return
	//}
	//

	//获取用户信息
	user, _ := ctx.Get("user")

	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	util.Success(ctx, gin.H{"post": post}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePost
	if error := ctx.ShouldBind(&requestPost); error != nil {
		util.Fail(ctx, nil, "文章名称必填")
		return
	}

	//获取 path中的postid
	postId := ctx.Params.ByName("id")
	var updatePost model.Post
	if p.DB.Where("id = ?", postId).First(&updatePost).RecordNotFound() {
		util.Fail(ctx, nil, "找不到该文章")
		return
	}

	//获取用户信息
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != updatePost.UserId {
		util.Fail(ctx, nil, "文章不属于您,请勿非法操作")
		return
	}

	//更新文章
	if err := p.DB.Model(&updatePost).Update(requestPost).Error; err != nil {
		util.Fail(ctx, nil, "文章更新失败")
		return
	}
	util.Success(ctx, gin.H{"category": updatePost}, "修改成功")

}

func (p PostController) Show(ctx *gin.Context) {
	//获取 path中的postid
	postId := ctx.Params.ByName("id")
	var updatePost model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&updatePost).RecordNotFound() {
		util.Fail(ctx, nil, "找不到该文章")
		return
	}

	//返回文章
	util.Success(ctx, gin.H{"post": updatePost}, "获取成功")

}

func (p PostController) Delete(ctx *gin.Context) {
	//获取 path中的postid
	postId := ctx.Params.ByName("id")
	var updatePost model.Post
	if p.DB.Where("id = ?", postId).First(&updatePost).RecordNotFound() {
		util.Fail(ctx, nil, "找不到该文章")
		return
	}

	//获取用户信息
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != updatePost.UserId {
		util.Fail(ctx, nil, "文章不属于您,请勿非法操作")
		return
	}

	if err := p.DB.Delete(updatePost).Error; err != nil {
		//fmt.Println("ss")
		//panic(err)
		util.Fail(ctx, nil, "删除失败")
		return
	}

	util.Success(ctx, nil, "删除成功")
}

func (p PostController) Patch(ctx *gin.Context) {
	panic("implement me")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}
