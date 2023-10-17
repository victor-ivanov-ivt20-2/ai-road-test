package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/auth"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/config"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/http/middleware"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/user"
)

func NewRouter(cfg *config.Config, log *slog.Logger,
	userRepository user.UserRepositoryType,
	authController *auth.AuthController,
	userController *user.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	service.Use(middleware.ErrorHandler(log))

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authController.Register)
	authenticationRouter.POST("/login", authController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", user.DeserializeUser(cfg, userRepository), userController.GetUsers)

	return service
}
