package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() {
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
			if i > 0 && i < lenX-1 && j > 0 && j < lenY-1 && matrix[i][j] == 'A' {
				count := 0
				if matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S' {
					count++
				} else if matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M' {
					count++
				}

				if matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' {
					count++
				} else if matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' {
					count++
				}

				if count == 2 {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}
