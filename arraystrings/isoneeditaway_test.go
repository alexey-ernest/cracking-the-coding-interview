package arraystrings

import "testing"

func TestIsOneEditAwayMid(t *testing.T) {
	s1 := "pale"
	s2 := "ple"
	if !IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}

func TestIsOneEditAwayEnd(t *testing.T) {
	s1 := "pale"
	s2 := "pal"
	if !IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}

func TestIsOneEditAwayBeg(t *testing.T) {
	s1 := "ale"
	s2 := "pale"
	if !IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}

func TestIsOneEditAwayEdit(t *testing.T) {
	s1 := "bale"
	s2 := "pale"
	if !IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}

func TestIsOneEditAwayEmpty(t *testing.T) {
	s1 := ""
	s2 := "pale"
	if IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}

func TestIsOneEditAway2LenDIff(t *testing.T) {
	s1 := "pal"
	s2 := "palee"
	if IsOneEditAway(s1, s2) {
		t.Errorf("s1=%s s2=%s\n", s1, s2)
	}
}
