package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
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

	scnr.Scan()

	buf := bytes.Buffer{}
	for scnr.Scan() {
		fmt.Fprintln(&buf, scnr.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

// Post : post type
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}
