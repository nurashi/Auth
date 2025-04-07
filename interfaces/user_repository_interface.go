package interfaces

import "attempt/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
	RegisterUser(user models.User) error
	FindByEmail(string) (*models.User, error)
	Login(email string, password string) (int, error)
}
