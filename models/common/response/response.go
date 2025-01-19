package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(200, Response{
		code,
		data,
		msg,
	})
}

func Success(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, nil, "操作失败", c)
}
