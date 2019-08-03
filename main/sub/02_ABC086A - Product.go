package sub

import (
	"fmt"
)

func OddOrEvenStr(a, b int) string {
	if c := a * b; c%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func Main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf(OddOrEvenStr(a, b))
}
