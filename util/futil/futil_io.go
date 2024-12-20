package futil

import (
	"io"

	"github.com/ZYallers/fine/internal/util/utils"
)

// NewReadCloser creates and returns a RepeatReadCloser object.
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return utils.NewReadCloser(content, repeatable)
}
