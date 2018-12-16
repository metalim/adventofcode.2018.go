package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

func _log1(a ...interface{}) {
	fmt.Println(a...)
	fmt.Scanln()
}

func sliceAtoi(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//
// Solution
//

type prep struct{}

func prepare(in string) prep {
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		s = s
	}
	return prep{}
}

func part1(p prep) int {
	return 1
}

func part2(p prep) int {
	return 2
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
		fmt.Println(Brown(fmt.Sprint("=== for ", i, " ===")))
		var t0, t1 time.Time

		t0 = time.Now()
		p := prepare(in)
		t1 = time.Now()
		fmt.Println(Gray("prepare:"), Black(t1.Sub(t0)).Bold())

		t0 = time.Now()
		v1 := part1(p)
		t1 = time.Now()
		fmt.Println(Gray("part 1:"), Black(t1.Sub(t0)).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := part2(p)
		t1 = time.Now()
		fmt.Println(Gray("part 2:"), Black(t1.Sub(t0)).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`}
