package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 空降到某一点
type point struct {
	i, j int
}

// 上, 左，下，右 四个方向(点)去探索(逆时针)
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 一律使用值类型，不容易出错
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 越界判断
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	// 维护一个数组，从 start 走了多少步，才走到了这一格
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 创建一个对列
	Q := []point{start}

	// 退出条件：已到终点 & 队列为空，走死路了
	for len(Q) > 0 {
		cur := Q[0] // 队头
		Q = Q[1:]   // 队头拿掉

		// 发现终点了，可以滚了
		if cur == end {
			break
		}

		// 开始探索一个点（循环4次，因为有4个方向）
		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0(迷宫下一步不能是墙)
			// and steps at next is 0(曾经到过，不能走了)
			// and next != start (steps一开始都是0，不能回到起点)
			val, ok := next.at(maze)
			// 把不能走剔除掉
			// !ok --> 越界， ok & 1--> 撞墙
			if !ok || val == 1 {
				continue // 继续探索下一个
			}

			val, ok = next.at(steps)
			// 把不能走剔除掉
			// !ok --> 越界， ok & !=0 --> 已经走过
			if !ok || val != 0 {
				continue // 继续探索下一个
			}
			// 不能回到原点
			if next == start {
				continue
			}
			// 能探索下一个点了，但是对于下一个点有两件事要做
			// 1--> 在下一个点中要填第几步(如：2)
			// 2--> 注意把下一点放入队列，以便下一次探索

			// 拿到当前位置的步数
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("12/maze/maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			// 输出 3 位对齐
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
