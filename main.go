package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-music/common"
	"go-music/config"
	"go-music/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化数据库 缓存
	common.InitDatabase()
	common.InitRedis()
	// 配置路由
	r := router.InitRouter()
	// 将数据库连接池与 Gin 上下文关联
	r.Use(func(c *gin.Context) {
		c.Set("db", common.DB)
		c.Next()
	})
	// 优雅地重启或停止
	srv := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
