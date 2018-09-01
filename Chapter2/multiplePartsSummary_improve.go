package main

import "fmt"

const max_k1 int = 100000
var dp3 [max_k1+1]int

func solve10(n, k int, a []int, m []int) {
	dp3[0] = 0;
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			if dp3[j] >= 0 {
				dp3[j] = m[i]
			} else if j < a[i] || dp3[j-a[i]] <= 0 {
				dp3[j] = -1
			} else {
				dp3[j] = dp3[j-a[i]] -1
			}
		}
	}

	if dp3[k] >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main()  {
	var n, k int
	a := []int{}
	m := []int{}

	fmt.Println("Please enter n:")
	fmt.Scan(&n)
	fmt.Println("Please enter the sum(K):")
	fmt.Scan(&k)

	for i := 0; i < n; i++ {
		var ai, mi int
		fmt.Printf("Please enter a[%d] and m[%d]:", i, i)
		fmt.Scan(&ai, &mi)
		a = append(a, ai)
		m = append(m, mi)
	}

	solve10(n, k, a, m)
}