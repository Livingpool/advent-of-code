package day20

import (
	"bufio"
	"fmt"
	"os"
)

// hint: https://www.reddit.com/r/adventofcode/comments/1hicdtb/comment/m2zgwf2/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
func Part2() {
	f, err := os.Open("day20/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var racetrack [][]rune
	var start Position

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
				row[j] = '.'
			}
		}
		racetrack = append(racetrack, row)
		i++
	}

	var path []Position
	var visited = make(map[Position]bool)
	dfs(start, racetrack, &path, visited)

	var minSave int = 100
	var cheats int = 0
	// var cheatsMap = make(map[int]int)

	for i := 0; i < len(path)-minSave; i++ { // candidates for cheat start
		for j := minSave; j < len(path); j++ { // candidates for cheat end
			manhattanDist := calcManhattanDistance(path[i], path[j])
			if manhattanDist <= 20 && (j-i)-manhattanDist >= minSave {
				cheats++
				// cheatsMap[j-i-manhattanDist]++
			}
		}
	}

	fmt.Println("part2:", cheats)

	// for k, v := range cheatsMap {
	// 	fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
	// }
}

func dfs(start Position, racetrack [][]rune, path *[]Position, visited map[Position]bool) {
	*path = append(*path, start)
	visited[start] = true

	for _, dir := range dirs {
		newPos := Position{start.i + dir[0], start.j + dir[1]}
		if newPos.isValid(racetrack) {
			if _, exists := visited[newPos]; !exists {
				dfs(newPos, racetrack, path, visited)
			}
		}
	}
}

func calcManhattanDistance(a, b Position) int {
	sum := 0
	if a.i-b.i >= 0 {
		sum += a.i - b.i
	} else {
		sum += b.i - a.i
	}

	if a.j-b.j >= 0 {
		sum += a.j - b.j
	} else {
		sum += b.j - a.j
	}

	return sum
}
