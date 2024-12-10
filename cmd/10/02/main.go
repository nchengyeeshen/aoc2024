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

	var (
		pathLengths = make(map[point]map[point]int)
		path        []point
		helper      func(r, c, prev int)
	)
	helper = func(r, c, prev int) {
		if r < 0 || r >= len(grid) {
			return
		}
		if c < 0 || c >= len(grid[0]) {
			return
		}

		curr := grid[r][c]

		if curr-prev != 1 {
			return
		}

		if curr == 9 {
			_, ok := pathLengths[point{r, c}]
			if !ok {
				pathLengths[point{r, c}] = make(map[point]int)
			}

			pathLengths[point{r, c}][path[0]]++
			return
		}

		path = append(path, point{r, c})
		defer func() {
			path = path[:len(path)-1]
		}()

		possible := []point{
			{r + 1, c},
			{r - 1, c},
			{r, c - 1},
			{r, c + 1},
		}
		for _, pt := range possible {
			if pt == path[len(path)-1] {
				continue
			}
			helper(pt.x, pt.y, curr)
		}
	}

	scores := make(map[point]int)
	for _, th := range trailheads {
		helper(th.x, th.y, -1)

		for _, connectedTrailheads := range pathLengths {
			for cth, count := range connectedTrailheads {
				if cth == th {
					scores[th] += count
				}
			}
		}
	}

	var total int
	for _, sc := range scores {
		total += sc
	}

	fmt.Println(total)
}

type point struct{ x, y int }

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
