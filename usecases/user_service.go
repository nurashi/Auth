package usecases

import (
	"attempt/interfaces"
	"attempt/models"
	"attempt/utils/hash"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

func (u *UserServiceImpl) Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error with registration, caused by ShouldBind"})
		return
	}

	if !strings.Contains(newUser.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email should contain @"})
		return
	}

	selectedID, err := u.UserRepo.FindByEmail(newUser.Email)
	if err == nil && selectedID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	hashedPassword, err := hash.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad password"})
		return
	}

	newUser.Password = hashedPassword

	if err := u.UserRepo.RegisterUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error with RegisterUser"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registreated == true"})
}
