package linkedlist

func Partition(l *Node, k int) *Node {
	if l == nil {
		return l
	}

	var lt, gte, lp, gp *Node
	for p := l; p != nil; p = p.Next {
		if p.Value < k {
			if lp != nil {
				lp.Next = p
			} else {
				lt = p
			}
			lp = p
		} else {
			if gp != nil {
				gp.Next = p
			} else {
				gte = p
			}
			gp = p
		}
	}
	if gp != nil {
		gp.Next = nil
	}
	if lp != nil {
		lp.Next = gte
	}
	if lt != nil {
		return lt
	}
	return gte
}