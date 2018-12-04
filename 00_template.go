package main

import (
	"fmt"
	"strconv"
	"strings"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

type prep struct{}

func prepare(ss []string) prep {
	for _, s := range ss {
		_log(s[0])
	}
	return prep{}
}

func part1(p prep) int {
	n, _ := strconv.Atoi("1")
	return n
}

func part2(p prep) int {
	n, _ := strconv.Atoi("2")
	return n
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		ss := strings.Split(in, "\n")
		p := prepare(ss)
		fmt.Println("part 1:", part1(p))
		fmt.Println("part 2:", part2(p))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`}
