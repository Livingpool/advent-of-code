package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	f, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	grid := [1000][1000]int{}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if tokens[0] == "toggle" {
			coord := strings.Split(tokens[1], ",")
			x1, _ := strconv.Atoi(coord[0])
			y1, _ := strconv.Atoi(coord[1])

			coord = strings.Split(tokens[3], ",")
			x2, _ := strconv.Atoi(coord[0])
			y2, _ := strconv.Atoi(coord[1])

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					grid[i][j] += 2
				}
			}
		} else {
			coord := strings.Split(tokens[2], ",")
			x1, _ := strconv.Atoi(coord[0])
			y1, _ := strconv.Atoi(coord[1])

			coord = strings.Split(tokens[4], ",")
			x2, _ := strconv.Atoi(coord[0])
			y2, _ := strconv.Atoi(coord[1])

			if tokens[1] == "on" {
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						grid[i][j] += 1
					}
				}
			} else {
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						grid[i][j] -= 1
						if grid[i][j] < 0 {
							grid[i][j] = 0
						}
					}
				}
			}
		}
	}

	var result int = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			result += grid[i][j]
		}
	}
	fmt.Println(result)
}
