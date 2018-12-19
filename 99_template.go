package main

import (
	"fmt"
	"log"
	"os"
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
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
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
		log.Output(2, fmt.Sprint(v, "!=", ex))
		os.Exit(1)
	}
}

func test() {
	log.SetPrefix("[test] ")
	log.SetFlags(log.Lshortfile)
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
		var t0 time.Time
		var t time.Duration

		t0 = time.Now()
		task := parse(in)
		t = time.Since(t0)
		fmt.Println(Gray("parse:"), Black(t).Bold())

		t0 = time.Now()
		task.process()
		t = time.Since(t0)
		fmt.Println(Gray("process:"), Black(t).Bold())

		t0 = time.Now()
		v1 := task.part1()
		t = time.Since(t0)
		fmt.Println(Gray("part 1:"), Black(t).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := task.part2()
		t = time.Since(t0)
		fmt.Println(Gray("part 2:"), Black(t).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`,
}
