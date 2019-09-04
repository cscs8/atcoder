package abs

import (
	"fmt"
	"strings"
)

// Test3 test function.
func Test3() {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println(strings.Count(a, "1"))
}
