package main

import "fmt"

const max_m7 int = 1000
const max_n7 int = 1000

var dp4 [max_m7 + 1][max_n7 + 1]int

func solve11(n,m,big_m int) {
	dp4[0][0] = 1
	for i := 1; i <= m; i++ {
		for j :=0; j <= n;j++ {
			if j - i >= 0 {
				dp4[i][j] = (dp4[i][j-i] + dp4[i-1][j]) % big_m
			} else {
				dp4[i][j] = dp4[i-1][j]
			}
		}
	}
	fmt.Println(dp4[m][n])
}

func main() {
	var n, m ,big_m int
	fmt.Scan(&n, &m, &big_m)
	solve11(n,m,big_m)
}
