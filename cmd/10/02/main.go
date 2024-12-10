package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	inputPath := flag.Arg(0)

	contents := must(os.ReadFile(inputPath))

	var grid [][]int
	for _, line := range strings.Split(
		strings.TrimSpace(string(contents)),
		"\n",
	) {
		var row []int
		for _, v := range strings.Split(line, "") {
			slope, err := strconv.Atoi(v)
			if err != nil {
				slope = 255
			}
			row = append(row, slope)
		}
		grid = append(grid, row)
	}

	var trailheads []point
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				trailheads = append(trailheads, point{r, c})
			}
		}
	}

	var total int
	for _, th := range trailheads {
		total += score(grid, th.x, th.y)
	}

	fmt.Println(total)
}

func score(grid [][]int, r, c int) int {
	var (
		summits int
		helper  func(r, c int)
	)

	helper = func(r, c int) {
		curr := grid[r][c]
		if curr == 9 {
			summits++
			return
		}
		possible := []point{
			{r + 1, c},
			{r - 1, c},
			{r, c - 1},
			{r, c + 1},
		}
		for _, pt := range possible {
			if pt.x < 0 || pt.x >= len(grid) || pt.y < 0 || pt.y >= len(grid[0]) {
				continue
			}
			if grid[pt.x][pt.y]-curr != 1 {
				continue
			}
			helper(pt.x, pt.y)
		}
	}

	helper(r, c)

	return summits
}

type point struct{ x, y int }

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
