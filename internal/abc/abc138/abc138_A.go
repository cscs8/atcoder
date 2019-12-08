package abc138

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var a int
	var s string
	fmt.Scan(&a, &s)
	if a >= 3200 {
		fmt.Print(s)
		return
	}

	fmt.Print("red")
}
