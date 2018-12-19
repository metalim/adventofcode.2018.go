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

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

//
// Solution
//

type task struct{}

func parse(in string) task {
	ss := strings.Split(in, "\n")
	for _, s := range ss {
		s = s
	}
	return task{}
}

func (task *task) process() {
}

func (task *task) part1() int {
	return 1
}

func (task *task) part2() int {
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
	task := parse(``)
	task.process()
	verify(task.part1(), 1)
	// verify(task.part2(), 2)
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
		task := parse(in)
		t1 = time.Now()
		fmt.Println(Gray("parse:"), Black(t1.Sub(t0)).Bold())

		t0 = time.Now()
		task.process()
		t1 = time.Now()
		fmt.Println(Gray("process:"), Black(t1.Sub(t0)).Bold())

		t0 = time.Now()
		v1 := task.part1()
		t1 = time.Now()
		fmt.Println(Gray("part 1:"), Black(t1.Sub(t0)).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := task.part2()
		t1 = time.Now()
		fmt.Println(Gray("part 2:"), Black(t1.Sub(t0)).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`,
}
