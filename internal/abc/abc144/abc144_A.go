package abc144

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var a, b int
	fmt.Scan(&a, &b)

	isValid := isValid(a) && isValid(b)
	if isValid {
		fmt.Print(a * b)
		return

	}

	fmt.Print(-1)
}

func isValid(a int) bool {
	return 0 <= a && a < 10
}
