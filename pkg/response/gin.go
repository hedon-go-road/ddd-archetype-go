package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinParamError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, HTTPResponse{
		RequestID: c.GetHeader(XRequestID),
		TraceID:   c.GetString(TraceIDKey),
		Code:      http.StatusBadRequest,
		Message:   err.Error(),
		Data:      nil,
	})
}

func GinError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, HTTPResponse{
		RequestID: c.GetHeader(XRequestID),
		TraceID:   c.GetString(TraceIDKey),
		Code:      http.StatusOK,
		Message:   err.Error(),
		Data:      nil,
	})
}

func GinSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, HTTPResponse{
		RequestID: c.GetHeader(XRequestID),
		TraceID:   c.GetString(TraceIDKey),
		Code:      http.StatusOK,
		Message:   "ok",
		Data:      data,
	})
}

func FromHTTPResponse[T any](response *HTTPResponse) *T {
	bs, _ := json.Marshal(response.Data)
	var res = new(T)
	_ = json.Unmarshal(bs, res)
	return res
}
