package interfaces

import "attempt/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
	RegisterUser(user models.User) error
	FindByEmail(string) (int, error)
}
