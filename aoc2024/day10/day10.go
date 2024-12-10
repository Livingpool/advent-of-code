package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	i int
	j int
}

var dirs = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func Day10() {
	f, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var matrix [][]int
	var trailHeads []*pos

	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for j := 0; j < len(line); j++ {
			num, _ := strconv.Atoi(string(line[j]))
			row[j] = num

			if num == 0 {
				trailHeads = append(trailHeads, &pos{
					i: i,
					j: j,
				})
			}
		}
		matrix = append(matrix, row)
		i++
	}

	var result1, result2 int = 0, 0
	for _, p := range trailHeads {
		uniqueHeads := make(map[pos]bool)
		result2 += dfs(matrix, p, uniqueHeads)
		result1 += len(uniqueHeads)
	}
	fmt.Println("part1:", result1)
	fmt.Println("part2:", result2)
}

func dfs(matrix [][]int, src *pos, uniqueHeads map[pos]bool) int {
	if matrix[src.i][src.j] == 9 {
		uniqueHeads[pos{i: src.i, j: src.j}] = true
		return 1
	}

	temp := matrix[src.i][src.j]
	matrix[src.i][src.j] = -100

	result := 0

	for _, dir := range dirs {
		i, j := src.i+dir[0], src.j+dir[1]
		if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && matrix[i][j] == temp+1 {
			result += dfs(matrix, &pos{i: i, j: j}, uniqueHeads)
		}
	}

	matrix[src.i][src.j] = temp
	return result
}
