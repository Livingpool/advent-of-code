package day23

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UnionFind struct {
	parent     []int
	size       []int
	components int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}
	size := make([]int, n)
	for i := range n {
		size[i] = 1
	}

	return &UnionFind{
		parent:     parent,
		size:       size,
		components: n,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// return false if noop; else true
func (uf *UnionFind) Unite(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}

	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}
	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

func (uf UnionFind) IsConnected() bool {
	return uf.components == 1
}

func Part1Failed() {
	f, err := os.Open("day23/input.txt")
	if err != nil {
		panic(err)
	}

	var scanner = bufio.NewScanner(f)
	var adj = make(map[string][]string)
	var cityToId = make(map[string]int)
	var idToCity = make(map[int]string)

	var id int = 0
	for scanner.Scan() {
		cities := strings.Split(strings.TrimSuffix(scanner.Text(), "\n"), "-")
		adj[cities[0]] = append(adj[cities[0]], cities[1])
		adj[cities[1]] = append(adj[cities[1]], cities[0])

		if _, exists := cityToId[cities[0]]; !exists {
			cityToId[cities[0]] = id
			idToCity[id] = cities[0]
			id++
		}
		if _, exists := cityToId[cities[1]]; !exists {
			cityToId[cities[1]] = id
			idToCity[id] = cities[1]
			id++
		}
	}

	uf := NewUnionFind(len(adj))
	for k, v := range adj {
		for _, neighbor := range v {
			uf.Unite(cityToId[k], cityToId[neighbor])
		}
	}

	var parents = make(map[int][]int)
	for i := range id {
		parents[uf.Find(i)] = append(parents[uf.Find(i)], i)
	}
	printComponents(parents, idToCity)

	var result int = 0
	for _, v := range parents {
		if len(v) >= 3 {
			tCount := 0
			for _, cityId := range v {
				if idToCity[cityId][:1] == "t" {
					tCount++
				}
			}
			result += calcComb(tCount, len(v)-tCount)
		}
	}
	fmt.Println("part1:", result)
}

func calcComb(t, r int) int {
	res := 0
	if t >= 1 && r >= 2 {
		res += t * (r * (r - 1) / 2)
	}
	if t >= 2 && r >= 1 {
		res += t * (t - 1) / 2 * r
	}
	if t >= 3 {
		res += t * (t - 1) * (t - 2) / 6
	}
	return res
}

func printComponents(parents map[int][]int, idToCity map[int]string) {
	for k, v := range parents {
		fmt.Printf("%s: ", idToCity[k])
		for i, cityId := range v {
			if i != len(v)-1 {
				fmt.Printf("%s, ", idToCity[cityId])
			} else {
				fmt.Printf("%s", idToCity[cityId])
			}
		}
		fmt.Println()
	}
}
