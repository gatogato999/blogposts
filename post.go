package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scnr := bufio.NewScanner(postFile)

	readline := func(separator string) string {
		scnr.Scan()
		return strings.TrimPrefix(scnr.Text(), separator)
	}

	title := readline(titleSeparator)
	description := readline(descriptionSeparator)
	tags := strings.Split(readline(tagSeparator), ",")
	_ = readline("")
	body := readline("")

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

// Post : post type
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}
