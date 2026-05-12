package app

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/oklog/ulid/v2"
)

type Fold struct {
	Out        string
	Root       string
	Files      []string
	Entries    []DirEntry
	EmptyDirs  int
	EmptyFiles int
}

type DirEntry struct {
	Path  string
	Depth int
}

func (f *Fold) WalkDir(root string) error {
	absPath, _ := filepath.Abs(root)
	homeDir, _ := os.UserHomeDir()
	cwdPath, _ := os.Getwd()

	f.Root = absPath
	rhs := strings.Count(root, PathSeparator)

	if !strings.HasPrefix(f.Root, homeDir) {
		return fmt.Errorf("%s is outside home directory.\n", root)
	}

	if !strings.HasPrefix(f.Root, cwdPath) {
		return fmt.Errorf("%s is outside current directory.\n", root)
	}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || root == path {
			return nil
		}

		if d.IsDir() {
			if d.Name() == f.Out {
				return filepath.SkipDir
			}

			lhs := strings.Count(path, PathSeparator)
			f.Entries = append(f.Entries, DirEntry{Path: path, Depth: lhs - rhs})
		} else {
			info, _ := d.Info()

			if info.Size() == 0 {
				f.EmptyFiles += 1
			}

			f.Files = append(f.Files, path)
		}

		return nil
	})
}

func (f *Fold) Unfold(root Root) {
	f.WalkDir(root.Path())
	f.Mkdir(f.Out)
	f.Sort()

	for _, entry := range f.Entries {
		entries, err := os.ReadDir(entry.Path)

		if err == nil && len(entries) == 0 {
			newPath := filepath.Join(f.Out, ulid.Make().String())
			os.Rename(entry.Path, newPath)
		}
	}

	for _, file := range f.Files {
		fmt.Println(file)
	}
}

func (f *Fold) Mkdir(name string) {
	info, err := os.Stat(name)

	if err == nil && info.IsDir() {
		return
	}

	os.Mkdir(name, 0755)
}

func (f *Fold) Sort() {
	sort.Slice(f.Entries, func(i, j int) bool {
		return f.Entries[i].Depth > f.Entries[j].Depth
	})
}

const OutDir = ".fold"
