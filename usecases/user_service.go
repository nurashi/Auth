package usecases

import (
	"attempt/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceImpl struct {
	UserRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
	}
}

func (u *UserServiceImpl) GetUsers(c *gin.Context) {
	users, err := u.UserRepo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
