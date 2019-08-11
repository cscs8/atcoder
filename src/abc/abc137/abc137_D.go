package abc137

import (
	"fmt"
	"sort"
)

type Job struct {
	A int
	B int
	C int
}
type SortBy []Job

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[j].C < a[i].C }

// TestD is ...
func TestD() {
	var n, m, a, b int

	fmt.Scanln(&n, &m)

	jobList := make([]Job, n)
	var cost int
	costMap := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		if m < a {
			continue
		}
		cost = -a + 2*b
		costMap[a] = cost

		jobList[i] = Job{a, b, cost}
		fmt.Println("cost : ", cost)
	}

	// DESC JobList sort
	sort.Sort(SortBy(jobList))

	maxDayCost := (m * (2 + (m - 1))) / 2
	dayCost := 0
	days := m
	ret := 0
	i := 0
	fmt.Println("maxDayCost : ", maxDayCost)
	size := len(jobList)
	if m < len(jobList) {
		size = m
	}
	for _, j := range jobList[:size] {
		if maxDayCost < j.A || maxDayCost < i+j.A {
			i++
			maxDayCost -= j.A
			days--
			break
		}

		dayCost += j.A
		ret += j.B
		maxDayCost -= j.A
		days--
		i++
	}

	fmt.Println(ret)
}
