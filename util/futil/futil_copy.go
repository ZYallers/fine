package futil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ZYallers/fine/internal/deepcopy"
)

var (
	// bufferPool is a pool of bytes.Buffer objects.
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 4096))
		},
	}
)

// Copy returns a deep copy of v.
//
// Copy is unable to copy unexported fields in a struct (lowercase field names).
// Unexported fields can't be reflected by the Go runtime and therefore
// they can't perform any data copies.
func Copy(src interface{}) (dst interface{}) {
	return deepcopy.Copy(src)
}

// CopyBytes reads all bytes from r and returns them as a byte slice.
func CopyBytes(r io.Reader) ([]byte, error) {
	dst := bufferPool.Get().(*bytes.Buffer)
	dst.Reset()
	defer func() {
		if dst != nil {
			bufferPool.Put(dst)
			dst = nil
		}
	}()

	if _, err := io.Copy(dst, r); err != nil {
		return nil, err
	}
	bodyB := dst.Bytes()

	bufferPool.Put(dst)
	dst = nil

	return bodyB, nil
}

// DrainBody reads all bytes from src and returns two bytes reader.
func DrainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
