package day22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day22/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var result int

	for scanner.Scan() {
		x, _ := strconv.Atoi(strings.TrimSuffix(scanner.Text(), "\n"))
		for range 2000 {
			x = calcNextSecretNumber(x)
		}
		result += x
	}
	fmt.Println("part1:", result)
}

func calcNextSecretNumber(x int) int {
	step1 := ((x * 64) ^ x) % 16777216
	step2 := ((step1 / 32) ^ step1) % 16777216
	step3 := ((step2 * 2048) ^ step2) % 16777216
	return step3

}
