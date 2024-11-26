package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	grid := [1000][1000]bool{}

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
					grid[i][j] = !grid[i][j]
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
						grid[i][j] = true
					}
				}
			} else {
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						grid[i][j] = false
					}
				}
			}
		}
	}

	var result int = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				result++
			}
		}
	}
	fmt.Println(result)
}
