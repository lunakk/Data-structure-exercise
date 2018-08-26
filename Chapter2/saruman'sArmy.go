package main

import (
	"fmt"
	"sort"
)

const max_n1 int = 1000

func solve5(n int, r int, x []int) {
	sort.Ints(x)

	i, ans := 0, 0
	for i < n {
		// s是没有被覆盖的最左的点的位置
		s := x[i]
		i++

		for i < n && x[i] <= s+r {
			i++
		}

		//一直向右前进直到距s的距离大于r的点
		p := x[i-1]
		//一直向右前进直到距p的距离大于r的点
		for i < n && x[i] <= p+r {
			i++
		}

		ans++
	}

	fmt.Println(ans)

}

func main() {
	var n, r int

	fmt.Println("Please enter a total of how many points:")
	fmt.Scan(&n)

	var x = make([]int, n, max_n1)

	fmt.Println("Please enter the specified distance:")
	fmt.Scan(&r)

	for i := 0; i < n; i++ {
		fmt.Printf("Please enter the location of point%d:\n", i+1)
		fmt.Scan(&x[i])
	}

	solve5(n, r, x)
}
