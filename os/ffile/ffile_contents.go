package ffile

import (
	"fmt"
	"io"
	"os"
)

// PutContents puts string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
func PutContents(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutContentsAppend appends string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
func PutContentsAppend(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// putContents puts binary content to file of `path`.
func putContents(path string, data []byte, flag int, perm os.FileMode) error {
	// It supports creating file of `path` recursively.
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	// Opening file with given `flag` and `perm`.
	f, err := OpenWithFlagPerm(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write data.
	var n int
	if n, err = f.Write(data); err != nil {
		err = fmt.Errorf(`Write data to file "%s" failed: %v`, path, err)
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}
