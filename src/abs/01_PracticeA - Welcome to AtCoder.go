package abs

import (
	"fmt"
)

// Test1 is ...
func Test1() {
	var a, b, c int
	var s string
	fmt.Scanln(&a)
	fmt.Scanln(&b, &c)
	fmt.Scanln(&s)
	fmt.Println(a+b+c, s)
}
