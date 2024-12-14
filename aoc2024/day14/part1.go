package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	px int
	py int
	vx int
	vy int
}

var (
	width  = 101
	height = 103
)

func Part1() {
	f, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var bots []*robot
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		position := strings.Split(tokens[0][2:], ",")
		velocity := strings.Split(tokens[1][2:], ",")
		px, _ := strconv.Atoi(position[0])
		py, _ := strconv.Atoi(position[1])
		vx, _ := strconv.Atoi(velocity[0])
		vy, _ := strconv.Atoi(velocity[1])

		bot := robot{px: px, py: py, vx: vx, vy: vy}
		bots = append(bots, &bot)
		calcBotPosition(&bot)
	}

	fmt.Println("part1:", calcQuadrants(bots))
}

func calcBotPosition(bot *robot) {
	duration := 100

	changeX, changeY := duration*bot.vx+bot.px, duration*bot.vy+bot.py
	if changeX >= 0 {
		changeX = changeX % width
	} else {
		times := (-changeX) / width
		if (-changeX)%width != 0 {
			times++
		}
		changeX = changeX + times*width
	}
	if changeY >= 0 {
		changeY = changeY % height
	} else {
		times := (-changeY) / height
		if (-changeY)%height != 0 {
			times++
		}
		changeY = changeY + times*height
	}

	bot.px = changeX
	bot.py = changeY
}

func calcQuadrants(bots []*robot) int {
	w, h := width/2, height/2
	nextW, nextH := w, h
	if width%2 == 1 {
		nextW++
	}
	if nextH%2 == 1 {
		nextH++
	}

	quadrants := make([]int, 4)
	for _, b := range bots {
		if b.px >= 0 && b.px < w && b.py >= 0 && b.py < h {
			quadrants[0]++
		} else if b.px >= nextW && b.px < width && b.py >= 0 && b.py < h {
			quadrants[1]++
		} else if b.px >= 0 && b.px < w && b.py >= nextH && b.py < height {
			quadrants[2]++
		} else if b.px >= nextW && b.px < width && b.py >= nextH && b.py < height {
			quadrants[3]++
		}
	}

	count := 1
	for i := range quadrants {
		count *= quadrants[i]
	}

	return count
}
