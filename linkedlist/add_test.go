package linkedlist

import "testing"

func TestAddReversedEven(t *testing.T) {
	l1 := &Node{
		Value: 7,
		Next: &Node {
			Value: 1,
			Next: &Node {
				Value: 6,
			},
		},
	}
	l2 := &Node{
		Value: 5,
		Next: &Node {
			Value: 9,
			Next: &Node {
				Value: 2,
			},
		},
	}
	res := AddReversed(l1, l2)
	exp := [...]int{2, 1, 9}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
}

func TestAddReversedOnlyOne(t *testing.T) {
	l1 := &Node{
		Value: 7,
		Next: &Node {
			Value: 1,
			Next: &Node {
				Value: 6,
			},
		},
	}
	res := AddReversed(l1, nil)
	exp := [...]int{7, 1, 6}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
}

func TestAddReversedSingleDigits(t *testing.T) {
	l1 := &Node{
		Value: 7,
	}
	l2 := &Node{
		Value: 2,
	}
	res := AddReversed(l1, l2)
	exp := [...]int{9}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
}

func TestAddReversedSingleDigitsOverflow(t *testing.T) {
	l1 := &Node{
		Value: 7,
	}
	l2 := &Node{
		Value: 3,
	}
	res := AddReversed(l1, l2)
	if res == nil {
		t.Errorf("addition failed")
		return
	}

	exp := [...]int{0, 1}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
	if i < len(exp) {
		t.Errorf("addition failed")
	}
}

func TestAddReversedNotEven(t *testing.T) {
	l1 := &Node{
		Value: 7,
		Next: &Node {
			Value: 1,
			Next: &Node {
				Value: 6,
			},
		},
	}
	l2 := &Node{
		Value: 5,
		Next: &Node {
			Value: 9,
			Next: &Node {
				Value: 2,
				Next: &Node {
					Value: 3,
				},
			},
		},
	}
	res := AddReversed(l1, l2)
	exp := [...]int{2, 1, 9, 3}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
	if i < len(exp) {
		t.Errorf("addition failed")
	}
}

func TestAddReversedNotEvenWithOverflow(t *testing.T) {
	l1 := &Node{
		Value: 7,
		Next: &Node {
			Value: 1,
			Next: &Node {
				Value: 7,
				Next: &Node {
					Value: 5,
					Next: &Node {
						Value: 6,
					},
				},
			},
		},
	}
	l2 := &Node{
		Value: 5,
		Next: &Node {
			Value: 9,
			Next: &Node {
				Value: 2,
			},
		},
	}
	res := AddReversed(l1, l2)
	exp := [...]int{2, 1, 0, 6, 6}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
	if i < len(exp) {
		t.Errorf("addition failed")
	}
}

func TestAddForwardOverflow(t *testing.T) {
	l1 := &Node{
		Value: 6,
		Next: &Node{
			Value: 1,
			Next: &Node{
				Value: 7,
			},
		},
	}
	l2 := &Node{
		Value: 2,
		Next: &Node{
			Value: 9,
			Next: &Node{
				Value: 5,
				Next: &Node{
					Value: 3,
				},
			},
		},
	}
	res := Add(l1, l2)
	if res == nil {
		t.Errorf("addition failed")
		return
	}

	exp := [...]int{3, 5, 7, 0}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
	if i < len(exp) {
		t.Errorf("addition failed")
	}
}

func TestAddForwardEven(t *testing.T) {
	l1 := &Node{
		Value: 6,
		Next: &Node{
			Value: 1,
			Next: &Node{
				Value: 7,
			},
		},
	}
	l2 := &Node{
		Value: 2,
		Next: &Node{
			Value: 9,
			Next: &Node{
				Value: 5,
			},
		},
	}
	res := Add(l1, l2)
	if res == nil {
		t.Errorf("addition failed")
		return
	}

	exp := [...]int{9, 1, 2}
	i := 0
	for p := res; p != nil; p = p.Next {
		if p.Value != exp[i] {
			t.Errorf("%d != %d", p.Value, exp[i])
			break
		}
		i++
	}
	if i < len(exp) {
		t.Errorf("addition failed")
	}
}
