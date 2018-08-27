package main

import "fmt"

const max_n2 int = 20000

func swap(x, y *int) {
	*x, *y = *y, *x
}

func solve6(n int, l []int) {
	ans := 0

	//直到计算到木板为1块时为止
	for n > 1 {
		// 求出最短的板mii1和次短的板mii2
		mii1, mii2 := 0, 1
		if l[mii1] > l[mii2] {
			swap(&mii1, &mii2)
		}

		for i := 2; i < n; i++ {
			if l[i] < l[mii1] {
				mii2 = mii1
				mii1 = i
			} else if l[i] < l[mii2] {
				mii2 = i
			}
		}

		// 将两块板拼合
		t := l[mii1] + l[mii2]
		ans += t

		if mii1 == n-1 {
			swap(&mii1, &mii2)
		}
		l[mii1] = t
		l[mii2] = l[n-1]
		n--

	}

	fmt.Println(ans)

}

func main() {
	var n int
	fmt.Println("Please enter how many pieces of wood to cut:")
	fmt.Scan(&n)
	l := make([]int, n, max_n2)
	for i := 0; i < n; i++ {
		fmt.Printf("Please enter the length of the board%d:\n", i+1)
		fmt.Scan(&l[i])
	}

	solve6(n, l)

}
