package interfaces

import "attempt/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
	RegisterUser(user models.User) error
	FindByEmail(string) (*models.User, error)
	Login(email string, password string) (int, error)
	GetRole(email string) (string, error)
	UpdateUserProfile(email string, updatedUser models.User) error
	RegisterUserWithVerification(user models.User, token string) error
	FindEmailByVerificationToken(token string) (string, error)
	MarkEmailAsVerified(email string) error
}
