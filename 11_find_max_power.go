package main

import (
	"fmt"
	"log"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

//
// Solution
//

type prep [][]int

//
// Inspired by hint given by https://github.com/petertseng
// Prepare Summed-area table, instead of just a power map.
// https://en.wikipedia.org/wiki/Summed-area_table
//
func prepare(in int) prep {
	m := make(prep, 0, 301)
	for y := 0; y <= 300; y++ {
		m = append(m, make([]int, 301))
	}
	for y := 1; y < 300; y++ {
		for x := 1; x < 300; x++ {
			rid := x + 10
			p := (rid*y+in)*rid/100%10 - 5
			m[y][x] = p + m[y][x-1] + m[y-1][x] - m[y-1][x-1]
		}
	}
	return m
}

func find(m prep, dMin, dMax int) [4]int {
	maxp := 0
	var out [4]int
	for d := dMin; d <= dMax; d++ {
		for y := 1; y < 300-d; y++ {
			for x := 1; x < 300-d; x++ {
				p := m[y+d-1][x+d-1] - m[y+d-1][x-1] - m[y-1][x+d-1] + m[y-1][x-1]
				if maxp < p {
					maxp = p
					out = [4]int{x, y, d, p}
				}
			}
		}
	}
	return out
}

func part1(m prep) string {
	p := find(m, 3, 3)
	return fmt.Sprintf("%d,%d power %d", p[0], p[1], p[3])
}

func part2(m prep) string {
	out := find(m, 1, 300)
	return fmt.Sprintf("%d,%d,%d power %d", out[0], out[1], out[2], out[3])
}

//
// tests
//

func verify(v, ex string) {
	if v != ex {
		log.Fatalf(`"%s" != "%s"`, v, ex)
	}
}

func test() {
	p := prepare(18)
	verify(part1(p), "33,45 power 29")
	verify(part2(p), "90,269,16 power 113")

	p = prepare(42)
	verify(part1(p), "21,61 power 30")
	verify(part2(p), "232,251,11 power 119")

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

var ins = map[string]int{
	"github": 9445,
	"google": 5177}
