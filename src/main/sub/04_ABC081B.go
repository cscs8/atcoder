package sub

import (
	"fmt"
	"strings"
)

// DivideOfIsEven is..
func DivideOfIsEven(a int) (int, bool) {
	return a / 2, a%2 == 0
}

// Test4 test function.
func Test4() {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println(strings.Count(a, "1"))
}
