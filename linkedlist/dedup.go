package linkedlist

func Dedup(l *Node) *Node {
	if l == nil {
		return l
	}
	m := make(map[int]bool)
	var prev *Node
	for p := l; p != nil; p = p.Next {
		if m[p.Value] {
			continue
		}
		m[p.Value] = true
		if prev != nil {
			prev.Next = p
		}
		prev = p
	}
	prev.Next = nil
	return l
}

func DedupSort(l *Node) *Node {
	if l == nil {
		return l
	}
	l = Sort(l)
	prev := l
	for p := l; p != nil; p = p.Next {
		if p.Value != prev.Value {
			prev.Next = p
			prev = p
		}
	}
	prev.Next = nil
	return l
}