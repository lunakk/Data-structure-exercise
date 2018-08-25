package main

import (
	"fmt"
	"math"
)

var v = [6]int{1, 5, 10, 50, 100, 500}

func solve2(c [6]int, A int) {
	ans := 0
	var t [6]int

	for i := 5; i >= 0; i-- {
		t[i] = int(math.Min(float64(A/v[i]),float64(c[i])))    //使用硬币i的枚数
		A -= t[i]*v[i]
		ans += t[i]
	}

	fmt.Printf("%d枚(",ans)
	for i := 0; i < 6; i++ {
		fmt.Printf("%d元硬币%d枚 ",v[i],t[i])
	}
	fmt.Println(")")
}

func main() {

	var c [6]int
	var A int

	for i, data := range v{
		fmt.Printf("Please enter the number of %d yuan silver coins:", data)
		fmt.Scan(&c[i])
		fmt.Println()
	}
	fmt.Printf("Please enter how much to pay:")
	fmt.Scan(&A)

	solve2(c,A)
}
