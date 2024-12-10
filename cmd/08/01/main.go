package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	inputPath := flag.Arg(0)

	contents := strings.Split(
		string(
			bytes.TrimSpace(
				must(os.ReadFile(inputPath)),
			),
		),
		"\n",
	)

	var grid [][]string
	for _, line := range contents {
		grid = append(grid, strings.Split(line, ""))
	}

	antennas := make(map[string][]point)
	for r, row := range grid {
		for c, cell := range row {
			if cell != "." {
				pt := point{r, c}
				antennas[cell] = append(antennas[cell], pt)
			}
		}
	}

	antinodes := make(map[point]struct{})
	for _, locations := range antennas {
		for _, src := range locations {
			for _, dst := range locations {
				if src == dst {
					continue
				}

				diffX, diffY := dst.x-src.x, dst.y-src.y
				antinodePt := point{src.x - diffX, src.y - diffY}
				if !withinBounds(grid, antinodePt.x, antinodePt.y) {
					continue
				}

				antinodes[antinodePt] = struct{}{}
			}
		}
	}

	fmt.Println(len(antinodes))
}

type point struct{ x, y int }

func withinBounds(grid [][]string, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
