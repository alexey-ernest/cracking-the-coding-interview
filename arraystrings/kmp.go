package arraystrings

const R int = 26

func KMPSubstring(s, sub string) int {
	if len(s) == 0 || len(sub) == 0 {
		return -1
	}
	m := len(sub)
	dfa := make([][]int, R)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}

	x := 0
	dfa[sub[0]-'a'][0] = 1
	for j := 1; j < m; j += 1 {
		for i := 0; i < R; i += 1 {
			dfa[i][j] = dfa[i][x] // miss
		}
		c := sub[j]-'a'
		dfa[c][j] = j+1 // match
		x = dfa[c][x]
	}

	j := 0
	for i := 0; i < len(s); i += 1 {
		c := s[i] - 'a'
		j = dfa[c][j]
		if j == m {
			return i-m+1
		}
	}
	return -1
}