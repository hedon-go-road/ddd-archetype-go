package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/rand"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/response"
)

func WithRequestAndTrace() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(response.RequestIDKey, c.GetHeader(response.XRequestID))
		c.Set(response.TraceIDKey, rand.UUIDV7())
	}
}
