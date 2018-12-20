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

type cell = int
type row []cell
type field []row

func makeField(dim int) field {
	m := make(field, 0, dim)
	for y := 0; y < dim; y++ {
		m = append(m, make(row, dim))
	}
	return m
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

func (t *task) process() {
}

func (t *task) part1() int {
	return 1
}

func (t *task) part2() int {
	return 2
}

//
// tests
//

func verify(v, ex int) {
	if v != ex {
		log.Output(3, fmt.Sprint(v, "!=", ex))
		os.Exit(1)
	}
}

func test() {
	t0 := time.Now()
	log.SetPrefix("[test] ")
	log.SetFlags(log.Lshortfile)
	test1 := func(in string, ex int) {
		t := parse(in)
		t.process()
		verify(t.part1(), ex)
	}
	test1(``, 1)
	fmt.Println("tests passed", Black(time.Since(t0)).Bold())
}

func main() {
	test()
	// delete(ins, "github")
	delete(ins, "google")
	for i, in := range ins {
		fmt.Println(Brown(fmt.Sprint("=== for ", i, " ===")))
		var t0 time.Time
		var d time.Duration

		t0 = time.Now()
		t := parse(in)
		d = time.Since(t0)
		fmt.Println(Gray("parse:"), Black(d).Bold())

		t0 = time.Now()
		t.process()
		d = time.Since(t0)
		fmt.Println(Gray("process:"), Black(d).Bold())

		t0 = time.Now()
		v1 := t.part1()
		d = time.Since(t0)
		fmt.Println(Gray("part 1:"), Black(d).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := t.part2()
		d = time.Since(t0)
		fmt.Println(Gray("part 2:"), Black(d).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `1`,
	"google": `2`,
}
