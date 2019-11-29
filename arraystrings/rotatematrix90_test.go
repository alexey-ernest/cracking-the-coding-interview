package arraystrings

import "testing"

func TestRotateMatrix90(t *testing.T) {
	n := 4
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = i*n + j
		}
	}

	RotateMatrix90(m)

	c := 0
	for j := n - 1; j >= 0; j -= 1 {
		for i := 0; i < n; i += 1 {
			if m[i][j] != c {
				t.Errorf("invalid rotated matrix: %+v\n", m)
				return
			}
			c++
		}
	}
}

func TestRotateMatrix903x3(t *testing.T) {
	n := 3
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = i*n + j
		}
	}

	RotateMatrix90(m)

	c := 0
	for j := n - 1; j >= 0; j -= 1 {
		for i := 0; i < n; i += 1 {
			if m[i][j] != c {
				t.Errorf("invalid rotated matrix: %+v\n", m)
				return
			}
			c++
		}
	}
}

func TestRotateMatrix902x2(t *testing.T) {
	n := 2
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = i*n + j
		}
	}

	RotateMatrix90(m)

	c := 0
	for j := n - 1; j >= 0; j -= 1 {
		for i := 0; i < n; i += 1 {
			if m[i][j] != c {
				t.Errorf("invalid rotated matrix: %+v\n", m)
				return
			}
			c++
		}
	}
}

func TestRotateMatrix901x1(t *testing.T) {
	n := 1
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = i*n + j
		}
	}

	RotateMatrix90(m)

	c := 0
	for j := n - 1; j >= 0; j -= 1 {
		for i := 0; i < n; i += 1 {
			if m[i][j] != c {
				t.Errorf("invalid rotated matrix: %+v\n", m)
				return
			}
			c++
		}
	}
}
