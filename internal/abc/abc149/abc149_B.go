package abc149

import (
	"fmt"
)

// TestB is ...
func TestB() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	fmt.Print(calc(a, b, k))
}

func calc(a, b, k int) (int, int) {
	if 0 < a {
		if k < a {
			a -= k
			return a, b
		}
		k -= a
		a = 0
	}

	if 0 < b {
		if b < k {
			return a, 0
		}
		return a, b - k
	}
	return a, b
}
