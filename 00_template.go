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

type def struct{}
type proc struct{}

func parse(in string) def {
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		s = s
	}
	return def{}
}

func process(def def) proc {
	return proc{}
}

func part1(def def, proc proc) int {
	return 1
}

func part2(def def, proc proc) int {
	return 2
}

//
// tests
//

func verify(v, ex int) {
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

func test() {
	def := parse(``)
	proc := process(def)
	verify(part1(def, proc), 1)
	// verify(part2(def, proc), 2)
	fmt.Println("tests passed")
}

func main() {
	test()
	delete(ins, "github")
	delete(ins, "google")
	for i, in := range ins {
		fmt.Println(Brown(fmt.Sprint("=== for ", i, " ===")))
		var t0, t1 time.Time

		t0 = time.Now()
		def := parse(in)
		t1 = time.Now()
		fmt.Println(Gray("parse:"), Black(t1.Sub(t0)).Bold())

		t0 = time.Now()
		proc := process(def)
		t1 = time.Now()
		fmt.Println(Gray("process:"), Black(t1.Sub(t0)).Bold())

		t0 = time.Now()
		v1 := part1(def, proc)
		t1 = time.Now()
		fmt.Println(Gray("part 1:"), Black(t1.Sub(t0)).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := part2(def, proc)
		t1 = time.Now()
		fmt.Println(Gray("part 2:"), Black(t1.Sub(t0)).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`,
}
