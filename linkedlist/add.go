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