package arraystrings

func RotateMatrix90(m [][]int) {
	n := len(m)
	if n <= 1 {
		return
	}

	l, r := 0, n-1
	for i := 0; i < n; i += 1 {
		for j := l; j < r; j += 1 {
			i1, j1 := i, j
			x1 := m[i1][j1]
			for k := 0; k < 4; k += 1 {
				i2, j2 := j1, n - i1 - 1
				x2 := m[i2][j2]
				m[i2][j2] = x1
				x1 = x2
				i1, j1 = i2, j2
			}
		}
		l++
		r--
		if l >= r {
			break
		}
	}
}