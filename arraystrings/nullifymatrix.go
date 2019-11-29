package arraystrings

func NullifyMatrix(m [][]int) {
	N := len(m)
	if N == 0 {
		return
	}
	M := len(m[0])
	if M == 0 {
		return
	}
	for i := 0; i < N; i += 1 {
		for j := 0; j < M; j += 1 {
			if m[i][j] == 0 {
				m[0][j] = 0
				m[i][0] = 0
			} else if m[0][j] == 0 || m[i][0] == 0 {
				m[i][j] = 0
			}
		}
	}

	for i := 0; i < N; i += 1 {
		if m[i][0] == 0 {
			nullifyRow(m, i)
		}
	}

	for j := 0; j < M; j += 1 {
		if m[0][j] == 0 {
			nullifyCol(m, j)
		}
	}
}

func nullifyRow(m [][]int, i int) {
	for j := 1; j < len(m[0]); j += 1 {
		if m[i][j] == 0 {
			return
		}
		m[i][j] = 0
	}
}

func nullifyCol(m [][]int, j int) {
	for i := 1; i < len(m); i += 1 {
		if m[i][j] == 0 {
			return
		}
		m[i][j] = 0
	}
}