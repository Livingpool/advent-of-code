package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	var cityMap = make(map[string]int)

	f, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var totalVertices int = 0
	var edgeList [][]int

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		city1, city2 := tokens[0], tokens[2]
		dist, err := strconv.Atoi(tokens[4])
		if err != nil {
			panic(err)
		}

		if _, exists := cityMap[city1]; !exists {
			cityMap[city1] = totalVertices
			totalVertices++
		}
		if _, exists := cityMap[city2]; !exists {
			cityMap[city2] = totalVertices
			totalVertices++
		}

		edgeList = append(edgeList, []int{cityMap[city1], cityMap[city2], dist})
	}

	// build adjacency list
	var adjacencyMatrix = make([][]int, totalVertices)
	for i := range totalVertices {
		adjacencyMatrix[i] = make([]int, totalVertices)
		for j := range totalVertices {
			adjacencyMatrix[i][j] = math.MaxInt32
		}
	}

	for _, edge := range edgeList {
		c1, c2, dist := edge[0], edge[1], edge[2]
		adjacencyMatrix[c1][c2] = dist
		adjacencyMatrix[c2][c1] = dist
	}

	var result = math.MaxInt32
	for i := range totalVertices {
		set := make(map[int]bool)
		result = min(result, search(adjacencyMatrix, i, set, 1))
	}

	fmt.Println("Shortest distance:", result)
}

func search(adj [][]int, src int, set map[int]bool, nodeCount int) int {
	if nodeCount == len(adj) {
		return 0
	}

	set[src] = true
	var result int = math.MaxInt32
	for i := range adj[src] {
		_, exists := set[i]
		if !exists && adj[src][i] < math.MaxInt32 {
			result = min(result, adj[src][i]+search(adj, i, set, nodeCount+1))
		}
	}
	delete(set, src)
	return result
}
