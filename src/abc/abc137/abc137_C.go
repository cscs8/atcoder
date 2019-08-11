package abc137

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

// TestC is ...
func TestC() {
	var n int

	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	var list = make([]string, n)
	for i := 0; i < n; i++ {
		if sc.Scan() {
			list[i] = sc.Text()
		}
	}

	m := map[string]int{}
	for _, s := range list {
		m[SortStringByCharacter(s)]++
	}

	ret := 0
	for _, v := range m {
		ret += v * (v - 1) / 2
	}

	fmt.Println(ret)
}
