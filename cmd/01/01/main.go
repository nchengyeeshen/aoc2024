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

	var (
		leftList  []int
		rightList []int
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
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var totalDist int
	for i := 0; i < len(leftList); i++ {
		dist := leftList[i] - rightList[i]
		if dist < 0 {
			dist *= -1
		}
		totalDist += dist
	}

	fmt.Println("Total distance:", totalDist)
}

func must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
