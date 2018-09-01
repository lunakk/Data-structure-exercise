package main

import (
	"fmt"
	"math"
)

const max_n5 int = 100

var dp1 = [max_n5 + 1][max_n5 + 1]int{}

func solve8(n, totalweight int, w []int, v []int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= totalweight; j++ {
			if j < w[i] {
				dp1[i+1][j] = dp1[i][j]
			} else {
				dp1[i+1][j] = int(math.Max(float64(dp1[i][j]), float64(dp1[i+1][j-w[i]]+v[i])))
			}
		}
	}

	fmt.Println(dp1[n][totalweight])

}

func main() {
	var n, totalWeight int
	fmt.Println("Please enter total weight:")
	fmt.Scan(&totalWeight)
	fmt.Println("Please enter n:")
	fmt.Scan(&n)
	w := make([]int, n, max_n5)
	v := make([]int, n, max_n5)

	for i := 0; i < n; i++ {
		fmt.Printf("Please enter the weight and value of item%d:", i+1)
		fmt.Scan(&w[i], &v[i])

	}

	for i := 0; i <= max_n5; i++ {
		for j := 0; j <= max_n5; j++ {
			dp1[i][j] = 0
		}
	}

	solve8(n, totalWeight, w, v)

}
