package main

import (
	"gin_app/docs"
	r "gin_app/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger 测试 API"
	docs.SwaggerInfo.Description = "测试 gin 服务器."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:80"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api"
	r := r.InitRoete()
	// swagger 界面的路由地址
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("127.0.0.1:80")
}
