package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/http/response"
)

type UserController struct {
	userRepository UserRepositoryType
}

func NewUsersController(repository UserRepositoryType) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(model.Users)
	users, err := controller.userRepository.FindAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
