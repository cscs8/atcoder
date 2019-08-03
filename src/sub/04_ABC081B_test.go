package sub

import (
	"fmt"
	"testing"
)

// Test test function.
func Test(t *testing.T) {
	req := 2
	resultI, resultB := DivideOfIsEven(req)
	fmt.Printf("req : %d, result : %d, %t", req, resultI, resultB)
	if resultI != 1 {
		t.Fatal("failed test")
	}

	if !resultB {
		t.Fatal("failed test")
	}
}
