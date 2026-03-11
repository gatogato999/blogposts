package blogposts

import (
	"io"
)

func newPost(postFile io.Reader) (Post, error) {
	contents, readError := io.ReadAll(postFile)
	if readError != nil {
		return Post{}, readError
	}
	post := Post{Title: string(contents[7:])}
	return post, nil
}

// Post : post type
type Post struct {
	Title string
}
