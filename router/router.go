package router

import (
	c "gin_app/controller"
	w "gin_app/websocket"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoete() *gin.Engine {
	route := gin.Default()
	// 设置代理为本机, 避免控制台的安全警告
	route.SetTrustedProxies([]string{"127.0.0.1"})
	// 重定向到swagger 界面
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:80/swagger/index.html")
	})

	r := route.Group("/api")
	// 路由表
	{
		go r.GET("/ping", w.Ping)
		r.GET("/test", c.FindTest)
		r.POST("/regis", c.RegisterUser)
		r.DELETE("/del/:name", c.DeleteByName)
		r.PATCH("/upd/:name", c.UpdateByName)
		r.GET("/find/:name", c.FindUserByName)
		r.POST("/add", c.AddUser)
	}

	return route
}
