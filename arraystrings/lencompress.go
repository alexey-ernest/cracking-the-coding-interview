package arraystrings

import (
	"fmt"
	"strings"
)

func LenCompress(s string) string {
	if len(s) <= 1 {
		return s
	}

	count := 1
	chr := s[0]
	var sb strings.Builder
	for i := 1; i < len(s); i += 1 {
		if s[i] != chr {
			sb.WriteString(fmt.Sprintf("%s%d", string(chr), count))
			count = 1
			chr = s[i]
		} else {
			count += 1
		}
	}
	sb.WriteString(fmt.Sprintf("%s%d", string(chr), count))
	res := sb.String()
	if len(res) >= len(s) {
		res = s
	}
	return res
}