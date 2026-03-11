package blogposts

import (
	"bufio"
	"io"
)

func newPost(postFile io.Reader) (Post, error) {
	scnr := bufio.NewScanner(postFile)

	scnr.Scan()
	titleLine := scnr.Text()
	scnr.Scan()
	descLine := scnr.Text()

	post := Post{Title: string(titleLine[7:]), Description: descLine[13:]}
	return post, nil
}

// Post : post type
type Post struct {
	Title       string
	Description string
}
