package interfaces

import (
	"attempt/models"
	"database/sql"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u UserRepositoryImpl) GetUsers() ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}

		users = append(users, user)
	}
	return users, nil
}
