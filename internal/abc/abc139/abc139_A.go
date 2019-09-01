package abc139

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var s, t string
	fmt.Scan(&s, &t)

	ret := 0
	for i := 0; i < 3; i++ {
		if s[i] == t[i] {
			ret++
		}

	}

	fmt.Print(ret)
}
