package main

import (
	"flag"
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	inputPath := flag.Arg(0)
	contents := must(os.ReadFile(inputPath))

	var (
		blocks    []*int
		fileID    int
		freeSpace bool
	)
	for _, b := range strings.TrimSpace(string(contents)) {
		num := must(strconv.Atoi(string(b)))

		var v *int
		if !freeSpace {
			tmp := fileID
			v = &tmp
			fileID++
		}

		freeSpace = !freeSpace

		blocks = slices.AppendSeq(
			blocks,
			repeat(v, num),
		)
	}

	left, right := 0, len(blocks)-1

	for left < right {
		for blocks[right] == nil {
			right--
		}
		for blocks[left] != nil {
			left++
		}

		if left >= right {
			break
		}

		blocks[left], blocks[right] = blocks[right], blocks[left]
	}

	var total float64
	for i, num := range blocks {
		if num == nil {
			break
		}
		total += float64(i * *num)
	}
	fmt.Println(strconv.FormatFloat(total, 'f', 0, 64))
}

func repeat[T any](v T, times int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for times > 0 {
			if !yield(v) {
				return
			}
			times--
		}
	}
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
