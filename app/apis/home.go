package apis

import (
	"go-app/app/middleware"
	"go-app/app/models"
	"go-app/app/vali"
	"go-app/pkg/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})

	r.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	// 定义一个处理GET请求的路由和处理程序
	r.GET("/hello", func(c *gin.Context) {
		c.Set("ResponseData", middleware.Response{
			Code:    200,
			Message: "Success",
			Data: "hello world",
		})
	})
}

func SetRegisterGroupROutes(r *gin.RouterGroup) {
	r.POST("/signup", func(c *gin.Context) {
		var u models.ParamRegister
		if err := c.ShouldBind(&u); err != nil {
			// 获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非validator.ValidationErrors类型错误直接返回
				c.Set("ResponseData", middleware.Response{
					Message: err.Error(),
				})
				return
			}
			// validator.ValidationErrors类型错误则进行翻译
			c.Set("ResponseData", middleware.Response{
				Message: vali.RemoveTopStruct(errs.Translate(common.Trans)),
			})
			return
		}
		// 保存入库等业务逻辑代码...

		c.Set("ResponseData", middleware.Response{
			Code:    200,
			Message: "Success",
			Data: map[string]string{
				"key": "value",
			},
		})
	})
}
