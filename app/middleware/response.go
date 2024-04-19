package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Response struct {
	Code    int         `json:"code"`
	Message interface{}      `json:"msg"`
	Data    interface{} `json:"data"`
}

// UnifiedResponseMiddleware是处理统一HTTP响应格式的中间件
func UnifiedResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 继续执行其他中间件和路由

		// 当处理完后，发送统一的JSON响应
		responseData := c.Keys["ResponseData"].(Response)
		c.JSON(http.StatusOK, responseData)
	}
}
