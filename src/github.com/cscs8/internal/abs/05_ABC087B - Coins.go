package abs

import (
	"fmt"
)

// Test5 test function.
func Test5() {
	var a, b, c, x int
	// 500
	fmt.Scanln(&a)
	// 100
	fmt.Scanln(&b)
	// 50
	fmt.Scanln(&c)
	// sum
	fmt.Scanln(&x)

	aMax := x / 500
	bMax := x / 100
	cMax := x / 50

	if aMax > a {
		aMax = a
	}
	if bMax > b {
		bMax = b
	}
	if cMax > c {
		cMax = c
	}

	ret := 0

	// 500円玉for
	for i := aMax; i >= 0; i-- {
		tmp := x - (500 * i)
		if tmp == 0 {
			ret++
			continue
		}

		j := bMax
		if bCnt := tmp / 100; bMax > bCnt {
			j = bCnt
		}
		// 100円玉for
		for ; j >= 0; j-- {
			tmp = x - (500 * i) - (100 * j)
			if tmp == 0 {
				ret++
				continue
			}

			// 50円玉
			if divideC := tmp / 50; cMax >= divideC {
				ret++
			}
		}
	}
	fmt.Println(ret)
}
