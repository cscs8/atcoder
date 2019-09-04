package abc137

import (
	"fmt"
)

// TestB is...
func TestB() {
	var k, x int
	fmt.Scanln(&k, &x)

	for i := x - k + 1; i < x+k; i++ {
		fmt.Printf("%d ", i)
	}

}
