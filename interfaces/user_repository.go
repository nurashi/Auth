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
	query := "SELECT id, name, age, email, phone, password, job, country FROM users"
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Age,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Job,
			&user.Country,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepositoryImpl) RegisterUser(user models.User) error {
	query := "INSERT INTO users (name, age, email, phone, password, job, country) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := u.DB.Exec(query, user.Name, user.Age, user.Email, user.Phone, user.Password, user.Job, user.Country)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) FindByEmail(email string) (int, error) {
	query := "SELECT id FROM users WHERE email = $1"

	row := u.DB.QueryRow(query, email)
	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // its just a case when email exactly not found
		}
		return 0, err // here can be any error with database
	}
	return id, nil
}
