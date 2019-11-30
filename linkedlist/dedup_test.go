package linkedlist

import "testing"

func TestDedupHashMap(t *testing.T) {
	l := &Node{
		Value: 2,
		Next: &Node {
			Value: 2,
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
		},
	}

	l = Dedup(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value == p.Value {
			t.Errorf("Dedup failed")
			break
		}
	}
}

func TestDedupHashMapSingle(t *testing.T) {
	l := &Node{
		Value: 1,
	}

	l = Dedup(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value == p.Value {
			t.Errorf("Dedup failed")
			break
		}
	}
}

func TestDedupHashMapUnique(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 2,
		},
	}

	l = Dedup(l)
	for p := l; p.Next != nil; p = p.Next {
		if p.Next.Value == p.Value {
			t.Errorf("Dedup failed")
			break
		}
	}
}

func TestDedupHashMapEmpty(t *testing.T) {
	l := Dedup(nil)
	if l != nil {
		t.Errorf("Dedup failed")
	}
}

func TestDedupSort(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 1,
			Next: &Node {
				Value: 1,
				Next: &Node {
					Value: 3,
					Next: &Node {
						Value: 4,
						Next: &Node {
							Value: 4,
							Next: &Node {
								Value: 4,
							},
						},
					},
				},
			},
		},
	}

	l = DedupSort(l)
	prev := l
	for p := l; p != nil; p = p.Next {
		if p != prev && p.Value == prev.Value {
			t.Errorf("Dedup failed")
			break
		}
		prev = p
	}
}
