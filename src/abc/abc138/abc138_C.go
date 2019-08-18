package abc138

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FloatSlice []float64

func (a FloatSlice) Len() int           { return len(a) }
func (a FloatSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FloatSlice) Less(i, j int) bool { return a[i] > a[j] }

type Value struct {
	value float64
	ok    bool
}

// TestC is...
func TestC() {
	var n int
	fmt.Scanln(&n)

	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		array := strings.Fields(sc.Text())[:n]
		fmt.Print(runC(array))
	}
}

// runC is run
func runC(arrayStr []string) float64 {
	n := len(arrayStr)
	var array = make([]float64, n)

	for idx, value := range arrayStr {
		valueInt, _ := strconv.Atoi(value)
		array[idx] = float64(valueInt)
	}

	sort.Float64s(array)
	var ret float64 = 0

	for idx, value := range array {
		if idx == 0 {
			ret = (value + array[idx+1]) / 2
			continue
		}
		if idx == 1 {
			continue
		}
		ret = (ret + value) / 2
	}
	return ret

	// arrayRet := make([]float64, Fact(n))
	// ret := float64(0)
	// c := 0

	// for i := 0; i < n-1; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		array2 := make([]Value, n)
	// 		for idx, value := range array {
	// 			array2[idx].value = value
	// 			array2[idx].ok = false
	// 		}
	// 		ret = (array[i] + array[j]) / 2

	// 		array2[i].ok = true
	// 		array2[j].ok = true
	// 		for idx, value := range array2 {
	// 			if array2[idx].ok {
	// 				continue
	// 			}
	// 			ret = (ret + value.value) / 2
	// 			array2[idx].ok = true
	// 		}

	// 		// arrayRet[c] = ret
	// 		arrayRet = append(arrayRet, ret)
	// 		c++
	// 	}
	// }

	// sort.Sort(FloatSlice(arrayRet))

	// return arrayRet[0]
}

// 階乗
// https://qiita.com/ozwk/items/4ffdba8e83df1171ac39
func Fact(n int) int {
	if n > 1 {
		return n * Fact(n-1)
	}
	return 1
}
