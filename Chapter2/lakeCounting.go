package main

import "fmt"

var n int = 10
var m int = 12

func dfs(x int, y int) {
	//将现在所在位置替换为.
	garden[x][y] = '.'

	//循环遍历移动的8个方向
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx := x + dx
			ny := y + dy
			if 0 <= nx && nx < n && 0 <= ny && ny < m && garden[nx][ny] == 'W' {
				dfs(nx, ny)
			}
		}
	}
	return
}

func solve() {
	var res int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if garden[i][j] == 'W' {
				dfs(i, j)
				res++
			}
		}
	}
	fmt.Printf("一共有%d个水洼\n", res)
}

var garden = [10][12]byte{{'W', '.', '.', '.', '.', '.', '.', '.', '.', 'W', 'W', '.'},
	{'.', 'W', 'W', 'W', '.', '.', '.', '.', '.', 'W', 'W', 'W'},
	{'.', '.', '.', '.', 'W', 'W', '.', '.', '.', 'W', 'W', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'W', 'W', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'W', '.', '.'},
	{'.', '.', 'W', '.', '.', '.', '.', '.', '.', 'W', '.', '.'},
	{'.', 'W', '.', 'W', '.', '.', '.', '.', '.', 'W', 'W', '.'},
	{'W', '.', 'W', '.', 'W', '.', '.', '.', '.', '.', 'W', '.'},
	{'.', 'W', '.', 'W', '.', '.', '.', '.', '.', '.', 'W', '.'},
	{'.', '.', 'W', '.', '.', '.', '.', '.', '.', '.', 'W', '.'}}

func main() {

	for _, data := range garden {
		for _, data1 := range data {
			fmt.Printf("%v\t", string(data1))
		}
		fmt.Println()
	}
	solve()
}
