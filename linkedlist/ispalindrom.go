package linkedlist

func IsPalindromRecursion(l *Node) bool {
	if l == nil {
		return false
	}
	res, _ := ispalindrom(l, l)
	return res
}

func ispalindrom(l *Node, start *Node) (bool, *Node) {
	if l == nil {
		return true, start
	}
	res, s1 := ispalindrom(l.Next, start)
	if !res {
		return false, nil
	}
	if s1 == nil {
		return true, nil
	}
	if s1.Value != l.Value {
		return false, nil
	}
	if s1 == l || s1.Next == l {
		return true, nil
	}
	return true, s1.Next
}