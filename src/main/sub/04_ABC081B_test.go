package sub

import (
	"fmt"
	"strings"
	"testing"
)

// Test test function.
func Test(t *testing.T) {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println(strings.Count(a, "1"))
}
