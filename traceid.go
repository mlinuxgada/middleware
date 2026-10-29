package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	traceIDHeader = http.CanonicalHeaderKey("TraceId")
)

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(traceIDHeader)
		if traceID == "" {
			traceID = uuid.Must(uuid.NewV7()).String()
		}

		c.Set("traceId", traceID)
		c.Writer.Header().Set(traceIDHeader, traceID)

		c.Next()
	}
}
