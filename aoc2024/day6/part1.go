package day6

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x      int
	y      int
	facing int
}

func Part1() {
	f, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	var matrix [][]rune
	var guard *pos

	var scanner = bufio.NewScanner(f)
	var i int = 0

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0, len(line))
		for j, t := range line {
			if t == '^' {
				guard = &pos{
					x:      i,
					y:      j,
					facing: 1,
				}
			}
			row = append(row, t)
		}
		matrix = append(matrix, row)
		i++
	}

	matrix[guard.x][guard.y] = 'X'

	var result, lenX, lenY int = 1, len(matrix), len(matrix[0])
	for {
		// guard tries to move
		if guard.facing == 1 {
			if isValidPos(guard.x-1, guard.y, lenX, lenY) {
				if matrix[guard.x-1][guard.y] == '.' {
					matrix[guard.x-1][guard.y] = 'X'
					result++
					guard.x--
				} else if matrix[guard.x-1][guard.y] == 'X' {
					guard.x--
				} else {
					guard.facing = 2
				}
			} else {
				break
			}
		} else if guard.facing == 2 {
			if isValidPos(guard.x, guard.y+1, lenX, lenY) {
				if matrix[guard.x][guard.y+1] == '.' {
					matrix[guard.x][guard.y+1] = 'X'
					result++
					guard.y++
				} else if matrix[guard.x][guard.y+1] == 'X' {
					guard.y++
				} else {
					guard.facing = 3
				}
			} else {
				break
			}
		} else if guard.facing == 3 {
			if isValidPos(guard.x+1, guard.y, lenX, lenY) {
				if matrix[guard.x+1][guard.y] == '.' {
					matrix[guard.x+1][guard.y] = 'X'
					result++
					guard.x++
				} else if matrix[guard.x+1][guard.y] == 'X' {
					guard.x++
				} else {
					guard.facing = 4
				}
			} else {
				break
			}
		} else if guard.facing == 4 {
			if isValidPos(guard.x, guard.y-1, lenX, lenY) {
				if matrix[guard.x][guard.y-1] == '.' {
					matrix[guard.x][guard.y-1] = 'X'
					result++
					guard.y--
				} else if matrix[guard.x][guard.y-1] == 'X' {
					guard.y--
				} else {
					guard.facing = 1
				}
			} else {
				break
			}
		}
		// fmt.Println(guard.x, guard.y)
	}
	fmt.Println(result)
}

func isValidPos(x, y, lenX, lenY int) bool {
	if x < 0 || y < 0 || x >= lenX || y >= lenY {
		return false
	}
	return true
}
