package main

import "fmt"

const (
	max_n8 int = 1000
	max_m8 int = 1000
)

var dp5 [max_m8 + 1][max_n8 + 1]int

func solve12(n, m, big_m int, a []int) {
	//一个都不取的方法总是只有一种
	for i := 0; i <= n; i++ {
		dp5[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1-a[i] >= 0 {
				//在有取余的情况下，要避免减法运算出现负值
				dp5[i+1][j] = (dp5[i+1][j-1] + dp5[i][j] - dp5[i][j-1-a[i]] + big_m) % big_m
			} else {
				dp5[i+1][j] = (dp5[i+1][j-1] + dp5[i][j]) % big_m
			}
		}
	}
	fmt.Println(dp5[n][m])
}

func main() {
	var n, m, big_m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&big_m)
	solve12(n, m, big_m, a)
}
