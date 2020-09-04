package main

import "fmt"

type point struct {
	x int
	y int
}

// 下一个点
func (p point) add(d point) point {
	return point{
		p.x + d.x,
		p.y + d.y,
	}
}

func (p point) at(grid [][]int) (int, bool) {
	// 上下越界
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	// 左右越界
	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}
	// 返回 grid[p.x][p.y] 的好处 --> 撞墙，已经探索等等统统放外层判断
	return grid[p.x][p.y], true
}

func walk(maze [][]int, start point, end point) [][]int {
	// 维护一个与 maze 相同的 Steps Slice
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	// 定义一个队列，并将起点入队
	// start: 已经发现但还没有探索（排队）
	Q := []point{start}

	// 定义 上，左，下，右 四个方向
	dirs := []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	// 经典写法，队列为空，走迷宫结束
	for len(Q) > 0 {
		cur := Q[0] // 要探索的点
		Q = Q[1:]   // 出队

		// 已到终点
		if cur == end {
			break
		}

		// 四个方向探索
		for _, d := range dirs {
			// 下一点
			next := cur.add(d)
			// val：用来判断在 maze 是否是墙(1)
			// ok: 用来判断在 maze 是否越界
			val, ok := next.at(maze)
			// 有墙不能探索，越界不能探索
			if val == 1 || !ok {
				continue
			}
			// val：用来判断点在 steps 是否是值，有值表明这个位置在 `maze` 中已经探索过了
			// ok: 用来判断在 steps 是否越界
			val, ok = next.at(steps)
			if val != 0 || !ok {
				continue
			}
			// 下一探索点不能是起点
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			// 格子放入起点到达它的步数
			steps[next.x][next.y] = curSteps + 1
			// 将发现的格子放入队列--> 已经发现但还没有探索（排队）
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	// 迷宫
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	// 入口
	start := point{0, 0}
	// 出口
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			// 3位对齐，打印结果
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
