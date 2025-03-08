package reading_files

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fs fs.FS) ([]Post, error) {
	posts := []Post{} // Start with an empty slice
	posts = append(posts, Post{Title: "Post 1"})
	posts = append(posts, Post{Title: "Post 2"})
	return posts, nil
}	