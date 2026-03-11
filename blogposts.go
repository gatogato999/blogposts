package blogposts

import (
	"io/fs"
)

// NewPostsFromFS : used to return posts
func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	file, readError := fs.ReadDir(fileSys, ".")
	if readError != nil {
		return nil, readError
	}
	var posts []Post
	for _, val := range file {
		post, err := getFileContents(fileSys, val.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getFileContents(fileSys fs.FS, file string) (Post, error) {
	postFile, openError := fileSys.Open(file)
	if openError != nil {
		return Post{}, openError
	}
	defer postFile.Close()
	return newPost(postFile)
}
