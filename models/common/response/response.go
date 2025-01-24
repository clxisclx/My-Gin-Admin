package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	FAIL            = 7
	SUCCESS         = 0
	Login_FAIL      = 10   // 登录失败
	REQ_PARAM_ERROR = 11   // 请求参数错误
	QUERY_NOT_DATA  = 1000 // 未查询到数据
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
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
	Result(FAIL, nil, "操作失败", c)
}

func FailWithMsg(c *gin.Context, msg string) {
	Result(FAIL, nil, msg, c)
}

func LoginFail(c *gin.Context, msg string) {
	Result(Login_FAIL, nil, msg, c)
}

func ParamError(c *gin.Context) {
	Result(REQ_PARAM_ERROR, nil, "参数错误", c)
}

func QueryNotData(c *gin.Context) {
	Result(QUERY_NOT_DATA, nil, "操作成功", c)
}

func Error(code int, data interface{}, msg string, c *gin.Context) {
	Result(code, data, msg, c)
}

func FailWithInternalServer(c *gin.Context) {
	c.JSON(200, Response{
		http.StatusInternalServerError,
		nil,
		"网络异常",
	})
}
