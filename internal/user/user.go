package user

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type UserRepositoryType interface {
	Save(user User) error
	Update(user User) error
	Delete(userId int) error
	FindById(userId int) (User, error)
	FindAll() ([]User, error)
	FindByUsername(username string) (User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepositoryType {
	return &UserRepository{Db: Db}
}

// Delete implements UsersRepository
func (u *UserRepository) Delete(userId int) error {
	var user User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements UsersRepository
func (u *UserRepository) FindAll() ([]User, error) {
	var users []User
	results := u.Db.Find(&users)
	if results.Error != nil {
		return nil, results.Error
	}
	return users, nil
}

// FindById implements UsersRepository
func (u *UserRepository) FindById(userId int) (User, error) {
	var user User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, result.Error
	}
}

// Save implements UsersRepository
func (u *UserRepository) Save(user User) error {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements UsersRepository
func (u *UserRepository) Update(user User) error {
	var updateUser = UpdateUserRequest{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindByUsername implements UsersRepository
func (u *UserRepository) FindByUsername(username string) (User, error) {
	var users User
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}
