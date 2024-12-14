package day14

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// cool! ref: https://www.reddit.com/r/adventofcode/comments/1hdw5op/2024_day_14_part_2_windows_explorer/
func Part2() {
	f, err := os.Open("day14/input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var bots []*robot
	var botArray = make([][]int, 0, height)

	for range height {
		row := make([]int, width)
		botArray = append(botArray, row)
	}

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
	}

	for i := range 10000 {
		for _, b := range bots {
			calcBotPositionPerSecond(b)
			botArray[b.py][b.px]++
		}
		if (i+1-23)%101 == 0 {
			createImage(bots, i+1)
		}
		clearArray(botArray)
	}
}

func calcBotPositionPerSecond(bot *robot) {
	changeX, changeY := bot.vx+bot.px, bot.vy+bot.py
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

func clearArray(arr [][]int) {
	for i := range len(arr) {
		for j := range len(arr[0]) {
			arr[i][j] = 0
		}
	}
}

func createImage(bots []*robot, secondsElapsed int) {
	img := image.NewGray(image.Rect(0, 0, width, height))

	file, err := os.Create(fmt.Sprintf("day14/img/%d.png", secondsElapsed))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// set base image as black
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.SetGray(j, i, color.Gray{Y: 0})
		}
	}

	// set bots as white
	for _, b := range bots {
		img.SetGray(b.px, b.py, color.Gray{Y: 255})
	}

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
