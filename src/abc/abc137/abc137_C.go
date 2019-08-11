package abc137

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// )

// type Comparable interface {
// 	LessThan(Comparable) bool
// }

// type ComparableSlice []Comparable

// func (c ComparableSlice) Len() int {
// 	return len(c)
// }

// func (c ComparableSlice) Less(i, j int) bool {
// 	return c[i].LessThan(c[j])
// }

// func (c ComparableSlice) Swap(i, j int) {
// 	c[i], c[j] = c[j], c[i]
// }

// func SortComparables(elts []Comparable) {
// 	sort.Sort(ComparableSlice(elts))
// }

// //////////////////////////////////////////////////////////////////////
// // Let's try using this:

// type ComparableRune rune

// func (r1 ComparableRune) LessThan(o Comparable) bool {
// 	return r1 < o.(ComparableRune)
// }

// // sort.type SortBy []Type

// // func (a SortBy) Len() int           { return len(a) }
// // func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// // func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
// // TestC is ...
// func TestC2() {
// 	var n int

// 	var sc = bufio.NewScanner(os.Stdin)
// 	if sc.Scan() {
// 		n, _ = strconv.Atoi(sc.Text())
// 	}
// 	var list = make([]string, n)
// 	for i := 0; i < n; i++ {
// 		if sc.Scan() {
// 			list[i] = sc.Text()
// 		}
// 	}

// 	fmt.Println("forここからだよ")

// 	m := map[string]int{}
// 	for i, s := range list {
// 		fmt.Println("sort前: ", s)
// 		sSlice := make([]string, len(s))
// 		sort.Strings(sSlice)

// 		fmt.Println("sort後: ", sSlice)
// 		comparables := make(ComparableSlice, len(s))
// 		for i, v := range s {
// 			comparables[i] = ComparableRune(v)
// 		}

// 		SortComparables(comparables)

// 		sortedRunes := make([]rune, len(s))
// 		for i, v := range comparables {
// 			sortedRunes[i] = rune(v.(ComparableRune))
// 		}

// 		fmt.Printf("result: %#v\n", string(sortedRunes))
// 		list[i] = string(sortedRunes)
// 		m[list[i]]++
// 	}

// 	fmt.Println("forここからだよ")
// 	// ret := 0
// 	// for i := 0; i < len(list)-1; i++ {
// 	// 	for j := i + 1; j < len(list); j++ {
// 	// 		fmt.Printf("%sと%sの比較:%t\n", list[i], list[j], list[i] == list[j])
// 	// 		if list[i] == list[j] {
// 	// 			ret++
// 	// 		}
// 	// 	}

// 	// }
// 	ret := 0
// 	for _, v := range m {
// 		ret += v * (v - 1) / 2
// 	}

// 	fmt.Println(ret)
// }

// // func removeDuplicate1(args []string) []string {
// // 	results := make([]string, 0, len(args))
// // 	encountered := map[string]bool{}
// // 	for i := 0; i < len(args); i++ {
// // 		if !encountered[args[i]] {
// // 			encountered[args[i]] = true
// // 			results = append(results, args[i])
// // 		}
// // 	}
// // 	return results
// // }
