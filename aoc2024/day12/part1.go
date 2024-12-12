package day12

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func Part1() {
	f, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var plots [][]rune
	var visited [][]bool

	for scanner.Scan() {
		line := scanner.Text()
		r1 := make([]rune, len(line))
		r2 := make([]bool, len(line))

		for i := range len(line) {
			r1[i] = rune(line[i])
		}

		plots = append(plots, r1)
		visited = append(visited, r2)
	}

	var result int = 0
	for i := range len(plots) {
		for j := range len(plots[0]) {
			if visited[i][j] == false {
				area, perimeter := calcPrice(plots, visited, i, j)
				result += area * perimeter
			}
		}
	}
	fmt.Println("part1:", result)
}

// dfs
func calcPrice(plots [][]rune, visited [][]bool, i, j int) (int, int) {
	region := plots[i][j]
	area, perimeter := 1, 0

	visited[i][j] = true

	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if newI < 0 || newI >= len(plots) || newJ < 0 || newJ >= len(plots[0]) || plots[newI][newJ] != region {
			perimeter++
		}
	}

	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < len(plots) && newJ >= 0 && newJ < len(plots[0]) && !visited[newI][newJ] && plots[newI][newJ] == region {
			a, p := calcPrice(plots, visited, newI, newJ)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}
