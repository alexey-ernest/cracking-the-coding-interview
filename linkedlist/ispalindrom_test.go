package linkedlist

import "testing"

func TestIsPalindromRec(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 2,
				Next: &Node {
					Value: 1,
				},
			},
		},
	}
	if !IsPalindromRecursion(l) {
		t.Errorf("it should be a palindrom!")
	}
}

func TestIsPalindromRecOdd(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 2,
				Next: &Node {
					Value: 2,
					Next: &Node {
						Value: 1,
					},
				},
			},
		},
	}
	if !IsPalindromRecursion(l) {
		t.Errorf("it should be a palindrom!")
	}
}

func TestIsPalindromRecursionEmpty(t *testing.T) {
	if IsPalindromRecursion(nil) {
		t.Errorf("it should not be a palindrom!")
	}
}

func TestIsPalindromRecSingle(t *testing.T) {
	l := &Node{
		Value: 1,
	}
	if !IsPalindromRecursion(l) {
		t.Errorf("it should be a palindrom!")
	}
}