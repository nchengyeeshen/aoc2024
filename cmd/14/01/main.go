package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	contents := must(io.ReadAll(os.Stdin))

	var inputs []inputLine
	for line := range slices.Values(
		strings.Split(strings.TrimSpace(string(contents)), "\n"),
	) {
		before, after, _ := strings.Cut(line, " ")
		beforeSplit := strings.Split(before, ",")
		afterSplit := strings.Split(after, ",")
		inputs = append(inputs, inputLine{
			px: must(strconv.Atoi(beforeSplit[0][len("p="):])),
			py: must(strconv.Atoi(beforeSplit[1])),
			vx: must(strconv.Atoi(afterSplit[0][len("v="):])),
			vy: must(strconv.Atoi(afterSplit[1])),
		})
	}

	height, width, t := 103, 101, 100

	quadrantCount := make(map[int]int)
	for _, in := range inputs {
		rx := in.px + in.vx*t
		ry := in.py + in.vy*t

		rx, ry = clip(rx, ry, height, width)
		quadrantCount[quadrant(rx, ry, height, width)]++
	}

	total := 1.0
	for i := 1; i <= 4; i++ {
		total *= float64(quadrantCount[i])
	}

	fmt.Printf("%.0f\n", total)
}

func clip(x, y, height, width int) (int, int) {
	x %= width
	y %= height
	if x < 0 {
		x += width
	}
	if y < 0 {
		y += height
	}
	return x, y
}

func quadrant(x, y, height, width int) int {
	heightM, widthM := height/2, width/2
	if x == widthM || y == heightM {
		// Non-existent
		return 0
	}

	if x < widthM {
		if y < heightM {
			return 1
		}
		return 2
	}
	if y < heightM {
		return 3
	}
	return 4
}

type inputLine struct {
	px, py, vx, vy int
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
