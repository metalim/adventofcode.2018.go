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
	ss := strings.Fields(in)
	fmt.Println("f1", f1(ss))
	fmt.Println("f2", f2(ss))
}

var in = `1`
