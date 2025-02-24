package ffile_test

import (
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_Home(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// user's home directory
		homePath, _ := ffile.Home()
		t.AssertNE(homePath, "")
	})
}
