package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

//
// Solution
//

type prep struct {
	p  int
	s  string
	rs map[string]string
}

func prepare(in string) prep {
	ss := strings.SplitN(in, "\n", 3)
	init := ss[0][15:]
	p := prep{0, init, map[string]string{}}
	rs := strings.Split(ss[2], "\n")
	sort.Sort(sort.Reverse(sort.StringSlice(rs)))
	for _, s := range rs {
		r := strings.Split(s, " => ")
		p.rs[r[0]] = r[1]
	}
	return p
}

func step(p prep) prep {
	s := "...." + p.s + "...."
	o := make([]string, 0, len(s))
	for i := 0; i < len(p.s)+4; i++ {
		m := s[i : i+5]
		o = append(o, p.rs[m])
	}

	// totally optional trims. But they help to print nice string out.
	for o[len(o)-1] == "." {
		o = o[:len(o)-1]
	}
	trim := 0
	for o[0] == "." {
		o = o[1:]
		trim++
	}

	return prep{p.p - 2 + trim, strings.Join(o, ""), p.rs}
}

func sum(p prep) int {
	n := 0
	for j, r := range p.s {
		if r == '#' {
			n += j + p.p
		}
	}
	return n
}

func part1(p prep) int {
	for i := 1; i <= 20; i++ {
		p = step(p)
	}
	return sum(p)
}

func part2(p prep) int {
	n0 := sum(p)
	d0 := 0
	nn := 0
	for i := 1; ; i++ {
		p = step(p)
		n1 := sum(p)
		d := n1 - n0
		n0 = n1
		if d0 == d {
			nn++
		} else {
			nn = 0
			d0 = d
		}
		//_log(i, n1, d, p.p, p.s) // watch convergence
		if nn == 10 {
			return n1 + (5e10-i)*d
		}
	}
}

func main() {
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
	"github": `initial state: ###..###....####.###...#..#...##...#..#....#.##.##.#..#.#..##.#####..######....#....##..#...#...#.#

..#.# => #
###.# => .
#.#.# => .
.#.#. => .
##... => #
...## => .
.##.# => .
.#... => #
####. => #
....# => .
.##.. => #
.#### => #
..### => .
.###. => #
##### => #
..#.. => #
#..#. => .
###.. => #
#..## => #
##.## => #
##..# => .
.#..# => #
#.#.. => #
#.### => #
#.##. => #
..... => .
.#.## => #
#...# => .
...#. => #
..##. => #
##.#. => #
#.... => .`,
	"google": `initial state: #.##.##.##.##.......###..####..#....#...#.##...##.#.####...#..##..###...##.#..#.##.#.#.#.#..####..#

..### => .
##..# => #
#..## => .
.#..# => .
#.##. => .
#.... => .
##... => #
#...# => .
###.# => #
##.## => .
....# => .
..##. => #
..#.. => .
##.#. => .
.##.# => #
#..#. => #
.##.. => #
###.. => #
.###. => #
##### => #
####. => .
.#.#. => .
...#. => #
#.### => .
.#... => #
.#### => .
#.#.# => #
...## => .
..... => .
.#.## => #
..#.# => #
#.#.. => #`}
