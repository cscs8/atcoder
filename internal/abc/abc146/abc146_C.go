package abc146

import (
	"fmt"
	"math"
	"strconv"
)

// TestC is ...
func TestC() {

	var a, b int
	var x int64

	max := int(math.Pow10(9))
	fmt.Scan(&a, &b, &x)

	if isValid(a, max, b, x) {
		fmt.Print(max)
		return
	}
	if !isValid(a, 1, b, x) {
		fmt.Print(0)
		return
	}

	min, mid := 0, 0
	for min <= max {
		mid = (min + max) / 2
		if isValid(a, mid, b, x) {
			min = mid + 1
			continue
		}
		max = mid - 1
	}

	fmt.Print(max)

}

func isValid(a, n, b int, x int64) bool {
	dN := len(strconv.Itoa(n))

	ret := int64(a*n + b*dN)
	return ret <= x
}
