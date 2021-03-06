package abs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DivideOfIsEven is..
func DivideOfIsEven(a int) (int, bool) {
	return a / 2, a%2 == 0
}

// Operation is ...
func Operation(array []int) ([]int, bool) {
	var resultArray = make([]int, len(array))
	var isEven bool
	for i := 0; i < len(array); i++ {
		resultArray[i], isEven = DivideOfIsEven(array[i])
		if !isEven {
			return nil, false
		}
		if resultArray[i] == 0 {
			return nil, false
		}
	}
	return resultArray, true
}

// run is run
func run(array []string) int {
	var arrayInt = make([]int, len(array))
	operationCount := 0

	// for i := 0; i < len(array); i++ {
	for idx, value := range array {
		valueInt, _ := strconv.Atoi(value)
		arrayInt[idx] = valueInt
	}

	arrayInt, hasNext := Operation(arrayInt)
	for hasNext {
		arrayInt, hasNext = Operation(arrayInt)
		operationCount++
	}
	return operationCount
}

// Test4 test function.
func Test4() {
	var n, an string
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		n = sc.Text()
	}
	if sc.Scan() {
		an = sc.Text()
	}

	ni, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	splitAn := strings.Fields(an)[:ni]
	fmt.Println(run(splitAn))
}
