package abc147

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var a1, a2, a3 int
	fmt.Scan(&a1, &a2, &a3)

	if a1+a2+a3 >= 22 {

		fmt.Print("bust")
		return
	}
	fmt.Print("win")
}
