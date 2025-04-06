package interfaces

import "attempt/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
}
