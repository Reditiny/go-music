package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应
func Response(c *gin.Context, httpStatus int, code int, data interface{}, message string, t string, s bool) {
	c.JSON(httpStatus, gin.H{"code": code, "message": message, "type": t, "success": s, "data": data})
}

// Success 返回前端-成功
func Success(c *gin.Context, data interface{}, message string) {
	Response(c, http.StatusOK, 200, data, message, "success", true)
}

// Fail 返回前端-失败
func Fail(c *gin.Context, data interface{}, message string) {
	Response(c, http.StatusBadRequest, 400, data, message, "error", false)
}
