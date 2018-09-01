package main

import "fmt"

const max_n6 int = 100
const max_k int = 100000

var dp2 [max_n6 + 1][max_k + 1]bool

func solve9(n, k int, a []int, m []int) {
	dp2[0][0] = true

	for i := 0; i < n; i++ {

		for j := 0; j <= k; j++ {
			for l := 0; l <= m[i] && l*a[i] <= j; l++ {

				dp2[ i+1 ][j] = dp2[i][j-l*a[i]] || dp2[ i+1 ][j]
				//fmt.Printf("dp2[ %d ][%d] = %v\t dp2[%d][%d] = %v \n", i+1, j, dp2[ i+1 ][j], i, j-l*a[i], dp2[i][j-l*a[i]])

			}
		}

	}

	if dp2[n][k] {
		fmt.Printf("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
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

	solve9(n, k, a, m)

}
