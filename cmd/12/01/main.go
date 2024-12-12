package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	inputPath := flag.Arg(0)
	contents := must(os.ReadFile(inputPath))

	var grid [][]string
	for _, row := range strings.Split(
		strings.TrimSpace(string(contents)),
		"\n",
	) {
		grid = append(grid, strings.Split(row, ""))
	}

	var (
		visited = make(map[point]struct{})
		total   float64
	)
	for r := range grid {
		for c := range grid[r] {
			if _, ok := visited[point{r, c}]; ok {
				continue
			}

			area, perimeter := compute(grid, point{r, c}, visited)
			total += float64(area) * float64(perimeter)
		}
	}

	fmt.Printf("%.0f\n", total)
}

func compute(grid [][]string, start point, visited map[point]struct{}) (int, int) {
	area, perimeter := 1, 4

	if _, ok := visited[start]; ok {
		return 0, 0
	}

	letter := grid[start.r][start.c]

	visited[start] = struct{}{}

	possible := []point{
		{start.r - 1, start.c},
		{start.r + 1, start.c},
		{start.r, start.c - 1},
		{start.r, start.c + 1},
	}
	for _, posPt := range possible {
		if posPt.r < 0 || posPt.r >= len(grid) || posPt.c < 0 || posPt.c >= len(grid[0]) {
			continue
		}
		if letter == grid[posPt.r][posPt.c] {
			resArea, resPeri := compute(grid, posPt, visited)
			area, perimeter = area+resArea, perimeter-1+resPeri
		}
	}

	return area, perimeter
}

type point struct{ r, c int }

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
