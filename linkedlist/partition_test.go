package linkedlist

import "testing"

func TestPartitionMiddle(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 5,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 2,
					Next: &Node {
						Value: 4,
					},
				},
			},
		},
	}

	k := 3
	l = Partition(l, k)
	less := true
	counter := 0
	for p := l; p != nil; p = p.Next {
		if p.Value >= k {
			less = false
		}
		if !less && p.Value < k {
			t.Errorf("partitioning failed")
			break
		}
		counter++
	}
	if counter != 5 {
		t.Errorf("partitioning failed")
	}
}

func TestPartitionLeft(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 5,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 2,
					Next: &Node {
						Value: 4,
					},
				},
			},
		},
	}

	k := 6
	l = Partition(l, k)
	less := true
	counter := 0
	for p := l; p != nil; p = p.Next {
		if p.Value >= k {
			less = false
		}
		if !less && p.Value < k {
			t.Errorf("partitioning failed")
			break
		}
		counter++
	}
	if counter != 5 {
		t.Errorf("partitioning failed")
	}
}

func TestPartitionRight(t *testing.T) {
	l := &Node{
		Value: 1,
		Next: &Node {
			Value: 5,
			Next: &Node {
				Value: 3,
				Next: &Node {
					Value: 2,
					Next: &Node {
						Value: 4,
					},
				},
			},
		},
	}

	k := 0
	l = Partition(l, k)
	less := true
	counter := 0
	for p := l; p != nil; p = p.Next {
		if p.Value >= k {
			less = false
		}
		if !less && p.Value < k {
			t.Errorf("partitioning failed")
			break
		}
		counter++
	}
	if counter != 5 {
		t.Errorf("partitioning failed")
	}
}