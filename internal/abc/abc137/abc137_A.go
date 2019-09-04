package abc137

import (
	"fmt"
	"sort"
)

// TestA is ...
func TestA() {
	var a, b int
	fmt.Scan(&a, &b)

	list := []int{a + b, a - b, a * b}

	sort.Ints(list)

	fmt.Print(list[len(list)-1])
}
