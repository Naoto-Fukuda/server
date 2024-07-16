package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: 1, Content: "First Content", Author: "Kevin"},
		{Id: 2, Content: "Second Content", Author: "Mike"},
		{Id: 3, Content: "Third Content", Author: "Kathy"},
		{Id: 4, Content: "Fourth Content", Author: "Catherine"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		post := Post{
			Id:      id,
			Content: record[1],
			Author:  record[2],
		}
		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Printf("Id: %d, Content: %s, Author: %s\n", post.Id, post.Content, post.Author)
	}
}
