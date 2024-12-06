package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
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

	matrix[guard.x][guard.y] = '.'

	var origin = [2]int{guard.x, guard.y}

	var result, lenX, lenY int = 0, len(matrix), len(matrix[0])
	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if matrix[i][j] == '.' && (i != origin[0] || j != origin[1]) {
				matrix[i][j] = '#'

				loopDetection := map[pos]int{*guard: 1}

				for {
					if guard.facing == 1 {
						if isValidPos(guard.x-1, guard.y, lenX, lenY) {
							if matrix[guard.x-1][guard.y] == '.' {
								matrix[guard.x-1][guard.y] = 'X'
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

					// loop detection
					if _, exists := loopDetection[pos{
						x:      guard.x,
						y:      guard.y,
						facing: guard.facing,
					}]; exists {
						result++
						break
					} else {
						loopDetection[pos{
							x:      guard.x,
							y:      guard.y,
							facing: guard.facing,
						}] = 1
					}
				}

				// reset state
				matrix[i][j] = '.'
				for k := 0; k < lenX; k++ {
					for l := 0; l < lenY; l++ {
						if matrix[k][l] == 'X' {
							matrix[k][l] = '.'
						}
					}
				}
				guard.x, guard.y = origin[0], origin[1]
				guard.facing = 1
			}
		}
	}
	fmt.Println(result)
}
