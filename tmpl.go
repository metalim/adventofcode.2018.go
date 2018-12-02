package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(ss []string) int {
	n, _ := strconv.Atoi(ss[0])
	return n
}

func part2(ss []string) int {
	n, _ := strconv.Atoi(ss[0])
	return n
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		ss := strings.Fields(in)
		fmt.Println("part 1:", part1(ss))
		fmt.Println("part 2:", part2(ss))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`}
