package linkedlist

func DeleteNode(x *Node) {
	if x == nil {
		return
	}
	prev := x
	for x.Next != nil {
		x.Value = x.Next.Value
		prev = x
		x = x.Next
	}
	prev.Next = nil
}