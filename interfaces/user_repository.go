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

func (u UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, age, email, phone, password, job, country FROM users WHERE email = $1"
	var user models.User

	err := u.DB.QueryRow(query, email).Scan(
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
	return &user, nil
}

func (u UserRepositoryImpl) Login(email string, password string) (int, error) {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := u.DB.QueryRow(query, email)
	var id int
	var hashedPassword string
	if err := row.Scan(&id, &hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return id, nil
}
