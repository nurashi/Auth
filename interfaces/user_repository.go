package interfaces

import (
	"attempt/models"
	"database/sql"
	"fmt"
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
	query := "SELECT id, name, age, email, phone, password, job, country, role FROM users WHERE email = $1"
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
		&user.Role,
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

func (u UserRepositoryImpl) GetRole(email string) (string, error) {
	query := "SELECT role FROM users WHERE email = $1"
	row := u.DB.QueryRow(query, email)

	var role string
	err := row.Scan(&role)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no user found with email %s", email)
	}

	if err != nil {
		return "", err
	}

	return role, nil
}

func (u *UserRepositoryImpl) UpdateUserProfile(email string, updatedUser models.User) error {
	query := `UPDATE users 
			  SET name = $1, age = $2, phone = $3, job = $4, country = $5
			  WHERE email = $6`
	_, err := u.DB.Exec(query, updatedUser.Name, updatedUser.Age, updatedUser.Phone, updatedUser.Job, updatedUser.Country, email)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) RegisterUserWithVerification(user models.User, token string) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO users (name, age, token, phone, password, job, country, role) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = tx.Exec(query, user.Name, user.Age, token, user.Phone, user.Password, user.Job, user.Country, user.Role)
	if err != nil {
		return err
	}

	verificationQuery := "INSERT INTO verification_tokens (email, token) VALUES ($1, $2)"
	_, err = tx.Exec(verificationQuery, user.Email, token)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *UserRepositoryImpl) FindEmailByVerificationToken(token string) (string, error) {
	var email string
	query := `SELECT email FROM verification_tokens WHERE token = $1`
	err := u.DB.QueryRow(query, token).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (u *UserRepositoryImpl) MarkEmailAsVerified(email string) error {
	query := `UPDATE users SET email_verified = TRUE WHERE email = $1`
	_, err := u.DB.Exec(query, email)
	return err
}
