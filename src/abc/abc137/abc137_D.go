package abc137

import (
	"fmt"
)

// TestD is ...
func TestD() {
	// var n int
	var n, m, a, b int

	fmt.Scanln(&n, &m)

	var listA, listB []int
	ret := map[int]int{}
	for ; 0 < n; n-- {
		fmt.Scanln(&a, &b)
		if m < a {
			continue
		}
		fmt.Println(a, b)

		listA = append(listA, a)
		listB = append(listB, b)
		if ret[a] < b {
			ret[a] = b
		}
	}
	fmt.Println(ret)
	fmt.Println(listA, listB)
	// var sc = bufio.NewScanner(os.Stdin)
	// if sc.Scan() {
	// 	n, _ = strconv.Atoi(sc.Text())
	// }
	// var list = make([]string, n)
	// for i := 0; i < n; i++ {
	// 	if sc.Scan() {
	// 		list[i] = sc.Text()
	// 	}
	// }

	// m := map[string]int{}
	// for _, s := range list {
	// 	m[SortStringByCharacter(s)]++
	// }

	// ret := 0
	// for _, v := range m {
	// 	ret += v * (v - 1) / 2
	// }

	// fmt.Println(ret)
}
