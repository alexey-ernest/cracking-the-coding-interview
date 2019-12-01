package linkedlist

import "testing"

func TestLoop(t *testing.T) {
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
							Next: &Node {
								Value: 7,
								Next: &Node{
									Value: 8,
								},
							},
						},
					},
				},
			},
		},
	}
	l.Next.Next.Next.Next.Next.Next.Next.Next = l.Next.Next.Next
	loop := Loop(l)
	if loop == nil || loop.Value != 4 {
		t.Errorf("loop is found incorrectly")
	}
}

func TestLoopNoLoop(t *testing.T) {
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
							Next: &Node {
								Value: 7,
								Next: &Node{
									Value: 8,
								},
							},
						},
					},
				},
			},
		},
	}

	loop := Loop(l)
	if loop != nil {
		t.Errorf("there is no loop actually")
	}
}