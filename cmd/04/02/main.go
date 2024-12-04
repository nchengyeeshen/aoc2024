package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"slices"
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
			if lines[r][c] != 'A' {
				continue
			}

			var (
				coords = [][2]int{
					{r - 1, c - 1}, // Top left.
					{r - 1, c + 1}, // Top right.
					{r + 1, c - 1}, // Bottom left.
					{r + 1, c + 1}, // Bottom right.
				}
				valid = true
			)
			for _, coord := range coords {
				if coord[0] < 0 || coord[0] >= m || coord[1] < 0 || coord[1] >= n {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			buf := make([]byte, 0, 4)
			for _, coord := range coords {
				buf = append(buf, lines[coord[0]][coord[1]])
			}
			if slices.Contains([]string{"MMSS", "MSMS", "SSMM", "SMSM"}, string(buf)) {
				total++
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
