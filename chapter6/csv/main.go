package main

import (
	"strconv"
	"os"
	"encoding/csv"
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string
}

func main () {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "B", Author: "Pierre"},
		Post{Id: 3, Content: "C", Author: "Pedro"},
		Post{Id: 4, Content: "D", Author: "Sau Sheong"},
	}

	write := csv.NewWriter(csvFile)
	for _, post := range allPosts{
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := write.Write(line)
		if err != nil {
			panic(err)
		}
	}
	write.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
}