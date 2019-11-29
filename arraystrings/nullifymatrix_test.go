package arraystrings

import "testing"

func TestNullifyMatrix(t *testing.T) {
	m := [][]int{
		[]int{1,2,0,4},
		[]int{0,2,3,4},
		[]int{1,2,3,0},
		[]int{0,0,3,4},
	}

	NullifyMatrix(m)

	for i := 0; i < 4; i += 1 {
		for j := 0; j < 4; j += 1 {
			if m[i][j] != 0 {
				t.Errorf("the matrix should has only 0s: %+v\n", m)
				return
			}
		}
	}
}