package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var result int = 0

	for scanner.Scan() {
		equation := strings.Split(scanner.Text(), ": ")
		testValue, _ := strconv.Atoi(equation[0])

		operands := strings.Split(equation[1], " ")
		opdInts := make([]int, len(operands))
		for i := 0; i < len(operands); i++ {
			val, _ := strconv.Atoi(operands[i])
			opdInts[i] = val
		}

		if checkEquation(testValue, opdInts) {
			result += testValue
		}
	}
	fmt.Println("result:", result)
}

func checkEquation(val int, operands []int) bool {
	possibleVals := make([]int, 0, int(math.Pow(float64(len(operands)-1), 2)))
	possibleVals = append(possibleVals, operands[0])

	for i := 1; i < len(operands); i++ {
		temp := make([]int, len(possibleVals))
		copy(temp, possibleVals)
		for j := range temp {
			possibleVals[j] += operands[i]
			if temp[j]*operands[i] <= val {
				possibleVals = append(possibleVals, temp[j]*operands[i])
			}
		}
		// if possibleVals[0] > val {
		// 	return false
		// }
	}
	// fmt.Println(val, possibleVals)

	for i := range possibleVals {
		if val == possibleVals[i] {
			return true
		}
	}

	return false
}
