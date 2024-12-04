package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	f, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var matrix []string

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		matrix = append(matrix, line)
	}

	var lenX, lenY int = len(matrix), len(matrix[0])
	var result int = 0

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if matrix[i][j] == 'X' {
				if j+3 < lenY {
					if matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
						result++
					}
					if i+3 < lenX && matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
						result++
					}
				}
				if i+3 < lenX {
					if matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
						result++
					}
					if j-3 >= 0 && matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
						result++
					}
				}
			} else if matrix[i][j] == 'S' {
				if j+3 < lenY {
					if matrix[i][j+1] == 'A' && matrix[i][j+2] == 'M' && matrix[i][j+3] == 'X' {
						result++
					}
					if i+3 < lenX && matrix[i+1][j+1] == 'A' && matrix[i+2][j+2] == 'M' && matrix[i+3][j+3] == 'X' {
						result++
					}
				}
				if i+3 < lenX {
					if matrix[i+1][j] == 'A' && matrix[i+2][j] == 'M' && matrix[i+3][j] == 'X' {
						result++
					}
					if j-3 >= 0 && matrix[i+1][j-1] == 'A' && matrix[i+2][j-2] == 'M' && matrix[i+3][j-3] == 'X' {
						result++
					}
				}
			}
		}
	}
	fmt.Println(result)
}
