package app

import (
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	OutDir string
	Prune  bool
	Root   Root
}

type Root struct {
	Value string
}

func (c *Config) Register(fs *flag.FlagSet) {
	fs.StringVar(&c.OutDir, "out", ".fold", "")
	fs.BoolVar(&c.Prune, "prune", false, "")
}

func (c *Config) SetRoot(path string) {
	c.Root.Value = path
}

func (r *Root) IsDir() bool {
	info, err := os.Stat(r.Value)

	if err == nil {
		return info.IsDir()
	}

	return false
}

func (r *Root) Path() string {
	return filepath.Clean(r.Value)
}

const PathSeparator = string(os.PathSeparator)
