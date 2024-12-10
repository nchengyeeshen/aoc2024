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
		blocks    []int
		fileID    int
		freeSpace bool

		fileRuns  [][2]int
		emptyRuns [][2]int
	)
	for _, b := range strings.TrimSpace(string(contents)) {
		num := must(strconv.Atoi(string(b)))

		var (
			v int = -1
			n     = len(blocks)
		)
		if !freeSpace {
			v = fileID
			fileID++
			fileRuns = append(fileRuns, [2]int{n, num})
		} else {
			emptyRuns = append(emptyRuns, [2]int{n, num})
		}

		freeSpace = !freeSpace

		blocks = slices.AppendSeq(
			blocks,
			repeat(v, num),
		)
	}

	for fri := len(fileRuns) - 1; fri >= 0; fri-- {
		fileStart, fileLen := fileRuns[fri][0], fileRuns[fri][1]

		for eri, er := range emptyRuns {
			emptyStart, emptyLen := er[0], er[1]
			if emptyLen >= fileLen && emptyStart < fileStart {
				i, j := emptyStart, fileStart+fileLen-1
				for j >= fileStart {
					blocks[i] = blocks[j]
					blocks[j] = -1
					i++
					j--
				}
				emptyRuns[eri] = [2]int{emptyStart + fileLen, emptyLen - fileLen}
				break
			}
		}
	}

	var total float64
	for i, num := range blocks {
		if num == -1 {
			continue
		}
		total += float64(i * num)
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
