package sub

import "fmt"

// DivideOfIsEven is..
func DivideOfIsEven(a int) (int, bool, error) {
	return a / 2, a%2 == 0, nil
}

// Animal ..
type Animal struct {
	Name string
	// Age  uint
	Age int32
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

// Test4 test function.
func Test4() {
	a := Animal{
		Name: "Gopher",
		Age:  16,
	}
	fmt.Printf("%#v\n", a)
	fmt.Println(a)
	// var n int
	// var s string
	// // N
	// fmt.Scanf("%d", &n)

	// // AN
	// fmt.Scanf("%s", &s)
	// strings.Split(s, strings.bla" ")
	// fmt.Println(strings.Count(a, "1"))
}
