package blogposts

import (
	"io"
	"io/fs"
)

// Post : post type
type Post struct {
	Title string
}

// NewPostsFromFS : used to return posts
func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	file, readError := fs.ReadDir(fileSys, ".")
	if readError != nil {
		return nil, readError
	}
	var posts []Post
	for _, val := range file {
		post, err := getFileContents(fileSys, val)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getFileContents(fileSys fs.FS, file fs.DirEntry) (Post, error) {
	postFile, openError := fileSys.Open(file.Name())
	if openError != nil {
		return Post{}, openError
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile fs.File) (Post, error) {
	contents, readError := io.ReadAll(postFile)
	if readError != nil {
		return Post{}, readError
	}
	post := Post{Title: string(contents[7:])}
	return post, nil
}
