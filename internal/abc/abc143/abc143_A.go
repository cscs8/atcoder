package abc143

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := a - (b * 2)
	if ans < 0 {
		ans = 0
	}
	fmt.Print(ans)
}
