package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
}

// ReturnSuccess for success 格式化返回
// stat=1成功，成功时code=0且msg=success
// stat=0失败，失败时有非0的code和具体msg信息
func ReturnSuccess(ctx *gin.Context, data interface{}, count int) {
	ctx.JSON(http.StatusOK, response{Code: 0, Message: "ok", Data: data, Count: count})
	ctx.Abort()
}

func ReturnSuccessNoCount(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, response{Code: 0, Message: "ok", Data: data})
	ctx.Abort()
}

//ReturnFail for format when fail
func ReturnFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, response{Code: 1, Message: err.Error(), Data: nil, Count: 0})
	ctx.Abort()
}

//ReturnHTML 输出html页面
func ReturnHTML(ctx *gin.Context, view string, data interface{}) {
	ctx.HTML(http.StatusOK, view, data)
}
