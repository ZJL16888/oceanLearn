package main

import (
	"github.com/gin-gonic/gin"
	"oceanLearn/controller"
	"oceanLearn/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoverMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info",middleware.AuthMiddleware(), controller.Info)
	
	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.PUT("/:id",categoryController.Update)
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE("/:id",categoryController.Delete)
	categoryRoutes.PATCH("/:id",categoryController.Patch)
	
	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("",postController.Create)
	postRoutes.PUT("/:id",postController.Update)
	postRoutes.GET("/:id",postController.Show)
	postRoutes.DELETE("/:id",postController.Delete)
	postRoutes.PATCH("/:id",postController.Patch)
	postRoutes.GET("/page/list",postController.PageList)
	
	return r
}
