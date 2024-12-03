package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	inputPath := flag.Arg(0)

	contents, err := os.ReadFile(inputPath)
	must(err)

	// mul(###,###)
	pat := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	var (
		expressions = pat.FindAllString(string(contents), -1)
		total       int
	)
	for _, expr := range expressions {
		// Remove `mul(`
		expr = expr[4:]
		// Remove `)`
		expr = expr[:len(expr)-1]

		pair := strings.Split(expr, ",")

		a, err := strconv.Atoi(pair[0])
		must(err)

		b, err := strconv.Atoi(pair[1])
		must(err)

		total += a * b
	}

	fmt.Println("What do you get if you add up all of the results of the multiplications?", total)
}

func must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
