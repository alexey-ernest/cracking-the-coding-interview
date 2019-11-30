package linkedlist

import "testing"

func TestLastNodes(t *testing.T) {
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
					},
				},
			},
		},
	}

	l = LastNodes(3, l)
	exp := 3
	if l.Value != exp {
		t.Errorf("3d node from tail should be %d, got %d", exp, l.Value)
	}
}

func TestLastNodesFullLength(t *testing.T) {
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
					},
				},
			},
		},
	}

	l = LastNodes(5, l)
	exp := 1
	if l.Value != exp {
		t.Errorf("3d node from tail should be %d, got %d", exp, l.Value)
	}
}

func TestLastNodesSingle(t *testing.T) {
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
					},
				},
			},
		},
	}

	l = LastNodes(1, l)
	exp := 5
	if l.Value != exp {
		t.Errorf("3d node from tail should be %d, got %d", exp, l.Value)
	}
}

func TestLastNodesZero(t *testing.T) {
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
					},
				},
			},
		},
	}

	l = LastNodes(0, l)
	if l != nil {
		t.Errorf("should be 0")
	}
}