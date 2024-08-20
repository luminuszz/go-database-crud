package main

import (
	"awesomeProject/service"
	"context"
	"github.com/jackc/pgx/v5"
)

type Handlers struct {
	usersService *service.UsersService
	postService  *service.PostsService
}

var connection = "postgresql://docker:docker@localhost:5432/mydb"

func main() {

	db, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		panic(err)
	}

	defer db.Close(context.Background())

	usersService := service.NewUserService(db)
	postService := service.NewPostService(db)

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

	postError := postService.CreatePost(&service.Post{
		Title:    "Title",
		Content:  "Content",
		AuthorId: 1,
	})

	if postError != nil {
		panic(postError)
	}

	err, postsWithAutorList := postService.FindAllPostByAuthor(1)

	if err != nil {
		panic(err)
	}

	for _, post := range postsWithAutorList {
		println(post.Id)
		println(post.Title)
		println(post.Content)
		println(post.Author)
		println("------------")

	}

}
