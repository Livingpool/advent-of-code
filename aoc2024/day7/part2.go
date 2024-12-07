package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	target   int
	operands []int
}

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func concat(a, b int) int {
	res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return res
}

func Part2() {
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

		var eq *Equation = &Equation{
			target:   testValue,
			operands: opdInts,
		}

		if dfs(eq, 1, opdInts[0], []func(a, b int) int{add, mul, concat}) {
			result += testValue
		}
	}
	fmt.Println("result:", result)
}

func dfs(equation *Equation, i int, partial int, ops []func(a, b int) int) bool {
	if partial > equation.target {
		return false
	} else if i == len(equation.operands) {
		return partial == equation.target
	} else {
		for _, op := range ops {
			if dfs(equation, i+1, op(partial, equation.operands[i]), ops) {
				return true
			}
		}
		return false
	}
}
