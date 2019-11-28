package arraystrings

import "testing"

func TestLenCompress(t *testing.T) {
	s := "aabbbcccaaa"
	e := "a2b3c3a3"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressOne(t *testing.T) {
	s := "a"
	e := "a"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressTwo(t *testing.T) {
	s := "ab"
	e := "ab"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressSingle(t *testing.T) {
	s := "abc"
	e := "abc"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressEmpty(t *testing.T) {
	s := ""
	e := ""
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressNotShorter(t *testing.T) {
	s := "abbc"
	e := "abbc"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}

func TestLenCompressOnlyOneChar(t *testing.T) {
	s := "aaaa"
	e := "a4"
	r := LenCompress(s)
	if r != e {
		t.Errorf("%s != %s\n", r, e)
	}
}
