package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/gatogato999/blogposts"
)

func TestNewPlogPosts(t *testing.T) {
	t.Run("fs with no errors", func(t *testing.T) {
		fs := fstest.MapFS{
			// "hw.md": {Data: []byte("Title: one")},
			"hw1.md": {Data: []byte(`Title: Post 1
Description: Description 1`)},
		}
		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatalf("\nexpect no errors but got : %v", err)
		}
		if len(posts) != len(fs) {
			t.Errorf("\ngot: %v posts\nwant: %v posts", len(posts), len(fs))
		}
		assertPostEquality(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
		})
	})

	t.Run("fs with error ", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Fatal("\nexpect errors but got non")
		}
	})
}

func assertPostEquality(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot: %v\nwant: %v", got, want)
	}
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
