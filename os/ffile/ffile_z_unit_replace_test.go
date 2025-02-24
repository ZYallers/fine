package ffile_test

import (
	"regexp"
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_ReplaceFile(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// init
		var (
			fileName = "ffile_replace.txt"
			tempDir  = ffile.Temp("ffile_replace_unit_test")
			tempFile = ffile.Join(tempDir, fileName)
		)

		// write contents
		content := "fine unit test content"
		err := ffile.PutContents(tempFile, content)
		t.AssertNil(err)

		// read contents
		t.AssertEQ(ffile.GetContents(tempFile), content)

		// It replaces content directly by file path
		err = ffile.ReplaceFile("content", "replace word", tempFile)
		t.AssertNil(err)
		t.AssertEQ(ffile.GetContents(tempFile), "fine unit test replace word")
	})
}

func Test_ReplaceFileFunc(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// init
		var (
			fileName = "ffile_replace.txt"
			tempDir  = ffile.Temp("ffile_replace_unit_test")
			tempFile = ffile.Join(tempDir, fileName)
		)

		// write contents
		content := "fine unit test 123"
		err := ffile.PutContents(tempFile, content)
		t.AssertNil(err)

		// read contents
		t.AssertEQ(ffile.GetContents(tempFile), content)

		// It replaces content directly by file path and callback function.
		err = ffile.ReplaceFileFunc(func(path, content string) string {
			// Replace with regular match
			reg, _ := regexp.Compile(`\d{3}`)
			return reg.ReplaceAllString(content, "???")
		}, tempFile)
		t.AssertNil(err)
		t.AssertEQ(ffile.GetContents(tempFile), "fine unit test ???")
	})
}

func Test_ReplaceDir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// init
		var (
			fileName = "ffile_replace.txt"
			tempDir  = ffile.Temp("ffile_replace_unit")
			tempFile = ffile.Join(tempDir, fileName)
		)

		// write contents
		err := ffile.PutContents(tempFile, "fine replace content")
		t.AssertNil(err)

		// read contents
		t.AssertEQ(ffile.GetContents(tempFile), "fine replace content")

		// It replaces content of all files under specified directory recursively.
		err = ffile.ReplaceDir("content", "word", tempDir, fileName, true)
		t.AssertNil(err)
		t.AssertEQ(ffile.GetContents(tempFile), "fine replace word")
	})
}

func Test_ReplaceDirFunc(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// init
		var (
			fileName = "ffile_replace.txt"
			tempDir  = ffile.Temp("ffile_replace_unit")
			tempFile = ffile.Join(tempDir, fileName)
		)

		// write contents
		err := ffile.PutContents(tempFile, "fine replace 123")
		t.AssertNil(err)

		// read contents
		t.AssertEQ(ffile.GetContents(tempFile), "fine replace 123")

		// It replaces content of all files under specified directory with custom callback function recursively.
		err = ffile.ReplaceDirFunc(func(path, content string) string {
			// Replace with regular match
			reg, _ := regexp.Compile(`\d{3}`)
			return reg.ReplaceAllString(content, "???")
		}, tempDir, fileName, true)

		t.AssertNil(err)
		t.AssertEQ(ffile.GetContents(tempFile), "fine replace ???")
	})
}
