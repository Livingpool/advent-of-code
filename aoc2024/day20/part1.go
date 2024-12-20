package day20

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var dirs = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

type Position struct {
	i int
	j int
}

func (p Position) isValid(racetrack [][]rune) bool {
	if p.i < 0 || p.j < 0 || p.i >= len(racetrack) || p.j >= len(racetrack[0]) || racetrack[p.i][p.j] == '#' {
		return false
	}
	return true
}

type Node struct {
	pos   Position
	dist  int
	index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Part1() {
	f, err := os.Open("day20/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var racetrack [][]rune
	var start, end Position

	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for j, c := range line {
			row[j] = c
			if row[j] == 'S' {
				start.i, start.j = i, j
				row[j] = '.'
			} else if row[j] == 'E' {
				end.i, end.j = i, j
				row[j] = '.'
			}
		}
		racetrack = append(racetrack, row)
		i++
	}

	var baseTime int = dijkstra(start, end, Position{-1, -1}, racetrack)
	var cheats int = 0
	// var cheatsMap = make(map[int]int)

	for i := 1; i < len(racetrack)-1; i++ {
		for j := 1; j < len(racetrack[0])-1; j++ {
			if racetrack[i][j] == '#' {
				racetrack[i][j] = '.'
				result := dijkstra(start, end, Position{i, j}, racetrack)
				if baseTime-result >= 100 {
					cheats++
				}
				// if baseTime-result != 0 {
				// 	cheatsMap[baseTime-result]++
				// }
				racetrack[i][j] = '#'
			}
		}
	}

	fmt.Println("part1:", cheats)

	// for k, v := range cheatsMap {
	// 	fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
	// }
}

func dijkstra(src, end, cheat Position, racetrack [][]rune) int {
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{src, 0, 0}
	heap.Init(&pq)
	visited := make(map[Position]bool)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if _, exists := visited[node.pos]; exists {
			continue
		}
		visited[node.pos] = true

		if node.pos == cheat {
			racetrack[cheat.i][cheat.j] = '#'
		}

		if node.pos == end {
			return node.dist
		}

		for _, dir := range dirs {
			newPosition := Position{node.pos.i + dir[0], node.pos.j + dir[1]}
			if newPosition.isValid(racetrack) {
				heap.Push(&pq, &Node{newPosition, node.dist + 1, 0})
			}
		}
	}
	return -1
}
