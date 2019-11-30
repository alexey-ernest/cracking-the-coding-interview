package linkedlist

func LastNodes(k int, l *Node) *Node {
	if l == nil || k == 0{
		return nil
	}
	counter := 0
	p1 := l
	for p2 := l; p2 != nil; p2 = p2.Next {
		counter++
		if counter > k {
			p1 = p1.Next
		}
	}

	return p1
}