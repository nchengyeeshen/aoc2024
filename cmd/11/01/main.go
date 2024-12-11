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

	for range 25 {
		stones = transform(stones)
	}

	fmt.Println(len(stones))
}

func transform(stones []string) []string {
	var result []string
	for _, stone := range stones {
		switch {
		case stone == "0":
			result = append(result, "1")
		case len(stone)%2 == 0:
			left, right := stone[:len(stone)/2], stone[len(stone)/2:]
			right = strings.TrimLeft(right, "0")
			result = append(result, left, cmp.Or(right, "0"))
		default:
			v := must(strconv.Atoi(stone)) * 2024
			result = append(result, strconv.Itoa(v))
		}
	}
	return result
}

// func digits(v int) []int {
// 	var result []int
// 	for v > 0 {
// 		digit := v % 10
// 		result = append(result, digit)
// 		v /= 10
// 	}
// 	slices.Reverse(result)
// 	return result
// }

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
