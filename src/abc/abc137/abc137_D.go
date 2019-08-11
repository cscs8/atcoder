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

		// if retMap[a] < b {
		// 	retMap[a] = b
		// }
	}

	// DESC JobList sort
	sort.Sort(SortBy(jobList))

	maxDayCost := (m * (2 + (m - 1))) / 2
	dayCost := 0
	days := m
	ret := 0
	sum := 0
	i := 0
	fmt.Println("maxDayCost : ", maxDayCost)
	// for _, j := range jobList {
	size := len(jobList)
	if m < len(jobList) {
		size = m
	}
	for _, j := range jobList[:size] {
		fmt.Println()
		fmt.Println("job : ", j)
		fmt.Println("maxDayCost : ", maxDayCost)
		fmt.Println("dayCost : ", dayCost)
		fmt.Println("j.A : ", j.A)
		fmt.Println("days : ", days)
		fmt.Println("sum : ", sum)
		// if (maxDayCost < j.A) || (days < j.A) {
		// if days < j.A {
		if maxDayCost < j.A || maxDayCost < i+j.A {
			fmt.Println("dayCost+j.A : ", dayCost+j.A)
			fmt.Println("Continue!")
			fmt.Println()
			i++
			maxDayCost -= j.A
			days--
			break
		}

		dayCost += j.A
		ret += j.B
		fmt.Println("ret : ", ret)
		maxDayCost -= j.A
		days--
		i++
	}

	fmt.Println("jobList : ", jobList)
	fmt.Println(ret)
	// var sc = bufio.NewScanner(os.Stdin)
	// if sc.Scan() {
	// 	n, _ = strconv.Atoi(sc.Text())
	// }
	// var list = make([]string, n)
	// for i := 0; i < n; i++ {
	// 	if sc.Scan() {
	// 		list[i] = sc.Text()
	// 	}
	// }

	// m := map[string]int{}
	// for _, s := range list {
	// 	m[SortStringByCharacter(s)]++
	// }

	// ret := 0
	// for _, v := range m {
	// 	ret += v * (v - 1) / 2
	// }

	// fmt.Println(ret)
}
