package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	// Implementation to read files from the provided FileSystem and create Post objects
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err //todo: handle error better, maybe skip the file instead of returning an error
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
