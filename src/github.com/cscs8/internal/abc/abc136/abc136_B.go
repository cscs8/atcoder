package abc136

import (
	"fmt"
	"math"
	"strconv"
)

// func calc(array []int) int {
// 	return array[2] - (array[0] - array[1])
// }

// // run is run
// func run(array []string) int {
// 	var arrayInt = make([]int, len(array))

// 	// for i := 0; i < len(array); i++ {
// 	for idx, value := range array {
// 		valueInt, _ := strconv.Atoi(value)
// 		arrayInt[idx] = valueInt
// 	}

// 	if ret := calc(arrayInt); ret > 0 {
// 		return ret
// 	}
// 	return 0
// }
func calcn(i, j int) (int, int) {

	fmt.Println("calcn : ", i)
	fmt.Println("i, j : ", i, j)
	if j%2 == 0 {
		return i, j - 1
	}

	pow := math.Pow(float64(10), float64(j-1))
	fmt.Println("calcn : ", i)
	i = int(pow)
	fmt.Println("pow : ", pow)
	ret := i % int(pow)
	fmt.Println("ret : ", ret)
	if ret == 0 {
		pow = math.Pow(float64(10), float64(j-2))
		fmt.Println("ret calc: ", i-int(pow)-1)
		return i - int(pow) - 1, j - 1
	}
	return ret + 1, j - 1

}

func calcOdd(i, j int) (int, int) {

	fmt.Println("calcOdd : ", i)
	fmt.Println("i, j : ", i, j)
	if j%2 == 0 {
		return i, j - 1
	}

	fmt.Println("calcOdd : ", i)
	pow := math.Pow(float64(10), float64(j-1))
	fmt.Println("pow : ", pow)
	ret := i % int(pow)
	fmt.Println("ret : ", ret)
	if ret == 0 {
		pow = math.Pow(float64(10), float64(j-2))
		fmt.Println("ret calc: ", i-int(pow)-1)
		return i - int(pow) - 1, j - 1
	}
	return ret + 1, j - 1

}

func runB(s string, i int) int {

	var sum, j int = 0, 0
	if len(s) == 1 {
		return i
	}
	if len(s) == 2 {
		return i
	}

	ret, j := calcOdd(i, len(s))
	sum += ret
	if sum == i {
		sum = 0
	}
	fmt.Println("sum : ", sum)

	for j > 1 {
		ret, j = calcOdd(ret, j)
		if j%2 == 0 {
			fmt.Println("sum plus: ")
			sum += ret
		}
		fmt.Println("sum : ", sum)
		fmt.Println("")
	}
	return sum + 9
}

// TestB is ...
func TestB() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(len(s))
	// array := make([]string, len(s))
	// array = append(array, r)
	i, _ := strconv.Atoi(s)
	fmt.Println(runB(s, i))
}
