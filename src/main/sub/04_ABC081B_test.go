package sub

import (
	"fmt"
	"testing"
)

// Test test function.
func Test(t *testing.T) {
	req := 2
	resultI, resultB, err := DivideOfIsEven(req)
	fmt.Printf("req : %d, result : %d, %t, %#v", req, resultI, resultB, err)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if resultI != 1 {
		t.Fatal("failed test")
	}

	if !resultB {
		t.Fatal("failed test")
	}
}
