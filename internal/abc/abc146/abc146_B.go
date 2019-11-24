package abc146

import (
	"fmt"
)

// TestB is ...
func TestB() {
	const (
		aCode  = 65
		zCode  = 90
		aZDiff = 26
	)
	var n rune // int32
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	for _, r := range s {
		if r+n > 'Z' {
			fmt.Printf("%c", r+n-aZDiff)
			continue

		}

		// rune
		fmt.Printf("%c", r+n)
	}
}
