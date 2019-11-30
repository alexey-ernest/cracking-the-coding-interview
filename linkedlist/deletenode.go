package linkedlist

func DeleteNode(x *Node) {
	if x == nil || x.Next == nil {
		return
	}
	x.Value = x.Next.Value
	x.Next = x.Next.Next
}