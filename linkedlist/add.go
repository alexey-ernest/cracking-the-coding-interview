package linkedlist

func AddReversed(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	overflow := 0
	res := l1
	var prev *Node
	for l1 != nil && l2 != nil {
		l1.Value += l2.Value + overflow
		if l1.Value >= 10 {
			overflow = 1
			l1.Value = l1.Value % 10
		} else {
			overflow = 0
		}

		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	tail := l1
	if l2 != nil {
		tail = l2
	}

	if tail != nil {
		// append tail
		tail.Value += overflow
		prev.Next = tail
	} else if overflow > 0 {
		// handle overflow at the end
		prev.Next = &Node{
			Value: overflow,
		}
	} else {
		// trim links at the end as they could be old links from original numbers
		prev.Next = nil
	}

	return res
}

func Add(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	len1 := length(l1)
	len2 := length(l2)
	if len1 < len2 {
		l1 = padd(l1, len2-len1)
	} else if len2 < len1 {
		l2 = padd(l2, len1-len2)
	}

	res, ovf := addEven(l1, l2)
	if ovf > 0 {
		n := &Node {
			Value: ovf,
			Next: res,
		}
		res = n
	}
	return res
}

func addEven(l1, l2 *Node) (*Node, int) {
	if l1 == nil {
		return nil, 0
	}
	res, ovf := addEven(l1.Next, l2.Next)
	n := &Node {
		Value: l1.Value + l2.Value + ovf,
		Next: res,
	}
	if n.Value < 10 {
		return n, 0
	}
	n.Value = n.Value % 10
	return n, 1
}

func length(l *Node) int {
	counter := 0
	for p := l; p != nil; p = p.Next {
		counter += 1
	}
	return counter
}

func padd(l *Node, k int) *Node {
	for i := 0; i < k; i += 1 {
		n := &Node {
			Value: 0,
			Next: l,
		}
		l = n
	}
	return l
}
