package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var match = regexp.MustCompile("[0-9]+")

func main() {
	contents := must(io.ReadAll(os.Stdin))

	var blocks [][]int
	for block := range slices.Values(strings.Split(string(contents), "\n\n")) {
		var value []int
		for line := range slices.Values(strings.Split(block, "\n")) {
			digits := match.FindAllString(line, -1)
			for _, digit := range digits {
				value = append(value, must(strconv.Atoi(digit)))
			}
		}
		blocks = append(blocks, value)
	}

	var total float64
	for block := range slices.Values(blocks) {
		a, b := compute(block)
		if a == 0 && b == 0 {
			continue
		}
		total += float64(a*3) + float64(b)
	}
	fmt.Printf("%.0f\n", total)
}

func compute(block []int) (int, int) {
	assert(len(block) == 6, "block must be a slice of 6")

	a, b, c, d := block[0], block[2], block[1], block[3]

	determinant := a*d - b*c

	inverse := []int{d, -b, -c, a}

	x, y := block[4], block[5]

	ansA := x*inverse[0] + y*inverse[1]
	ansB := x*inverse[2] + y*inverse[3]
	if ansA%determinant != 0 && ansB%determinant != 0 {
		return 0, 0
	}

	return ansA / determinant, ansB / determinant
}

func assert(cond bool, msg string) {
	if cond {
		return
	}
	panic(msg)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
