package router

import (
	"github.com/gin-gonic/gin"
	"go-music/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件路由 TODO MinIO

	// 使用中间件处理跨域问题
	r.Use(middlewares.Cors())

	// 路由分组
	group := r.Group("/")
	InitAdmin(group)
	InitBanner(group)
	InitSong(group)
	InitSinger(group)

	return r
}
