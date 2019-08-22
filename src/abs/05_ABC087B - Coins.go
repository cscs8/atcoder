package abs

import (
	"fmt"
)

// Test5 test function.
func Test5() {
	length := 5
	capacity := 10
	array := make([]int, length, capacity)

	fmt.Printf("初期の長さ          -> %d\n", len(array))
	fmt.Printf("初期の容量          -> %d\n", cap(array))
	fmt.Printf("初期のアドレス      -> %p\n", array)

	appendCount := 5
	for i := 0; i < appendCount; i++ {
		array = append(array, i+1)
	}

	fmt.Println()
	fmt.Printf("append 後の長さ     -> %d\n", len(array))
	fmt.Printf("append 後の容量     -> %d\n", cap(array))
	fmt.Printf("append 後のアドレス -> %p\n", array)
}
