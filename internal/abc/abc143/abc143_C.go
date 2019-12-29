package abc143

import (
	"fmt"
	"strings"
)

// TestC is...
func TestC() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)
	var listB = make([]string, n)
	listB = strings.Split(s, "")

	ret := 1
	for i := 0; i < len(listB)-1; i++ {
		if listB[i] != listB[i+1] {
			ret++
		}
	}

	fmt.Println(ret)

}
