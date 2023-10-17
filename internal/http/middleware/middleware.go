package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			logger.Error("whoops", slog.String("error", ginErr.Error()))
		}
	}
}
