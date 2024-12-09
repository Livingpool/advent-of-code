package day9

import (
	"fmt"
	"os"
	"strconv"
)

func Part2() {
	input, err := os.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}

	var inputStr string = string(input)

	var disk = make([]int, 0, len(input))
	var fileMap = make([]int, 0, len(input))
	var id int = 0

	for i, b := range inputStr {
		num, _ := strconv.Atoi(string(b))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk = append(disk, id)
			}
			fileMap = append(fileMap, num)
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
			rightLen := fileMap[disk[r]]

			l2, r2 := l, r
			leftLen := 0
			for l2 <= r-rightLen {
				if disk[l2] == -1 {
					l2++
					leftLen++
					if leftLen >= rightLen {
						l2 -= leftLen
						for range rightLen {
							disk[l2] = disk[r2]
							disk[r2] = -1
							l2++
							r2--
						}
						break
					}
				} else {
					leftLen = 0
					l2++
				}
			}
			r -= rightLen
		} else {
			l++
		}
	}

	var checkSum int = 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checkSum += i * disk[i]
		}
	}
	fmt.Println("part2:", checkSum)
}
