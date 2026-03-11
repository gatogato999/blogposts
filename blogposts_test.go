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
			"hw1.md": {Data: []byte("Title: two")},
		}
		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatalf("\nexpect no errors but got : %v", err)
		}
		if len(posts) != len(fs) {
			t.Errorf("\ngot: %v posts\nwant: %v posts", len(posts), len(fs))
		}
		got := posts[0]
		want := blogposts.Post{"two"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})

	t.Run("fs with error ", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Fatal("\nexpect errors but got non")
		}
	})
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
