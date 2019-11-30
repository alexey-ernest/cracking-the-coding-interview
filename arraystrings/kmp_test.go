package arraystrings

import "testing"

func TestKMPSubstring(t *testing.T) {
	s := "ababace"
	sub := "abac"
	res := KMPSubstring(s, sub)
	exp := 2
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestKMPSubstringMiss(t *testing.T) {
	s := "ababace"
	sub := "abad"
	res := KMPSubstring(s, sub)
	exp := -1
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestKMPSubstringOneChar(t *testing.T) {
	s := "ababace"
	sub := "c"
	res := KMPSubstring(s, sub)
	exp := 5
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}