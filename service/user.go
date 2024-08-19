package service

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id           int
	Name         string
	Age          int
	Email        string
	PasswordHash string
}

type UsersService struct {
	db *pgx.Conn
}

var ctx = context.Background()

func NewUserService(db *pgx.Conn) *UsersService {
	return &UsersService{db: db}
}

func (u *UsersService) CreateUser(user *User) error {
	query := "INSERT INTO users (name, age, email, password_hash) VALUES ($1, $2, $3, $4)"

	_, err := u.db.Exec(ctx, query, user.Name, user.Age, user.Email, user.PasswordHash)

	if err != nil {
		return error(err)
	}
	return nil
}

func (u *UsersService) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE  id = $1"

	_, err := u.db.Exec(ctx, query, id)

	if err != nil {
		return error(err)
	}

	return nil
}

func (u *UsersService) GetAllUsers() ([]User, error) {

	query := "SELECT * FROM users"

	rows, err := u.db.Query(ctx, query)

	if err != nil {
		return nil, error(err)
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.PasswordHash)

		if err != nil {
			return nil, error(err)
		}
		users = append(users, user)
	}

	return users, nil

}
