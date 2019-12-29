package abc149

import (
	"fmt"
)

// TestC is ...
func TestC() {
	var x int
	fmt.Scan(&x)

	var jTmp int
	for i := x; ; i++ {
		for j := 2; j <= i; j++ {
			jTmp = j
			if i%j == 0 {
				break
			}
		}
		if i == jTmp {
			fmt.Print(i)
			return

		}
	}

}
