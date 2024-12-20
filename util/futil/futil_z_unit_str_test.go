package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_IsLetterUpper(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsLetterUpper('a'), false)
		t.Assert(futil.IsLetterUpper('A'), true)
		t.Assert(futil.IsLetterUpper('1'), false)
	})
}

func Test_IsLetterLower(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsLetterLower('a'), true)
		t.Assert(futil.IsLetterLower('A'), false)
		t.Assert(futil.IsLetterLower('1'), false)
	})
}

func Test_IsLetter(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsLetter('a'), true)
		t.Assert(futil.IsLetter('A'), true)
		t.Assert(futil.IsLetter('1'), false)
	})
}

func Test_IsNumeric(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsNumeric("1a我"), false)
		t.Assert(futil.IsNumeric("0123"), true)
		t.Assert(futil.IsNumeric("我是中国人"), false)
		t.Assert(futil.IsNumeric("1.2.3.4"), false)
	})
}

func Test_UcFirst(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "AbcdEFG乱入的中文abcdefg"
		t.Assert(futil.UcFirst(""), "")
		t.Assert(futil.UcFirst(s1), e1)
		t.Assert(futil.UcFirst(e1), e1)
	})
}

func Test_ReplaceByMap(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(futil.ReplaceByMap(s1, f.MapStrStr{
			"a": "A",
			"G": "g",
		}), "AbcdEFg乱入的中文Abcdefg")
	})
}

func Test_RemoveSymbols(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`), `abac101`)
		t.Assert(futil.RemoveSymbols(`-a-b我._a c1!@#$%^&*是()_+:帅";'.,哥'01`), `ab我ac1是帅哥01`)
	})
}

func Test_EqualFoldWithoutChars(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.EqualFoldWithoutChars("a", "A"), true)
		t.Assert(futil.EqualFoldWithoutChars("a", "a-"), true)
		t.Assert(futil.EqualFoldWithoutChars("a.", "a-"), true)
	})
}

func Test_SplitAndTrim(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := `

010    

020  

`
		a := futil.SplitAndTrim(s, "\n", "0")
		t.Assert(len(a), 2)
		t.Assert(a[0], "1")
		t.Assert(a[1], "2")
	})
}

func Test_Trim(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.Trim(" 123456\n "), "123456")
		t.Assert(futil.Trim("#123456#;", "#;"), "123456")
	})
}

func Test_FormatCmdKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.FormatCmdKey("a-b_c"), "a-b.c")
		t.Assert(futil.FormatCmdKey("a-b_c-d"), "a-b.c-d")
		t.Assert(futil.FormatCmdKey("a-b_c-d-e"), "a-b.c-d-e")
	})
}

func Test_FormatEnvKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.FormatEnvKey("a.b"), "A_B")
		t.Assert(futil.FormatEnvKey("a.b.c"), "A_B_C")
		t.Assert(futil.FormatEnvKey("a.b-c.d"), "A_B-C_D")
	})
}

func Test_StripSlashes(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.StripSlashes(`\\a\\b\\c\\d`), `\a\b\c\d`)
	})
}
