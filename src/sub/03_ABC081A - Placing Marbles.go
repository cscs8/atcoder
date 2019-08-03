package sub

import (
	"fmt"
	"strings"
)

// Test2 test function.
func Test2() {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println(strings.Count(a, "1"))
}
