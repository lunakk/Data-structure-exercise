package main

import (
	"fmt"
	"math"
)

const max_n3 int = 100

func rec(i int, j int, w []int, v []int, n int, dp *[max_n3+1][max_n3+1]int) (int) {
	if dp[i][j] >= 0 {
		return dp[i][j]
	}
	var res int
	if i == n {
		res = 0
	} else if j < w[i] {
		res = rec(i+1, j, w, v, n, dp)
	} else {
		res = int(math.Max(float64(rec(i+1, j, w, v, n, dp)), float64(rec(i+1, j-w[i], w, v, n, dp)+v[i])))
	}
	dp[i][j] = res

	return dp[i][j]
}

func main() {
	var n, totalWeight int
	fmt.Println("Please enter total weight:")
	fmt.Scan(&totalWeight)
	fmt.Println("Please enter n:")
	fmt.Scan(&n)
	w := make([]int, n, max_n3)
	v := make([]int, n, max_n3)

	var dp = [max_n3 + 1][max_n3 + 1]int{}

	for i := 0; i < n; i++ {
		fmt.Printf("Please enter the weight and value of item%d:", i+1)
		fmt.Scan(&w[i], &v[i])

	}

	for i := 0; i <= max_n3; i++ {
		for j := 0; j <= max_n3; j++ {
			dp[i][j] = -1
		}
	}

	fmt.Printf("%d\n", rec(0, totalWeight, w, v, n, &dp))

}
