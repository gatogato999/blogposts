package blogposts

import (
	"bufio"
	"io"
)

func newPost(postFile io.Reader) (Post, error) {
	scnr := bufio.NewScanner(postFile)

	readline := func() string {
		scnr.Scan()
		return scnr.Text()
	}
	title := readline()[7:]
	description := readline()[13:]

	post := Post{Title: title, Description: description}
	return post, nil
}

// Post : post type
type Post struct {
	Title       string
	Description string
}
