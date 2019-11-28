package arraystrings

func IsOneEditAway(s1, s2 string) bool {
	if abs(len(s1) - len(s2)) > 1 {
		return false
	}

	c := 0
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			c++
			if c > 1 {
				return false
			}
			if i + 1 < len(s1) && s1[i + 1] == s2[j] {
				i++
			} else if j + 1 < len(s2) && s2[j + 1] == s1[i] {
				j++
			}
		}
		i++
		j++
	}

	return true
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}