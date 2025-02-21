package fhtml_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fhtml"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_StripTags(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
		dst := `Test paragraph.  Other text`
		t.Assert(fhtml.StripTags(src), dst)
	})
}

func Test_Entities(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(fhtml.Entities(src), dst)
		t.Assert(fhtml.EntitiesDecode(dst), src)
	})
}

func Test_SpecialChars(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(fhtml.SpecialChars(src), dst)
		t.Assert(fhtml.SpecialCharsDecode(dst), src)
	})
}

func Test_SpecialCharsMapOrStruct_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		a := f.Map{
			"Title":   "<h1>T</h1>",
			"Content": "<div>C</div>",
		}
		err := fhtml.SpecialCharsMapOrStruct(a)
		t.AssertNil(err)
		t.Assert(a["Title"], `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a["Content"], `&lt;div&gt;C&lt;/div&gt;`)
	})
	ftest.C(t, func(t *ftest.T) {
		a := f.MapStrStr{
			"Title":   "<h1>T</h1>",
			"Content": "<div>C</div>",
		}
		err := fhtml.SpecialCharsMapOrStruct(a)
		t.AssertNil(err)
		t.Assert(a["Title"], `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a["Content"], `&lt;div&gt;C&lt;/div&gt;`)
	})
}

func Test_SpecialCharsMapOrStruct_Struct(t *testing.T) {
	type A struct {
		Title   string
		Content string
	}
	ftest.C(t, func(t *ftest.T) {
		a := &A{
			Title:   "<h1>T</h1>",
			Content: "<div>C</div>",
		}
		err := fhtml.SpecialCharsMapOrStruct(a)
		t.AssertNil(err)
		t.Assert(a.Title, `&lt;h1&gt;T&lt;/h1&gt;`)
		t.Assert(a.Content, `&lt;div&gt;C&lt;/div&gt;`)
	})

	// should error
	ftest.C(t, func(t *ftest.T) {
		a := A{
			Title:   "<h1>T</h1>",
			Content: "<div>C</div>",
		}
		err := fhtml.SpecialCharsMapOrStruct(a)
		t.AssertNE(err, nil)
	})
}
