package routers

import (
	"blog/common/confsetting"
	_ "blog/docs"
	"blog/middleware"
	v1 "blog/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Group("/user")
	{
		r.POST("/login", v1.Login)
	}

	r.Group("/tag")
	r.Use(middleware.TokenCheck())
	{
		r.POST("/add", v1.AddTag)
		r.POST("/delete", v1.DeleteTag)
	}

	//实现热更新

	r.Run("0.0.0.0:" + confsetting.ServerPort)
}
