package service

import "github.com/jackc/pgx/v5"

type Post struct {
	Id       int
	Title    string
	Content  string
	AuthorId int
}

type PostWithAuthor struct {
	Id       int
	Title    string
	Content  string
	AuthorId int
	Author   string
}

type PostsService struct {
	db *pgx.Conn
}

func NewPostService(db *pgx.Conn) *PostsService {
	return &PostsService{db: db}
}

func (p *PostsService) CreatePost(post *Post) error {
	query := "INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)"

	_, err := p.db.Exec(ctx, query, post.Title, post.Content, post.AuthorId)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostsService) FindPostById(id int) (error, *Post) {
	query := "SELECT * FROM posts WHERE id = $1"

	results := p.db.QueryRow(ctx, query, id)

	var post Post

	err := results.Scan(&post.Id, &post.Title, &post.Content, &post.AuthorId)

	if err != nil {
		return err, nil
	}

	return nil, &post
}

func (p *PostsService) FindAllPostByAuthor(authorId int) (error, []PostWithAuthor) {

	query := "SELECT p.id as id,title,content,author_id, u.name FROM posts as p INNER JOIN users as u ON u.id = p.author_id"

	rows, err := p.db.Query(ctx, query)

	if err != nil {
		return err, nil
	}

	var list []PostWithAuthor

	for rows.Next() {
		var postWithAutor PostWithAuthor
		err = rows.Scan(&postWithAutor.Id, &postWithAutor.Title, &postWithAutor.Content, &postWithAutor.AuthorId, &postWithAutor.Author)
		if err != nil {
			return err, nil
		}

		list = append(list, postWithAutor)
	}

	return nil, list
}
