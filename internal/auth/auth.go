package auth

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/config"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/jwt"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/user"
)

type AuthServiceType interface {
	Login(user user.LoginRequest) (string, error)
	Register(user user.CreateUserRequest) error
}

type AuthService struct {
	UserRepository user.UserRepositoryType
	Validate       *validator.Validate
	cfg            *config.Config
}

func NewAuthService(cfg *config.Config, userRepository user.UserRepositoryType, validate *validator.Validate) AuthServiceType {
	return &AuthService{
		cfg:            cfg,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService
func (a *AuthService) Login(user user.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UserRepository.FindByUsername(user.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	verify_error := VerifyPassword(new_users.Password, user.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := jwt.GenerateToken(a.cfg.Jwt.TokenExpiresIn, new_users.Id, a.cfg.Jwt.TokenSecret)
	if err_token != nil {
		return "", err_token
	}
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthService) Register(users user.CreateUserRequest) error {

	hashedPassword, err := HashPassword(users.Password)
	if err != nil {
		return err
	}

	newUser := user.User{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	userErr := a.UserRepository.Save(newUser)
	if userErr != nil {
		return userErr
	}
	return nil
}
