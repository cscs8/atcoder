package abc139

import (
	"fmt"
)

// TestB is...
func TestB() {
	var a, b int
	fmt.Scan(&a, &b)

	cnt := b / a

	for b > (a*cnt)-(cnt-1) {
		cnt++
	}
	fmt.Print(cnt)

}
