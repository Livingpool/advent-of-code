package day14

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var distances []int
	const timeLimit int = 2503
	var currTime int = 0
	var currDist int = 0
	var flying bool = true

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		speed, err := strconv.Atoi(tokens[3])
		if err != nil {
			panic(err)
		}

		flyTime, err := strconv.Atoi(tokens[6])
		if err != nil {
			panic(err)
		}

		restTime, err := strconv.Atoi(tokens[13])
		if err != nil {
			panic(err)
		}

		for {
			if flying {
				if currTime+flyTime <= timeLimit {
					currDist += speed * flyTime
					currTime += flyTime
					flying = false
				} else {
					currDist += speed * (timeLimit - currTime)
					break
				}
			} else {
				currTime += restTime
				flying = true
				if currTime >= timeLimit {
					break
				}
			}
		}

		distances = append(distances, currDist)
		currTime = 0
		currDist = 0
		flying = true
	}

	fmt.Println(slices.Max(distances))
}
