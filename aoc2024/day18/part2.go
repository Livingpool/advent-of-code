package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	f, err := os.Open("day18/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	src, end := Position{0, 0}, Position{memX - 1, memY - 1}

	// we know from part1 that there is a path up to 1024 bytes
	for range numOfBytes {
		scanner.Scan()
		coordinates := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(coordinates[0]) // distance from the left
		x, _ := strconv.Atoi(coordinates[1]) // distance from the top
		grid[x][y] = 1
	}

	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(coordinates[0]) // distance from the left
		x, _ := strconv.Atoi(coordinates[1]) // distance from the top
		grid[x][y] = 1
		if dijkstra(src, end) == -1 {
			fmt.Printf("part2: %d, %d\n", y, x)
			break
		}
	}
}
