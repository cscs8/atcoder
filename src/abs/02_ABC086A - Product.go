package abs

import (
	"fmt"
)

// OddOrEvenStr is ...
func OddOrEvenStr(a, b int) string {
	if c := a * b; c%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// Test2 is ...
func Test2() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf(OddOrEvenStr(a, b))
}
