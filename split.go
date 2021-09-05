package funcname

import "strings"

// Split splits s into two strings by the first '.' after the last '/'.
func Split(s string) (lhs, rhs string) {
	i := strings.LastIndexByte(s, '/') + 1
	j := strings.IndexByte(s[i:], '.') + 1
	if j == 0 {
		return s, ""
	}
	k := i + j
	return s[:k-1], s[k:]
}
