package arraystrings

import "testing"

func TestIsRotation(t *testing.T) {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	if !IsRotation(s1, s2) {
		t.Errorf("%s is rotation of %s", s1, s2)
	}
}