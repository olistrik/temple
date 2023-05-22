package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

type FileNotFound struct {
	Path string
}

func (fnf FileNotFound) Error() string {
	return fmt.Sprintf("file not found: %s", filepath.Base(fnf.Path))
}

func NewFileNotFound(path string) error {
	return &FileNotFound{Path: path}
}

func foobar() {
	err := errors.New("foo bar")

	if errors.Is(err, NewFileNotFound("")) {
		//...
	}

}
