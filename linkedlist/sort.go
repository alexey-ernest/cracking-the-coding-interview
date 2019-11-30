package linkedlist

func Sort(l *Node) *Node {
	if l == nil || l.Next == nil {
		return l
	}

	p1 := l
	c := 0
	for p2 := p1; p2.Next != nil; p2 = p2.Next {
		c++
		if c % 2 == 0 {
			p1 = p1.Next
		}
	}
	p2 := p1.Next
	p1.Next = nil
	p1 = Sort(p1)
	p2 = Sort(p2)
	return merge(p1, p2)
}

func merge(l1 *Node, l2 *Node) *Node {
	var head *Node
	var prev *Node
	for l1 != nil && l2 != nil {
		var n *Node
		if l2.Value < l1.Value {
			n = l2
			l2 = l2.Next
		} else {
			n = l1
			l1 = l1.Next
		}
		if head == nil {
			head = n
		}
		if prev != nil {
			prev.Next = n
		}

		prev = n
	}

	tail := l1
	if l2 != nil {
		tail = l2
	}
	if head == nil {
		head = tail
	} else {
		prev.Next = tail
	}
	
	return head
}