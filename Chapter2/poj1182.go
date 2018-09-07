package main

import (
	"fmt"
)

const max_k2 int = 100000

var par [max_k2]int  //父亲
var rank [max_k2]int //树的高度

//初始化n个元素
func init1(n int) {
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
}

// 查询树的根
func find(x int) int{
	if par[x] == x {
		return x
	} else {
		par[x] = find(par[x])
		return par[x]
	}
}

//合并x和y所属的集合
func unite(x,y int) {
	x = find(x)
	y = find(y)
	if x == y {
		return
	}

	if rank[x] < rank[y] {
		par[x] = y
	} else {
		par[y] = x
		if rank[x] == rank[y] {
			rank[x]++
		}
	}
}

//判断x和y是否属于同一个集合
func same(x,y int) bool {
	return find(x) == find(y)
}

func solve14(n, k int, t, x, y []int) {
	// 初始化并查集
	// 元素x,x+n,x+2*n分别代表x-A,x-B,x-C
	init1(n*3)

	ans := 0
	for i := 0; i < k; i++ {
		t1 := t[i]
		x1 := x[i] - 1
		y1 := y[i] - 1 // 把输入变成0,...,n-1的范围

		//不正确的编号
		if x1 < 0 || n <= x1 || y1 <0 || n <= y1 {
			ans++
			continue
		}

		if t1 == 1 {
			if same(x1,y1+n) || same(x1,y1+2*n) {
				ans++
			}else {
				unite(x1,y1)
				unite(x1+n,y1+n)
				unite(x1+n*2,y1+n*2)
			}
		}else {
			// "x吃y"的信息
			if same(x1,y1) || same(x1,y1+2*n) {
				ans++
			}else {
				unite(x1,y1+n)
				unite(x1+n,y1+2*n)
				unite(x1+2*n,y1)
			}
		}
	}

	fmt.Println(ans)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	t := make([]int, max_k2)
	x := make([]int, max_k2)
	y := make([]int, max_k2)
	for i := 0; i < k; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}

	solve14(n, k, t, x, y)
}
