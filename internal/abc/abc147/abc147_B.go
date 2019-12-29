package abc147

import (
	"fmt"
)

// TestB is ...
func TestB() {
	var s string
	fmt.Scan(&s)

	sLen := len(s) / 2
	var s1, s2 string

	for idx, r := range s {
		if len(s)%2 != 0 && idx == sLen {
			continue
		}
		if idx >= sLen {
			s2 += string(r)
			continue
		}
		s1 += string(r)
	}
	var sTmp string
	for i := sLen - 1; i >= 0; i-- {
		sTmp += string(s2[i])
	}
	s2 = sTmp
	if s1 == s2 {
		fmt.Print(0)
		return
	}

	count := 0
	for i := 0; i < sLen; i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	fmt.Print(count)
}
