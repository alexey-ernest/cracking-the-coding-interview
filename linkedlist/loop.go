package linkedlist

func Loop(l *Node) *Node {
	if l == nil {
		return nil
	}
	p1 := l
	c := 0
	for p2 := l.Next; p2 != p1 && p2 != nil; p2 = p2.Next {
		c++
		if c % 2 == 0 {
			p1 = p1.Next
		}
		if p2.Next == nil {
			return nil // no cycles
		}
	}

	// now p1 == p2, breaking links in cycle
	for p1 != nil {
		next := p1.Next
		p1.Next = nil
		p1 = next
	}
	for l.Next != nil {
		l = l.Next
	}
	return l
}