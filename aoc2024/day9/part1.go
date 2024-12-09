package day9

import (
	"fmt"
	"os"
	"strconv"
)

func Part1() {
	input, err := os.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}

	var inputStr string = string(input)

	var disk = make([]int, 0, len(input))
	var id int = 0

	for i, b := range inputStr {
		num, _ := strconv.Atoi(string(b))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				disk = append(disk, -1)
			}
		}
	}

	var l, r int = 0, len(disk) - 1
	for l < r {
		if disk[r] == -1 {
			r--
		} else if disk[l] == -1 {
			disk[l] = disk[r]
			disk[r] = -1
			l++
			r--
		} else {
			l++
		}
	}

	var checkSum int = 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}
		// fmt.Println(i, "*", disk[i])
		checkSum += i * disk[i]
	}
	fmt.Println("part1:", checkSum)
}
