package day2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day2/input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := strings.Split(line, "x")

		x, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(arr[2])
		if err != nil {
			panic(err)
		}

		result += calc([]int{x, y, z})
	}

	fmt.Println(result)
}

func calc(arr []int) int {
	sort.Ints(arr)
	total := arr[0] * arr[1]

	total += 2 * (arr[0]*arr[1] + arr[1]*arr[2] + arr[0]*arr[2])
	return total
}
