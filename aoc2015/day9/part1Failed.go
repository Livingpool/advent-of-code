package day9

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type UnionFind struct {
	parent     []int
	size       []int
	components int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n+1)
	size := make([]int, n+1)
	for i := range n + 1 {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{
		parent:     parent,
		size:       size,
		components: n,
	}
}

func (u *UnionFind) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFind) Unite(x, y int) bool {
	rootX, rootY := u.Find(x), u.Find(y)
	if rootX == rootY {
		return false
	}

	if u.size[rootX] < u.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	u.parent[rootY] = rootX
	u.size[rootX] += u.size[rootY]
	u.components--

	return true
}

func (u *UnionFind) IsConnected() bool {
	return u.components == 1
}

// idea: kruskal's mst algo
// ah shit i realised this is not mst
func Part1Failed() {
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

	slices.SortFunc(edgeList, func(a, b []int) int {
		return cmp.Compare(a[2], b[2])
	})

	var res int = 0
	var u *UnionFind = NewUnionFind(totalVertices)

	for _, edge := range edgeList {
		x, y, dist := edge[0], edge[1], edge[2]
		// cycle detection
		if u.Find(x) != u.Find(y) {
			fmt.Println(x, y, dist)
			u.Unite(x, y)
			res += dist
		}
	}

	fmt.Println("Minimum spanning distance: ", res)
}
