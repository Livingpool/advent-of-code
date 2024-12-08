package day8

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	i int
	j int
}

func Part1() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)

	var antennas = make(map[rune][]*Pos)
	var antinodes [][]int

	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for j, c := range line {
			if c != '.' {
				pos := Pos{
					i: i,
					j: j,
				}
				antennas[c] = append(antennas[c], &pos)
			}
		}

		antinodes = append(antinodes, row)
		i++
	}

	for _, v := range antennas {
		addAntinodes(antinodes, v)
	}

	var result int = 0
	for i := 0; i < len(antinodes); i++ {
		for j := 0; j < len(antinodes[0]); j++ {
			if antinodes[i][j] == 1 {
				result++
			}
		}
	}
	fmt.Println(result)
}

func addAntinodes(antinodes [][]int, antennas []*Pos) {
	lenI, lenJ := len(antinodes), len(antinodes[0])
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			i1 := 2*antennas[i].i - antennas[j].i
			j1 := 2*antennas[i].j - antennas[j].j
			if i1 >= 0 && i1 < lenI && j1 >= 0 && j1 < lenJ {
				antinodes[i1][j1] = 1
			}

			i2 := 2*antennas[j].i - antennas[i].i
			j2 := 2*antennas[j].j - antennas[i].j
			if i2 >= 0 && i2 < lenI && j2 >= 0 && j2 < lenJ {
				antinodes[i2][j2] = 1
			}
		}
	}
}
