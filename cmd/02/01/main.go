package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	inputPath := flag.Arg(0)

	inputFile, err := os.Open(inputPath)
	must(err)
	defer inputFile.Close()

	var safeCount int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		text := scanner.Text()
		report := strings.Fields(strings.TrimSpace(text))

		levels := make([]int, 0, len(report))
		for _, levelValue := range report {
			level, err := strconv.Atoi(levelValue)
			must(err)
			levels = append(levels, level)
		}

		if safe(levels) {
			safeCount++
		}
	}

	fmt.Println("How many reports are safe?", safeCount)
}

func safe(levels []int) bool {
	diffs := make([]int, 0, len(levels)-1)
	for i := 1; i < len(levels); i++ {
		diffs = append(diffs, levels[i]-levels[i-1])
	}

	var posCount, negCount int
	for _, diff := range diffs {
		// Differ by at least 1 and at most 3.
		if diff == 0 || (diff < -3 || diff > 3) {
			return false
		}

		if diff > 0 {
			posCount++
		} else {
			negCount++
		}

		if posCount > 0 && negCount > 0 {
			// All increasing or decreasing, but not both.
			return false
		}
	}

	return true
}

func must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
