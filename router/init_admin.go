package router

import "github.com/gin-gonic/gin"

func InitAdmin(r *gin.RouterGroup) {
	r.Group("/admin")
	{
		//router.POST("/login")
	}
}
