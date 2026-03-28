package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/oklog/ulid/v2"
)

func main() {
	var paths []string

	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		paths = append(paths, path)
		return nil
	})

	fmt.Println(paths)
	fmt.Println(ulid.Make())
}
