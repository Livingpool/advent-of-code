package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stat struct {
	speed    int
	flyTime  int
	restTime int
	currDist int
	points   int
}

func Part2() {
	f, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var stats []*Stat
	const timeLimit int = 2503

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

		stats = append(stats, &Stat{
			speed:    speed,
			flyTime:  flyTime,
			restTime: restTime,
		})
	}

	for t := 1; t <= timeLimit; t++ {
		leadingDist := 0
		for _, ptr := range stats {
			// this equation took me a lot of time lol
			if (t-1)%(ptr.flyTime+ptr.restTime) < ptr.flyTime {
				ptr.currDist += ptr.speed
			}
			if ptr.currDist > leadingDist {
				leadingDist = ptr.currDist
			}
		}

		for _, ptr := range stats {
			if ptr.currDist == leadingDist {
				ptr.points++
			}
		}
	}

	result := 0
	for _, ptr := range stats {
		if ptr.points > result {
			result = ptr.points
		}
	}
	fmt.Println(result)
}
