package abc137

import (
	"fmt"
	"sort"
)

// Job is ...
type Job struct {
	A int
	B int
	C int
}

// SortBy is ...
type SortBy []Job

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[j].C < a[i].C }

// TestD is ...
func TestD() {
	var n, m, a, b int

	fmt.Scanln(&n, &m)

	jobList := make([]Job, n)
	cost := 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		if m < a {
			continue
		}
		cost = -a + 2*b

		jobList[i] = Job{a, b, cost}
	}

	// DESC JobList sort
	sort.Sort(SortBy(jobList))

	// SUURETSU
	maxDayCost := (m * (2 + (m - 1))) / 2
	ret, i := 0, 0
	size := len(jobList)
	if m < len(jobList) {
		size = m
	}
	for _, j := range jobList[:size] {
		if maxDayCost < i+j.A {
			break
		}
		ret += j.B
		maxDayCost -= j.A
		i++
	}
	fmt.Println(ret)
}
