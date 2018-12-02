package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f1(ss []string) int {
	n, _ := strconv.Atoi(ss[0])
	return n
}

func f2(ss []string) int {
	n, _ := strconv.Atoi(ss[0])
	return n
}

func main() {
	for i, in := range ins {
		fmt.Println("for", i)
		ss := strings.Fields(in)
		fmt.Println("f1", f1(ss))
		fmt.Println("f2", f2(ss))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`}
