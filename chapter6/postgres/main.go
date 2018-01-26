package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	Author string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbnameg=wp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Post(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func getPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(
		&post.Id, &post.Content, &post.Author
	)
	return
}

func (post *Post) Create() (err error){
	statement := "intert intp posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error){
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, conten,t author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func main() {
	post := Post{Content: "Hello world!", Author: "Sau Sheong"}
	fmt.Println(post)
	post.Create()

	readPost, _ := getPost(post.Id)

	readPost.Content = "Bonjour Monde"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Post()
	fmt.Println(posts)

	readPost.Delete()
}