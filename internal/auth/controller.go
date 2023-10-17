package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/http/response"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/user"
)

type AuthController struct {
	authService AuthServiceType
}

func NewAuthController(service AuthServiceType) *AuthController {
	return &AuthController{authService: service}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := user.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err_token := controller.authService.Login(loginRequest)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	resp := user.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := user.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	controller.authService.Register(createUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
