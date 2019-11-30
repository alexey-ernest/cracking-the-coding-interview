package arraystrings

func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	return isSubstring(s1+s1, s2)
}

// Boyer-Moore
func isSubstring(s, sub string) bool {
	ind := make([]int, 26)
	for i := 0; i < len(sub); i += 1 {
		ind[sub[i]-'a'] = i + 1
	}
	skip := 0
	n := len(s)
	m := len(sub)
	for i := 0; i <= n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j -= 1 {
			if sub[j] != s[i+j] {
				skip = j - ind[s[i+j]-'a'] + 1
				if skip < 0 {
					// avoiding back skipping
					skip = 1
				}
				break
			}
		}
		if skip == 0 {
			return true
		}
	}

	return false
}