package main

import (
	"cmp"
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

	stones := strings.Split(
		strings.TrimSpace(string(contents)),
		" ",
	)

	stoneCounts := make(map[string]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for range 75 {
		stoneCounts = transform(stoneCounts)
	}

	var total int
	for _, count := range stoneCounts {
		total += count
	}
	fmt.Println(total)
}

func transform(counts map[string]int) map[string]int {
	result := make(map[string]int)
	for stone, count := range counts {
		if count == 0 {
			continue
		}

		switch {
		case stone == "0":
			result["1"] += count
		case len(stone)%2 == 0:
			left, right := stone[:len(stone)/2], stone[len(stone)/2:]
			right = cmp.Or(strings.TrimLeft(right, "0"), "0")

			result[left] += count
			result[right] += count
		default:
			v := strconv.Itoa(must(strconv.Atoi(stone)) * 2024)
			result[v] += count
		}
	}
	return result
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
