# Blogpost

A TDD example used to convert `.md` file into a social media post.

## Functionality

- given a `.md` file :

```md
Title: Post 1
Description: Description 1
Tags: rust,borrow-checker
---
this is body
and this is the reminder of the file
it may be as long as ever
how to get it ?`
```

- Use the `NewPostsFromFS` Function to get the `post` with this properties :

```go
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}
```

## Haven't Handled Yet:

- [ ] incorrect file formats.
- [ ] non `.md` files.
- [ ] unordered file meta data.

## Used in Testing Environment

1. [gotestsum](https://github.com/gotestyourself/gotestsum)
