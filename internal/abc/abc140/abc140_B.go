package abc140

import (
	"fmt"
)

// TestB is...
func TestB() {
	// IO start
	var n int
	fmt.Scan(&n)

	var listA = make([]int, n)
	var listB = make([]int, n)
	var listC = make([]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Scan(&listA[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&listB[i])
	}
	for i := 0; i < n-1; i++ {
		fmt.Scan(&listC[i])
	}

	ret := 0
	// process start
	ret += listB[listA[0]-1]
	for i := 1; i < n; i++ {
		value := listA[i] - 1
		ret += listB[value]
		if previous := listA[i-1]; previous == value {
			ret += listC[value-1]
		}
	}
	fmt.Print(ret)
}
