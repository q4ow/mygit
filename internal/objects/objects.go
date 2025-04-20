package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Object struct {
	Type    string
	Content []byte
	Size    int64
}

type ObjectInterface interface {
	Hash() string
	Serialize() string
}

type Blob struct {
	Object
}

type Tree struct {
	Object
	Entries []TreeEntry
}

type TreeEntry struct {
	Name string
	Type string
	Hash string
}

type Commit struct {
	Object
	Tree      *Tree
	Parent    *Commit
	Author    string
	Committer string
	Message   string
}

func NewBlob(content []byte) *Blob {
	return &Blob{Object{Type: "blob", Content: content, Size: int64(len(content))}}
}

func NewTree(entries []TreeEntry) *Tree {
	return &Tree{Object{Type: "tree", Content: nil, Size: 0}, entries}
}

func NewCommit(tree *Tree, parent *Commit, message string) *Commit {
	return &Commit{
		Object{Type: "commit", Content: nil, Size: 0},
		tree,
		parent,
		"Your Name",
		"Your Name",
		message,
	}
}

func (o *Object) Hash() string {
	hasher := sha1.New()
	hasher.Write(o.Content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (o *Object) Serialize() string {
	return fmt.Sprintf("%02x%04x%s", len(o.Content), o.Type, hex.EncodeToString(o.Content))
}
