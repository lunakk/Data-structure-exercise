package main

import (
	"fmt"
	"math"
)

const max_n4 int = 1000
const max_m4 int = 1000

var dp [max_n4+1][max_m4+1]int

func solve7(s string, t string) {
	n, m := len(s), len(t)

	for i := 0; i <n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1;
			}else {
				dp[i+1][j+1] = int(math.Max(float64(dp[i][j+1]),float64(dp[i+1][j])))
			}
		}
	}
	fmt.Println(dp[n][m])
}

func main() {
	var s, t string

	fmt.Println("Please enter string s:")
	fmt.Scan(&s)

	fmt.Println("Please enter string t:")
	fmt.Scan(&t)

	solve7(s, t)
}