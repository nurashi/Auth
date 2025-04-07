package usecases

import (
	"attempt/adapters/httpAuth"
	"attempt/interfaces"
	"attempt/models"
	"attempt/utils/hash"
	"attempt/utils/jwtAuth"
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

	_, err := u.UserRepo.FindByEmail(newUser.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	hashedPassword, err := hash.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad password"})
		return
	}

	newUser.Password = hashedPassword

	verificationToken, err := jwtAuth.GenerateEmailVerificationToken()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "eerror with verification of email"})
		return
	}

	err = u.UserRepo.RegisterUserWithVerification(newUser, verificationToken)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "eerror with verification of email SECOND"})
		return
	}

	go httpAuth.SendVerificationEmail(newUser.Email, verificationToken)

	//if err := u.UserRepo.RegisterUser(newUser); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"message": "Error with RegisterUser"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"message": "registreated == true"})
}

func (u *UserServiceImpl) VerifyEmail(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "token is empty"})
		return
	}

	email, err := u.UserRepo.FindEmailByVerificationToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email does not exist"})
		return
	}

	err = u.UserRepo.MarkEmailAsVerified(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with MarkEmailAsVerified"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email verified"})
}

func (u *UserServiceImpl) Login(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error with login"})
		return
	}

	user, err := u.UserRepo.FindByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such email"})
		return
	}

	if !hash.CheckPasswordHash(input.Pass, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	token, err := jwtAuth.GenerateToken(user.Email, user.Role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error generating token"})
		return
	}

	_, err = u.UserRepo.GetRole(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Role error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserServiceImpl) GetProfile(c *gin.Context) {
	email, _ := c.Get("email")

	user, err := u.UserRepo.FindByEmail(email.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such email"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserServiceImpl) UpdateProfile(c *gin.Context) {
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body request"})
		return
	}

	email, _ := c.Get("email")

	err := u.UserRepo.UpdateUserProfile(email.(string), updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error updating user profile"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated == true"})
}
