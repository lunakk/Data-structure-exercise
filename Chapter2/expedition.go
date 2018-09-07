package main

import (
	"container/list"
	"fmt"
)

const max_n9 int = 10000

func solve13(l, p, n int, a, b []int) {
	// 为了写起来方便，我们把终点也认为是加油站
	a[n] = l
	b[n] = 0
	n++

	//维护加油站的优先队列
	que := list.New()

	// ans：加油次数，pos：现在所在位置,tank:油箱中汽油的量
	ans, pos, tank := 0,0,p

	for i := 0; i < n; i++ {
		//接下去要前进的距离
		d := a[i] - pos

		//不断加油直到油量足够行驶到下一个加油站
		for tank - d < 0 {
			if que.Len() == 0 {
				fmt.Println("-1")
				return
			}
			g := que.Front()
			que.Remove(g)
			v,_ := g.Value.(int)
			tank += v
			ans++
		}

		tank -= d
		pos = a[i]
		que.PushBack(b[i])
	}

	fmt.Println(ans)

}

func main() {
	var l, p, n int
	fmt.Scan(&n, &l, &p)
	a := make([]int, max_n9)
	b := make([]int, max_n9)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	solve13(l, p, n, a, b)
}
