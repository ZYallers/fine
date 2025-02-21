package furl_test

import (
	"net/url"
	"testing"

	"github.com/ZYallers/fine/encoding/furl"
	"github.com/ZYallers/fine/test/ftest"
)

var (
	urlStr       = `https://golang.org/x/crypto?go-get=1 +`
	urlEncode    = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1+%2B`
	rawUrlEncode = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1%20%2B`
)

func TestEncodeAndDecode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(furl.Encode(urlStr), urlEncode)

		res, err := furl.Decode(urlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestRowEncodeAndDecode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(furl.RawEncode(urlStr), rawUrlEncode)

		res, err := furl.RawDecode(rawUrlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestBuildQuery(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := url.Values{
			"a": {"a2", "a1"},
			"b": {"b2", "b1"},
			"c": {"c1", "c2"},
		}
		expect := "a=a2&a=a1&b=b2&b=b1&c=c1&c=c2"
		t.Assert(furl.BuildQuery(src), expect)
	})
}

func TestParseURL(t *testing.T) {
	src := `http://username:password@hostname:9090/path?arg=value#anchor`
	expect := map[string]string{
		"scheme":   "http",
		"host":     "hostname",
		"port":     "9090",
		"user":     "username",
		"pass":     "password",
		"path":     "/path",
		"query":    "arg=value",
		"fragment": "anchor",
	}

	ftest.C(t, func(t *ftest.T) {
		component := 0
		for k, v := range []string{"all", "scheme", "host", "port", "user", "pass", "path", "query", "fragment"} {
			if v == "all" {
				component = -1
			} else {
				component = 1 << (uint(k - 1))
			}

			res, err := furl.ParseURL(src, component)
			if err != nil {
				t.Errorf("ParseURL failed. component:%v, err:%v", component, err)
				return
			}

			if v == "all" {
				t.Assert(res, expect)
			} else {
				t.Assert(res[v], expect[v])
			}

		}
	})
}
