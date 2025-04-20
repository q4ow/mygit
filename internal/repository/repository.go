package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/q4ow/mygit/internal/objects"
	"github.com/q4ow/mygit/internal/utils"
)

const gitDir = ".git"

var objectsDir = filepath.Join(gitDir, "objects")

func InitRepository() error {
	if utils.IsGitRepository(".") {
		return fmt.Errorf("Already a Git repository")
	}

	if err := os.MkdirAll(gitDir, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		return err
	}

	fmt.Println("Initialized empty repository")
	return nil
}

func AddFiles(files []string) error {
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		blob := objects.NewBlob(content)
		fmt.Printf("Added file: %s\n", file)

		if err := storeObject(blob); err != nil {
			return err
		}
	}
	return nil
}

func Commit(message string) error {
	tree := objects.NewTree([]objects.TreeEntry{
		{Name: "test.txt", Type: "blob", Hash: ""},
	})

	commit := objects.NewCommit(tree, nil, message)

	if err := storeObject(commit); err != nil {
		return err
	}

	fmt.Printf("[main (root)]: Created initial commit\n")
	return nil
}

func storeObject(obj objects.ObjectInterface) error {
	objDir := filepath.Join(objectsDir, obj.Hash()[:2])
	if err := os.MkdirAll(objDir, 0755); err != nil {
		return err
	}

	objPath := filepath.Join(objDir, obj.Hash()[2:])
	if err := os.WriteFile(objPath, []byte(obj.Serialize()), 0644); err != nil {
		return err
	}
	return nil
}
