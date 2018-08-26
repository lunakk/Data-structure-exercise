package main

import (
	"fmt"
	"sort"
)

const max_n int = 100000

type Pair struct {
	first int
	second int
}

type PairSlice []Pair

func (a PairSlice) Len() int {
	return len(a)
}

func (a PairSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PairSlice) Less(i, j int) bool {
	return a[j].second > a[i].second
}

func solve3(s []int, t []int, itv []Pair, n int) {
	//对Pair进行的是字典序比较
	//为了让结束时间早的工作排在前面，把t存入first，把s存入second
	for i := 0; i < n; i++ {
		itv[i].first = t[i]
		itv[i].second = s[i]
	}

	sort.Sort(PairSlice(itv))

	ans, time := 0, 0
	optJob := []int{}
	for i := 0; i < n; i++ {
		if time < itv[i].second {
			ans++
			time = itv[i].first
			optJob = append(optJob,i+1)
		}
	}

	fmt.Printf("%d (选取工作",ans)
	for _, data := range optJob {
		fmt.Printf("%d ",data)
	}
	fmt.Printf(")")
}

func main() {
	var n int     //工作的组数

	fmt.Println("Please enter how many jobs:")
	fmt.Scan(&n)

	s := make([]int, n, max_n)	  //开始时间
	t := make([]int, n, max_n)	  //结束时间
	itv := make([]Pair, n ,max_n)    //用于对工作排序的Pair数组

	for i := 0; i < n ; i++ {

		fmt.Printf("Please enter the start time and end time of job%d:\n",i+1)
		fmt.Scan(&s[i],&t[i])

	}

	solve3(s,t,itv,n)
}
