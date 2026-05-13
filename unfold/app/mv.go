package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Move(src string, dest string) error {
	base := filepath.Base(src)
	suffix := filepath.Ext(base)

	destPath := filepath.Join(dest, base)
	_, err := os.Stat(destPath)

	if os.IsNotExist(err) {
		return os.Rename(src, destPath)
	}

	for i := range 1000 {
		baseName := strings.TrimSuffix(base, suffix)
		alternative := fmt.Sprintf("%s(%d)%s", baseName, i+1, suffix)

		destPath = filepath.Join(dest, alternative)
		_, err := os.Stat(destPath)

		if os.IsNotExist(err) {
			return os.Rename(src, destPath)
		}
	}

	return nil
}
