package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	f, err := os.Open("day15/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var warehouse [][]rune
	var i, x, y int = 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 1 {
			break
		}
		row := make([]rune, len(line))
		for j, b := range line {
			row[j] = b
			if b == '@' {
				x = i
				y = j
			}
		}
		warehouse = append(warehouse, row)
		i++
	}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		for _, b := range line {
			x, y = moveRobot(warehouse, x, y, b)
		}
	}
	// printWarehouse(warehouse)

	var result int = 0
	for i := 1; i < len(warehouse)-1; i++ {
		for j := 0; j < len(warehouse[0])-1; j++ {
			if warehouse[i][j] == 'O' {
				result += 100*i + j
			}
		}
	}
	fmt.Println("part1:", result)
}

func moveRobot(warehouse [][]rune, x, y int, move rune) (int, int) {
	switch move {
	case '^':
		pos := x - 1
		if warehouse[pos][y] == '.' {
			warehouse[pos][y] = '@'
			warehouse[x][y] = '.'
			return x - 1, y
		} else if warehouse[pos][y] == '#' {
			return x, y
		}

		for i := pos; i >= 1; i-- {
			if warehouse[i][y] != 'O' {
				pos = i
				break
			}
		}
		if warehouse[pos][y] == '.' {
			warehouse[pos][y] = 'O'
			warehouse[x][y] = '.'
			warehouse[x-1][y] = '@'
			return x - 1, y
		}
		return x, y
	case '>':
		pos := y + 1
		if warehouse[x][pos] == '.' {
			warehouse[x][pos] = '@'
			warehouse[x][y] = '.'
			return x, y + 1
		} else if warehouse[pos][y] == '#' {
			return x, y
		}

		for j := pos; j <= len(warehouse[0])-2; j++ {
			if warehouse[x][j] != 'O' {
				pos = j
				break
			}
		}
		if warehouse[x][pos] == '.' {
			warehouse[x][pos] = 'O'
			warehouse[x][y] = '.'
			warehouse[x][y+1] = '@'
			return x, y + 1
		}
		return x, y
	case '<':
		pos := y - 1
		if warehouse[x][pos] == '.' {
			warehouse[x][pos] = '@'
			warehouse[x][y] = '.'
			return x, y - 1
		} else if warehouse[x][pos] == '#' {
			return x, y
		}

		for j := pos; j >= 1; j-- {
			if warehouse[x][j] != 'O' {
				pos = j
				break
			}
		}
		if warehouse[x][pos] == '.' {
			warehouse[x][pos] = 'O'
			warehouse[x][y] = '.'
			warehouse[x][y-1] = '@'
			return x, y - 1
		}
		return x, y
	case 'v':
		pos := x + 1
		if warehouse[pos][y] == '.' {
			warehouse[pos][y] = '@'
			warehouse[x][y] = '.'
			return x + 1, y
		} else if warehouse[pos][y] == '#' {
			return x, y
		}

		for i := pos; i <= len(warehouse)-2; i++ {
			if warehouse[i][y] != 'O' {
				pos = i
				break
			}
		}
		if warehouse[pos][y] == '.' {
			warehouse[pos][y] = 'O'
			warehouse[x][y] = '.'
			warehouse[x+1][y] = '@'
			return x + 1, y
		}
		return x, y
	}

	fmt.Println("unexpected move:", string(move))
	return -1, -1
}

func printWarehouse(warehouse [][]rune) {
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[0]); j++ {
			fmt.Printf("%c", warehouse[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
