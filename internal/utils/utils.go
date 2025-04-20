package utils

import (
	"os"
	"path/filepath"
)

func IsGitRepository(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, ".git"))
	return err == nil
}
