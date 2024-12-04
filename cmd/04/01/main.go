package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	inputPath := flag.Arg(0)

	contents, err := os.ReadFile(inputPath)
	must(err)

	lines := bytes.Split(bytes.TrimSpace(contents), []byte("\n"))

	m, n := len(lines), len(lines[0])

	var total int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if lines[r][c] != 'X' {
				continue
			}

			for dr := -1; dr <= 1; dr++ { // [-1, 0, 1]
				for dc := -1; dc <= 1; dc++ { // [-1, 0, 1]
					if dc == 0 && dr == 0 {
						continue
					}

					newR := r + 3*dr
					if newR < 0 || newR >= m {
						continue
					}

					newC := c + 3*dc
					if newC < 0 || newC >= n {
						continue
					}

					if lines[r+dr][c+dc] != 'M' || lines[r+2*dr][c+2*dc] != 'A' || lines[r+3*dr][c+3*dc] != 'S' {
						continue
					}

					total++
				}
			}
		}
	}

	fmt.Println("How many times does XMAS appear?", total)
}

func must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
