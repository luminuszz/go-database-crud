package main

import (
	"awesomeProject/service"
	"context"
	"github.com/jackc/pgx/v5"
)

type Handlers struct {
	usersService *service.UsersService
}

var connection = "postgresql://docker:docker@localhost:5432/mydb"

func main() {

	db, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		panic(err)
	}

	defer db.Close(context.Background())

	usersService := service.NewUserService(db)

	result, err := usersService.GetAllUsers()

	usersError := usersService.CreateUser(&service.User{
		Name:         "Carlos",
		Age:          10,
		Email:        "davi@gmail.com",
		PasswordHash: "123456",
	})
	if usersError != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	for _, user := range result {
		println(user.Name)
		println(user.Age)
		println(user.Email)
		println(user.PasswordHash)
		println("------------")
	}

}
