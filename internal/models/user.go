package models

import (
	"BookingService/internal/db"
	"BookingService/internal/utills"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int
	Email    string `binding:"required" valid:"email"`
	Password string `binding:"required" valid:"stringlength(8|15)"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES ($1,$2)"
	hashedPassword, err := utills.HashPassword(u.Password)
	if err != nil {
		return err
	}
	res, err := db.ConnectionPool.Query(context.Background(), query, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	res.Close()
	return nil
}

func GetUserById(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	rows, err := db.ConnectionPool.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	rows, err := db.ConnectionPool.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateUserBy(user User) error {
	_, err := GetEventById(user.ID)
	if err != nil {
		return err
	}
	query := "UPDATE users SET email = $1, password = $2 WHERE id = $3"
	rows, err := db.ConnectionPool.Query(context.Background(), query, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func DeleteUser(id int) error {

	query := "DELETE FROM users WHERE id = $1"
	rows, err := db.ConnectionPool.Query(context.Background(), query, id)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func (u User) ValidateCredentials() (int, error) {
	user, err := GetUserByEmail(u.Email)
	if err != nil {
		return 0, errors.New("invalid credentials")
	}
	if !utills.CheckPasswords(u.Password, user.Password) {
		return 0, errors.New("invalid credentials")
	}
	return user.ID, nil
}
