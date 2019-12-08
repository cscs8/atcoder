package abc147

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TestC is ...
func TestC() {
	var n int
	fmt.Scanln(&n)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	// var list = make([]string, n)
	var listA []int
	// for i := 0; i < n; i++ {
	for i := n; i >= 0; i-- {
		if sc.Scan() {
			i, _ := strconv.Atoi(sc.Text())
			listA = append(listA, i)
			for ; i >= 0; i-- {
				if sc.Scan() {
					i, _ := strconv.Atoi(sc.Text())
					j, _ := strconv.Atoi(sc.Text())
					listA = append(listA, i)
					listA = append(listA, j)
				}
			}
		}
	}
	fmt.Println("n : ", n)
	fmt.Println("list : ", listA)
	// }

	// スペースで区切られた各単語を、スペース削除してスキャナーを分割する
	// e.g. 1 2 3 4 5
	//  -> 1
	//     2
	//     3
	//     4
	// //     5
	// sc.Split(bufio.ScanWords)

	// for idx := range list {
	// 	if sc.Scan() {
	// 		valueInt, _ := strconv.Atoi(sc.Text())
	// 		list[idx] = valueInt
	// 	}
	// }

	// max, tmp := 0, 0
	// for i := n - 1; i > 0; i-- {
	// 	if list[i-1] >= list[i] {
	// 		tmp++
	// 	} else {
	// 		tmp = 0
	// 	}

	// 	if tmp > max {
	// 		max = tmp
	// 	}
	// }

	// fmt.Println(max)

}
