package abc137

import (
	"fmt"
)

// TestA is ...
func TestA() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	tmp := a + b

	if tmp2 := a - b; tmp < tmp2 {
		tmp = tmp2
	}

	if tmp2 := a * b; tmp < tmp2 {
		tmp = tmp2
	}

	fmt.Println(tmp)
}
