package abs

import (
	"fmt"
)

// Test5 test function.
func Test5() {
	var a, b, c, x int
	// // 500
	// fmt.Scanln(&a)
	// // 100
	// fmt.Scanln(&b)
	// // 50
	// fmt.Scanln(&c)
	// // sum
	// fmt.Scanln(&x)
	// a = 2
	// b = 2
	// c = 2
	// x = 100
	// a = 30
	// b = 40
	// c = 50
	// x = 6000
	a = 5
	b = 1
	c = 0
	x = 150

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
	cStart := 0
	if mod := x % 100; mod != 0 {
		cStart = 1
	}
	// if aMax != 0 {
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
		for ; j >= 0; j-- {
			tmp = x - (500 * i) - (100 * j)
			if tmp == 0 {
				ret++
				continue
			}

			cCnt := cMax
			if cTmp := tmp / 50; cMax > cCnt {
				cCnt = cTmp
			}
			for k := cStart; cCnt >= k; k += 2 {
				tmp = x - (500 * i) - (100 * j) - (50 * k)
				if tmp == 0 {
					ret++
					break
				}
			}
		}
	}
	// }

	// // 100円玉for
	// for j := bMax; j >= 0; j-- {
	// 	tmp := x - (100 * j)
	// 	if tmp == 0 {
	// 		ret++
	// 		continue
	// 	}

	// 	cCnt := cMax
	// 	if cTmp := tmp / 50; cMax > cCnt {
	// 		cCnt = cTmp
	// 	}
	// 	for k := cStart; cCnt >= k; k += 2 {
	// 		tmp = x - (100 * j) - (50 * k)
	// 		if tmp == 0 {
	// 			ret++
	// 		}
	// 	}
	// }

	fmt.Println(ret)
	// ret := 0
	// if ret = 50 / x; 1 >= ret {
	// 	if c == 0 {
	// 		fmt.Println(0)
	// 		return

	// 	}
	// 	fmt.Println(1)
	// 	return
	// }

	// if ret = 100 / x; 1 >= ret {

	// }
	// if 0 == c {
	// 	if 100 > x {
	// 		fmt.Println(0)
	// 		return
	// 	}

	// 	if 0 == b {
	// 		if 500 > x {
	// 			fmt.Println(0)
	// 			return
	// 		}
	// 	}
	// }

}
