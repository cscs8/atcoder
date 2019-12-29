package abc138

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TestB is...
func TestB() {
	var n int
	fmt.Scanln(&n)

	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		array := strings.Fields(sc.Text())[:n]
		fmt.Print(run(array))
	}
}

func calc(array []int) int {
	return array[2] - (array[0] - array[1])
}

// run is run
func run(array []string) float32 {
	// var arrayInt = make([]float32, len(array))
	var sum float32

	// for i := 0; i < len(array); i++ {
	for _, value := range array {
		valueInt, _ := strconv.Atoi(value)
		// arrayInt[idx] = 1 / float32(valueInt)
		sum += 1 / float32(valueInt)
	}

	// if ret := calc(arrayInt); ret > 0 {
	// 	return ret
	// }
	return 1 / sum
}
