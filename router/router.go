package router

import (
	"github.com/gin-gonic/gin"
	"go-music/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 使用中间件处理跨域问题
	r.Use(middlewares.Cors())

	// 静态文件路由 TODO MinIO
	r.Static("/img", "./static/img")
	// 前端请求中 /song/music.mp3 与 /song/detail 路由有歧义 因此歌曲请求尚未解决
	// r.Static("/song", "./static/song")

	// 路由分组
	group := r.Group("/")
	InitAdmin(group)
	InitBanner(group)
	InitSong(group)
	InitSinger(group)
	InitUser(group)
	InitRank(group)
	InitComment(group)
	InitCollect(group)

	return r
}
