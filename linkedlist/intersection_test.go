package linkedlist

import "testing"

func TestIntersection(t *testing.T) {
	l1 := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 5,
					},
				},
			},
		},
	}

	l2 := &Node{
		Value: 6,
		Next: &Node {
			Value: 7,
			Next: l1.Next.Next.Next,
		},
	}

	res := Intersection(l1, l2)
	if res.Value != 4 {
		t.Errorf("invalid intersection found %d", res.Value)
	}
}

func TestIntersectionNoIntersection(t *testing.T) {
	l1 := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 5,
					},
				},
			},
		},
	}

	l2 := &Node{
		Value: 6,
		Next: &Node {
			Value: 7,
		},
	}

	res := Intersection(l1, l2)
	if res != nil {
		t.Errorf("there is no intersection in reality")
	}
}

func TestIntersectionByTheLastNode(t *testing.T) {
	l1 := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 5,
					},
				},
			},
		},
	}

	l2 := &Node{
		Value: 6,
		Next: &Node {
			Value: 7,
			Next: l1.Next.Next.Next.Next,
		},
	}

	res := Intersection(l1, l2)
	if res == nil || res.Value != 5 {
		t.Errorf("no intersection or invalid intersection found")
	}
}