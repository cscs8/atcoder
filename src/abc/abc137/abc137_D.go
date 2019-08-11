package abc137

import (
	"fmt"
	"sort"
)

// Job is ...
type Job struct {
	A int
	B int
}

// AscABy is ...
type AscABy []Job

func (a AscABy) Len() int      { return len(a) }
func (a AscABy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a AscABy) Less(i, j int) bool {
	return a[i].A < a[j].A
}

// DescBBy is ...
type DescBBy []Job

func (a DescBBy) Len() int      { return len(a) }
func (a DescBBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DescBBy) Less(i, j int) bool {
	return a[i].B < a[j].B
}

// DescBy is ...
type DescBy []int

func (a DescBy) Len() int      { return len(a) }
func (a DescBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DescBy) Less(i, j int) bool {
	return a[j] < a[i]
}

// TestD is ...
func TestD() {
	var n, m, a, b int

	fmt.Scanln(&n, &m)

	jobList := make([]Job, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		if m < a {
			continue
		}
		jobList[i] = Job{a, b}
	}

	// A Asc JobList sort
	sort.Sort(AscABy(jobList))

	retList := []int{}
	ret := 0
	size := len(jobList)
	if m < len(jobList) {
		size = m
	}
	// for _, j := range jobList[:size] {
	for i := 1; i <= m; i++ {
		for _, j := range jobList {
			if j.A == i {
				retList = append(retList, j.B)
				continue
			}
			if j.A == i+1 {
				break
			}
		}

		if 0 == len(retList) {
			continue
		}
		sort.Sort(DescBy(retList))
		if len(retList) > size {
			retList = retList[:size]
		}
		fmt.Println(retList)

		ret += retList[0]
		retList[0] = 0
	}
	fmt.Print(ret)
}
