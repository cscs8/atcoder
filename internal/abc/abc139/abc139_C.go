package abc139

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TestC is...
func TestC() {
	var n int
	fmt.Scanln(&n)

	var sc = bufio.NewScanner(os.Stdin)
	var list = make([]string, n)
	var listInt = make([]int, n)

	if sc.Scan() {
		list = strings.Fields(sc.Text())
	}

	for idx, value := range list {
		valueInt, _ := strconv.Atoi(value)
		listInt[idx] = valueInt
	}

	max, tmp := 0, 0
	fmt.Println(listInt)
	for i := 0; i < n-1; i++ {
		fmt.Println("list[i", i, "] : ", list[i])
		tmp = 0

		for j := i + 1; j < n; j++ {
			fmt.Println("  list[j:", j, "] : ", list[j])
			fmt.Println("  list[j-1", j-1, "] : ", list[j-1])
			if listInt[j] > listInt[j-1] {
				break
			}
			tmp++
			fmt.Println("tmp++")
		}
		if tmp > max {
			max = tmp
			fmt.Println("max upd")
		}

	}
	fmt.Println(max)
}
