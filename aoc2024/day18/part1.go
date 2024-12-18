package day18

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	numOfBytes = 1024
	memX       = 71
	memY       = 71
)

var grid [memX][memY]int

var dirs = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

type Position struct {
	x int
	y int
}

func (p Position) isValid() bool {
	if p.x >= 0 && p.x < memX && p.y >= 0 && p.y < memY && grid[p.x][p.y] == 0 {
		return true
	}
	return false
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
	f, err := os.Open("day18/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	for range numOfBytes {
		scanner.Scan()
		coordinates := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(coordinates[0]) // distance from the left
		x, _ := strconv.Atoi(coordinates[1]) // distance from the top
		grid[x][y] = 1
	}
	src, end := Position{0, 0}, Position{memX - 1, memY - 1}

	fmt.Println("part1:", dijkstra(src, end))
}

func dijkstra(src Position, end Position) int {
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

		if node.pos == end {
			return node.dist
		}

		for _, dir := range dirs {
			newPosition := Position{node.pos.x + dir[0], node.pos.y + dir[1]}
			if newPosition.isValid() {
				heap.Push(&pq, &Node{newPosition, node.dist + 1, 0})
			}
		}
	}
	return -1
}
