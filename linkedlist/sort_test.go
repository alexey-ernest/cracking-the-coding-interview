package linkedlist

import "testing"

func TestSort(t *testing.T) {
	l := &Node{
		Value: 3,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 5,
				Next: &Node {
					Value: 6,
					Next: &Node {
						Value: 1,
						Next: &Node {
							Value: 4,
						},
					},
				},
			},
		},
	}

	l = Sort(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value < p.Value {
			t.Errorf("Sort failed")
			break
		}
	}
}

func TestSortInverse(t *testing.T) {
	l := &Node{
		Value: 6,
		Next: &Node {
			Value: 5,
			Next: &Node {
				Value: 4,
				Next: &Node {
					Value: 3,
					Next: &Node {
						Value: 2,
						Next: &Node {
							Value: 1,
						},
					},
				},
			},
		},
	}

	l = Sort(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value < p.Value {
			t.Errorf("Sort failed")
			break
		}
	}
}

func TestSortOrdered(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 5,
						Next: &Node {
							Value: 6,
						},
					},
				},
			},
		},
	}

	l = Sort(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value < p.Value {
			t.Errorf("Sort failed")
			break
		}
	}
}

func TestSortSingle(t *testing.T) {
	l := &Node{
		Value: 1,
	}

	l = Sort(l)
	if l == nil {
		t.Errorf("Sort failed")
	}
}

func TestSortTwo(t *testing.T) {
	l := &Node{
		Value: 2,
		Next: &Node {
			Value: 1,
		},
	}

	l = Sort(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value < p.Value {
			t.Errorf("Sort failed")
			break
		}
	}
}