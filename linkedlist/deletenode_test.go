package linkedlist

import "testing"

func TestDeleteNodeMiddle(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 2,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 5,
					},
				},
			},
		},
	}

	DeleteNode(l.Next.Next)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value == p.Value {
			t.Errorf("Deletion failed")
			break
		}
	}
}

func TestDeleteNodePrevLast(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 4,
					Next: &Node {
						Value: 4,
					},
				},
			},
		},
	}

	DeleteNode(l.Next.Next.Next)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value == p.Value {
			t.Errorf("Deletion failed")
			break
		}
	}
}