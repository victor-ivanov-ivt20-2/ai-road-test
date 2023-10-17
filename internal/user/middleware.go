package user

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/config"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/jwt"
)

func DeserializeUser(cfg *config.Config, userRepository UserRepositoryType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := jwt.ValidateToken(token, cfg.Jwt.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		if err_id != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err_id)
			return
		}
		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()

	}
}
