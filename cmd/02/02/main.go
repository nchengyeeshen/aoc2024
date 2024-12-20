package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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
SCANNER:
	for scanner.Scan() {
		text := scanner.Text()
		report := strings.Fields(strings.TrimSpace(text))

		levels := make([]int, 0, len(report))
		for _, levelValue := range report {
			level, err := strconv.Atoi(levelValue)
			must(err)
			levels = append(levels, level)
		}

		// The whole report is safe.
		if safe(levels) {
			safeCount++
			continue SCANNER
		}

		// Problem dampener.
		for i := 0; i < len(levels); i++ {
			if safe(slices.Concat(levels[:i], levels[i+1:])) {
				safeCount++
				continue SCANNER
			}
		}
	}

	fmt.Println("How many reports are safe?", safeCount)
}

func safe(levels []int) bool {
	var posCount, negCount int
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

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
