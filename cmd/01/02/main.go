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

	var (
		leftList   []int
		rightCount = make(map[int]int)
	)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Fields(strings.TrimSpace(text))

		left, err := strconv.Atoi(values[0])
		must(err)

		right, err := strconv.Atoi(values[1])
		must(err)

		leftList = append(leftList, left)
		rightCount[right]++
	}

	var totalDist int
	for _, num := range leftList {
		totalDist += num * rightCount[num]
	}

	fmt.Println("Total distance:", totalDist)
}

func must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
