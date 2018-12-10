package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

//
// Solution
//

type prep struct{}

func prepare(in string) prep {
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		_log(s)
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

//
// tests
//

func verify(p prep, ex int) {
	v := part1(p)
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

func test() {
	verify(prep{}, 1)
	fmt.Println("tests passed")
}

func main() {
	test()
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		p := prepare(in)
		fmt.Println("part 1:", part1(p))
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		fmt.Println("part 2:", part2(p))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`}
