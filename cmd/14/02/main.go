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

	height, width := 103, 101

	for t := 1; ; t++ {
		robotPositions := make(map[[2]int]int)
		for _, in := range inputs {
			rx := in.px + in.vx*t
			ry := in.py + in.vy*t

			rx, ry = clip(rx, ry, height, width)
			robotPositions[[2]int{rx, ry}]++
		}

		picture := make([]byte, 0, height*width+height)
		for i := range height {
			for j := range width {
				if robotPositions[[2]int{i, j}] != 0 {
					picture = append(picture, 'x')
					continue
				}
				picture = append(picture, '.')
			}
			picture = append(picture, '\n')
		}
		fmt.Println("===")
		fmt.Println("t:", t)
		fmt.Println(string(picture))
	}
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

type inputLine struct {
	px, py, vx, vy int
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
