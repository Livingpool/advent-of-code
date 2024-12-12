package day12

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
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
				area, perimeter := calcPrice2(plots, visited, i, j)
				// fmt.Println(string(plots[i][j]), area, perimeter)
				result += area * perimeter
			}
		}
	}
	fmt.Println("part2:", result)
}

// dfs
// hint: https://www.reddit.com/r/adventofcode/comments/1hchskj/2024_day_12_its_been_fun/
func calcPrice2(plots [][]rune, visited [][]bool, i, j int) (int, int) {
	region := plots[i][j]
	area, corner := 1, 0
	lenX, lenY := len(plots), len(plots[0])

	visited[i][j] = true

	// outward corners
	if (!isValidPos(i+1, j, lenX, lenY) || plots[i+1][j] != region) && (!isValidPos(i, j+1, lenX, lenY) || plots[i][j+1] != region) {
		corner++
	}
	if (!isValidPos(i+1, j, lenX, lenY) || plots[i+1][j] != region) && (!isValidPos(i, j-1, lenX, lenY) || plots[i][j-1] != region) {
		corner++
	}
	if (!isValidPos(i-1, j, lenX, lenY) || plots[i-1][j] != region) && (!isValidPos(i, j+1, lenX, lenY) || plots[i][j+1] != region) {
		corner++
	}
	if (!isValidPos(i-1, j, lenX, lenY) || plots[i-1][j] != region) && (!isValidPos(i, j-1, lenX, lenY) || plots[i][j-1] != region) {
		corner++
	}

	// inward corners
	if (isValidPos(i+1, j, lenX, lenY) && plots[i+1][j] == region) && (isValidPos(i, j+1, lenX, lenY) && plots[i][j+1] == region) {
		if plots[i+1][j+1] != region {
			corner++
		}
	}
	if (isValidPos(i+1, j, lenX, lenY) && plots[i+1][j] == region) && (isValidPos(i, j-1, lenX, lenY) && plots[i][j-1] == region) {
		if plots[i+1][j-1] != region {
			corner++
		}
	}
	if (isValidPos(i-1, j, lenX, lenY) && plots[i-1][j] == region) && (isValidPos(i, j+1, lenX, lenY) && plots[i][j+1] == region) {
		if plots[i-1][j+1] != region {
			corner++
		}
	}
	if (isValidPos(i-1, j, lenX, lenY) && plots[i-1][j] == region) && (isValidPos(i, j-1, lenX, lenY) && plots[i][j-1] == region) {
		if plots[i-1][j-1] != region {
			corner++
		}
	}

	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if isValidPos(newI, newJ, lenX, lenY) && !visited[newI][newJ] && plots[newI][newJ] == region {
			a, c := calcPrice2(plots, visited, newI, newJ)
			area += a
			corner += c
		}
	}
	return area, corner
}

func isValidPos(i, j, maxI, maxJ int) bool {
	if i >= 0 && i < maxI && j >= 0 && j < maxJ {
		return true
	} else {
		return false
	}
}
