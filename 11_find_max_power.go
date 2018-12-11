package main

import (
	"fmt"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

//
// Solution
//

type prep [][]int

func prepare(in int) prep {
	m := make(prep, 0, 300)
	for y := 1; y <= 300; y++ {
		r := make([]int, 0, 300)
		for x := 1; x <= 300; x++ {
			rid := x + 10
			p := (rid*y+in)*rid/100%10 - 5
			r = append(r, p)
		}
		m = append(m, r)
	}
	return m
}

func find(m prep, d int) [4]int {
	maxp := 0
	var out [4]int
	for y := 0; y <= 300-d; y++ {
		for x := 0; x <= 300-d; x++ {
			p := 0
			for dy := 0; dy < d; dy++ {
				for dx := 0; dx < d; dx++ {
					p += m[y+dy][x+dx]
				}
			}
			if maxp < p {
				maxp = p
				out = [4]int{x, y, d, p}
			}
		}
	}
	return out
}

func part1(m prep) string {
	p := find(m, 3)
	return fmt.Sprintf("%d,%d power %d", p[0]+1, p[1]+1, p[3])
}

func (sums *prep) add(m prep, d int) [4]int {
	maxp := 0
	var out [4]int
	for y := 0; y <= 300-d; y++ {
		for x := 0; x <= 300-d; x++ {
			p := 0
			for dy := 0; dy < d; dy++ {
				p += m[y+dy][x+d-1]
			}
			for dx := 0; dx < d-1; dx++ {
				p += m[y+d-1][x+dx]
			}
			p += (*sums)[y][x]
			(*sums)[y][x] = p
			if maxp < p {
				maxp = p
				out = [4]int{x, y, d, p}
			}
		}
	}
	return out
}

func part2(m prep) string {
	out := find(m, 1)
	sums := make(prep, 300)
	for i := 0; i < 300; i++ {
		sums[i] = make([]int, 300)
		copy(sums[i], m[i])
	}
	for d := 2; d <= 300; d++ {
		p := sums.add(m, d)
		if out[3] < p[3] {
			out = p
		}
	}
	return fmt.Sprintf("%d,%d,%d power %d", out[0]+1, out[1]+1, out[2], out[3])
}

//
// tests
//

/*
func verify(p prep, ex int) {
	v := part1(p)
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

func test() {
	// verify(prep{}, 1)
	fmt.Println("tests passed")
}
*/

func main() {
	// test()
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
