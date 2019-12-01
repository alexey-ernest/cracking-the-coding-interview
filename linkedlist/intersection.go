package linkedlist

func Intersection(l1, l2 *Node) *Node {
	if l1 == nil || l2 == nil {
		return nil
	}
	
	// save length of the second list
	len2 := length(l2)

	var tail1 *Node
	for l1 != nil {
		tail1 = l1
		next := l1.Next
		l1.Next = nil
		l1 = next
	}

	newlen2 := 1
	for l2.Next != nil {
		l2 = l2.Next
		newlen2++
	}

	if newlen2 < len2 {
		// if the length changed then we removed some nodes from the l2 then we have an intersection
		return l2
	}
	if newlen2 == len2 && l2 == tail1 {
		// if the length didn't change but both tails are the same, then we have an intersection
		// as the last node
		return l2
	}

	return nil
}
