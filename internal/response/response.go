package response

import (
	"net/http"

	"github.com/hubogle/chatcode-server/internal/code"

	"github.com/gin-gonic/gin"
)

// UnifiedResponse 统一返回
type UnifiedResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"msg"`
}

// BadResponse 错误返回
type BadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// HandlerParamsResponse 处理参数响应错误
func HandlerParamsResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, BadResponse{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	})
}

// HandleResponse 统一返回处理 {"msg":"","data":{},"code":200}
func HandleResponse(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, UnifiedResponse{
			Code:    http.StatusOK,
			Data:    data,
			Message: "success",
		})
	} else {
		coder := code.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), BadResponse{
			Code:    coder.Code(),
			Message: coder.String(),
		})
	}
}
