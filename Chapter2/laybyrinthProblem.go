package main

import (
	"fmt"
	"container/list"
)

const inf = 1000000000

var maze = [10][10]byte{{'#', 'S', '#', '#', '#', '#', '#', '#', '.', '#'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '#'},
	{'.', '#', '.', '#', '#', '.', '#', '#', '.', '#'},
	{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'#', '#', '.', '#', '#', '.', '#', '#', '#', '#'},
	{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	{'.', '#', '#', '#', '#', '#', '#', '#', '.', '#'},
	{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	{'.', '#', '#', '#', '#', '.', '#', '#', '#', '.'},
	{'.', '.', '.', '.', '#', '.', '.', '.', 'G', '#'}}
var n1, m1 int = 10, 10  
var sx, sy int = 0, 1		//起点位置
var gx, gy int = 9, 8		//终点位置

var d [10][10]int			//到各个位置的最短距离的数组

// 4个方向移动的向量
var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

//求从(sx, sy)到(gx, gy)的最短距离
//如果无法到达，则是inf
func bfs()  (result int){
	que := list.New()
	for i := 0; i < n1; i++ {
		for j := 0; j < m1; j++ {
			d[i][j] = inf
		}
	}
	var p = [2]int{sx, sy}  //表示状态，即坐标
	//fmt.Println("p[0]=",p[0])
	//fmt.Println("p[1]=",p[1])
	que.PushBack(p[0])
	que.PushBack(p[1])
	d[sx][sy] = 0

	//不断循环直到队列的长度为0
	for que.Len() != 0 {
		//从队列的最前端取出元素
		p1 := que.Front()
		que.Remove(p1)
		v1 , _ :=p1.Value.(int)
		//fmt.Println("v1 = ",v1)
		p2 := que.Front()
		que.Remove(p2)
		v2, _ := p2.Value.(int)
		//fmt.Println("v2 = ",v2)


		//如果取出的状态已经时终点，则结束搜索
		if v1 == gx && v2 == gy {
			break
		}

		//四个方向的循环
		for i := 0; i < 4; i++ {
			//移动之后的位置记为(nx,ny)
			nx ,ny := v1 + dx[i], v2 +dy[i]

			//判断是否可以移动以及是否已经访问过(d[nx][ny]!=inf即为访问过)
			if 0 <= nx && nx < n1 && 0 <= ny && ny < m1 && maze[nx][ny] != '#' && d[nx][ny] == inf {
				//可以移动的话，则加入队列，并且到该位置的距离确定为到p的距离+1
				p[0], p[1] = nx, ny
				que.PushBack(p[0])
				que.PushBack(p[1])
				d[nx][ny] = d[v1][v2] +1
				//fmt.Printf("d[%d][%d] = %d\n",nx, ny, d[nx][ny])

			}
		}
	}
	result = d[gx][gy]
	return
}

func solve1()  {
	res := bfs()
	fmt.Printf("%d\n",res)
}

func main() {
	for _, data := range maze {
		for _, data1 := range data {
			fmt.Printf("%v\t", string(data1))
		}
		fmt.Println()
	}

	solve1()

}
