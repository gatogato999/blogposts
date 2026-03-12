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

	return Post{
		Title:       readline(titleSeparator),
		Description: readline(descriptionSeparator),
		Tags:        strings.Split(readline(tagSeparator), ","),
		Body:        readBody(scnr),
	}, nil
}

func readBody(scnr *bufio.Scanner) string {
	scnr.Scan()

	buf := bytes.Buffer{}
	for scnr.Scan() {
		fmt.Fprintln(&buf, scnr.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

// Post : post type
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}
