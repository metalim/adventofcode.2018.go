package main

import (
	"fmt"
	"regexp"
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

//
// Solution
//

type sample struct {
	b, c, a [4]int
}

type prep struct {
	ss []sample
	p  [][4]int
}

func prepare(in string) prep {
	p := prep{}
	ps := strings.Split(in, "\n\n\n\n")
	r := regexp.MustCompile("\\d+")

	p1 := strings.Split(ps[0], "\n\n")
	for _, tr := range p1 {
		ns := sliceAtoi(r.FindAllString(tr, -1))
		var s sample
		copy(s.b[:], ns[:4])
		copy(s.c[:], ns[4:8])
		copy(s.a[:], ns[8:])
		p.ss = append(p.ss, s)
	}

	p2 := strings.Split(ps[1], "\n")
	for _, l := range p2 {
		ns := sliceAtoi(r.FindAllString(l, -1))
		p.p = append(p.p, [4]int{ns[0], ns[1], ns[2], ns[3]})
	}

	return p
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func exec(reg [4]int, code, a, b, c int) [4]int {
	switch code {
	case 0: // addr
		reg[c] = reg[a] + reg[b]
	case 1: // addi
		reg[c] = reg[a] + b
	case 2: // mulr
		reg[c] = reg[a] * reg[b]
	case 3: // muli
		reg[c] = reg[a] * b // there was a typo here: + instead of *. FFFFUUUUU....
	case 4: // banr
		reg[c] = reg[a] & reg[b]
	case 5: // bani
		reg[c] = reg[a] & b
	case 6: // borr
		reg[c] = reg[a] | reg[b]
	case 7: // bori
		reg[c] = reg[a] | b
	case 8: // setr
		reg[c] = reg[a]
	case 9: // seti
		reg[c] = a
	case 10: // gtir
		reg[c] = b2i(a > reg[b])
	case 11: // gtri
		reg[c] = b2i(reg[a] > b)
	case 12: // gtrr
		reg[c] = b2i(reg[a] > reg[b])
	case 13: // eqir
		reg[c] = b2i(a == reg[b])
	case 14: // eqri
		reg[c] = b2i(reg[a] == b)
	case 15: // eqrr
		reg[c] = b2i(reg[a] == reg[b])
	}
	return reg
}

func part1(p prep) int {
	var n int
	for _, s := range p.ss {
		var m int
		for v := 0; v <= 15; v++ {
			r := exec(s.b, v, s.c[1], s.c[2], s.c[3])
			if r == s.a {
				m++
			}
		}
		if m >= 3 {
			n++
		}
	}
	return n
}

func part2(p prep) int {
	var bs [16][16]bool
	for _, s := range p.ss {
		k := s.c[0]
		for v := 0; v <= 15; v++ {
			reg := exec(s.b, v, s.c[1], s.c[2], s.c[3])
			if reg != s.a {
				bs[k][v] = true
			}
		}
	}

	ops := [16]int{}
	found := 0
	for found < 16 {
		for k := range bs {
			var notBad, v int
			for vv, bad := range bs[k] {
				if !bad {
					notBad++
					v = vv
				}
			}
			if notBad == 1 {
				found++
				ops[k] = v
				for k2 := 0; k2 < 16; k2++ {
					bs[k2][v] = true
				}
			}
		}
	}

	var reg [4]int
	for _, c := range p.p {
		reg = exec(reg, ops[c[0]], c[1], c[2], c[3])
	}
	return reg[0]
}

func main() {
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
	"github": `Before: [2, 3, 2, 2]
0 3 3 0
After:  [0, 3, 2, 2]

Before: [1, 1, 2, 3]
6 0 2 0
After:  [0, 1, 2, 3]

Before: [1, 0, 2, 2]
6 0 2 0
After:  [0, 0, 2, 2]

Before: [1, 1, 1, 1]
11 2 1 0
After:  [2, 1, 1, 1]

Before: [3, 0, 0, 2]
0 3 3 2
After:  [3, 0, 0, 2]

Before: [1, 1, 2, 2]
9 1 0 2
After:  [1, 1, 1, 2]

Before: [3, 2, 1, 1]
5 2 1 1
After:  [3, 2, 1, 1]

Before: [1, 1, 0, 3]
7 1 3 0
After:  [0, 1, 0, 3]

Before: [1, 2, 1, 3]
5 2 1 0
After:  [2, 2, 1, 3]

Before: [0, 2, 2, 0]
8 0 0 0
After:  [0, 2, 2, 0]

Before: [2, 0, 0, 1]
3 0 3 0
After:  [1, 0, 0, 1]

Before: [3, 1, 2, 2]
4 1 3 1
After:  [3, 0, 2, 2]

Before: [2, 2, 1, 1]
5 2 1 1
After:  [2, 2, 1, 1]

Before: [1, 1, 2, 2]
6 0 2 2
After:  [1, 1, 0, 2]

Before: [1, 1, 1, 2]
4 1 3 0
After:  [0, 1, 1, 2]

Before: [2, 1, 3, 1]
13 1 3 0
After:  [1, 1, 3, 1]

Before: [0, 1, 2, 1]
13 1 3 1
After:  [0, 1, 2, 1]

Before: [2, 1, 0, 2]
4 1 3 1
After:  [2, 0, 0, 2]

Before: [2, 1, 0, 1]
2 0 1 3
After:  [2, 1, 0, 1]

Before: [3, 1, 2, 1]
12 1 2 2
After:  [3, 1, 0, 1]

Before: [1, 1, 3, 2]
4 1 3 3
After:  [1, 1, 3, 0]

Before: [2, 2, 1, 3]
7 1 3 0
After:  [0, 2, 1, 3]

Before: [1, 3, 2, 1]
6 0 2 1
After:  [1, 0, 2, 1]

Before: [2, 1, 2, 1]
13 1 3 1
After:  [2, 1, 2, 1]

Before: [2, 1, 3, 0]
14 2 0 3
After:  [2, 1, 3, 1]

Before: [1, 1, 2, 3]
6 0 2 3
After:  [1, 1, 2, 0]

Before: [1, 1, 1, 3]
11 2 1 2
After:  [1, 1, 2, 3]

Before: [2, 2, 3, 2]
0 3 3 0
After:  [0, 2, 3, 2]

Before: [1, 2, 0, 2]
1 0 2 3
After:  [1, 2, 0, 0]

Before: [2, 1, 0, 0]
2 0 1 3
After:  [2, 1, 0, 1]

Before: [0, 2, 1, 1]
5 2 1 3
After:  [0, 2, 1, 2]

Before: [0, 3, 2, 1]
10 3 2 3
After:  [0, 3, 2, 1]

Before: [3, 3, 2, 2]
0 3 3 0
After:  [0, 3, 2, 2]

Before: [1, 1, 2, 0]
12 1 2 0
After:  [0, 1, 2, 0]

Before: [0, 2, 1, 3]
5 2 1 0
After:  [2, 2, 1, 3]

Before: [0, 3, 2, 1]
8 0 0 0
After:  [0, 3, 2, 1]

Before: [1, 1, 1, 3]
11 2 1 1
After:  [1, 2, 1, 3]

Before: [0, 1, 1, 2]
11 2 1 2
After:  [0, 1, 2, 2]

Before: [1, 1, 1, 1]
13 1 3 1
After:  [1, 1, 1, 1]

Before: [1, 3, 0, 0]
1 0 2 1
After:  [1, 0, 0, 0]

Before: [2, 2, 3, 1]
14 2 0 1
After:  [2, 1, 3, 1]

Before: [0, 3, 0, 3]
8 0 0 3
After:  [0, 3, 0, 0]

Before: [0, 0, 1, 1]
8 0 0 2
After:  [0, 0, 0, 1]

Before: [0, 3, 2, 1]
8 0 0 2
After:  [0, 3, 0, 1]

Before: [2, 1, 2, 3]
12 1 2 1
After:  [2, 0, 2, 3]

Before: [3, 2, 2, 3]
14 2 1 2
After:  [3, 2, 1, 3]

Before: [2, 2, 3, 0]
15 2 2 2
After:  [2, 2, 1, 0]

Before: [2, 3, 3, 2]
15 2 2 0
After:  [1, 3, 3, 2]

Before: [1, 1, 0, 0]
1 0 2 3
After:  [1, 1, 0, 0]

Before: [3, 2, 2, 2]
0 3 3 3
After:  [3, 2, 2, 0]

Before: [1, 3, 2, 2]
6 0 2 3
After:  [1, 3, 2, 0]

Before: [2, 1, 0, 1]
3 0 3 1
After:  [2, 1, 0, 1]

Before: [3, 3, 1, 3]
7 2 3 0
After:  [0, 3, 1, 3]

Before: [0, 2, 1, 0]
5 2 1 3
After:  [0, 2, 1, 2]

Before: [1, 1, 1, 2]
4 1 3 2
After:  [1, 1, 0, 2]

Before: [0, 3, 1, 2]
8 0 0 1
After:  [0, 0, 1, 2]

Before: [2, 1, 3, 3]
7 1 3 0
After:  [0, 1, 3, 3]

Before: [3, 2, 2, 1]
10 3 2 0
After:  [1, 2, 2, 1]

Before: [2, 1, 0, 1]
3 0 3 3
After:  [2, 1, 0, 1]

Before: [2, 1, 1, 1]
13 1 3 2
After:  [2, 1, 1, 1]

Before: [2, 2, 0, 3]
7 1 3 1
After:  [2, 0, 0, 3]

Before: [2, 2, 0, 1]
3 0 3 0
After:  [1, 2, 0, 1]

Before: [2, 2, 3, 1]
3 0 3 3
After:  [2, 2, 3, 1]

Before: [1, 2, 0, 0]
1 0 2 1
After:  [1, 0, 0, 0]

Before: [2, 2, 2, 2]
14 3 2 1
After:  [2, 0, 2, 2]

Before: [3, 1, 1, 2]
4 1 3 1
After:  [3, 0, 1, 2]

Before: [2, 1, 1, 1]
2 0 1 3
After:  [2, 1, 1, 1]

Before: [1, 1, 0, 0]
1 0 2 1
After:  [1, 0, 0, 0]

Before: [1, 3, 0, 2]
1 0 2 1
After:  [1, 0, 0, 2]

Before: [1, 1, 1, 3]
9 1 0 2
After:  [1, 1, 1, 3]

Before: [3, 1, 2, 2]
12 1 2 2
After:  [3, 1, 0, 2]

Before: [0, 1, 2, 1]
12 1 2 2
After:  [0, 1, 0, 1]

Before: [3, 2, 0, 3]
7 1 3 3
After:  [3, 2, 0, 0]

Before: [2, 1, 2, 3]
7 2 3 2
After:  [2, 1, 0, 3]

Before: [3, 1, 3, 1]
13 1 3 0
After:  [1, 1, 3, 1]

Before: [2, 1, 1, 1]
11 2 1 0
After:  [2, 1, 1, 1]

Before: [0, 1, 1, 0]
11 2 1 3
After:  [0, 1, 1, 2]

Before: [2, 1, 3, 3]
7 1 3 2
After:  [2, 1, 0, 3]

Before: [2, 3, 2, 1]
10 3 2 1
After:  [2, 1, 2, 1]

Before: [1, 1, 2, 2]
4 1 3 1
After:  [1, 0, 2, 2]

Before: [1, 3, 0, 1]
1 0 2 0
After:  [0, 3, 0, 1]

Before: [1, 3, 0, 3]
1 0 2 3
After:  [1, 3, 0, 0]

Before: [2, 3, 3, 1]
3 0 3 1
After:  [2, 1, 3, 1]

Before: [2, 1, 1, 2]
11 2 1 3
After:  [2, 1, 1, 2]

Before: [2, 1, 1, 1]
2 0 1 1
After:  [2, 1, 1, 1]

Before: [3, 1, 2, 2]
4 1 3 0
After:  [0, 1, 2, 2]

Before: [2, 0, 2, 1]
10 3 2 1
After:  [2, 1, 2, 1]

Before: [1, 3, 0, 1]
1 0 2 1
After:  [1, 0, 0, 1]

Before: [1, 1, 0, 2]
9 1 0 0
After:  [1, 1, 0, 2]

Before: [2, 3, 2, 1]
3 0 3 2
After:  [2, 3, 1, 1]

Before: [1, 2, 2, 1]
0 3 3 3
After:  [1, 2, 2, 0]

Before: [3, 1, 2, 2]
12 1 2 1
After:  [3, 0, 2, 2]

Before: [0, 2, 3, 1]
8 0 0 1
After:  [0, 0, 3, 1]

Before: [0, 0, 2, 1]
10 3 2 2
After:  [0, 0, 1, 1]

Before: [3, 2, 1, 3]
15 0 0 3
After:  [3, 2, 1, 1]

Before: [1, 3, 2, 2]
6 0 2 2
After:  [1, 3, 0, 2]

Before: [1, 2, 2, 3]
6 0 2 3
After:  [1, 2, 2, 0]

Before: [1, 1, 3, 2]
4 1 3 2
After:  [1, 1, 0, 2]

Before: [1, 2, 2, 1]
10 3 2 3
After:  [1, 2, 2, 1]

Before: [1, 2, 2, 1]
6 0 2 2
After:  [1, 2, 0, 1]

Before: [1, 2, 1, 3]
7 2 3 1
After:  [1, 0, 1, 3]

Before: [1, 2, 2, 1]
10 3 2 0
After:  [1, 2, 2, 1]

Before: [2, 3, 3, 1]
3 0 3 3
After:  [2, 3, 3, 1]

Before: [2, 3, 2, 3]
14 2 0 2
After:  [2, 3, 1, 3]

Before: [2, 1, 3, 1]
2 0 1 3
After:  [2, 1, 3, 1]

Before: [0, 3, 3, 0]
8 0 0 1
After:  [0, 0, 3, 0]

Before: [2, 1, 1, 3]
7 2 3 2
After:  [2, 1, 0, 3]

Before: [0, 2, 2, 1]
10 3 2 3
After:  [0, 2, 2, 1]

Before: [3, 2, 1, 3]
5 2 1 3
After:  [3, 2, 1, 2]

Before: [3, 1, 1, 2]
0 3 3 2
After:  [3, 1, 0, 2]

Before: [0, 3, 1, 3]
7 2 3 3
After:  [0, 3, 1, 0]

Before: [2, 0, 2, 1]
10 3 2 3
After:  [2, 0, 2, 1]

Before: [2, 2, 1, 0]
5 2 1 2
After:  [2, 2, 2, 0]

Before: [2, 1, 2, 2]
4 1 3 3
After:  [2, 1, 2, 0]

Before: [1, 3, 1, 1]
0 2 3 2
After:  [1, 3, 0, 1]

Before: [1, 1, 0, 3]
1 0 2 3
After:  [1, 1, 0, 0]

Before: [1, 0, 0, 3]
1 0 2 2
After:  [1, 0, 0, 3]

Before: [2, 1, 1, 0]
11 2 1 0
After:  [2, 1, 1, 0]

Before: [2, 0, 0, 1]
3 0 3 3
After:  [2, 0, 0, 1]

Before: [3, 3, 0, 1]
14 0 2 2
After:  [3, 3, 1, 1]

Before: [0, 1, 2, 0]
8 0 0 1
After:  [0, 0, 2, 0]

Before: [2, 0, 1, 1]
3 0 3 2
After:  [2, 0, 1, 1]

Before: [1, 3, 2, 0]
6 0 2 1
After:  [1, 0, 2, 0]

Before: [3, 3, 2, 0]
2 0 2 3
After:  [3, 3, 2, 1]

Before: [2, 1, 0, 1]
13 1 3 2
After:  [2, 1, 1, 1]

Before: [1, 1, 2, 1]
13 1 3 2
After:  [1, 1, 1, 1]

Before: [1, 3, 2, 0]
6 0 2 2
After:  [1, 3, 0, 0]

Before: [3, 1, 3, 2]
4 1 3 1
After:  [3, 0, 3, 2]

Before: [2, 3, 2, 2]
15 0 0 3
After:  [2, 3, 2, 1]

Before: [2, 3, 2, 1]
3 0 3 3
After:  [2, 3, 2, 1]

Before: [2, 1, 1, 2]
4 1 3 0
After:  [0, 1, 1, 2]

Before: [1, 1, 1, 1]
13 1 3 0
After:  [1, 1, 1, 1]

Before: [3, 1, 1, 0]
11 2 1 2
After:  [3, 1, 2, 0]

Before: [3, 1, 1, 1]
11 2 1 0
After:  [2, 1, 1, 1]

Before: [3, 1, 0, 2]
4 1 3 0
After:  [0, 1, 0, 2]

Before: [3, 3, 1, 3]
15 0 0 3
After:  [3, 3, 1, 1]

Before: [1, 2, 2, 1]
10 3 2 1
After:  [1, 1, 2, 1]

Before: [1, 1, 1, 0]
11 2 1 3
After:  [1, 1, 1, 2]

Before: [1, 1, 1, 2]
11 2 1 0
After:  [2, 1, 1, 2]

Before: [3, 2, 2, 2]
14 2 1 2
After:  [3, 2, 1, 2]

Before: [0, 0, 3, 3]
15 2 2 3
After:  [0, 0, 3, 1]

Before: [0, 3, 2, 2]
0 3 3 0
After:  [0, 3, 2, 2]

Before: [1, 0, 2, 1]
10 3 2 1
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 2]
14 3 2 1
After:  [2, 0, 2, 2]

Before: [1, 0, 0, 3]
1 0 2 1
After:  [1, 0, 0, 3]

Before: [3, 2, 1, 3]
7 2 3 1
After:  [3, 0, 1, 3]

Before: [3, 1, 1, 2]
11 2 1 0
After:  [2, 1, 1, 2]

Before: [1, 3, 2, 1]
6 0 2 0
After:  [0, 3, 2, 1]

Before: [2, 0, 3, 1]
3 0 3 0
After:  [1, 0, 3, 1]

Before: [3, 1, 2, 2]
12 1 2 0
After:  [0, 1, 2, 2]

Before: [3, 1, 2, 0]
12 1 2 3
After:  [3, 1, 2, 0]

Before: [2, 1, 2, 0]
2 0 1 3
After:  [2, 1, 2, 1]

Before: [1, 1, 3, 1]
14 2 3 2
After:  [1, 1, 0, 1]

Before: [1, 3, 2, 3]
6 0 2 0
After:  [0, 3, 2, 3]

Before: [1, 1, 2, 3]
12 1 2 0
After:  [0, 1, 2, 3]

Before: [3, 0, 2, 1]
10 3 2 1
After:  [3, 1, 2, 1]

Before: [1, 0, 2, 0]
6 0 2 1
After:  [1, 0, 2, 0]

Before: [2, 3, 1, 3]
7 2 3 2
After:  [2, 3, 0, 3]

Before: [1, 1, 1, 1]
11 2 1 3
After:  [1, 1, 1, 2]

Before: [2, 1, 2, 2]
2 0 1 0
After:  [1, 1, 2, 2]

Before: [1, 2, 1, 3]
7 2 3 3
After:  [1, 2, 1, 0]

Before: [1, 1, 2, 2]
12 1 2 0
After:  [0, 1, 2, 2]

Before: [2, 0, 2, 1]
10 3 2 2
After:  [2, 0, 1, 1]

Before: [0, 1, 2, 3]
12 1 2 2
After:  [0, 1, 0, 3]

Before: [2, 1, 1, 3]
11 2 1 0
After:  [2, 1, 1, 3]

Before: [2, 1, 3, 1]
13 1 3 3
After:  [2, 1, 3, 1]

Before: [0, 2, 1, 1]
8 0 0 1
After:  [0, 0, 1, 1]

Before: [1, 0, 0, 2]
1 0 2 1
After:  [1, 0, 0, 2]

Before: [2, 1, 3, 3]
2 0 1 1
After:  [2, 1, 3, 3]

Before: [0, 1, 2, 2]
4 1 3 2
After:  [0, 1, 0, 2]

Before: [1, 1, 2, 1]
13 1 3 0
After:  [1, 1, 2, 1]

Before: [1, 1, 3, 0]
9 1 0 1
After:  [1, 1, 3, 0]

Before: [1, 1, 0, 1]
1 0 2 1
After:  [1, 0, 0, 1]

Before: [2, 2, 3, 1]
3 0 3 1
After:  [2, 1, 3, 1]

Before: [3, 2, 1, 2]
5 2 1 0
After:  [2, 2, 1, 2]

Before: [1, 1, 2, 0]
12 1 2 1
After:  [1, 0, 2, 0]

Before: [3, 0, 2, 3]
2 0 2 3
After:  [3, 0, 2, 1]

Before: [2, 1, 3, 3]
2 0 1 2
After:  [2, 1, 1, 3]

Before: [3, 1, 3, 1]
15 0 0 0
After:  [1, 1, 3, 1]

Before: [0, 1, 3, 2]
4 1 3 1
After:  [0, 0, 3, 2]

Before: [3, 2, 3, 3]
15 2 0 0
After:  [1, 2, 3, 3]

Before: [1, 3, 3, 1]
0 3 3 0
After:  [0, 3, 3, 1]

Before: [0, 0, 2, 3]
7 2 3 0
After:  [0, 0, 2, 3]

Before: [0, 2, 1, 3]
7 2 3 2
After:  [0, 2, 0, 3]

Before: [3, 0, 2, 1]
2 0 2 0
After:  [1, 0, 2, 1]

Before: [2, 2, 2, 1]
10 3 2 2
After:  [2, 2, 1, 1]

Before: [1, 2, 0, 1]
1 0 2 0
After:  [0, 2, 0, 1]

Before: [1, 2, 0, 0]
1 0 2 2
After:  [1, 2, 0, 0]

Before: [3, 1, 2, 1]
2 0 2 1
After:  [3, 1, 2, 1]

Before: [0, 0, 3, 1]
8 0 0 1
After:  [0, 0, 3, 1]

Before: [0, 1, 1, 2]
11 2 1 3
After:  [0, 1, 1, 2]

Before: [0, 1, 3, 1]
13 1 3 2
After:  [0, 1, 1, 1]

Before: [1, 1, 1, 2]
11 2 1 2
After:  [1, 1, 2, 2]

Before: [2, 0, 3, 1]
3 0 3 3
After:  [2, 0, 3, 1]

Before: [0, 2, 1, 2]
8 0 0 0
After:  [0, 2, 1, 2]

Before: [1, 0, 2, 1]
6 0 2 1
After:  [1, 0, 2, 1]

Before: [1, 1, 0, 2]
4 1 3 3
After:  [1, 1, 0, 0]

Before: [2, 2, 1, 1]
3 0 3 2
After:  [2, 2, 1, 1]

Before: [1, 2, 1, 2]
5 2 1 2
After:  [1, 2, 2, 2]

Before: [2, 0, 2, 1]
3 0 3 3
After:  [2, 0, 2, 1]

Before: [2, 1, 0, 1]
3 0 3 2
After:  [2, 1, 1, 1]

Before: [2, 2, 1, 2]
5 2 1 1
After:  [2, 2, 1, 2]

Before: [1, 1, 2, 2]
9 1 0 3
After:  [1, 1, 2, 1]

Before: [2, 2, 1, 3]
15 0 0 3
After:  [2, 2, 1, 1]

Before: [3, 1, 0, 1]
13 1 3 3
After:  [3, 1, 0, 1]

Before: [3, 3, 2, 1]
10 3 2 2
After:  [3, 3, 1, 1]

Before: [0, 1, 3, 2]
4 1 3 3
After:  [0, 1, 3, 0]

Before: [0, 1, 1, 0]
11 2 1 2
After:  [0, 1, 2, 0]

Before: [3, 1, 3, 1]
14 3 1 0
After:  [0, 1, 3, 1]

Before: [0, 1, 3, 3]
8 0 0 3
After:  [0, 1, 3, 0]

Before: [0, 1, 2, 1]
10 3 2 0
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 1]
3 0 3 2
After:  [2, 1, 1, 1]

Before: [0, 2, 1, 3]
5 2 1 3
After:  [0, 2, 1, 2]

Before: [1, 0, 0, 3]
1 0 2 0
After:  [0, 0, 0, 3]

Before: [2, 3, 0, 1]
3 0 3 0
After:  [1, 3, 0, 1]

Before: [2, 1, 2, 1]
12 1 2 1
After:  [2, 0, 2, 1]

Before: [2, 1, 3, 2]
4 1 3 0
After:  [0, 1, 3, 2]

Before: [1, 2, 1, 0]
5 2 1 3
After:  [1, 2, 1, 2]

Before: [3, 1, 3, 1]
13 1 3 1
After:  [3, 1, 3, 1]

Before: [1, 2, 1, 0]
5 2 1 1
After:  [1, 2, 1, 0]

Before: [3, 1, 2, 1]
10 3 2 1
After:  [3, 1, 2, 1]

Before: [1, 1, 1, 1]
13 1 3 2
After:  [1, 1, 1, 1]

Before: [2, 1, 2, 1]
13 1 3 2
After:  [2, 1, 1, 1]

Before: [1, 2, 1, 3]
7 1 3 1
After:  [1, 0, 1, 3]

Before: [0, 0, 2, 2]
14 3 2 3
After:  [0, 0, 2, 0]

Before: [2, 2, 1, 3]
15 0 0 1
After:  [2, 1, 1, 3]

Before: [2, 1, 3, 2]
4 1 3 1
After:  [2, 0, 3, 2]

Before: [1, 2, 1, 3]
5 2 1 2
After:  [1, 2, 2, 3]

Before: [2, 2, 1, 0]
5 2 1 3
After:  [2, 2, 1, 2]

Before: [2, 0, 2, 1]
3 0 3 2
After:  [2, 0, 1, 1]

Before: [1, 0, 0, 1]
1 0 2 0
After:  [0, 0, 0, 1]

Before: [2, 1, 1, 0]
15 0 0 0
After:  [1, 1, 1, 0]

Before: [0, 0, 3, 3]
8 0 0 0
After:  [0, 0, 3, 3]

Before: [1, 1, 1, 2]
4 1 3 3
After:  [1, 1, 1, 0]

Before: [1, 2, 0, 3]
1 0 2 1
After:  [1, 0, 0, 3]

Before: [1, 1, 0, 2]
9 1 0 1
After:  [1, 1, 0, 2]

Before: [3, 1, 1, 1]
11 2 1 3
After:  [3, 1, 1, 2]

Before: [1, 1, 0, 3]
7 1 3 1
After:  [1, 0, 0, 3]

Before: [1, 1, 1, 3]
7 1 3 2
After:  [1, 1, 0, 3]

Before: [1, 1, 2, 3]
6 0 2 1
After:  [1, 0, 2, 3]

Before: [2, 1, 1, 2]
4 1 3 3
After:  [2, 1, 1, 0]

Before: [2, 2, 2, 3]
7 1 3 2
After:  [2, 2, 0, 3]

Before: [1, 3, 2, 1]
0 3 3 3
After:  [1, 3, 2, 0]

Before: [0, 0, 3, 3]
8 0 0 3
After:  [0, 0, 3, 0]

Before: [3, 1, 3, 1]
15 0 0 1
After:  [3, 1, 3, 1]

Before: [1, 0, 0, 2]
1 0 2 2
After:  [1, 0, 0, 2]

Before: [0, 0, 0, 1]
0 3 3 1
After:  [0, 0, 0, 1]

Before: [1, 1, 1, 2]
9 1 0 0
After:  [1, 1, 1, 2]

Before: [1, 3, 0, 1]
1 0 2 2
After:  [1, 3, 0, 1]

Before: [1, 1, 3, 3]
9 1 0 0
After:  [1, 1, 3, 3]

Before: [2, 1, 3, 1]
13 1 3 2
After:  [2, 1, 1, 1]

Before: [2, 1, 3, 2]
4 1 3 3
After:  [2, 1, 3, 0]

Before: [2, 1, 2, 1]
13 1 3 3
After:  [2, 1, 2, 1]

Before: [1, 0, 2, 2]
6 0 2 1
After:  [1, 0, 2, 2]

Before: [1, 1, 2, 1]
10 3 2 2
After:  [1, 1, 1, 1]

Before: [3, 2, 1, 3]
5 2 1 2
After:  [3, 2, 2, 3]

Before: [0, 1, 2, 0]
12 1 2 2
After:  [0, 1, 0, 0]

Before: [2, 1, 1, 3]
2 0 1 0
After:  [1, 1, 1, 3]

Before: [1, 2, 2, 3]
14 2 1 2
After:  [1, 2, 1, 3]

Before: [1, 2, 0, 3]
1 0 2 0
After:  [0, 2, 0, 3]

Before: [0, 1, 2, 2]
8 0 0 2
After:  [0, 1, 0, 2]

Before: [0, 2, 1, 0]
5 2 1 1
After:  [0, 2, 1, 0]

Before: [2, 0, 0, 1]
15 0 0 2
After:  [2, 0, 1, 1]

Before: [2, 2, 1, 3]
5 2 1 0
After:  [2, 2, 1, 3]

Before: [3, 2, 2, 1]
10 3 2 2
After:  [3, 2, 1, 1]

Before: [0, 3, 2, 2]
14 3 2 2
After:  [0, 3, 0, 2]

Before: [1, 2, 0, 1]
1 0 2 2
After:  [1, 2, 0, 1]

Before: [0, 1, 1, 0]
11 2 1 0
After:  [2, 1, 1, 0]

Before: [1, 2, 2, 3]
14 2 1 3
After:  [1, 2, 2, 1]

Before: [2, 1, 3, 1]
3 0 3 3
After:  [2, 1, 3, 1]

Before: [0, 1, 2, 3]
7 1 3 3
After:  [0, 1, 2, 0]

Before: [2, 1, 2, 2]
2 0 1 1
After:  [2, 1, 2, 2]

Before: [2, 2, 1, 0]
5 2 1 1
After:  [2, 2, 1, 0]

Before: [3, 2, 1, 3]
5 2 1 0
After:  [2, 2, 1, 3]

Before: [1, 1, 2, 1]
0 3 3 1
After:  [1, 0, 2, 1]

Before: [1, 0, 2, 1]
6 0 2 3
After:  [1, 0, 2, 0]

Before: [1, 3, 0, 2]
1 0 2 0
After:  [0, 3, 0, 2]

Before: [0, 1, 1, 3]
11 2 1 2
After:  [0, 1, 2, 3]

Before: [1, 1, 3, 3]
9 1 0 1
After:  [1, 1, 3, 3]

Before: [3, 1, 2, 3]
12 1 2 1
After:  [3, 0, 2, 3]

Before: [0, 1, 1, 1]
13 1 3 0
After:  [1, 1, 1, 1]

Before: [1, 1, 2, 3]
9 1 0 1
After:  [1, 1, 2, 3]

Before: [0, 3, 1, 3]
7 2 3 0
After:  [0, 3, 1, 3]

Before: [3, 1, 2, 1]
13 1 3 2
After:  [3, 1, 1, 1]

Before: [1, 0, 1, 3]
7 2 3 1
After:  [1, 0, 1, 3]

Before: [1, 1, 0, 3]
1 0 2 0
After:  [0, 1, 0, 3]

Before: [2, 1, 2, 2]
12 1 2 2
After:  [2, 1, 0, 2]

Before: [3, 0, 1, 3]
14 3 0 0
After:  [1, 0, 1, 3]

Before: [3, 1, 3, 3]
7 1 3 3
After:  [3, 1, 3, 0]

Before: [1, 1, 0, 0]
1 0 2 0
After:  [0, 1, 0, 0]

Before: [1, 1, 1, 1]
0 2 3 2
After:  [1, 1, 0, 1]

Before: [2, 1, 0, 1]
2 0 1 2
After:  [2, 1, 1, 1]

Before: [1, 1, 2, 1]
14 3 1 1
After:  [1, 0, 2, 1]

Before: [0, 0, 2, 3]
7 2 3 3
After:  [0, 0, 2, 0]

Before: [3, 2, 0, 0]
14 0 2 1
After:  [3, 1, 0, 0]

Before: [0, 0, 2, 3]
8 0 0 0
After:  [0, 0, 2, 3]

Before: [3, 1, 1, 0]
11 2 1 1
After:  [3, 2, 1, 0]

Before: [1, 2, 1, 1]
5 2 1 2
After:  [1, 2, 2, 1]

Before: [0, 2, 1, 3]
7 2 3 3
After:  [0, 2, 1, 0]

Before: [3, 1, 2, 2]
15 0 0 3
After:  [3, 1, 2, 1]

Before: [0, 0, 0, 2]
8 0 0 2
After:  [0, 0, 0, 2]

Before: [3, 1, 3, 1]
13 1 3 2
After:  [3, 1, 1, 1]

Before: [1, 1, 2, 3]
9 1 0 2
After:  [1, 1, 1, 3]

Before: [1, 2, 0, 2]
1 0 2 2
After:  [1, 2, 0, 2]

Before: [2, 1, 2, 3]
2 0 1 3
After:  [2, 1, 2, 1]

Before: [1, 2, 0, 3]
1 0 2 2
After:  [1, 2, 0, 3]

Before: [1, 0, 2, 0]
6 0 2 3
After:  [1, 0, 2, 0]

Before: [1, 0, 3, 1]
0 3 3 2
After:  [1, 0, 0, 1]

Before: [1, 3, 2, 1]
6 0 2 3
After:  [1, 3, 2, 0]

Before: [1, 1, 1, 1]
9 1 0 3
After:  [1, 1, 1, 1]

Before: [0, 3, 2, 1]
0 3 3 1
After:  [0, 0, 2, 1]

Before: [1, 1, 3, 1]
13 1 3 3
After:  [1, 1, 3, 1]

Before: [2, 2, 0, 3]
7 1 3 0
After:  [0, 2, 0, 3]

Before: [0, 3, 2, 1]
0 3 3 0
After:  [0, 3, 2, 1]

Before: [1, 0, 0, 1]
1 0 2 2
After:  [1, 0, 0, 1]

Before: [2, 1, 2, 1]
2 0 1 2
After:  [2, 1, 1, 1]

Before: [1, 2, 2, 2]
6 0 2 2
After:  [1, 2, 0, 2]

Before: [0, 1, 1, 1]
13 1 3 3
After:  [0, 1, 1, 1]

Before: [2, 1, 1, 0]
11 2 1 2
After:  [2, 1, 2, 0]

Before: [0, 1, 3, 1]
13 1 3 1
After:  [0, 1, 3, 1]

Before: [3, 2, 0, 2]
0 3 3 1
After:  [3, 0, 0, 2]

Before: [1, 1, 2, 1]
10 3 2 3
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 1]
13 1 3 0
After:  [1, 1, 2, 1]

Before: [2, 1, 0, 1]
13 1 3 1
After:  [2, 1, 0, 1]

Before: [2, 1, 2, 2]
12 1 2 3
After:  [2, 1, 2, 0]

Before: [0, 1, 2, 0]
12 1 2 1
After:  [0, 0, 2, 0]

Before: [3, 1, 2, 2]
4 1 3 2
After:  [3, 1, 0, 2]

Before: [1, 1, 0, 2]
1 0 2 1
After:  [1, 0, 0, 2]

Before: [0, 2, 1, 1]
0 2 3 2
After:  [0, 2, 0, 1]

Before: [1, 1, 2, 0]
6 0 2 0
After:  [0, 1, 2, 0]

Before: [0, 3, 1, 2]
8 0 0 3
After:  [0, 3, 1, 0]

Before: [1, 3, 0, 0]
1 0 2 2
After:  [1, 3, 0, 0]

Before: [1, 1, 2, 0]
12 1 2 2
After:  [1, 1, 0, 0]

Before: [2, 1, 0, 2]
0 3 3 1
After:  [2, 0, 0, 2]

Before: [0, 3, 3, 3]
8 0 0 1
After:  [0, 0, 3, 3]

Before: [3, 3, 0, 1]
0 3 3 0
After:  [0, 3, 0, 1]

Before: [3, 1, 1, 2]
4 1 3 3
After:  [3, 1, 1, 0]

Before: [2, 1, 2, 3]
12 1 2 3
After:  [2, 1, 2, 0]

Before: [3, 1, 2, 1]
12 1 2 3
After:  [3, 1, 2, 0]

Before: [1, 0, 2, 2]
6 0 2 3
After:  [1, 0, 2, 0]

Before: [1, 1, 0, 1]
0 3 3 1
After:  [1, 0, 0, 1]

Before: [1, 1, 0, 3]
9 1 0 2
After:  [1, 1, 1, 3]

Before: [3, 0, 2, 1]
10 3 2 3
After:  [3, 0, 2, 1]

Before: [2, 2, 3, 3]
14 3 2 3
After:  [2, 2, 3, 1]

Before: [3, 1, 2, 2]
12 1 2 3
After:  [3, 1, 2, 0]

Before: [0, 1, 2, 1]
10 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 1, 3, 0]
8 0 0 2
After:  [0, 1, 0, 0]

Before: [3, 1, 2, 0]
12 1 2 1
After:  [3, 0, 2, 0]

Before: [1, 3, 2, 0]
6 0 2 3
After:  [1, 3, 2, 0]

Before: [2, 0, 1, 3]
7 2 3 3
After:  [2, 0, 1, 0]

Before: [3, 2, 2, 1]
10 3 2 3
After:  [3, 2, 2, 1]

Before: [1, 2, 0, 0]
1 0 2 3
After:  [1, 2, 0, 0]

Before: [2, 1, 1, 1]
0 2 3 0
After:  [0, 1, 1, 1]

Before: [3, 2, 1, 1]
5 2 1 3
After:  [3, 2, 1, 2]

Before: [3, 1, 3, 1]
14 2 3 0
After:  [0, 1, 3, 1]

Before: [2, 1, 1, 3]
14 2 1 1
After:  [2, 0, 1, 3]

Before: [0, 1, 1, 2]
8 0 0 0
After:  [0, 1, 1, 2]

Before: [2, 3, 3, 2]
15 2 2 2
After:  [2, 3, 1, 2]

Before: [0, 1, 2, 3]
7 2 3 1
After:  [0, 0, 2, 3]

Before: [1, 1, 0, 2]
4 1 3 2
After:  [1, 1, 0, 2]

Before: [0, 2, 3, 0]
8 0 0 2
After:  [0, 2, 0, 0]

Before: [0, 1, 1, 1]
11 2 1 1
After:  [0, 2, 1, 1]

Before: [2, 1, 1, 1]
13 1 3 0
After:  [1, 1, 1, 1]

Before: [2, 3, 1, 3]
7 2 3 0
After:  [0, 3, 1, 3]

Before: [2, 1, 2, 3]
12 1 2 2
After:  [2, 1, 0, 3]

Before: [2, 2, 1, 3]
5 2 1 3
After:  [2, 2, 1, 2]

Before: [3, 1, 1, 3]
11 2 1 0
After:  [2, 1, 1, 3]

Before: [0, 0, 1, 3]
7 2 3 1
After:  [0, 0, 1, 3]

Before: [1, 3, 2, 1]
10 3 2 2
After:  [1, 3, 1, 1]

Before: [3, 2, 1, 2]
15 0 0 2
After:  [3, 2, 1, 2]

Before: [1, 2, 1, 1]
0 2 3 1
After:  [1, 0, 1, 1]

Before: [1, 1, 1, 3]
9 1 0 3
After:  [1, 1, 1, 1]

Before: [1, 1, 0, 3]
9 1 0 3
After:  [1, 1, 0, 1]

Before: [0, 1, 1, 1]
11 2 1 2
After:  [0, 1, 2, 1]

Before: [0, 1, 2, 1]
13 1 3 2
After:  [0, 1, 1, 1]

Before: [1, 1, 2, 2]
4 1 3 2
After:  [1, 1, 0, 2]

Before: [3, 1, 1, 2]
11 2 1 3
After:  [3, 1, 1, 2]

Before: [2, 2, 3, 2]
0 3 3 3
After:  [2, 2, 3, 0]

Before: [0, 0, 1, 1]
0 2 3 1
After:  [0, 0, 1, 1]

Before: [0, 1, 2, 2]
12 1 2 1
After:  [0, 0, 2, 2]

Before: [2, 0, 3, 1]
3 0 3 2
After:  [2, 0, 1, 1]

Before: [1, 0, 2, 0]
6 0 2 0
After:  [0, 0, 2, 0]

Before: [0, 2, 1, 1]
5 2 1 0
After:  [2, 2, 1, 1]

Before: [1, 3, 3, 0]
15 2 2 0
After:  [1, 3, 3, 0]

Before: [0, 3, 2, 0]
8 0 0 2
After:  [0, 3, 0, 0]

Before: [2, 2, 2, 1]
0 3 3 1
After:  [2, 0, 2, 1]

Before: [3, 1, 1, 2]
4 1 3 0
After:  [0, 1, 1, 2]

Before: [1, 2, 1, 0]
5 2 1 0
After:  [2, 2, 1, 0]

Before: [2, 2, 3, 3]
15 0 0 0
After:  [1, 2, 3, 3]

Before: [2, 1, 0, 0]
2 0 1 1
After:  [2, 1, 0, 0]

Before: [1, 2, 2, 3]
6 0 2 2
After:  [1, 2, 0, 3]

Before: [1, 0, 0, 1]
1 0 2 1
After:  [1, 0, 0, 1]

Before: [2, 2, 0, 1]
3 0 3 1
After:  [2, 1, 0, 1]

Before: [3, 2, 1, 2]
5 2 1 1
After:  [3, 2, 1, 2]

Before: [2, 1, 3, 2]
14 2 0 1
After:  [2, 1, 3, 2]

Before: [1, 1, 0, 0]
9 1 0 2
After:  [1, 1, 1, 0]

Before: [2, 2, 3, 3]
15 2 2 2
After:  [2, 2, 1, 3]

Before: [0, 2, 1, 0]
8 0 0 2
After:  [0, 2, 0, 0]

Before: [1, 1, 0, 1]
9 1 0 0
After:  [1, 1, 0, 1]

Before: [0, 1, 2, 2]
4 1 3 0
After:  [0, 1, 2, 2]

Before: [1, 1, 0, 0]
9 1 0 0
After:  [1, 1, 0, 0]

Before: [2, 3, 2, 1]
3 0 3 1
After:  [2, 1, 2, 1]

Before: [1, 2, 1, 3]
5 2 1 3
After:  [1, 2, 1, 2]

Before: [2, 1, 1, 3]
11 2 1 2
After:  [2, 1, 2, 3]

Before: [1, 1, 3, 0]
9 1 0 2
After:  [1, 1, 1, 0]

Before: [2, 1, 1, 3]
11 2 1 1
After:  [2, 2, 1, 3]

Before: [2, 1, 3, 2]
2 0 1 2
After:  [2, 1, 1, 2]

Before: [0, 2, 1, 3]
5 2 1 2
After:  [0, 2, 2, 3]

Before: [1, 0, 0, 2]
1 0 2 3
After:  [1, 0, 0, 0]

Before: [1, 1, 1, 2]
9 1 0 3
After:  [1, 1, 1, 1]

Before: [2, 1, 3, 2]
4 1 3 2
After:  [2, 1, 0, 2]

Before: [1, 0, 2, 2]
6 0 2 2
After:  [1, 0, 0, 2]

Before: [3, 1, 1, 3]
11 2 1 1
After:  [3, 2, 1, 3]

Before: [3, 1, 2, 3]
2 0 2 0
After:  [1, 1, 2, 3]

Before: [1, 2, 0, 2]
1 0 2 0
After:  [0, 2, 0, 2]

Before: [3, 1, 2, 1]
10 3 2 2
After:  [3, 1, 1, 1]

Before: [1, 0, 2, 3]
7 2 3 0
After:  [0, 0, 2, 3]

Before: [3, 1, 2, 3]
12 1 2 0
After:  [0, 1, 2, 3]

Before: [2, 1, 1, 3]
7 2 3 1
After:  [2, 0, 1, 3]

Before: [0, 2, 1, 2]
5 2 1 3
After:  [0, 2, 1, 2]

Before: [3, 1, 1, 0]
11 2 1 0
After:  [2, 1, 1, 0]

Before: [1, 1, 3, 1]
9 1 0 0
After:  [1, 1, 3, 1]

Before: [1, 1, 2, 2]
9 1 0 1
After:  [1, 1, 2, 2]

Before: [2, 1, 1, 3]
11 2 1 3
After:  [2, 1, 1, 2]

Before: [1, 1, 1, 2]
4 1 3 1
After:  [1, 0, 1, 2]

Before: [3, 1, 0, 1]
13 1 3 0
After:  [1, 1, 0, 1]

Before: [1, 2, 2, 3]
6 0 2 0
After:  [0, 2, 2, 3]

Before: [1, 3, 0, 3]
1 0 2 0
After:  [0, 3, 0, 3]

Before: [2, 1, 1, 0]
2 0 1 2
After:  [2, 1, 1, 0]

Before: [0, 1, 2, 1]
12 1 2 3
After:  [0, 1, 2, 0]

Before: [2, 3, 1, 1]
3 0 3 3
After:  [2, 3, 1, 1]

Before: [2, 1, 3, 3]
2 0 1 3
After:  [2, 1, 3, 1]

Before: [1, 3, 2, 1]
10 3 2 3
After:  [1, 3, 2, 1]

Before: [1, 1, 3, 3]
9 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 1, 3, 2]
9 1 0 1
After:  [1, 1, 3, 2]

Before: [1, 1, 0, 1]
13 1 3 2
After:  [1, 1, 1, 1]

Before: [3, 0, 2, 0]
2 0 2 1
After:  [3, 1, 2, 0]

Before: [2, 0, 0, 0]
14 0 1 2
After:  [2, 0, 1, 0]

Before: [0, 1, 2, 1]
13 1 3 3
After:  [0, 1, 2, 1]

Before: [2, 1, 3, 0]
14 2 0 1
After:  [2, 1, 3, 0]

Before: [2, 1, 0, 1]
13 1 3 0
After:  [1, 1, 0, 1]

Before: [2, 1, 0, 1]
2 0 1 1
After:  [2, 1, 0, 1]

Before: [0, 3, 2, 1]
10 3 2 0
After:  [1, 3, 2, 1]

Before: [0, 1, 3, 1]
0 3 3 2
After:  [0, 1, 0, 1]

Before: [0, 2, 1, 1]
5 2 1 1
After:  [0, 2, 1, 1]

Before: [2, 1, 1, 2]
15 0 0 3
After:  [2, 1, 1, 1]

Before: [1, 1, 2, 0]
6 0 2 2
After:  [1, 1, 0, 0]

Before: [1, 1, 2, 1]
6 0 2 0
After:  [0, 1, 2, 1]

Before: [0, 2, 1, 3]
7 1 3 1
After:  [0, 0, 1, 3]

Before: [1, 0, 0, 0]
1 0 2 3
After:  [1, 0, 0, 0]

Before: [2, 1, 2, 3]
2 0 1 2
After:  [2, 1, 1, 3]

Before: [0, 2, 0, 2]
0 3 3 1
After:  [0, 0, 0, 2]

Before: [0, 2, 3, 0]
15 2 2 2
After:  [0, 2, 1, 0]

Before: [1, 2, 2, 2]
14 2 1 3
After:  [1, 2, 2, 1]

Before: [0, 1, 3, 1]
8 0 0 2
After:  [0, 1, 0, 1]

Before: [3, 3, 3, 2]
15 0 0 3
After:  [3, 3, 3, 1]

Before: [3, 3, 0, 2]
14 0 2 1
After:  [3, 1, 0, 2]

Before: [0, 1, 1, 3]
11 2 1 0
After:  [2, 1, 1, 3]

Before: [1, 1, 0, 1]
9 1 0 2
After:  [1, 1, 1, 1]

Before: [0, 1, 2, 1]
10 3 2 3
After:  [0, 1, 2, 1]

Before: [2, 2, 2, 1]
10 3 2 1
After:  [2, 1, 2, 1]

Before: [0, 1, 2, 2]
4 1 3 3
After:  [0, 1, 2, 0]

Before: [1, 2, 2, 1]
10 3 2 2
After:  [1, 2, 1, 1]

Before: [2, 1, 1, 2]
11 2 1 1
After:  [2, 2, 1, 2]

Before: [1, 1, 2, 1]
12 1 2 3
After:  [1, 1, 2, 0]

Before: [3, 3, 1, 1]
0 2 3 1
After:  [3, 0, 1, 1]

Before: [0, 1, 2, 2]
4 1 3 1
After:  [0, 0, 2, 2]

Before: [0, 3, 2, 2]
8 0 0 3
After:  [0, 3, 2, 0]

Before: [2, 1, 2, 1]
2 0 1 0
After:  [1, 1, 2, 1]

Before: [1, 1, 0, 3]
1 0 2 1
After:  [1, 0, 0, 3]

Before: [3, 3, 3, 2]
15 0 0 0
After:  [1, 3, 3, 2]

Before: [0, 1, 1, 2]
4 1 3 2
After:  [0, 1, 0, 2]

Before: [1, 3, 0, 3]
1 0 2 1
After:  [1, 0, 0, 3]

Before: [1, 1, 0, 1]
1 0 2 2
After:  [1, 1, 0, 1]

Before: [2, 1, 0, 2]
4 1 3 0
After:  [0, 1, 0, 2]

Before: [3, 2, 2, 2]
2 0 2 2
After:  [3, 2, 1, 2]

Before: [0, 2, 2, 1]
10 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 1, 0, 2]
4 1 3 2
After:  [0, 1, 0, 2]

Before: [0, 1, 0, 2]
4 1 3 3
After:  [0, 1, 0, 0]

Before: [1, 1, 2, 1]
10 3 2 1
After:  [1, 1, 2, 1]

Before: [1, 1, 0, 1]
13 1 3 0
After:  [1, 1, 0, 1]

Before: [1, 3, 2, 2]
6 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 1, 2, 1]
13 1 3 0
After:  [1, 1, 2, 1]

Before: [0, 1, 1, 3]
11 2 1 1
After:  [0, 2, 1, 3]

Before: [3, 2, 1, 0]
5 2 1 3
After:  [3, 2, 1, 2]

Before: [2, 1, 2, 3]
7 2 3 3
After:  [2, 1, 2, 0]

Before: [1, 1, 1, 1]
11 2 1 2
After:  [1, 1, 2, 1]

Before: [2, 1, 1, 1]
3 0 3 2
After:  [2, 1, 1, 1]

Before: [0, 1, 1, 3]
8 0 0 1
After:  [0, 0, 1, 3]

Before: [3, 2, 3, 3]
7 1 3 3
After:  [3, 2, 3, 0]

Before: [0, 3, 0, 0]
8 0 0 2
After:  [0, 3, 0, 0]

Before: [1, 1, 2, 1]
6 0 2 1
After:  [1, 0, 2, 1]

Before: [0, 1, 1, 2]
4 1 3 0
After:  [0, 1, 1, 2]

Before: [1, 1, 2, 1]
9 1 0 1
After:  [1, 1, 2, 1]

Before: [3, 1, 2, 0]
12 1 2 0
After:  [0, 1, 2, 0]

Before: [1, 3, 0, 3]
1 0 2 2
After:  [1, 3, 0, 3]

Before: [1, 1, 0, 3]
9 1 0 1
After:  [1, 1, 0, 3]

Before: [0, 2, 2, 2]
8 0 0 1
After:  [0, 0, 2, 2]

Before: [0, 1, 1, 1]
13 1 3 1
After:  [0, 1, 1, 1]

Before: [1, 1, 3, 1]
13 1 3 0
After:  [1, 1, 3, 1]

Before: [0, 1, 2, 1]
8 0 0 0
After:  [0, 1, 2, 1]

Before: [2, 1, 2, 1]
12 1 2 2
After:  [2, 1, 0, 1]

Before: [1, 0, 2, 3]
6 0 2 1
After:  [1, 0, 2, 3]

Before: [3, 0, 3, 1]
15 2 0 2
After:  [3, 0, 1, 1]

Before: [0, 1, 1, 1]
0 2 3 0
After:  [0, 1, 1, 1]

Before: [3, 0, 0, 3]
14 0 2 1
After:  [3, 1, 0, 3]

Before: [3, 1, 1, 1]
0 2 3 1
After:  [3, 0, 1, 1]

Before: [0, 1, 2, 3]
7 2 3 3
After:  [0, 1, 2, 0]

Before: [3, 1, 0, 1]
13 1 3 1
After:  [3, 1, 0, 1]

Before: [0, 0, 3, 0]
8 0 0 1
After:  [0, 0, 3, 0]

Before: [1, 1, 0, 2]
1 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 1, 1, 2]
4 1 3 1
After:  [2, 0, 1, 2]

Before: [3, 2, 3, 0]
15 2 2 3
After:  [3, 2, 3, 1]

Before: [0, 2, 0, 3]
7 1 3 0
After:  [0, 2, 0, 3]

Before: [1, 1, 3, 2]
9 1 0 2
After:  [1, 1, 1, 2]

Before: [0, 3, 1, 3]
8 0 0 1
After:  [0, 0, 1, 3]

Before: [3, 1, 2, 1]
2 0 2 0
After:  [1, 1, 2, 1]

Before: [1, 1, 3, 1]
9 1 0 2
After:  [1, 1, 1, 1]

Before: [2, 1, 3, 0]
2 0 1 3
After:  [2, 1, 3, 1]

Before: [2, 1, 1, 0]
11 2 1 1
After:  [2, 2, 1, 0]

Before: [3, 1, 1, 1]
13 1 3 0
After:  [1, 1, 1, 1]

Before: [2, 2, 1, 3]
5 2 1 1
After:  [2, 2, 1, 3]

Before: [0, 0, 2, 1]
10 3 2 3
After:  [0, 0, 2, 1]

Before: [3, 3, 0, 2]
0 3 3 1
After:  [3, 0, 0, 2]

Before: [0, 2, 1, 0]
8 0 0 0
After:  [0, 2, 1, 0]

Before: [3, 3, 0, 2]
15 0 0 3
After:  [3, 3, 0, 1]

Before: [1, 0, 2, 3]
6 0 2 0
After:  [0, 0, 2, 3]

Before: [0, 0, 1, 1]
8 0 0 1
After:  [0, 0, 1, 1]

Before: [1, 0, 2, 1]
10 3 2 0
After:  [1, 0, 2, 1]

Before: [1, 2, 1, 2]
5 2 1 1
After:  [1, 2, 1, 2]

Before: [2, 1, 3, 1]
14 2 0 1
After:  [2, 1, 3, 1]

Before: [2, 1, 2, 0]
2 0 1 0
After:  [1, 1, 2, 0]

Before: [1, 1, 2, 2]
6 0 2 3
After:  [1, 1, 2, 0]

Before: [2, 1, 1, 3]
2 0 1 2
After:  [2, 1, 1, 3]

Before: [2, 3, 3, 2]
14 2 0 2
After:  [2, 3, 1, 2]

Before: [1, 0, 0, 2]
1 0 2 0
After:  [0, 0, 0, 2]

Before: [3, 3, 2, 2]
15 0 0 0
After:  [1, 3, 2, 2]

Before: [0, 1, 1, 2]
4 1 3 3
After:  [0, 1, 1, 0]

Before: [2, 2, 1, 2]
5 2 1 3
After:  [2, 2, 1, 2]

Before: [2, 1, 2, 0]
12 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 1, 0, 1]
13 1 3 2
After:  [3, 1, 1, 1]

Before: [1, 2, 1, 1]
5 2 1 1
After:  [1, 2, 1, 1]

Before: [2, 1, 2, 2]
4 1 3 2
After:  [2, 1, 0, 2]

Before: [0, 1, 0, 2]
4 1 3 0
After:  [0, 1, 0, 2]

Before: [3, 1, 0, 2]
4 1 3 2
After:  [3, 1, 0, 2]

Before: [1, 1, 3, 2]
4 1 3 1
After:  [1, 0, 3, 2]

Before: [3, 1, 1, 1]
13 1 3 2
After:  [3, 1, 1, 1]

Before: [0, 0, 2, 0]
8 0 0 3
After:  [0, 0, 2, 0]

Before: [1, 1, 3, 2]
9 1 0 0
After:  [1, 1, 3, 2]

Before: [3, 2, 1, 0]
5 2 1 1
After:  [3, 2, 1, 0]

Before: [1, 1, 0, 2]
1 0 2 0
After:  [0, 1, 0, 2]

Before: [2, 1, 0, 1]
13 1 3 3
After:  [2, 1, 0, 1]

Before: [3, 1, 2, 0]
12 1 2 2
After:  [3, 1, 0, 0]

Before: [3, 2, 2, 3]
2 0 2 0
After:  [1, 2, 2, 3]

Before: [1, 1, 1, 0]
11 2 1 1
After:  [1, 2, 1, 0]

Before: [0, 0, 1, 2]
8 0 0 3
After:  [0, 0, 1, 0]

Before: [1, 1, 0, 0]
9 1 0 3
After:  [1, 1, 0, 1]

Before: [1, 1, 3, 0]
9 1 0 3
After:  [1, 1, 3, 1]

Before: [1, 1, 1, 1]
11 2 1 1
After:  [1, 2, 1, 1]

Before: [3, 0, 0, 0]
14 0 2 3
After:  [3, 0, 0, 1]

Before: [2, 1, 1, 3]
7 1 3 3
After:  [2, 1, 1, 0]

Before: [0, 3, 3, 2]
8 0 0 2
After:  [0, 3, 0, 2]

Before: [3, 1, 2, 1]
12 1 2 1
After:  [3, 0, 2, 1]

Before: [3, 0, 2, 3]
7 2 3 0
After:  [0, 0, 2, 3]

Before: [3, 1, 1, 1]
14 3 1 1
After:  [3, 0, 1, 1]

Before: [1, 1, 1, 3]
9 1 0 0
After:  [1, 1, 1, 3]

Before: [0, 0, 3, 3]
8 0 0 2
After:  [0, 0, 0, 3]

Before: [3, 1, 3, 3]
7 1 3 1
After:  [3, 0, 3, 3]

Before: [1, 1, 2, 2]
12 1 2 1
After:  [1, 0, 2, 2]

Before: [1, 1, 0, 1]
1 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 2, 2, 1]
3 0 3 2
After:  [2, 2, 1, 1]

Before: [2, 0, 3, 0]
14 0 1 1
After:  [2, 1, 3, 0]

Before: [1, 1, 2, 2]
4 1 3 3
After:  [1, 1, 2, 0]

Before: [1, 1, 2, 3]
12 1 2 2
After:  [1, 1, 0, 3]

Before: [1, 2, 1, 3]
7 2 3 2
After:  [1, 2, 0, 3]

Before: [3, 0, 0, 1]
14 0 2 2
After:  [3, 0, 1, 1]

Before: [3, 2, 1, 0]
5 2 1 0
After:  [2, 2, 1, 0]

Before: [2, 3, 2, 1]
3 0 3 0
After:  [1, 3, 2, 1]

Before: [0, 1, 3, 2]
8 0 0 3
After:  [0, 1, 3, 0]

Before: [2, 2, 1, 1]
3 0 3 3
After:  [2, 2, 1, 1]

Before: [3, 2, 3, 1]
0 3 3 3
After:  [3, 2, 3, 0]

Before: [2, 1, 1, 0]
14 2 1 3
After:  [2, 1, 1, 0]

Before: [2, 2, 1, 3]
7 2 3 1
After:  [2, 0, 1, 3]

Before: [2, 3, 3, 1]
3 0 3 2
After:  [2, 3, 1, 1]

Before: [1, 1, 2, 1]
9 1 0 2
After:  [1, 1, 1, 1]

Before: [0, 3, 2, 1]
10 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 1, 0, 1]
13 1 3 3
After:  [0, 1, 0, 1]

Before: [1, 1, 1, 3]
11 2 1 3
After:  [1, 1, 1, 2]

Before: [3, 1, 1, 2]
11 2 1 2
After:  [3, 1, 2, 2]

Before: [1, 3, 2, 3]
6 0 2 3
After:  [1, 3, 2, 0]

Before: [0, 1, 2, 3]
8 0 0 2
After:  [0, 1, 0, 3]

Before: [3, 0, 1, 3]
14 3 0 2
After:  [3, 0, 1, 3]

Before: [2, 1, 2, 0]
12 1 2 3
After:  [2, 1, 2, 0]

Before: [0, 1, 1, 1]
11 2 1 0
After:  [2, 1, 1, 1]

Before: [2, 3, 2, 1]
0 3 3 2
After:  [2, 3, 0, 1]

Before: [1, 1, 0, 2]
0 3 3 3
After:  [1, 1, 0, 0]

Before: [1, 0, 0, 1]
1 0 2 3
After:  [1, 0, 0, 0]

Before: [3, 2, 1, 3]
7 2 3 3
After:  [3, 2, 1, 0]

Before: [3, 1, 1, 3]
11 2 1 2
After:  [3, 1, 2, 3]

Before: [0, 1, 2, 2]
12 1 2 3
After:  [0, 1, 2, 0]

Before: [3, 3, 2, 1]
10 3 2 0
After:  [1, 3, 2, 1]

Before: [1, 1, 3, 1]
13 1 3 1
After:  [1, 1, 3, 1]

Before: [2, 2, 1, 1]
3 0 3 1
After:  [2, 1, 1, 1]

Before: [2, 1, 2, 2]
4 1 3 0
After:  [0, 1, 2, 2]

Before: [1, 1, 1, 1]
9 1 0 2
After:  [1, 1, 1, 1]

Before: [1, 3, 2, 1]
10 3 2 0
After:  [1, 3, 2, 1]

Before: [2, 0, 2, 1]
10 3 2 0
After:  [1, 0, 2, 1]

Before: [1, 1, 0, 3]
1 0 2 2
After:  [1, 1, 0, 3]

Before: [1, 2, 0, 1]
1 0 2 3
After:  [1, 2, 0, 0]

Before: [1, 3, 0, 0]
1 0 2 0
After:  [0, 3, 0, 0]

Before: [2, 1, 1, 3]
14 2 1 0
After:  [0, 1, 1, 3]

Before: [1, 1, 1, 2]
9 1 0 1
After:  [1, 1, 1, 2]

Before: [1, 1, 0, 1]
13 1 3 1
After:  [1, 1, 0, 1]

Before: [2, 0, 0, 2]
15 0 0 0
After:  [1, 0, 0, 2]

Before: [2, 3, 1, 1]
3 0 3 0
After:  [1, 3, 1, 1]

Before: [0, 1, 2, 0]
12 1 2 3
After:  [0, 1, 2, 0]

Before: [1, 2, 1, 2]
5 2 1 0
After:  [2, 2, 1, 2]

Before: [2, 0, 2, 2]
14 3 2 2
After:  [2, 0, 0, 2]

Before: [0, 2, 2, 1]
10 3 2 0
After:  [1, 2, 2, 1]

Before: [2, 1, 0, 2]
4 1 3 3
After:  [2, 1, 0, 0]

Before: [1, 3, 0, 2]
1 0 2 2
After:  [1, 3, 0, 2]

Before: [0, 0, 2, 3]
8 0 0 1
After:  [0, 0, 2, 3]

Before: [2, 1, 1, 3]
7 1 3 0
After:  [0, 1, 1, 3]

Before: [3, 1, 2, 1]
13 1 3 0
After:  [1, 1, 2, 1]

Before: [2, 0, 1, 1]
3 0 3 1
After:  [2, 1, 1, 1]

Before: [1, 1, 2, 1]
13 1 3 1
After:  [1, 1, 2, 1]

Before: [0, 1, 2, 1]
12 1 2 1
After:  [0, 0, 2, 1]

Before: [2, 2, 3, 3]
14 3 2 2
After:  [2, 2, 1, 3]

Before: [3, 1, 1, 1]
13 1 3 3
After:  [3, 1, 1, 1]

Before: [3, 3, 3, 2]
15 2 0 1
After:  [3, 1, 3, 2]

Before: [2, 1, 2, 1]
3 0 3 1
After:  [2, 1, 2, 1]

Before: [3, 1, 2, 0]
2 0 2 3
After:  [3, 1, 2, 1]

Before: [1, 2, 1, 2]
5 2 1 3
After:  [1, 2, 1, 2]

Before: [3, 2, 1, 1]
5 2 1 0
After:  [2, 2, 1, 1]

Before: [0, 1, 2, 1]
12 1 2 0
After:  [0, 1, 2, 1]

Before: [2, 1, 1, 1]
3 0 3 0
After:  [1, 1, 1, 1]

Before: [3, 1, 1, 2]
11 2 1 1
After:  [3, 2, 1, 2]

Before: [1, 1, 1, 3]
11 2 1 0
After:  [2, 1, 1, 3]

Before: [1, 1, 2, 0]
9 1 0 3
After:  [1, 1, 2, 1]

Before: [0, 2, 2, 3]
8 0 0 3
After:  [0, 2, 2, 0]

Before: [0, 0, 2, 1]
10 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 2, 3, 3]
14 3 2 0
After:  [1, 2, 3, 3]

Before: [2, 1, 0, 3]
2 0 1 2
After:  [2, 1, 1, 3]

Before: [3, 1, 2, 0]
2 0 2 0
After:  [1, 1, 2, 0]

Before: [3, 1, 0, 2]
14 0 2 0
After:  [1, 1, 0, 2]

Before: [2, 1, 3, 0]
2 0 1 1
After:  [2, 1, 3, 0]

Before: [1, 1, 1, 0]
9 1 0 3
After:  [1, 1, 1, 1]

Before: [1, 0, 0, 0]
1 0 2 1
After:  [1, 0, 0, 0]

Before: [0, 3, 2, 2]
8 0 0 2
After:  [0, 3, 0, 2]

Before: [3, 3, 2, 2]
2 0 2 0
After:  [1, 3, 2, 2]

Before: [0, 2, 1, 2]
5 2 1 1
After:  [0, 2, 1, 2]

Before: [3, 3, 2, 2]
2 0 2 3
After:  [3, 3, 2, 1]

Before: [0, 2, 1, 2]
5 2 1 0
After:  [2, 2, 1, 2]

Before: [1, 0, 2, 1]
10 3 2 2
After:  [1, 0, 1, 1]

Before: [0, 1, 0, 1]
13 1 3 2
After:  [0, 1, 1, 1]

Before: [3, 1, 1, 1]
14 2 1 1
After:  [3, 0, 1, 1]

Before: [0, 1, 0, 1]
13 1 3 1
After:  [0, 1, 0, 1]

Before: [2, 2, 0, 1]
3 0 3 2
After:  [2, 2, 1, 1]

Before: [3, 2, 1, 3]
14 3 0 0
After:  [1, 2, 1, 3]

Before: [1, 1, 2, 2]
4 1 3 0
After:  [0, 1, 2, 2]

Before: [3, 1, 2, 3]
7 1 3 1
After:  [3, 0, 2, 3]

Before: [3, 0, 3, 0]
15 2 2 1
After:  [3, 1, 3, 0]

Before: [0, 2, 2, 2]
14 2 1 3
After:  [0, 2, 2, 1]

Before: [1, 1, 2, 3]
12 1 2 1
After:  [1, 0, 2, 3]

Before: [3, 1, 1, 1]
13 1 3 1
After:  [3, 1, 1, 1]

Before: [2, 1, 1, 1]
13 1 3 3
After:  [2, 1, 1, 1]

Before: [2, 2, 2, 3]
7 2 3 3
After:  [2, 2, 2, 0]

Before: [2, 3, 3, 3]
15 0 0 2
After:  [2, 3, 1, 3]

Before: [3, 1, 2, 1]
13 1 3 3
After:  [3, 1, 2, 1]

Before: [3, 3, 3, 2]
15 0 2 0
After:  [1, 3, 3, 2]

Before: [3, 1, 0, 2]
0 3 3 0
After:  [0, 1, 0, 2]

Before: [2, 0, 3, 2]
14 0 1 1
After:  [2, 1, 3, 2]

Before: [1, 0, 2, 1]
10 3 2 3
After:  [1, 0, 2, 1]

Before: [1, 3, 3, 1]
0 3 3 2
After:  [1, 3, 0, 1]

Before: [0, 2, 2, 1]
10 3 2 2
After:  [0, 2, 1, 1]

Before: [2, 2, 1, 0]
5 2 1 0
After:  [2, 2, 1, 0]

Before: [2, 3, 0, 1]
3 0 3 2
After:  [2, 3, 1, 1]

Before: [1, 2, 2, 2]
6 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 1, 2, 2]
12 1 2 0
After:  [0, 1, 2, 2]

Before: [1, 1, 0, 2]
9 1 0 2
After:  [1, 1, 1, 2]

Before: [0, 1, 2, 2]
12 1 2 2
After:  [0, 1, 0, 2]

Before: [2, 1, 2, 0]
12 1 2 2
After:  [2, 1, 0, 0]

Before: [2, 3, 3, 0]
15 0 0 3
After:  [2, 3, 3, 1]

Before: [2, 2, 0, 1]
15 0 0 0
After:  [1, 2, 0, 1]

Before: [2, 0, 3, 2]
0 3 3 2
After:  [2, 0, 0, 2]

Before: [3, 0, 3, 2]
15 2 2 3
After:  [3, 0, 3, 1]

Before: [2, 3, 2, 1]
10 3 2 3
After:  [2, 3, 2, 1]

Before: [2, 1, 2, 1]
3 0 3 3
After:  [2, 1, 2, 1]

Before: [1, 3, 0, 0]
1 0 2 3
After:  [1, 3, 0, 0]

Before: [3, 1, 2, 3]
2 0 2 3
After:  [3, 1, 2, 1]

Before: [2, 1, 1, 2]
11 2 1 2
After:  [2, 1, 2, 2]

Before: [1, 3, 2, 3]
7 2 3 1
After:  [1, 0, 2, 3]

Before: [0, 0, 0, 0]
8 0 0 3
After:  [0, 0, 0, 0]

Before: [1, 0, 3, 1]
14 2 3 2
After:  [1, 0, 0, 1]

Before: [3, 2, 0, 3]
14 0 2 3
After:  [3, 2, 0, 1]

Before: [3, 2, 2, 1]
2 0 2 1
After:  [3, 1, 2, 1]

Before: [2, 1, 2, 1]
3 0 3 0
After:  [1, 1, 2, 1]

Before: [2, 2, 0, 1]
3 0 3 3
After:  [2, 2, 0, 1]

Before: [0, 3, 3, 2]
8 0 0 0
After:  [0, 3, 3, 2]

Before: [3, 2, 0, 1]
14 0 2 1
After:  [3, 1, 0, 1]

Before: [1, 1, 1, 3]
9 1 0 1
After:  [1, 1, 1, 3]

Before: [0, 1, 0, 1]
13 1 3 0
After:  [1, 1, 0, 1]

Before: [1, 1, 1, 0]
9 1 0 1
After:  [1, 1, 1, 0]

Before: [1, 3, 2, 2]
6 0 2 0
After:  [0, 3, 2, 2]

Before: [2, 1, 1, 1]
14 3 1 0
After:  [0, 1, 1, 1]

Before: [1, 1, 3, 0]
9 1 0 0
After:  [1, 1, 3, 0]

Before: [2, 1, 3, 1]
3 0 3 2
After:  [2, 1, 1, 1]

Before: [2, 1, 1, 1]
3 0 3 1
After:  [2, 1, 1, 1]

Before: [3, 2, 1, 3]
7 1 3 3
After:  [3, 2, 1, 0]

Before: [2, 0, 3, 3]
15 0 0 2
After:  [2, 0, 1, 3]

Before: [3, 0, 2, 1]
10 3 2 2
After:  [3, 0, 1, 1]

Before: [1, 1, 2, 3]
9 1 0 0
After:  [1, 1, 2, 3]

Before: [1, 2, 1, 1]
5 2 1 0
After:  [2, 2, 1, 1]

Before: [0, 1, 2, 3]
12 1 2 1
After:  [0, 0, 2, 3]

Before: [1, 3, 0, 1]
1 0 2 3
After:  [1, 3, 0, 0]

Before: [2, 1, 0, 1]
2 0 1 0
After:  [1, 1, 0, 1]

Before: [3, 2, 2, 3]
2 0 2 1
After:  [3, 1, 2, 3]

Before: [1, 2, 0, 1]
1 0 2 1
After:  [1, 0, 0, 1]

Before: [1, 2, 2, 0]
6 0 2 0
After:  [0, 2, 2, 0]

Before: [2, 1, 1, 2]
11 2 1 0
After:  [2, 1, 1, 2]

Before: [3, 1, 1, 3]
7 1 3 2
After:  [3, 1, 0, 3]

Before: [2, 2, 1, 3]
5 2 1 2
After:  [2, 2, 2, 3]

Before: [3, 1, 1, 1]
11 2 1 1
After:  [3, 2, 1, 1]

Before: [2, 1, 2, 2]
4 1 3 1
After:  [2, 0, 2, 2]

Before: [1, 1, 2, 1]
12 1 2 0
After:  [0, 1, 2, 1]

Before: [1, 1, 0, 2]
9 1 0 3
After:  [1, 1, 0, 1]

Before: [3, 3, 2, 3]
2 0 2 0
After:  [1, 3, 2, 3]

Before: [1, 1, 2, 3]
9 1 0 3
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 1]
12 1 2 0
After:  [0, 1, 2, 1]

Before: [1, 1, 0, 2]
4 1 3 1
After:  [1, 0, 0, 2]

Before: [1, 2, 2, 0]
6 0 2 3
After:  [1, 2, 2, 0]

Before: [2, 1, 1, 0]
11 2 1 3
After:  [2, 1, 1, 2]

Before: [1, 1, 0, 1]
13 1 3 3
After:  [1, 1, 0, 1]

Before: [3, 1, 2, 3]
7 1 3 3
After:  [3, 1, 2, 0]

Before: [0, 2, 1, 3]
8 0 0 0
After:  [0, 2, 1, 3]

Before: [3, 2, 1, 3]
7 1 3 0
After:  [0, 2, 1, 3]

Before: [1, 2, 2, 2]
6 0 2 3
After:  [1, 2, 2, 0]

Before: [1, 1, 1, 1]
13 1 3 3
After:  [1, 1, 1, 1]

Before: [2, 1, 3, 2]
15 2 2 1
After:  [2, 1, 3, 2]

Before: [2, 1, 0, 3]
2 0 1 0
After:  [1, 1, 0, 3]

Before: [1, 1, 2, 1]
12 1 2 2
After:  [1, 1, 0, 1]

Before: [1, 1, 3, 2]
4 1 3 0
After:  [0, 1, 3, 2]

Before: [2, 3, 2, 3]
7 2 3 2
After:  [2, 3, 0, 3]

Before: [2, 2, 1, 1]
5 2 1 3
After:  [2, 2, 1, 2]

Before: [0, 0, 2, 1]
10 3 2 0
After:  [1, 0, 2, 1]

Before: [3, 1, 0, 3]
7 1 3 2
After:  [3, 1, 0, 3]

Before: [2, 1, 3, 2]
2 0 1 1
After:  [2, 1, 3, 2]

Before: [2, 3, 1, 1]
3 0 3 1
After:  [2, 1, 1, 1]

Before: [2, 2, 1, 3]
7 1 3 3
After:  [2, 2, 1, 0]

Before: [3, 3, 3, 1]
15 0 2 1
After:  [3, 1, 3, 1]

Before: [0, 1, 1, 0]
11 2 1 1
After:  [0, 2, 1, 0]

Before: [1, 1, 1, 0]
11 2 1 0
After:  [2, 1, 1, 0]

Before: [3, 1, 3, 1]
13 1 3 3
After:  [3, 1, 3, 1]

Before: [0, 1, 3, 2]
8 0 0 1
After:  [0, 0, 3, 2]

Before: [2, 2, 0, 3]
7 1 3 3
After:  [2, 2, 0, 0]

Before: [1, 0, 2, 1]
6 0 2 2
After:  [1, 0, 0, 1]

Before: [1, 3, 0, 2]
0 3 3 3
After:  [1, 3, 0, 0]

Before: [1, 1, 0, 1]
9 1 0 3
After:  [1, 1, 0, 1]

Before: [1, 2, 2, 3]
7 1 3 1
After:  [1, 0, 2, 3]

Before: [1, 1, 2, 2]
12 1 2 3
After:  [1, 1, 2, 0]

Before: [1, 1, 2, 0]
12 1 2 3
After:  [1, 1, 2, 0]

Before: [0, 1, 0, 2]
4 1 3 1
After:  [0, 0, 0, 2]

Before: [1, 1, 1, 0]
9 1 0 0
After:  [1, 1, 1, 0]

Before: [1, 1, 2, 0]
9 1 0 0
After:  [1, 1, 2, 0]

Before: [1, 2, 1, 1]
5 2 1 3
After:  [1, 2, 1, 2]

Before: [3, 0, 3, 2]
15 2 2 0
After:  [1, 0, 3, 2]

Before: [2, 2, 1, 3]
7 2 3 3
After:  [2, 2, 1, 0]

Before: [3, 1, 2, 2]
4 1 3 3
After:  [3, 1, 2, 0]

Before: [3, 1, 2, 1]
15 0 0 1
After:  [3, 1, 2, 1]

Before: [2, 3, 2, 1]
10 3 2 0
After:  [1, 3, 2, 1]

Before: [2, 1, 2, 2]
0 3 3 1
After:  [2, 0, 2, 2]

Before: [1, 2, 0, 2]
1 0 2 1
After:  [1, 0, 0, 2]

Before: [3, 3, 2, 0]
2 0 2 0
After:  [1, 3, 2, 0]

Before: [0, 1, 1, 2]
11 2 1 1
After:  [0, 2, 1, 2]

Before: [3, 1, 2, 1]
13 1 3 1
After:  [3, 1, 2, 1]

Before: [3, 1, 3, 3]
15 2 0 3
After:  [3, 1, 3, 1]

Before: [0, 1, 0, 1]
8 0 0 3
After:  [0, 1, 0, 0]

Before: [2, 3, 2, 1]
0 3 3 3
After:  [2, 3, 2, 0]

Before: [2, 1, 1, 2]
4 1 3 2
After:  [2, 1, 0, 2]

Before: [0, 1, 3, 1]
13 1 3 0
After:  [1, 1, 3, 1]

Before: [2, 2, 1, 1]
5 2 1 0
After:  [2, 2, 1, 1]

Before: [3, 1, 2, 0]
15 0 0 0
After:  [1, 1, 2, 0]

Before: [1, 1, 1, 1]
9 1 0 0
After:  [1, 1, 1, 1]

Before: [1, 1, 2, 2]
12 1 2 2
After:  [1, 1, 0, 2]

Before: [1, 1, 2, 1]
10 3 2 0
After:  [1, 1, 2, 1]

Before: [2, 0, 1, 1]
3 0 3 3
After:  [2, 0, 1, 1]



8 0 0 2
5 2 2 2
6 3 1 1
8 0 0 3
5 3 0 3
9 2 3 1
8 1 3 1
8 1 2 1
11 0 1 0
10 0 0 1
8 0 0 2
5 2 3 2
6 3 0 3
8 2 0 0
5 0 1 0
12 3 2 0
8 0 3 0
11 1 0 1
10 1 1 3
6 1 2 0
8 1 0 1
5 1 0 1
6 0 0 2
5 0 1 1
8 1 3 1
11 3 1 3
10 3 3 2
6 3 0 3
6 2 1 1
4 3 1 0
8 0 3 0
11 0 2 2
10 2 0 1
8 1 0 0
5 0 1 0
8 3 0 2
5 2 0 2
6 2 1 3
8 0 2 3
8 3 3 3
8 3 3 3
11 3 1 1
10 1 1 3
6 3 3 2
6 0 0 1
8 0 2 0
8 0 1 0
8 0 1 0
11 3 0 3
10 3 0 0
6 2 2 1
6 1 0 3
13 1 2 1
8 1 2 1
11 1 0 0
10 0 3 2
6 2 0 0
6 3 1 1
6 2 0 3
9 0 3 1
8 1 1 1
8 1 3 1
11 1 2 2
10 2 0 0
6 3 1 1
6 2 1 2
6 0 2 3
7 3 2 1
8 1 2 1
8 1 2 1
11 1 0 0
10 0 3 3
6 3 1 2
6 1 2 0
6 0 1 1
6 2 1 1
8 1 1 1
11 3 1 3
10 3 3 1
6 2 0 3
8 0 0 2
5 2 0 2
6 2 2 0
15 0 3 0
8 0 1 0
8 0 2 0
11 0 1 1
10 1 0 0
6 3 1 1
6 0 0 3
6 2 1 2
7 3 2 2
8 2 2 2
11 2 0 0
10 0 1 1
6 1 0 3
8 0 0 0
5 0 2 0
8 0 0 2
5 2 0 2
3 0 3 2
8 2 3 2
8 2 2 2
11 1 2 1
10 1 1 3
8 2 0 2
5 2 3 2
6 3 2 0
6 1 3 1
8 1 2 2
8 2 1 2
11 2 3 3
6 1 1 0
8 1 0 2
5 2 0 2
6 2 0 0
8 0 3 0
11 3 0 3
10 3 3 2
6 2 1 0
6 2 2 3
6 0 1 1
9 0 3 0
8 0 1 0
8 0 1 0
11 0 2 2
10 2 3 3
6 3 1 1
8 3 0 2
5 2 1 2
6 1 3 0
5 0 1 0
8 0 2 0
11 0 3 3
10 3 0 0
6 1 3 3
8 0 0 2
5 2 0 2
6 0 1 1
5 3 1 2
8 2 2 2
11 0 2 0
10 0 2 3
6 2 1 1
6 2 0 2
6 3 3 0
13 1 0 1
8 1 2 1
11 1 3 3
10 3 2 2
6 1 1 3
6 0 2 1
6 0 2 0
5 3 1 3
8 3 2 3
11 2 3 2
10 2 3 3
6 1 3 0
6 0 0 2
6 3 1 1
8 0 2 1
8 1 3 1
11 3 1 3
10 3 0 1
6 2 2 0
6 2 1 3
6 3 3 2
9 0 3 2
8 2 2 2
8 2 3 2
11 2 1 1
10 1 3 3
6 0 3 2
8 2 0 1
5 1 3 1
2 0 1 1
8 1 1 1
11 1 3 3
10 3 3 2
6 3 0 1
6 1 1 3
3 0 3 1
8 1 3 1
11 2 1 2
10 2 1 0
6 2 1 1
6 1 3 2
6 2 0 3
9 1 3 3
8 3 3 3
11 3 0 0
10 0 3 2
6 2 3 0
6 1 0 3
3 0 3 3
8 3 2 3
11 2 3 2
10 2 2 1
6 1 2 3
6 2 3 2
3 0 3 2
8 2 1 2
8 2 2 2
11 2 1 1
10 1 0 2
6 3 2 1
6 3 3 0
11 3 3 1
8 1 2 1
8 1 2 1
11 2 1 2
6 2 1 1
13 1 0 0
8 0 3 0
8 0 1 0
11 0 2 2
10 2 1 3
8 1 0 0
5 0 1 0
6 0 0 2
8 0 2 0
8 0 3 0
8 0 2 0
11 3 0 3
10 3 3 1
6 3 3 2
6 2 1 0
6 2 2 3
15 0 3 3
8 3 2 3
11 3 1 1
10 1 3 2
6 3 1 1
8 2 0 3
5 3 0 3
4 1 0 1
8 1 2 1
11 2 1 2
10 2 0 1
6 3 1 2
6 1 1 3
6 1 0 0
8 3 2 3
8 3 2 3
11 1 3 1
10 1 2 3
6 2 1 1
6 2 1 2
10 0 2 2
8 2 1 2
11 3 2 3
6 0 0 1
6 2 1 2
10 0 2 0
8 0 3 0
11 3 0 3
10 3 1 2
8 2 0 3
5 3 0 3
6 1 1 0
11 0 0 3
8 3 2 3
11 3 2 2
10 2 3 3
6 2 1 1
6 3 2 2
6 2 2 0
0 0 2 0
8 0 1 0
11 0 3 3
10 3 3 2
6 2 1 3
6 2 0 0
6 3 0 1
15 0 3 3
8 3 1 3
8 3 3 3
11 3 2 2
10 2 3 1
8 1 0 0
5 0 1 0
6 3 3 2
8 0 0 3
5 3 1 3
8 0 2 3
8 3 1 3
11 3 1 1
10 1 1 3
6 0 0 1
6 2 1 2
10 0 2 1
8 1 2 1
11 1 3 3
10 3 2 0
8 0 0 2
5 2 0 2
6 3 2 3
6 3 3 1
12 3 2 1
8 1 1 1
11 1 0 0
10 0 0 3
6 3 2 1
6 2 2 0
6 3 3 2
0 0 2 0
8 0 3 0
8 0 1 0
11 0 3 3
10 3 3 2
6 1 3 1
8 1 0 3
5 3 2 3
8 0 0 0
5 0 2 0
1 1 3 0
8 0 1 0
11 2 0 2
6 3 1 1
6 2 0 0
4 1 0 1
8 1 2 1
11 2 1 2
10 2 1 1
6 3 0 2
6 1 2 3
3 0 3 0
8 0 1 0
11 0 1 1
8 3 0 2
5 2 2 2
6 2 0 3
8 2 0 0
5 0 3 0
2 2 0 2
8 2 1 2
11 1 2 1
10 1 2 0
6 3 0 1
8 0 0 2
5 2 0 2
14 2 3 1
8 1 2 1
11 1 0 0
10 0 2 3
6 3 3 2
6 1 3 0
8 3 0 1
5 1 1 1
11 1 0 2
8 2 3 2
8 2 3 2
11 3 2 3
10 3 3 1
6 2 2 2
6 0 2 3
7 3 2 0
8 0 2 0
8 0 2 0
11 1 0 1
10 1 2 3
6 3 2 0
8 2 0 2
5 2 0 2
8 1 0 1
5 1 3 1
0 2 0 0
8 0 3 0
11 0 3 3
10 3 2 1
6 1 1 0
8 3 0 3
5 3 0 3
6 1 1 2
6 3 0 2
8 2 2 2
11 2 1 1
10 1 3 2
8 3 0 1
5 1 1 1
6 3 1 3
11 0 0 1
8 1 1 1
11 1 2 2
10 2 1 1
6 1 2 2
6 3 2 0
6 2 1 3
12 0 2 0
8 0 3 0
11 1 0 1
10 1 1 2
6 2 3 0
8 2 0 3
5 3 1 3
6 3 1 1
5 3 1 3
8 3 2 3
11 2 3 2
6 0 3 3
6 3 0 0
8 2 0 1
5 1 2 1
13 1 0 0
8 0 1 0
11 0 2 2
10 2 3 1
6 2 1 2
6 1 2 3
6 3 0 0
2 2 0 3
8 3 3 3
11 1 3 1
10 1 2 3
6 1 3 0
6 3 0 1
5 0 1 1
8 1 2 1
11 1 3 3
10 3 3 2
6 1 3 3
8 0 0 1
5 1 0 1
8 1 0 0
5 0 2 0
3 0 3 1
8 1 2 1
11 2 1 2
6 1 1 0
6 0 2 1
8 3 0 3
5 3 2 3
5 0 1 0
8 0 3 0
8 0 2 0
11 0 2 2
6 1 0 1
6 2 1 0
15 0 3 3
8 3 3 3
11 3 2 2
6 0 2 1
6 2 3 3
6 3 3 0
4 0 3 1
8 1 2 1
11 2 1 2
6 0 1 3
8 0 0 0
5 0 2 0
8 3 0 1
5 1 2 1
6 3 0 0
8 0 2 0
8 0 3 0
11 0 2 2
10 2 1 3
8 2 0 2
5 2 3 2
6 1 1 1
6 2 3 0
1 1 0 2
8 2 3 2
11 3 2 3
10 3 0 0
6 0 3 3
8 3 0 1
5 1 0 1
6 2 0 2
7 3 2 2
8 2 2 2
11 0 2 0
10 0 0 3
6 2 3 1
8 3 0 2
5 2 0 2
6 3 3 0
13 1 0 2
8 2 3 2
8 2 3 2
11 3 2 3
8 3 0 2
5 2 2 2
2 2 0 1
8 1 3 1
11 3 1 3
10 3 2 2
6 1 3 3
6 3 3 1
6 2 0 0
3 0 3 3
8 3 3 3
11 3 2 2
10 2 2 0
6 1 1 1
6 2 1 3
8 2 0 2
5 2 0 2
14 2 3 2
8 2 2 2
8 2 3 2
11 2 0 0
10 0 1 3
6 3 0 1
6 3 2 2
6 2 0 0
6 2 1 0
8 0 2 0
11 3 0 3
10 3 0 1
6 0 1 3
6 2 3 2
8 3 0 0
5 0 0 0
7 3 2 3
8 3 3 3
8 3 2 3
11 1 3 1
10 1 1 2
6 2 0 0
8 0 0 3
5 3 1 3
6 1 1 1
1 3 0 1
8 1 2 1
11 1 2 2
10 2 2 3
6 0 2 1
6 3 1 0
6 2 2 2
2 2 0 1
8 1 3 1
8 1 1 1
11 1 3 3
10 3 1 2
6 1 0 1
6 2 2 3
1 1 3 0
8 0 2 0
8 0 3 0
11 2 0 2
10 2 3 1
6 2 0 0
6 2 0 2
15 0 3 3
8 3 1 3
11 3 1 1
10 1 0 3
6 3 0 2
6 1 3 1
0 0 2 2
8 2 3 2
11 3 2 3
10 3 3 0
6 1 2 3
6 3 1 1
6 0 0 2
12 1 2 3
8 3 3 3
8 3 3 3
11 3 0 0
10 0 3 1
8 1 0 0
5 0 2 0
8 1 0 2
5 2 2 2
6 0 2 3
7 3 2 2
8 2 2 2
11 2 1 1
10 1 2 3
6 2 3 1
6 3 1 2
13 1 2 0
8 0 3 0
11 0 3 3
10 3 1 1
6 1 2 3
6 3 1 0
8 3 2 0
8 0 1 0
11 1 0 1
6 1 3 0
6 2 2 2
6 3 0 3
10 0 2 0
8 0 1 0
11 0 1 1
6 3 0 0
2 2 0 3
8 3 1 3
11 3 1 1
10 1 2 3
6 1 2 2
8 1 0 1
5 1 1 1
6 2 2 0
1 1 0 0
8 0 3 0
11 3 0 3
10 3 1 0
6 2 0 3
6 2 0 2
1 1 3 3
8 3 1 3
11 3 0 0
10 0 0 3
6 3 2 1
8 3 0 0
5 0 2 0
6 3 2 2
0 0 2 2
8 2 1 2
11 3 2 3
10 3 3 0
8 0 0 3
5 3 0 3
6 1 3 1
6 2 2 2
7 3 2 3
8 3 1 3
8 3 3 3
11 0 3 0
10 0 0 3
6 3 2 2
6 2 1 0
0 0 2 1
8 1 3 1
8 1 2 1
11 1 3 3
8 1 0 2
5 2 2 2
6 3 1 0
6 3 3 1
2 2 1 1
8 1 1 1
11 3 1 3
10 3 1 1
6 3 2 3
6 2 0 0
6 3 0 2
13 0 2 3
8 3 3 3
8 3 1 3
11 3 1 1
10 1 2 3
8 2 0 0
5 0 1 0
8 3 0 2
5 2 2 2
8 0 0 1
5 1 0 1
11 0 0 1
8 1 2 1
11 3 1 3
10 3 3 0
6 0 1 3
6 3 3 1
6 3 1 2
14 3 2 3
8 3 3 3
11 0 3 0
10 0 2 3
6 1 0 1
6 1 0 0
6 2 3 2
10 0 2 1
8 1 1 1
8 1 2 1
11 1 3 3
6 1 1 2
6 0 3 1
5 0 1 2
8 2 3 2
11 3 2 3
10 3 2 1
8 3 0 3
5 3 2 3
8 1 0 0
5 0 2 0
6 3 3 2
0 0 2 3
8 3 2 3
8 3 1 3
11 3 1 1
10 1 2 0
6 1 3 3
8 3 0 2
5 2 1 2
6 0 1 1
5 3 1 3
8 3 1 3
11 0 3 0
10 0 3 1
6 1 3 0
6 0 3 3
8 1 0 2
5 2 2 2
7 3 2 3
8 3 2 3
11 3 1 1
6 0 0 2
6 2 3 3
6 0 3 0
14 2 3 0
8 0 1 0
8 0 1 0
11 0 1 1
10 1 0 3
6 3 0 2
6 2 2 0
6 2 0 1
13 0 2 1
8 1 3 1
8 1 1 1
11 3 1 3
10 3 2 1
6 3 1 3
6 1 2 0
8 0 2 2
8 2 3 2
11 1 2 1
6 1 0 3
6 3 3 2
11 0 0 2
8 2 2 2
11 1 2 1
10 1 2 3
6 0 3 0
8 1 0 2
5 2 2 2
6 3 2 1
2 2 1 2
8 2 1 2
11 3 2 3
6 0 0 2
6 1 3 1
8 1 2 0
8 0 3 0
8 0 2 0
11 0 3 3
10 3 2 2
6 2 2 0
6 1 2 3
3 0 3 0
8 0 2 0
11 2 0 2
10 2 3 3
6 0 0 1
6 2 0 0
6 3 3 2
0 0 2 2
8 2 1 2
8 2 1 2
11 2 3 3
10 3 1 0
6 2 3 1
6 2 3 3
8 3 0 2
5 2 0 2
14 2 3 2
8 2 3 2
11 2 0 0
10 0 0 1
8 2 0 2
5 2 0 2
6 1 0 0
1 0 3 3
8 3 2 3
8 3 2 3
11 1 3 1
6 2 2 0
6 2 0 3
6 1 3 2
15 0 3 0
8 0 2 0
11 1 0 1
10 1 2 0
6 3 0 1
6 0 3 2
6 1 1 3
12 1 2 1
8 1 1 1
11 1 0 0
10 0 1 2
6 3 1 1
6 3 1 0
5 3 1 3
8 3 3 3
11 2 3 2
10 2 1 0
6 0 2 3
6 2 3 2
8 1 0 1
5 1 0 1
7 3 2 3
8 3 1 3
8 3 1 3
11 3 0 0
6 2 1 1
6 3 0 2
6 0 1 3
14 3 2 2
8 2 2 2
11 2 0 0
8 2 0 2
5 2 0 2
6 1 1 3
6 2 3 1
8 1 3 1
11 0 1 0
10 0 0 1
6 2 1 0
6 3 1 2
1 3 0 0
8 0 2 0
11 0 1 1
10 1 3 0
6 0 3 3
6 1 1 1
14 3 2 1
8 1 2 1
11 0 1 0
10 0 3 2
6 2 3 1
6 2 0 0
6 1 1 3
1 3 0 0
8 0 1 0
11 0 2 2
10 2 0 1
6 0 2 2
6 2 0 3
6 2 1 0
15 0 3 3
8 3 2 3
8 3 3 3
11 1 3 1
6 3 0 0
8 0 0 3
5 3 1 3
11 3 3 3
8 3 1 3
8 3 1 3
11 3 1 1
10 1 3 3
6 2 0 0
6 3 0 2
8 0 0 1
5 1 2 1
0 0 2 1
8 1 1 1
8 1 1 1
11 1 3 3
10 3 1 2
6 2 0 1
6 2 3 3
15 0 3 1
8 1 1 1
8 1 1 1
11 1 2 2
10 2 1 0
6 0 2 3
6 3 3 1
6 2 1 2
7 3 2 3
8 3 3 3
11 3 0 0
10 0 0 2
6 2 3 0
6 0 0 3
2 0 1 3
8 3 3 3
11 2 3 2
10 2 2 3
6 1 0 2
6 1 0 1
1 1 0 1
8 1 1 1
11 3 1 3
8 1 0 1
5 1 3 1
2 0 1 1
8 1 3 1
11 3 1 3
10 3 0 2
6 2 2 3
8 2 0 1
5 1 3 1
15 0 3 1
8 1 1 1
8 1 1 1
11 2 1 2
10 2 3 0
6 3 0 1
6 1 0 3
6 2 1 2
2 2 1 2
8 2 2 2
8 2 3 2
11 0 2 0
10 0 0 1
6 1 0 0
6 2 2 2
10 0 2 2
8 2 1 2
11 1 2 1
10 1 2 0
6 3 1 1
6 2 0 3
6 2 0 2
9 2 3 2
8 2 1 2
11 0 2 0
10 0 2 2
6 1 1 3
6 0 2 1
8 0 0 0
5 0 1 0
5 3 1 3
8 3 2 3
11 3 2 2
10 2 0 1
6 1 2 2
6 2 1 3
1 0 3 2
8 2 2 2
11 2 1 1
6 3 3 2
6 2 3 0
15 0 3 2
8 2 2 2
11 2 1 1
10 1 2 0
6 2 1 1
6 3 3 2
9 1 3 3
8 3 3 3
11 0 3 0
10 0 2 1
6 0 0 2
6 2 0 0
6 1 1 3
1 3 0 2
8 2 1 2
8 2 1 2
11 2 1 1
10 1 0 0
6 2 2 1
6 0 0 3
6 3 2 2
13 1 2 2
8 2 3 2
11 0 2 0
10 0 0 2
6 2 2 3
8 2 0 0
5 0 0 0
6 3 1 1
6 3 0 0
8 0 3 0
11 0 2 2
10 2 0 1
6 0 1 2
6 3 2 0
6 0 1 3
12 0 2 3
8 3 1 3
8 3 2 3
11 3 1 1
10 1 3 2
6 2 2 1
6 2 3 0
6 3 2 3
4 3 1 1
8 1 2 1
11 2 1 2
10 2 0 0
8 3 0 3
5 3 1 3
6 1 1 1
6 3 3 2
8 3 2 3
8 3 2 3
11 0 3 0
6 2 3 2
6 2 0 1
6 2 0 3
9 1 3 3
8 3 1 3
11 0 3 0
10 0 1 2
6 1 3 3
6 2 2 0
11 3 3 3
8 3 2 3
11 2 3 2
6 2 0 3
6 3 1 1
15 0 3 0
8 0 2 0
8 0 2 0
11 2 0 2
10 2 2 3
8 2 0 0
5 0 1 0
8 3 0 1
5 1 1 1
6 0 0 2
8 1 2 2
8 2 2 2
8 2 1 2
11 3 2 3
10 3 2 0`,
	"google": `Before: [1, 2, 3, 2]
3 1 3 0
After:  [1, 2, 3, 2]

Before: [1, 1, 1, 3]
5 1 3 0
After:  [3, 1, 1, 3]

Before: [2, 3, 0, 2]
0 1 0 2
After:  [2, 3, 6, 2]

Before: [1, 2, 2, 3]
11 0 3 3
After:  [1, 2, 2, 0]

Before: [0, 0, 3, 3]
9 0 0 1
After:  [0, 0, 3, 3]

Before: [1, 0, 1, 2]
10 1 2 0
After:  [1, 0, 1, 2]

Before: [0, 2, 0, 2]
13 1 1 1
After:  [0, 1, 0, 2]

Before: [3, 1, 1, 1]
6 1 0 1
After:  [3, 3, 1, 1]

Before: [2, 3, 2, 0]
4 1 2 2
After:  [2, 3, 1, 0]

Before: [1, 2, 0, 2]
3 1 3 2
After:  [1, 2, 1, 2]

Before: [0, 1, 3, 3]
2 1 0 2
After:  [0, 1, 1, 3]

Before: [0, 3, 1, 3]
6 2 1 2
After:  [0, 3, 3, 3]

Before: [3, 1, 1, 1]
6 1 0 2
After:  [3, 1, 3, 1]

Before: [0, 2, 3, 2]
3 1 3 0
After:  [1, 2, 3, 2]

Before: [0, 0, 1, 3]
7 2 1 0
After:  [1, 0, 1, 3]

Before: [3, 1, 2, 3]
15 2 3 1
After:  [3, 3, 2, 3]

Before: [3, 2, 3, 2]
3 1 3 3
After:  [3, 2, 3, 1]

Before: [0, 2, 1, 2]
3 1 3 0
After:  [1, 2, 1, 2]

Before: [1, 0, 1, 1]
1 1 0 1
After:  [1, 1, 1, 1]

Before: [3, 2, 2, 2]
3 1 3 1
After:  [3, 1, 2, 2]

Before: [2, 3, 2, 3]
8 2 2 3
After:  [2, 3, 2, 4]

Before: [2, 1, 0, 2]
8 0 1 0
After:  [3, 1, 0, 2]

Before: [1, 1, 2, 3]
5 0 3 1
After:  [1, 3, 2, 3]

Before: [1, 0, 3, 1]
1 1 0 0
After:  [1, 0, 3, 1]

Before: [0, 1, 3, 1]
14 3 2 1
After:  [0, 3, 3, 1]

Before: [3, 3, 1, 1]
0 1 0 1
After:  [3, 9, 1, 1]

Before: [2, 3, 0, 3]
15 0 3 2
After:  [2, 3, 3, 3]

Before: [2, 2, 3, 3]
15 1 3 2
After:  [2, 2, 3, 3]

Before: [1, 1, 3, 2]
12 0 2 3
After:  [1, 1, 3, 3]

Before: [2, 1, 1, 3]
15 0 3 3
After:  [2, 1, 1, 3]

Before: [0, 2, 3, 2]
9 0 0 2
After:  [0, 2, 0, 2]

Before: [3, 2, 1, 0]
6 2 0 0
After:  [3, 2, 1, 0]

Before: [3, 3, 3, 2]
13 1 0 0
After:  [1, 3, 3, 2]

Before: [3, 2, 2, 3]
8 2 2 3
After:  [3, 2, 2, 4]

Before: [1, 0, 2, 3]
1 1 0 2
After:  [1, 0, 1, 3]

Before: [1, 0, 2, 0]
8 2 2 2
After:  [1, 0, 4, 0]

Before: [1, 3, 3, 2]
12 0 2 1
After:  [1, 3, 3, 2]

Before: [0, 1, 2, 2]
2 1 0 2
After:  [0, 1, 1, 2]

Before: [3, 0, 2, 3]
0 3 0 0
After:  [9, 0, 2, 3]

Before: [2, 0, 3, 1]
7 3 1 3
After:  [2, 0, 3, 1]

Before: [1, 2, 1, 3]
10 1 0 0
After:  [1, 2, 1, 3]

Before: [2, 3, 3, 0]
4 1 0 2
After:  [2, 3, 1, 0]

Before: [0, 1, 1, 2]
2 1 0 3
After:  [0, 1, 1, 1]

Before: [1, 0, 0, 2]
7 0 1 1
After:  [1, 1, 0, 2]

Before: [2, 3, 2, 2]
4 1 2 1
After:  [2, 1, 2, 2]

Before: [1, 2, 1, 0]
13 1 1 2
After:  [1, 2, 1, 0]

Before: [1, 0, 2, 2]
1 1 0 2
After:  [1, 0, 1, 2]

Before: [2, 2, 1, 3]
11 0 3 2
After:  [2, 2, 0, 3]

Before: [0, 1, 3, 1]
2 1 0 0
After:  [1, 1, 3, 1]

Before: [0, 1, 2, 2]
2 1 0 1
After:  [0, 1, 2, 2]

Before: [2, 3, 1, 1]
6 2 1 1
After:  [2, 3, 1, 1]

Before: [1, 2, 1, 0]
10 1 0 1
After:  [1, 1, 1, 0]

Before: [3, 1, 0, 3]
5 1 3 3
After:  [3, 1, 0, 3]

Before: [2, 1, 3, 1]
12 1 2 0
After:  [3, 1, 3, 1]

Before: [0, 0, 2, 3]
10 0 1 0
After:  [1, 0, 2, 3]

Before: [0, 1, 2, 0]
2 1 0 1
After:  [0, 1, 2, 0]

Before: [0, 1, 0, 3]
5 1 3 1
After:  [0, 3, 0, 3]

Before: [0, 1, 1, 1]
2 1 0 0
After:  [1, 1, 1, 1]

Before: [2, 0, 2, 1]
7 3 1 2
After:  [2, 0, 1, 1]

Before: [1, 0, 1, 0]
7 2 1 1
After:  [1, 1, 1, 0]

Before: [3, 2, 3, 3]
15 1 3 1
After:  [3, 3, 3, 3]

Before: [1, 2, 0, 3]
10 1 0 0
After:  [1, 2, 0, 3]

Before: [1, 2, 0, 0]
13 1 1 3
After:  [1, 2, 0, 1]

Before: [2, 3, 2, 3]
0 1 0 2
After:  [2, 3, 6, 3]

Before: [3, 3, 3, 1]
13 1 0 1
After:  [3, 1, 3, 1]

Before: [1, 2, 3, 1]
14 3 2 2
After:  [1, 2, 3, 1]

Before: [0, 0, 1, 3]
8 0 3 1
After:  [0, 3, 1, 3]

Before: [0, 2, 2, 3]
0 3 1 2
After:  [0, 2, 6, 3]

Before: [1, 1, 3, 3]
5 0 3 2
After:  [1, 1, 3, 3]

Before: [0, 1, 1, 0]
2 1 0 2
After:  [0, 1, 1, 0]

Before: [1, 0, 3, 0]
1 1 0 0
After:  [1, 0, 3, 0]

Before: [0, 1, 2, 0]
2 1 0 0
After:  [1, 1, 2, 0]

Before: [1, 1, 3, 2]
14 3 1 1
After:  [1, 3, 3, 2]

Before: [1, 3, 2, 2]
6 0 1 1
After:  [1, 3, 2, 2]

Before: [1, 0, 0, 2]
1 1 0 1
After:  [1, 1, 0, 2]

Before: [2, 3, 3, 0]
4 1 0 3
After:  [2, 3, 3, 1]

Before: [1, 0, 2, 1]
10 1 3 2
After:  [1, 0, 1, 1]

Before: [2, 1, 1, 1]
0 3 0 3
After:  [2, 1, 1, 2]

Before: [1, 0, 0, 1]
7 0 1 3
After:  [1, 0, 0, 1]

Before: [2, 0, 0, 1]
7 3 1 3
After:  [2, 0, 0, 1]

Before: [2, 3, 3, 2]
13 1 2 0
After:  [1, 3, 3, 2]

Before: [2, 2, 0, 1]
13 1 0 3
After:  [2, 2, 0, 1]

Before: [2, 3, 1, 3]
5 2 3 3
After:  [2, 3, 1, 3]

Before: [3, 2, 3, 2]
3 1 3 2
After:  [3, 2, 1, 2]

Before: [1, 3, 3, 1]
12 0 2 2
After:  [1, 3, 3, 1]

Before: [1, 0, 1, 3]
7 2 1 0
After:  [1, 0, 1, 3]

Before: [1, 3, 2, 0]
6 0 1 1
After:  [1, 3, 2, 0]

Before: [2, 0, 3, 1]
14 3 2 3
After:  [2, 0, 3, 3]

Before: [0, 1, 3, 1]
8 0 1 1
After:  [0, 1, 3, 1]

Before: [1, 1, 2, 3]
15 2 3 2
After:  [1, 1, 3, 3]

Before: [1, 2, 2, 2]
3 1 3 3
After:  [1, 2, 2, 1]

Before: [0, 1, 0, 2]
2 1 0 1
After:  [0, 1, 0, 2]

Before: [1, 2, 3, 3]
5 0 3 2
After:  [1, 2, 3, 3]

Before: [0, 0, 2, 0]
8 2 2 1
After:  [0, 4, 2, 0]

Before: [2, 3, 3, 1]
0 0 2 1
After:  [2, 6, 3, 1]

Before: [1, 0, 3, 2]
12 0 2 1
After:  [1, 3, 3, 2]

Before: [3, 1, 0, 3]
5 1 3 0
After:  [3, 1, 0, 3]

Before: [0, 0, 3, 1]
10 0 1 1
After:  [0, 1, 3, 1]

Before: [3, 3, 0, 2]
13 1 0 2
After:  [3, 3, 1, 2]

Before: [1, 1, 0, 3]
5 0 3 1
After:  [1, 3, 0, 3]

Before: [3, 3, 3, 1]
14 3 2 0
After:  [3, 3, 3, 1]

Before: [1, 3, 2, 3]
5 0 3 0
After:  [3, 3, 2, 3]

Before: [1, 0, 0, 3]
7 0 1 2
After:  [1, 0, 1, 3]

Before: [2, 2, 3, 1]
14 3 2 3
After:  [2, 2, 3, 3]

Before: [0, 1, 3, 1]
12 1 2 0
After:  [3, 1, 3, 1]

Before: [0, 1, 2, 3]
9 0 0 3
After:  [0, 1, 2, 0]

Before: [1, 3, 2, 3]
5 0 3 3
After:  [1, 3, 2, 3]

Before: [0, 2, 2, 3]
8 1 2 2
After:  [0, 2, 4, 3]

Before: [2, 3, 3, 0]
0 2 0 3
After:  [2, 3, 3, 6]

Before: [2, 3, 3, 1]
14 3 2 3
After:  [2, 3, 3, 3]

Before: [3, 1, 2, 2]
8 2 2 0
After:  [4, 1, 2, 2]

Before: [2, 3, 3, 3]
11 0 3 2
After:  [2, 3, 0, 3]

Before: [2, 2, 3, 2]
13 1 0 3
After:  [2, 2, 3, 1]

Before: [3, 0, 3, 3]
0 3 0 0
After:  [9, 0, 3, 3]

Before: [2, 3, 2, 3]
4 1 2 2
After:  [2, 3, 1, 3]

Before: [0, 0, 1, 1]
7 2 1 0
After:  [1, 0, 1, 1]

Before: [1, 0, 3, 0]
1 1 0 2
After:  [1, 0, 1, 0]

Before: [0, 1, 2, 3]
8 2 2 1
After:  [0, 4, 2, 3]

Before: [0, 1, 3, 0]
2 1 0 3
After:  [0, 1, 3, 1]

Before: [3, 3, 3, 1]
13 1 0 2
After:  [3, 3, 1, 1]

Before: [0, 1, 3, 0]
2 1 0 2
After:  [0, 1, 1, 0]

Before: [0, 0, 3, 1]
7 3 1 3
After:  [0, 0, 3, 1]

Before: [1, 0, 3, 0]
1 1 0 3
After:  [1, 0, 3, 1]

Before: [0, 3, 1, 2]
6 2 1 2
After:  [0, 3, 3, 2]

Before: [2, 1, 1, 3]
5 2 3 3
After:  [2, 1, 1, 3]

Before: [2, 0, 0, 3]
11 0 3 3
After:  [2, 0, 0, 0]

Before: [1, 0, 2, 1]
1 1 0 3
After:  [1, 0, 2, 1]

Before: [1, 3, 2, 3]
4 1 2 0
After:  [1, 3, 2, 3]

Before: [0, 2, 1, 2]
3 1 3 3
After:  [0, 2, 1, 1]

Before: [0, 0, 0, 0]
9 0 0 3
After:  [0, 0, 0, 0]

Before: [1, 0, 3, 3]
1 1 0 1
After:  [1, 1, 3, 3]

Before: [1, 1, 3, 1]
12 0 2 3
After:  [1, 1, 3, 3]

Before: [2, 3, 0, 0]
4 1 0 2
After:  [2, 3, 1, 0]

Before: [0, 1, 2, 3]
9 0 0 0
After:  [0, 1, 2, 3]

Before: [1, 0, 0, 3]
1 1 0 3
After:  [1, 0, 0, 1]

Before: [2, 2, 2, 0]
13 1 0 3
After:  [2, 2, 2, 1]

Before: [3, 3, 3, 3]
13 1 0 2
After:  [3, 3, 1, 3]

Before: [1, 2, 1, 3]
15 1 3 3
After:  [1, 2, 1, 3]

Before: [1, 3, 2, 1]
4 1 2 2
After:  [1, 3, 1, 1]

Before: [3, 3, 3, 3]
13 1 0 1
After:  [3, 1, 3, 3]

Before: [1, 0, 1, 1]
7 2 1 0
After:  [1, 0, 1, 1]

Before: [2, 1, 0, 3]
5 1 3 2
After:  [2, 1, 3, 3]

Before: [1, 0, 1, 2]
7 0 1 2
After:  [1, 0, 1, 2]

Before: [2, 1, 3, 2]
8 1 3 3
After:  [2, 1, 3, 3]

Before: [1, 2, 2, 3]
15 1 3 0
After:  [3, 2, 2, 3]

Before: [1, 0, 2, 3]
11 0 3 3
After:  [1, 0, 2, 0]

Before: [1, 3, 2, 2]
4 1 2 0
After:  [1, 3, 2, 2]

Before: [2, 2, 2, 3]
15 0 3 2
After:  [2, 2, 3, 3]

Before: [3, 1, 3, 1]
12 1 2 2
After:  [3, 1, 3, 1]

Before: [1, 0, 2, 1]
1 1 0 2
After:  [1, 0, 1, 1]

Before: [0, 3, 2, 0]
0 2 1 2
After:  [0, 3, 6, 0]

Before: [0, 2, 3, 2]
13 1 1 0
After:  [1, 2, 3, 2]

Before: [2, 3, 3, 2]
4 1 0 2
After:  [2, 3, 1, 2]

Before: [3, 3, 2, 1]
4 1 2 1
After:  [3, 1, 2, 1]

Before: [1, 3, 1, 2]
8 2 3 3
After:  [1, 3, 1, 3]

Before: [1, 3, 1, 1]
6 0 1 0
After:  [3, 3, 1, 1]

Before: [0, 3, 2, 1]
14 3 2 2
After:  [0, 3, 3, 1]

Before: [1, 0, 0, 3]
1 1 0 0
After:  [1, 0, 0, 3]

Before: [0, 0, 1, 1]
10 0 1 2
After:  [0, 0, 1, 1]

Before: [1, 1, 2, 2]
8 2 2 0
After:  [4, 1, 2, 2]

Before: [1, 2, 3, 2]
3 1 3 1
After:  [1, 1, 3, 2]

Before: [1, 2, 1, 2]
10 1 0 1
After:  [1, 1, 1, 2]

Before: [2, 3, 2, 0]
4 1 0 1
After:  [2, 1, 2, 0]

Before: [1, 0, 1, 0]
1 1 0 1
After:  [1, 1, 1, 0]

Before: [2, 0, 1, 2]
10 1 2 1
After:  [2, 1, 1, 2]

Before: [3, 3, 3, 0]
13 1 0 2
After:  [3, 3, 1, 0]

Before: [2, 0, 3, 3]
11 0 3 3
After:  [2, 0, 3, 0]

Before: [0, 1, 2, 3]
5 1 3 3
After:  [0, 1, 2, 3]

Before: [1, 0, 2, 3]
1 1 0 0
After:  [1, 0, 2, 3]

Before: [0, 0, 3, 3]
10 0 1 0
After:  [1, 0, 3, 3]

Before: [1, 0, 2, 2]
1 1 0 0
After:  [1, 0, 2, 2]

Before: [1, 0, 0, 2]
1 1 0 2
After:  [1, 0, 1, 2]

Before: [2, 3, 1, 1]
6 2 1 2
After:  [2, 3, 3, 1]

Before: [2, 2, 3, 1]
14 3 2 2
After:  [2, 2, 3, 1]

Before: [3, 3, 1, 3]
6 2 1 3
After:  [3, 3, 1, 3]

Before: [0, 1, 1, 2]
2 1 0 0
After:  [1, 1, 1, 2]

Before: [2, 0, 3, 3]
0 2 2 3
After:  [2, 0, 3, 9]

Before: [0, 1, 0, 0]
9 0 0 1
After:  [0, 0, 0, 0]

Before: [0, 1, 2, 1]
2 1 0 3
After:  [0, 1, 2, 1]

Before: [2, 2, 0, 3]
15 0 3 3
After:  [2, 2, 0, 3]

Before: [0, 3, 2, 1]
14 3 2 3
After:  [0, 3, 2, 3]

Before: [1, 2, 1, 3]
15 1 3 0
After:  [3, 2, 1, 3]

Before: [0, 3, 0, 2]
9 0 0 3
After:  [0, 3, 0, 0]

Before: [1, 2, 0, 1]
13 1 1 1
After:  [1, 1, 0, 1]

Before: [2, 3, 3, 1]
14 3 2 1
After:  [2, 3, 3, 1]

Before: [0, 2, 1, 3]
15 1 3 3
After:  [0, 2, 1, 3]

Before: [1, 0, 2, 1]
7 3 1 2
After:  [1, 0, 1, 1]

Before: [0, 2, 3, 2]
3 1 3 3
After:  [0, 2, 3, 1]

Before: [3, 3, 1, 3]
13 1 0 3
After:  [3, 3, 1, 1]

Before: [3, 2, 2, 2]
3 1 3 3
After:  [3, 2, 2, 1]

Before: [0, 1, 1, 1]
9 0 0 1
After:  [0, 0, 1, 1]

Before: [0, 1, 2, 3]
9 0 0 1
After:  [0, 0, 2, 3]

Before: [0, 1, 3, 2]
9 0 0 0
After:  [0, 1, 3, 2]

Before: [1, 1, 3, 3]
11 0 3 2
After:  [1, 1, 0, 3]

Before: [1, 1, 1, 3]
5 0 3 1
After:  [1, 3, 1, 3]

Before: [3, 1, 1, 2]
14 3 1 1
After:  [3, 3, 1, 2]

Before: [0, 3, 1, 0]
0 1 1 2
After:  [0, 3, 9, 0]

Before: [0, 1, 1, 1]
2 1 0 1
After:  [0, 1, 1, 1]

Before: [2, 2, 2, 1]
8 2 2 3
After:  [2, 2, 2, 4]

Before: [1, 2, 3, 0]
12 0 2 0
After:  [3, 2, 3, 0]

Before: [3, 3, 2, 1]
4 1 2 2
After:  [3, 3, 1, 1]

Before: [2, 2, 2, 1]
14 3 2 0
After:  [3, 2, 2, 1]

Before: [1, 1, 2, 1]
8 2 1 2
After:  [1, 1, 3, 1]

Before: [0, 3, 2, 2]
4 1 2 1
After:  [0, 1, 2, 2]

Before: [1, 0, 0, 0]
1 1 0 0
After:  [1, 0, 0, 0]

Before: [0, 1, 3, 2]
2 1 0 1
After:  [0, 1, 3, 2]

Before: [3, 3, 2, 0]
0 0 1 1
After:  [3, 9, 2, 0]

Before: [2, 3, 3, 3]
15 0 3 2
After:  [2, 3, 3, 3]

Before: [1, 0, 0, 2]
1 1 0 0
After:  [1, 0, 0, 2]

Before: [0, 3, 2, 3]
4 1 2 3
After:  [0, 3, 2, 1]

Before: [0, 1, 1, 2]
2 1 0 2
After:  [0, 1, 1, 2]

Before: [1, 2, 1, 2]
3 1 3 0
After:  [1, 2, 1, 2]

Before: [2, 2, 0, 3]
15 0 3 2
After:  [2, 2, 3, 3]

Before: [1, 3, 3, 0]
6 0 1 3
After:  [1, 3, 3, 3]

Before: [3, 2, 2, 3]
15 2 3 2
After:  [3, 2, 3, 3]

Before: [2, 2, 0, 2]
3 1 3 0
After:  [1, 2, 0, 2]

Before: [0, 1, 3, 1]
12 1 2 1
After:  [0, 3, 3, 1]

Before: [2, 0, 3, 3]
15 0 3 1
After:  [2, 3, 3, 3]

Before: [1, 1, 0, 2]
14 3 1 0
After:  [3, 1, 0, 2]

Before: [3, 3, 1, 2]
0 1 1 0
After:  [9, 3, 1, 2]

Before: [3, 1, 0, 0]
6 1 0 1
After:  [3, 3, 0, 0]

Before: [2, 2, 3, 1]
13 1 1 0
After:  [1, 2, 3, 1]

Before: [1, 0, 3, 2]
7 0 1 1
After:  [1, 1, 3, 2]

Before: [0, 1, 2, 1]
8 2 2 0
After:  [4, 1, 2, 1]

Before: [1, 0, 3, 1]
7 0 1 2
After:  [1, 0, 1, 1]

Before: [0, 3, 1, 1]
6 2 1 0
After:  [3, 3, 1, 1]

Before: [1, 0, 3, 3]
12 0 2 1
After:  [1, 3, 3, 3]

Before: [1, 3, 3, 3]
12 0 2 2
After:  [1, 3, 3, 3]

Before: [2, 1, 3, 1]
12 1 2 3
After:  [2, 1, 3, 3]

Before: [0, 2, 3, 2]
3 1 3 1
After:  [0, 1, 3, 2]

Before: [3, 2, 0, 1]
0 0 0 3
After:  [3, 2, 0, 9]

Before: [2, 0, 1, 3]
15 0 3 2
After:  [2, 0, 3, 3]

Before: [3, 1, 1, 2]
6 2 0 3
After:  [3, 1, 1, 3]

Before: [0, 0, 3, 1]
7 3 1 2
After:  [0, 0, 1, 1]

Before: [2, 2, 3, 3]
15 1 3 0
After:  [3, 2, 3, 3]

Before: [0, 0, 1, 1]
7 2 1 3
After:  [0, 0, 1, 1]

Before: [0, 1, 3, 2]
12 1 2 2
After:  [0, 1, 3, 2]

Before: [0, 1, 3, 3]
0 2 2 2
After:  [0, 1, 9, 3]

Before: [0, 0, 2, 3]
10 0 1 3
After:  [0, 0, 2, 1]

Before: [1, 3, 1, 1]
6 0 1 3
After:  [1, 3, 1, 3]

Before: [1, 0, 1, 2]
10 1 2 1
After:  [1, 1, 1, 2]

Before: [0, 1, 0, 2]
14 3 1 0
After:  [3, 1, 0, 2]

Before: [0, 1, 3, 2]
9 0 0 2
After:  [0, 1, 0, 2]

Before: [1, 0, 1, 2]
1 1 0 0
After:  [1, 0, 1, 2]

Before: [3, 0, 2, 1]
7 3 1 0
After:  [1, 0, 2, 1]

Before: [2, 2, 1, 3]
5 2 3 2
After:  [2, 2, 3, 3]

Before: [1, 3, 0, 3]
5 0 3 2
After:  [1, 3, 3, 3]

Before: [0, 1, 3, 1]
12 1 2 2
After:  [0, 1, 3, 1]

Before: [2, 3, 2, 3]
15 2 3 1
After:  [2, 3, 2, 3]

Before: [2, 1, 1, 3]
11 0 3 2
After:  [2, 1, 0, 3]

Before: [2, 2, 0, 2]
3 1 3 3
After:  [2, 2, 0, 1]

Before: [2, 3, 2, 3]
15 0 3 2
After:  [2, 3, 3, 3]

Before: [3, 2, 1, 3]
5 2 3 1
After:  [3, 3, 1, 3]

Before: [1, 0, 1, 3]
5 2 3 2
After:  [1, 0, 3, 3]

Before: [0, 1, 3, 1]
2 1 0 2
After:  [0, 1, 1, 1]

Before: [1, 0, 3, 2]
7 0 1 3
After:  [1, 0, 3, 1]

Before: [1, 2, 2, 2]
3 1 3 1
After:  [1, 1, 2, 2]

Before: [2, 3, 0, 3]
4 1 0 0
After:  [1, 3, 0, 3]

Before: [3, 3, 1, 1]
6 2 0 2
After:  [3, 3, 3, 1]

Before: [2, 3, 3, 1]
0 0 2 3
After:  [2, 3, 3, 6]

Before: [1, 1, 2, 1]
14 3 2 0
After:  [3, 1, 2, 1]

Before: [1, 3, 1, 3]
11 0 3 0
After:  [0, 3, 1, 3]

Before: [2, 2, 0, 3]
15 0 3 1
After:  [2, 3, 0, 3]

Before: [1, 2, 2, 0]
10 1 0 1
After:  [1, 1, 2, 0]

Before: [1, 0, 1, 0]
10 1 2 1
After:  [1, 1, 1, 0]

Before: [1, 2, 0, 3]
10 1 0 2
After:  [1, 2, 1, 3]

Before: [0, 1, 1, 1]
2 1 0 3
After:  [0, 1, 1, 1]

Before: [2, 0, 1, 0]
7 2 1 2
After:  [2, 0, 1, 0]

Before: [0, 3, 1, 3]
6 2 1 1
After:  [0, 3, 1, 3]

Before: [3, 1, 3, 3]
5 1 3 1
After:  [3, 3, 3, 3]

Before: [0, 3, 3, 1]
0 1 2 2
After:  [0, 3, 9, 1]

Before: [2, 3, 2, 2]
8 3 2 0
After:  [4, 3, 2, 2]

Before: [2, 3, 1, 2]
4 1 0 3
After:  [2, 3, 1, 1]

Before: [1, 0, 3, 2]
12 0 2 0
After:  [3, 0, 3, 2]

Before: [2, 1, 2, 2]
8 2 1 0
After:  [3, 1, 2, 2]

Before: [0, 0, 0, 0]
10 0 1 3
After:  [0, 0, 0, 1]

Before: [1, 3, 0, 3]
11 0 3 2
After:  [1, 3, 0, 3]

Before: [2, 0, 3, 1]
7 3 1 1
After:  [2, 1, 3, 1]

Before: [0, 0, 3, 0]
10 0 1 2
After:  [0, 0, 1, 0]

Before: [2, 3, 1, 3]
4 1 0 0
After:  [1, 3, 1, 3]

Before: [3, 3, 2, 0]
13 1 0 3
After:  [3, 3, 2, 1]

Before: [1, 2, 2, 2]
10 1 0 1
After:  [1, 1, 2, 2]

Before: [3, 1, 3, 3]
12 1 2 1
After:  [3, 3, 3, 3]

Before: [3, 3, 2, 1]
14 3 2 0
After:  [3, 3, 2, 1]

Before: [2, 3, 2, 1]
4 1 2 1
After:  [2, 1, 2, 1]

Before: [0, 2, 2, 3]
15 2 3 3
After:  [0, 2, 2, 3]

Before: [3, 2, 3, 1]
0 2 0 1
After:  [3, 9, 3, 1]

Before: [0, 2, 0, 2]
3 1 3 3
After:  [0, 2, 0, 1]

Before: [3, 2, 3, 0]
0 2 2 3
After:  [3, 2, 3, 9]

Before: [0, 2, 1, 3]
9 0 0 1
After:  [0, 0, 1, 3]

Before: [0, 0, 0, 2]
9 0 0 1
After:  [0, 0, 0, 2]

Before: [1, 0, 3, 1]
1 1 0 1
After:  [1, 1, 3, 1]

Before: [1, 0, 1, 1]
7 3 1 3
After:  [1, 0, 1, 1]

Before: [0, 3, 3, 2]
9 0 0 0
After:  [0, 3, 3, 2]

Before: [1, 0, 1, 1]
1 1 0 2
After:  [1, 0, 1, 1]

Before: [3, 2, 1, 3]
6 2 0 0
After:  [3, 2, 1, 3]

Before: [2, 3, 2, 0]
4 1 2 0
After:  [1, 3, 2, 0]

Before: [1, 0, 2, 2]
1 1 0 1
After:  [1, 1, 2, 2]

Before: [1, 2, 1, 3]
11 0 3 1
After:  [1, 0, 1, 3]

Before: [0, 0, 2, 1]
7 3 1 2
After:  [0, 0, 1, 1]

Before: [1, 3, 1, 0]
6 2 1 1
After:  [1, 3, 1, 0]

Before: [1, 0, 1, 3]
1 1 0 0
After:  [1, 0, 1, 3]

Before: [3, 2, 2, 3]
15 2 3 0
After:  [3, 2, 2, 3]

Before: [1, 0, 1, 3]
11 0 3 1
After:  [1, 0, 1, 3]

Before: [3, 2, 1, 1]
13 1 1 2
After:  [3, 2, 1, 1]

Before: [1, 3, 0, 3]
5 0 3 1
After:  [1, 3, 0, 3]

Before: [2, 2, 2, 3]
8 0 2 1
After:  [2, 4, 2, 3]

Before: [0, 0, 3, 0]
9 0 0 3
After:  [0, 0, 3, 0]

Before: [1, 0, 2, 2]
1 1 0 3
After:  [1, 0, 2, 1]

Before: [0, 0, 3, 2]
10 0 1 3
After:  [0, 0, 3, 1]

Before: [0, 0, 0, 1]
10 1 3 1
After:  [0, 1, 0, 1]

Before: [0, 3, 3, 2]
0 3 1 2
After:  [0, 3, 6, 2]

Before: [2, 1, 3, 3]
12 1 2 1
After:  [2, 3, 3, 3]

Before: [0, 0, 2, 0]
8 0 2 1
After:  [0, 2, 2, 0]

Before: [3, 1, 1, 2]
14 3 1 3
After:  [3, 1, 1, 3]

Before: [0, 3, 2, 1]
4 1 2 2
After:  [0, 3, 1, 1]

Before: [3, 0, 1, 3]
5 2 3 0
After:  [3, 0, 1, 3]

Before: [1, 0, 3, 2]
0 3 2 2
After:  [1, 0, 6, 2]

Before: [2, 1, 3, 3]
11 0 3 3
After:  [2, 1, 3, 0]

Before: [1, 0, 0, 2]
1 1 0 3
After:  [1, 0, 0, 1]

Before: [3, 0, 1, 1]
7 3 1 0
After:  [1, 0, 1, 1]

Before: [2, 1, 2, 3]
5 1 3 3
After:  [2, 1, 2, 3]

Before: [1, 2, 0, 1]
10 1 0 3
After:  [1, 2, 0, 1]

Before: [2, 1, 0, 3]
15 0 3 0
After:  [3, 1, 0, 3]

Before: [1, 3, 3, 1]
12 0 2 3
After:  [1, 3, 3, 3]

Before: [2, 2, 3, 2]
3 1 3 1
After:  [2, 1, 3, 2]

Before: [2, 3, 3, 2]
4 1 0 1
After:  [2, 1, 3, 2]

Before: [3, 0, 1, 0]
7 2 1 1
After:  [3, 1, 1, 0]

Before: [1, 1, 0, 2]
14 3 1 2
After:  [1, 1, 3, 2]

Before: [2, 3, 1, 2]
6 2 1 3
After:  [2, 3, 1, 3]

Before: [2, 0, 0, 1]
10 1 3 2
After:  [2, 0, 1, 1]

Before: [1, 3, 1, 3]
5 0 3 0
After:  [3, 3, 1, 3]

Before: [0, 0, 1, 0]
7 2 1 2
After:  [0, 0, 1, 0]

Before: [1, 2, 1, 3]
11 0 3 2
After:  [1, 2, 0, 3]

Before: [2, 3, 0, 0]
4 1 0 3
After:  [2, 3, 0, 1]

Before: [1, 3, 1, 1]
6 2 1 0
After:  [3, 3, 1, 1]

Before: [0, 2, 2, 2]
3 1 3 2
After:  [0, 2, 1, 2]

Before: [3, 0, 3, 1]
10 1 3 3
After:  [3, 0, 3, 1]

Before: [2, 2, 1, 2]
13 1 1 1
After:  [2, 1, 1, 2]

Before: [0, 1, 3, 0]
9 0 0 3
After:  [0, 1, 3, 0]

Before: [2, 3, 0, 3]
4 1 0 3
After:  [2, 3, 0, 1]

Before: [0, 3, 2, 2]
4 1 2 2
After:  [0, 3, 1, 2]

Before: [2, 3, 3, 0]
4 1 0 0
After:  [1, 3, 3, 0]

Before: [0, 2, 3, 3]
15 1 3 1
After:  [0, 3, 3, 3]

Before: [0, 0, 2, 2]
8 0 2 3
After:  [0, 0, 2, 2]

Before: [1, 2, 2, 3]
11 0 3 0
After:  [0, 2, 2, 3]

Before: [0, 1, 0, 3]
2 1 0 3
After:  [0, 1, 0, 1]

Before: [3, 2, 3, 2]
3 1 3 0
After:  [1, 2, 3, 2]

Before: [3, 3, 1, 0]
0 0 0 2
After:  [3, 3, 9, 0]

Before: [1, 0, 0, 3]
7 0 1 1
After:  [1, 1, 0, 3]

Before: [0, 0, 0, 1]
7 3 1 1
After:  [0, 1, 0, 1]

Before: [1, 0, 0, 3]
1 1 0 1
After:  [1, 1, 0, 3]

Before: [1, 0, 2, 1]
14 3 2 2
After:  [1, 0, 3, 1]

Before: [1, 3, 2, 1]
14 3 2 2
After:  [1, 3, 3, 1]

Before: [2, 3, 3, 2]
0 2 2 3
After:  [2, 3, 3, 9]

Before: [0, 1, 2, 1]
2 1 0 2
After:  [0, 1, 1, 1]

Before: [1, 0, 3, 1]
1 1 0 3
After:  [1, 0, 3, 1]

Before: [3, 0, 3, 0]
0 2 2 2
After:  [3, 0, 9, 0]

Before: [0, 1, 2, 0]
2 1 0 3
After:  [0, 1, 2, 1]

Before: [2, 2, 3, 3]
11 0 3 0
After:  [0, 2, 3, 3]

Before: [0, 1, 0, 0]
2 1 0 1
After:  [0, 1, 0, 0]

Before: [0, 1, 0, 3]
2 1 0 0
After:  [1, 1, 0, 3]

Before: [3, 0, 2, 3]
15 2 3 2
After:  [3, 0, 3, 3]

Before: [0, 3, 1, 3]
9 0 0 1
After:  [0, 0, 1, 3]

Before: [3, 3, 2, 2]
0 2 0 3
After:  [3, 3, 2, 6]

Before: [1, 0, 2, 3]
1 1 0 3
After:  [1, 0, 2, 1]

Before: [1, 3, 3, 0]
12 0 2 2
After:  [1, 3, 3, 0]

Before: [3, 1, 3, 1]
12 1 2 3
After:  [3, 1, 3, 3]

Before: [2, 0, 1, 1]
7 3 1 0
After:  [1, 0, 1, 1]

Before: [0, 2, 0, 2]
3 1 3 1
After:  [0, 1, 0, 2]

Before: [1, 2, 1, 2]
8 0 3 3
After:  [1, 2, 1, 3]

Before: [3, 1, 1, 3]
6 1 0 2
After:  [3, 1, 3, 3]

Before: [1, 1, 3, 3]
5 0 3 1
After:  [1, 3, 3, 3]

Before: [0, 1, 3, 0]
2 1 0 1
After:  [0, 1, 3, 0]

Before: [0, 1, 2, 3]
2 1 0 2
After:  [0, 1, 1, 3]

Before: [0, 2, 0, 2]
13 1 1 3
After:  [0, 2, 0, 1]

Before: [1, 0, 1, 3]
1 1 0 2
After:  [1, 0, 1, 3]

Before: [1, 1, 0, 3]
5 0 3 3
After:  [1, 1, 0, 3]

Before: [1, 2, 3, 1]
14 3 2 1
After:  [1, 3, 3, 1]

Before: [2, 3, 2, 3]
15 2 3 2
After:  [2, 3, 3, 3]

Before: [1, 1, 2, 2]
8 2 1 1
After:  [1, 3, 2, 2]

Before: [1, 2, 1, 2]
3 1 3 2
After:  [1, 2, 1, 2]

Before: [0, 0, 1, 3]
5 2 3 2
After:  [0, 0, 3, 3]

Before: [2, 3, 0, 1]
4 1 0 2
After:  [2, 3, 1, 1]

Before: [0, 1, 3, 0]
9 0 0 2
After:  [0, 1, 0, 0]

Before: [3, 0, 0, 1]
7 3 1 0
After:  [1, 0, 0, 1]

Before: [1, 1, 3, 2]
12 1 2 1
After:  [1, 3, 3, 2]

Before: [1, 2, 3, 2]
3 1 3 3
After:  [1, 2, 3, 1]

Before: [2, 0, 1, 2]
10 1 2 3
After:  [2, 0, 1, 1]

Before: [3, 1, 3, 1]
6 1 0 2
After:  [3, 1, 3, 1]

Before: [3, 2, 2, 3]
13 1 1 2
After:  [3, 2, 1, 3]

Before: [1, 2, 0, 2]
3 1 3 3
After:  [1, 2, 0, 1]

Before: [0, 1, 0, 1]
2 1 0 3
After:  [0, 1, 0, 1]

Before: [3, 1, 1, 1]
6 1 0 3
After:  [3, 1, 1, 3]

Before: [3, 0, 3, 0]
0 2 0 3
After:  [3, 0, 3, 9]

Before: [1, 0, 1, 2]
1 1 0 2
After:  [1, 0, 1, 2]

Before: [0, 1, 2, 1]
8 2 1 1
After:  [0, 3, 2, 1]

Before: [3, 0, 2, 1]
7 3 1 2
After:  [3, 0, 1, 1]

Before: [0, 1, 1, 3]
2 1 0 1
After:  [0, 1, 1, 3]

Before: [1, 1, 2, 2]
8 1 2 2
After:  [1, 1, 3, 2]

Before: [3, 1, 3, 3]
6 1 0 0
After:  [3, 1, 3, 3]

Before: [1, 0, 3, 1]
12 0 2 3
After:  [1, 0, 3, 3]

Before: [1, 3, 3, 2]
6 0 1 1
After:  [1, 3, 3, 2]

Before: [0, 0, 3, 1]
10 1 3 1
After:  [0, 1, 3, 1]

Before: [3, 2, 3, 2]
3 1 3 1
After:  [3, 1, 3, 2]

Before: [3, 3, 1, 2]
6 2 0 2
After:  [3, 3, 3, 2]

Before: [0, 1, 0, 2]
2 1 0 2
After:  [0, 1, 1, 2]

Before: [2, 1, 2, 3]
8 0 1 1
After:  [2, 3, 2, 3]

Before: [3, 3, 1, 0]
6 2 0 3
After:  [3, 3, 1, 3]

Before: [1, 3, 3, 2]
13 1 2 1
After:  [1, 1, 3, 2]

Before: [0, 0, 2, 3]
15 2 3 0
After:  [3, 0, 2, 3]

Before: [0, 1, 1, 1]
2 1 0 2
After:  [0, 1, 1, 1]

Before: [0, 0, 0, 2]
10 0 1 2
After:  [0, 0, 1, 2]

Before: [0, 1, 1, 0]
2 1 0 1
After:  [0, 1, 1, 0]

Before: [0, 0, 2, 2]
8 3 2 2
After:  [0, 0, 4, 2]

Before: [2, 3, 1, 1]
4 1 0 3
After:  [2, 3, 1, 1]

Before: [1, 0, 0, 0]
1 1 0 3
After:  [1, 0, 0, 1]

Before: [1, 0, 3, 3]
7 0 1 2
After:  [1, 0, 1, 3]

Before: [1, 0, 3, 2]
1 1 0 1
After:  [1, 1, 3, 2]

Before: [0, 1, 1, 0]
9 0 0 0
After:  [0, 1, 1, 0]

Before: [3, 0, 3, 1]
14 3 2 2
After:  [3, 0, 3, 1]

Before: [2, 2, 2, 3]
15 1 3 1
After:  [2, 3, 2, 3]

Before: [0, 3, 3, 0]
9 0 0 0
After:  [0, 3, 3, 0]

Before: [0, 1, 3, 3]
2 1 0 3
After:  [0, 1, 3, 1]

Before: [1, 2, 2, 3]
15 2 3 2
After:  [1, 2, 3, 3]

Before: [3, 2, 2, 2]
0 3 0 2
After:  [3, 2, 6, 2]

Before: [0, 2, 3, 3]
13 1 1 3
After:  [0, 2, 3, 1]

Before: [0, 3, 2, 3]
4 1 2 1
After:  [0, 1, 2, 3]

Before: [0, 0, 3, 1]
10 0 1 3
After:  [0, 0, 3, 1]

Before: [2, 0, 2, 3]
15 0 3 2
After:  [2, 0, 3, 3]

Before: [1, 1, 3, 2]
12 1 2 3
After:  [1, 1, 3, 3]

Before: [1, 0, 2, 1]
7 0 1 3
After:  [1, 0, 2, 1]

Before: [0, 1, 3, 0]
0 2 2 3
After:  [0, 1, 3, 9]

Before: [2, 2, 2, 1]
14 3 2 2
After:  [2, 2, 3, 1]

Before: [0, 1, 0, 1]
2 1 0 0
After:  [1, 1, 0, 1]

Before: [0, 0, 2, 3]
15 2 3 1
After:  [0, 3, 2, 3]

Before: [1, 3, 0, 1]
6 0 1 2
After:  [1, 3, 3, 1]

Before: [1, 2, 2, 3]
10 1 0 3
After:  [1, 2, 2, 1]

Before: [0, 2, 3, 3]
13 1 1 2
After:  [0, 2, 1, 3]

Before: [0, 0, 0, 1]
10 1 3 0
After:  [1, 0, 0, 1]

Before: [2, 2, 3, 2]
3 1 3 0
After:  [1, 2, 3, 2]

Before: [0, 1, 2, 2]
14 3 1 1
After:  [0, 3, 2, 2]

Before: [0, 1, 3, 3]
2 1 0 1
After:  [0, 1, 3, 3]

Before: [3, 3, 3, 3]
0 2 0 0
After:  [9, 3, 3, 3]

Before: [1, 3, 2, 2]
8 0 3 1
After:  [1, 3, 2, 2]

Before: [2, 2, 0, 3]
15 1 3 2
After:  [2, 2, 3, 3]

Before: [1, 1, 2, 3]
5 0 3 0
After:  [3, 1, 2, 3]

Before: [2, 3, 3, 3]
11 0 3 3
After:  [2, 3, 3, 0]

Before: [2, 3, 2, 3]
4 1 0 1
After:  [2, 1, 2, 3]

Before: [2, 0, 0, 3]
11 0 3 1
After:  [2, 0, 0, 3]

Before: [3, 3, 1, 3]
5 2 3 1
After:  [3, 3, 1, 3]

Before: [1, 0, 3, 0]
12 0 2 2
After:  [1, 0, 3, 0]

Before: [1, 1, 3, 3]
5 1 3 2
After:  [1, 1, 3, 3]

Before: [1, 2, 0, 2]
3 1 3 1
After:  [1, 1, 0, 2]

Before: [2, 2, 3, 2]
3 1 3 2
After:  [2, 2, 1, 2]

Before: [3, 1, 2, 1]
6 1 0 1
After:  [3, 3, 2, 1]

Before: [0, 1, 3, 0]
2 1 0 0
After:  [1, 1, 3, 0]

Before: [1, 0, 1, 3]
8 1 3 2
After:  [1, 0, 3, 3]

Before: [1, 1, 2, 3]
11 0 3 2
After:  [1, 1, 0, 3]

Before: [0, 3, 2, 3]
9 0 0 0
After:  [0, 3, 2, 3]

Before: [3, 2, 2, 0]
13 1 1 0
After:  [1, 2, 2, 0]

Before: [1, 0, 0, 1]
1 1 0 1
After:  [1, 1, 0, 1]

Before: [2, 0, 0, 3]
15 0 3 0
After:  [3, 0, 0, 3]

Before: [3, 2, 0, 2]
3 1 3 3
After:  [3, 2, 0, 1]

Before: [1, 0, 1, 3]
5 2 3 1
After:  [1, 3, 1, 3]

Before: [2, 3, 3, 1]
4 1 0 1
After:  [2, 1, 3, 1]

Before: [2, 2, 1, 3]
5 2 3 1
After:  [2, 3, 1, 3]

Before: [0, 1, 0, 0]
9 0 0 2
After:  [0, 1, 0, 0]

Before: [0, 1, 0, 1]
2 1 0 2
After:  [0, 1, 1, 1]

Before: [0, 1, 0, 2]
2 1 0 3
After:  [0, 1, 0, 1]

Before: [1, 2, 2, 2]
13 1 1 1
After:  [1, 1, 2, 2]

Before: [1, 2, 1, 2]
3 1 3 3
After:  [1, 2, 1, 1]

Before: [1, 0, 2, 3]
11 0 3 1
After:  [1, 0, 2, 3]

Before: [0, 2, 0, 3]
15 1 3 0
After:  [3, 2, 0, 3]

Before: [1, 2, 3, 3]
15 1 3 2
After:  [1, 2, 3, 3]

Before: [0, 0, 1, 1]
10 1 2 2
After:  [0, 0, 1, 1]

Before: [1, 3, 1, 3]
5 2 3 3
After:  [1, 3, 1, 3]

Before: [1, 3, 1, 3]
6 0 1 2
After:  [1, 3, 3, 3]

Before: [0, 1, 2, 2]
2 1 0 0
After:  [1, 1, 2, 2]

Before: [2, 2, 2, 2]
3 1 3 0
After:  [1, 2, 2, 2]

Before: [0, 1, 3, 2]
2 1 0 3
After:  [0, 1, 3, 1]

Before: [2, 2, 3, 1]
13 1 0 1
After:  [2, 1, 3, 1]

Before: [2, 0, 2, 1]
7 3 1 0
After:  [1, 0, 2, 1]

Before: [1, 0, 3, 3]
12 0 2 0
After:  [3, 0, 3, 3]

Before: [0, 0, 3, 3]
10 0 1 3
After:  [0, 0, 3, 1]

Before: [0, 2, 0, 2]
9 0 0 0
After:  [0, 2, 0, 2]

Before: [1, 2, 0, 3]
11 0 3 0
After:  [0, 2, 0, 3]

Before: [2, 0, 0, 3]
11 0 3 0
After:  [0, 0, 0, 3]

Before: [0, 2, 0, 2]
3 1 3 0
After:  [1, 2, 0, 2]

Before: [1, 3, 1, 3]
6 0 1 1
After:  [1, 3, 1, 3]

Before: [0, 1, 3, 2]
2 1 0 2
After:  [0, 1, 1, 2]

Before: [1, 0, 3, 1]
1 1 0 2
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 0]
7 0 1 1
After:  [1, 1, 2, 0]

Before: [2, 3, 1, 3]
15 0 3 3
After:  [2, 3, 1, 3]

Before: [1, 2, 0, 3]
11 0 3 1
After:  [1, 0, 0, 3]

Before: [1, 3, 0, 1]
6 0 1 0
After:  [3, 3, 0, 1]

Before: [2, 2, 1, 2]
3 1 3 2
After:  [2, 2, 1, 2]

Before: [1, 0, 3, 3]
1 1 0 2
After:  [1, 0, 1, 3]

Before: [3, 1, 0, 3]
5 1 3 2
After:  [3, 1, 3, 3]

Before: [0, 1, 2, 3]
8 2 1 1
After:  [0, 3, 2, 3]

Before: [0, 2, 2, 2]
3 1 3 1
After:  [0, 1, 2, 2]

Before: [1, 1, 2, 2]
0 0 3 1
After:  [1, 2, 2, 2]

Before: [2, 3, 1, 3]
11 0 3 1
After:  [2, 0, 1, 3]

Before: [3, 3, 2, 3]
4 1 2 0
After:  [1, 3, 2, 3]

Before: [3, 1, 3, 3]
5 1 3 2
After:  [3, 1, 3, 3]

Before: [2, 2, 1, 0]
13 1 1 1
After:  [2, 1, 1, 0]

Before: [0, 1, 3, 2]
2 1 0 0
After:  [1, 1, 3, 2]

Before: [3, 2, 0, 3]
15 1 3 0
After:  [3, 2, 0, 3]

Before: [2, 3, 0, 2]
0 1 0 0
After:  [6, 3, 0, 2]

Before: [2, 3, 0, 2]
4 1 0 1
After:  [2, 1, 0, 2]

Before: [1, 0, 1, 2]
1 1 0 3
After:  [1, 0, 1, 1]

Before: [2, 3, 3, 2]
0 3 2 1
After:  [2, 6, 3, 2]

Before: [0, 1, 3, 1]
2 1 0 3
After:  [0, 1, 3, 1]

Before: [2, 0, 0, 3]
15 0 3 3
After:  [2, 0, 0, 3]

Before: [2, 1, 3, 3]
5 1 3 3
After:  [2, 1, 3, 3]

Before: [3, 1, 3, 2]
12 1 2 3
After:  [3, 1, 3, 3]

Before: [0, 2, 1, 2]
3 1 3 1
After:  [0, 1, 1, 2]

Before: [0, 1, 0, 0]
2 1 0 3
After:  [0, 1, 0, 1]

Before: [0, 2, 3, 1]
14 3 2 0
After:  [3, 2, 3, 1]

Before: [3, 2, 2, 1]
13 1 1 1
After:  [3, 1, 2, 1]

Before: [3, 3, 1, 3]
6 2 0 3
After:  [3, 3, 1, 3]

Before: [3, 0, 3, 1]
14 3 2 3
After:  [3, 0, 3, 3]

Before: [2, 3, 2, 3]
11 0 3 0
After:  [0, 3, 2, 3]

Before: [2, 0, 1, 3]
5 2 3 2
After:  [2, 0, 3, 3]

Before: [1, 3, 1, 3]
11 0 3 1
After:  [1, 0, 1, 3]

Before: [1, 0, 1, 1]
7 3 1 1
After:  [1, 1, 1, 1]

Before: [1, 2, 2, 1]
14 3 2 0
After:  [3, 2, 2, 1]

Before: [1, 0, 2, 3]
8 1 2 1
After:  [1, 2, 2, 3]

Before: [1, 0, 1, 0]
7 2 1 0
After:  [1, 0, 1, 0]

Before: [2, 2, 1, 0]
13 1 1 2
After:  [2, 2, 1, 0]

Before: [1, 0, 1, 2]
7 2 1 2
After:  [1, 0, 1, 2]

Before: [2, 2, 2, 2]
3 1 3 2
After:  [2, 2, 1, 2]

Before: [2, 1, 0, 3]
11 0 3 1
After:  [2, 0, 0, 3]

Before: [2, 1, 2, 3]
11 0 3 1
After:  [2, 0, 2, 3]

Before: [0, 0, 1, 1]
7 3 1 1
After:  [0, 1, 1, 1]

Before: [1, 1, 0, 2]
14 3 1 1
After:  [1, 3, 0, 2]

Before: [1, 0, 1, 0]
7 0 1 3
After:  [1, 0, 1, 1]

Before: [0, 1, 2, 2]
8 3 2 1
After:  [0, 4, 2, 2]

Before: [1, 3, 0, 2]
6 0 1 0
After:  [3, 3, 0, 2]

Before: [1, 0, 1, 2]
7 0 1 0
After:  [1, 0, 1, 2]

Before: [3, 0, 3, 1]
7 3 1 2
After:  [3, 0, 1, 1]

Before: [0, 2, 0, 0]
9 0 0 0
After:  [0, 2, 0, 0]

Before: [1, 1, 2, 3]
8 1 2 2
After:  [1, 1, 3, 3]

Before: [0, 1, 2, 1]
2 1 0 1
After:  [0, 1, 2, 1]

Before: [1, 3, 3, 3]
11 0 3 1
After:  [1, 0, 3, 3]

Before: [3, 3, 1, 0]
6 2 0 0
After:  [3, 3, 1, 0]

Before: [2, 2, 0, 3]
11 0 3 3
After:  [2, 2, 0, 0]

Before: [3, 0, 1, 0]
7 2 1 0
After:  [1, 0, 1, 0]

Before: [1, 0, 3, 3]
1 1 0 0
After:  [1, 0, 3, 3]

Before: [1, 3, 3, 1]
14 3 2 3
After:  [1, 3, 3, 3]

Before: [2, 2, 0, 0]
13 1 0 1
After:  [2, 1, 0, 0]

Before: [2, 2, 2, 3]
15 2 3 0
After:  [3, 2, 2, 3]

Before: [3, 3, 0, 3]
0 0 0 3
After:  [3, 3, 0, 9]

Before: [2, 2, 1, 3]
15 1 3 0
After:  [3, 2, 1, 3]

Before: [2, 1, 3, 2]
0 2 0 3
After:  [2, 1, 3, 6]

Before: [0, 1, 1, 3]
2 1 0 2
After:  [0, 1, 1, 3]

Before: [2, 0, 1, 1]
7 2 1 0
After:  [1, 0, 1, 1]

Before: [1, 0, 2, 3]
8 1 3 2
After:  [1, 0, 3, 3]

Before: [0, 1, 0, 3]
2 1 0 1
After:  [0, 1, 0, 3]

Before: [0, 1, 2, 2]
2 1 0 3
After:  [0, 1, 2, 1]

Before: [1, 0, 1, 3]
1 1 0 3
After:  [1, 0, 1, 1]

Before: [2, 3, 1, 3]
4 1 0 3
After:  [2, 3, 1, 1]

Before: [2, 3, 2, 1]
4 1 0 3
After:  [2, 3, 2, 1]

Before: [1, 1, 1, 2]
0 0 3 0
After:  [2, 1, 1, 2]

Before: [0, 0, 0, 1]
9 0 0 0
After:  [0, 0, 0, 1]

Before: [1, 0, 3, 3]
5 0 3 1
After:  [1, 3, 3, 3]

Before: [0, 1, 1, 0]
2 1 0 3
After:  [0, 1, 1, 1]

Before: [1, 2, 1, 3]
10 1 0 3
After:  [1, 2, 1, 1]

Before: [0, 2, 0, 2]
3 1 3 2
After:  [0, 2, 1, 2]

Before: [2, 2, 1, 2]
3 1 3 3
After:  [2, 2, 1, 1]

Before: [2, 1, 3, 2]
12 1 2 3
After:  [2, 1, 3, 3]

Before: [2, 1, 2, 3]
15 2 3 3
After:  [2, 1, 2, 3]

Before: [0, 1, 2, 2]
9 0 0 2
After:  [0, 1, 0, 2]

Before: [0, 2, 2, 2]
8 0 2 0
After:  [2, 2, 2, 2]

Before: [0, 3, 2, 3]
4 1 2 0
After:  [1, 3, 2, 3]

Before: [1, 1, 3, 2]
8 0 3 0
After:  [3, 1, 3, 2]

Before: [1, 1, 3, 1]
12 1 2 1
After:  [1, 3, 3, 1]

Before: [2, 1, 1, 3]
5 1 3 3
After:  [2, 1, 1, 3]

Before: [3, 1, 1, 2]
6 1 0 1
After:  [3, 3, 1, 2]

Before: [1, 2, 2, 3]
15 2 3 0
After:  [3, 2, 2, 3]

Before: [3, 1, 3, 2]
6 1 0 1
After:  [3, 3, 3, 2]

Before: [2, 2, 0, 2]
3 1 3 1
After:  [2, 1, 0, 2]

Before: [3, 0, 1, 3]
5 2 3 3
After:  [3, 0, 1, 3]

Before: [3, 3, 1, 2]
8 2 3 2
After:  [3, 3, 3, 2]

Before: [2, 1, 2, 0]
8 3 2 2
After:  [2, 1, 2, 0]

Before: [1, 2, 3, 0]
12 0 2 2
After:  [1, 2, 3, 0]

Before: [3, 2, 0, 2]
3 1 3 1
After:  [3, 1, 0, 2]

Before: [1, 2, 1, 3]
5 0 3 2
After:  [1, 2, 3, 3]

Before: [0, 1, 0, 3]
9 0 0 0
After:  [0, 1, 0, 3]

Before: [1, 2, 3, 3]
10 1 0 2
After:  [1, 2, 1, 3]

Before: [2, 2, 1, 1]
13 1 0 3
After:  [2, 2, 1, 1]

Before: [2, 0, 2, 3]
15 2 3 0
After:  [3, 0, 2, 3]

Before: [2, 0, 0, 1]
10 1 3 1
After:  [2, 1, 0, 1]

Before: [1, 0, 2, 1]
1 1 0 1
After:  [1, 1, 2, 1]

Before: [3, 0, 1, 3]
10 1 2 3
After:  [3, 0, 1, 1]

Before: [1, 1, 3, 3]
5 1 3 1
After:  [1, 3, 3, 3]

Before: [1, 1, 2, 3]
5 0 3 3
After:  [1, 1, 2, 3]

Before: [3, 0, 1, 0]
6 2 0 3
After:  [3, 0, 1, 3]

Before: [1, 2, 0, 0]
10 1 0 3
After:  [1, 2, 0, 1]

Before: [0, 1, 0, 2]
2 1 0 0
After:  [1, 1, 0, 2]

Before: [0, 1, 2, 0]
9 0 0 1
After:  [0, 0, 2, 0]

Before: [1, 1, 0, 3]
8 2 3 1
After:  [1, 3, 0, 3]

Before: [2, 3, 1, 3]
4 1 0 1
After:  [2, 1, 1, 3]

Before: [0, 1, 3, 3]
9 0 0 3
After:  [0, 1, 3, 0]

Before: [0, 0, 1, 3]
10 0 1 0
After:  [1, 0, 1, 3]

Before: [2, 0, 1, 1]
10 1 3 3
After:  [2, 0, 1, 1]

Before: [3, 3, 0, 0]
0 0 1 1
After:  [3, 9, 0, 0]

Before: [1, 3, 2, 0]
4 1 2 0
After:  [1, 3, 2, 0]

Before: [0, 1, 1, 3]
9 0 0 0
After:  [0, 1, 1, 3]

Before: [0, 1, 2, 0]
2 1 0 2
After:  [0, 1, 1, 0]

Before: [2, 3, 3, 0]
0 2 1 1
After:  [2, 9, 3, 0]

Before: [0, 0, 1, 2]
10 1 2 3
After:  [0, 0, 1, 1]

Before: [2, 2, 0, 2]
13 1 0 3
After:  [2, 2, 0, 1]

Before: [1, 0, 1, 1]
10 1 2 0
After:  [1, 0, 1, 1]

Before: [3, 3, 2, 2]
13 1 0 1
After:  [3, 1, 2, 2]

Before: [1, 0, 2, 1]
8 0 2 1
After:  [1, 3, 2, 1]

Before: [1, 1, 2, 2]
8 2 2 2
After:  [1, 1, 4, 2]

Before: [0, 0, 1, 3]
5 2 3 3
After:  [0, 0, 1, 3]

Before: [2, 3, 1, 0]
6 2 1 3
After:  [2, 3, 1, 3]

Before: [1, 3, 1, 0]
6 2 1 3
After:  [1, 3, 1, 3]

Before: [1, 0, 2, 3]
7 0 1 2
After:  [1, 0, 1, 3]

Before: [2, 0, 3, 1]
0 3 0 0
After:  [2, 0, 3, 1]

Before: [2, 2, 3, 2]
3 1 3 3
After:  [2, 2, 3, 1]

Before: [0, 1, 0, 3]
2 1 0 2
After:  [0, 1, 1, 3]

Before: [3, 1, 1, 3]
6 1 0 1
After:  [3, 3, 1, 3]

Before: [1, 2, 2, 3]
13 1 1 1
After:  [1, 1, 2, 3]

Before: [2, 2, 1, 3]
15 0 3 2
After:  [2, 2, 3, 3]

Before: [1, 0, 0, 0]
1 1 0 1
After:  [1, 1, 0, 0]

Before: [2, 2, 0, 3]
11 0 3 0
After:  [0, 2, 0, 3]

Before: [1, 0, 3, 2]
1 1 0 0
After:  [1, 0, 3, 2]

Before: [1, 0, 3, 0]
1 1 0 1
After:  [1, 1, 3, 0]

Before: [0, 2, 2, 0]
9 0 0 3
After:  [0, 2, 2, 0]

Before: [0, 3, 3, 3]
0 3 1 0
After:  [9, 3, 3, 3]

Before: [3, 2, 1, 2]
3 1 3 2
After:  [3, 2, 1, 2]

Before: [3, 1, 2, 2]
14 3 1 3
After:  [3, 1, 2, 3]

Before: [3, 1, 3, 1]
12 1 2 1
After:  [3, 3, 3, 1]

Before: [3, 3, 3, 3]
0 3 0 0
After:  [9, 3, 3, 3]

Before: [1, 3, 3, 2]
12 0 2 2
After:  [1, 3, 3, 2]

Before: [2, 0, 1, 3]
11 0 3 1
After:  [2, 0, 1, 3]

Before: [1, 0, 1, 0]
1 1 0 2
After:  [1, 0, 1, 0]

Before: [1, 2, 0, 2]
3 1 3 0
After:  [1, 2, 0, 2]

Before: [1, 0, 1, 3]
11 0 3 2
After:  [1, 0, 0, 3]

Before: [2, 0, 2, 1]
7 3 1 3
After:  [2, 0, 2, 1]

Before: [0, 2, 3, 2]
3 1 3 2
After:  [0, 2, 1, 2]

Before: [3, 1, 2, 1]
14 3 2 3
After:  [3, 1, 2, 3]

Before: [1, 0, 2, 3]
1 1 0 1
After:  [1, 1, 2, 3]

Before: [0, 1, 3, 1]
2 1 0 1
After:  [0, 1, 3, 1]

Before: [1, 0, 2, 0]
1 1 0 3
After:  [1, 0, 2, 1]

Before: [1, 2, 0, 3]
13 1 1 3
After:  [1, 2, 0, 1]

Before: [0, 2, 2, 2]
9 0 0 3
After:  [0, 2, 2, 0]

Before: [0, 2, 0, 1]
13 1 1 3
After:  [0, 2, 0, 1]

Before: [1, 2, 2, 3]
13 1 1 0
After:  [1, 2, 2, 3]

Before: [1, 0, 1, 2]
7 0 1 3
After:  [1, 0, 1, 1]

Before: [2, 3, 2, 3]
15 2 3 3
After:  [2, 3, 2, 3]

Before: [2, 3, 0, 3]
4 1 0 1
After:  [2, 1, 0, 3]

Before: [1, 0, 0, 1]
7 0 1 0
After:  [1, 0, 0, 1]

Before: [0, 1, 3, 2]
12 1 2 3
After:  [0, 1, 3, 3]

Before: [2, 2, 1, 3]
11 0 3 1
After:  [2, 0, 1, 3]

Before: [0, 3, 2, 2]
0 2 1 3
After:  [0, 3, 2, 6]

Before: [3, 0, 3, 1]
14 3 2 1
After:  [3, 3, 3, 1]

Before: [3, 3, 2, 0]
4 1 2 0
After:  [1, 3, 2, 0]

Before: [2, 0, 3, 3]
15 0 3 3
After:  [2, 0, 3, 3]

Before: [1, 0, 2, 0]
1 1 0 0
After:  [1, 0, 2, 0]

Before: [2, 3, 1, 1]
6 2 1 0
After:  [3, 3, 1, 1]

Before: [0, 1, 0, 0]
2 1 0 0
After:  [1, 1, 0, 0]

Before: [0, 2, 1, 2]
9 0 0 2
After:  [0, 2, 0, 2]

Before: [0, 2, 2, 1]
14 3 2 3
After:  [0, 2, 2, 3]

Before: [2, 3, 2, 1]
4 1 2 0
After:  [1, 3, 2, 1]

Before: [3, 3, 0, 3]
8 2 3 1
After:  [3, 3, 0, 3]

Before: [3, 3, 3, 0]
13 1 0 0
After:  [1, 3, 3, 0]

Before: [1, 3, 0, 3]
0 3 1 0
After:  [9, 3, 0, 3]

Before: [1, 0, 1, 0]
1 1 0 3
After:  [1, 0, 1, 1]

Before: [0, 0, 1, 3]
5 2 3 1
After:  [0, 3, 1, 3]

Before: [3, 1, 3, 2]
12 1 2 1
After:  [3, 3, 3, 2]

Before: [1, 3, 2, 1]
6 0 1 3
After:  [1, 3, 2, 3]

Before: [2, 0, 2, 1]
8 0 2 1
After:  [2, 4, 2, 1]

Before: [2, 0, 2, 2]
8 1 2 2
After:  [2, 0, 2, 2]

Before: [0, 0, 1, 1]
7 3 1 2
After:  [0, 0, 1, 1]

Before: [3, 0, 2, 2]
8 1 2 1
After:  [3, 2, 2, 2]

Before: [3, 1, 3, 3]
5 1 3 3
After:  [3, 1, 3, 3]

Before: [0, 1, 0, 2]
9 0 0 0
After:  [0, 1, 0, 2]

Before: [3, 1, 2, 2]
6 1 0 0
After:  [3, 1, 2, 2]

Before: [3, 0, 1, 0]
6 2 0 0
After:  [3, 0, 1, 0]

Before: [0, 1, 1, 1]
9 0 0 2
After:  [0, 1, 0, 1]

Before: [1, 2, 0, 3]
15 1 3 2
After:  [1, 2, 3, 3]

Before: [2, 3, 1, 0]
4 1 0 2
After:  [2, 3, 1, 0]

Before: [2, 1, 2, 2]
14 3 1 1
After:  [2, 3, 2, 2]

Before: [3, 2, 0, 3]
15 1 3 3
After:  [3, 2, 0, 3]

Before: [1, 3, 0, 0]
6 0 1 0
After:  [3, 3, 0, 0]

Before: [0, 1, 1, 3]
2 1 0 0
After:  [1, 1, 1, 3]

Before: [0, 3, 3, 1]
14 3 2 0
After:  [3, 3, 3, 1]

Before: [1, 3, 3, 3]
13 1 2 1
After:  [1, 1, 3, 3]

Before: [0, 3, 2, 0]
4 1 2 1
After:  [0, 1, 2, 0]

Before: [1, 0, 3, 1]
12 0 2 0
After:  [3, 0, 3, 1]

Before: [1, 1, 2, 2]
14 3 1 2
After:  [1, 1, 3, 2]

Before: [2, 1, 1, 3]
5 2 3 2
After:  [2, 1, 3, 3]

Before: [0, 2, 2, 3]
15 1 3 1
After:  [0, 3, 2, 3]

Before: [3, 3, 2, 2]
4 1 2 3
After:  [3, 3, 2, 1]

Before: [0, 1, 3, 1]
9 0 0 2
After:  [0, 1, 0, 1]

Before: [3, 1, 1, 3]
5 2 3 0
After:  [3, 1, 1, 3]

Before: [1, 0, 1, 3]
1 1 0 1
After:  [1, 1, 1, 3]

Before: [1, 0, 0, 3]
1 1 0 2
After:  [1, 0, 1, 3]

Before: [3, 2, 1, 2]
3 1 3 1
After:  [3, 1, 1, 2]

Before: [3, 2, 3, 3]
15 1 3 0
After:  [3, 2, 3, 3]

Before: [1, 2, 3, 3]
0 2 2 2
After:  [1, 2, 9, 3]

Before: [3, 0, 0, 1]
7 3 1 3
After:  [3, 0, 0, 1]

Before: [0, 1, 0, 1]
2 1 0 1
After:  [0, 1, 0, 1]

Before: [2, 1, 1, 3]
5 2 3 0
After:  [3, 1, 1, 3]

Before: [0, 1, 1, 2]
2 1 0 1
After:  [0, 1, 1, 2]

Before: [0, 0, 2, 3]
9 0 0 0
After:  [0, 0, 2, 3]

Before: [0, 2, 1, 2]
8 2 3 3
After:  [0, 2, 1, 3]

Before: [1, 0, 1, 1]
1 1 0 3
After:  [1, 0, 1, 1]

Before: [1, 0, 1, 3]
8 1 3 0
After:  [3, 0, 1, 3]

Before: [3, 1, 3, 2]
14 3 1 0
After:  [3, 1, 3, 2]

Before: [2, 2, 0, 0]
13 1 0 2
After:  [2, 2, 1, 0]

Before: [3, 1, 3, 1]
12 1 2 0
After:  [3, 1, 3, 1]

Before: [0, 3, 3, 0]
13 1 2 1
After:  [0, 1, 3, 0]

Before: [3, 1, 3, 3]
12 1 2 0
After:  [3, 1, 3, 3]

Before: [1, 2, 2, 3]
8 0 2 1
After:  [1, 3, 2, 3]

Before: [3, 0, 1, 0]
7 2 1 3
After:  [3, 0, 1, 1]

Before: [0, 1, 2, 3]
2 1 0 3
After:  [0, 1, 2, 1]

Before: [0, 0, 1, 0]
9 0 0 1
After:  [0, 0, 1, 0]

Before: [3, 2, 3, 1]
13 1 1 1
After:  [3, 1, 3, 1]

Before: [2, 2, 0, 2]
3 1 3 2
After:  [2, 2, 1, 2]

Before: [3, 3, 3, 2]
0 2 0 0
After:  [9, 3, 3, 2]

Before: [0, 0, 2, 3]
10 0 1 1
After:  [0, 1, 2, 3]

Before: [2, 3, 2, 2]
8 0 2 3
After:  [2, 3, 2, 4]

Before: [3, 3, 2, 1]
14 3 2 1
After:  [3, 3, 2, 1]

Before: [0, 0, 1, 1]
9 0 0 1
After:  [0, 0, 1, 1]

Before: [0, 2, 2, 2]
3 1 3 0
After:  [1, 2, 2, 2]

Before: [2, 2, 1, 2]
3 1 3 1
After:  [2, 1, 1, 2]

Before: [2, 1, 2, 3]
11 0 3 2
After:  [2, 1, 0, 3]

Before: [1, 2, 2, 2]
3 1 3 2
After:  [1, 2, 1, 2]

Before: [1, 3, 2, 0]
4 1 2 2
After:  [1, 3, 1, 0]

Before: [1, 1, 3, 3]
12 1 2 2
After:  [1, 1, 3, 3]

Before: [1, 2, 0, 3]
15 1 3 3
After:  [1, 2, 0, 3]

Before: [1, 2, 3, 3]
12 0 2 3
After:  [1, 2, 3, 3]

Before: [0, 2, 1, 2]
3 1 3 2
After:  [0, 2, 1, 2]

Before: [1, 1, 1, 2]
8 1 3 1
After:  [1, 3, 1, 2]

Before: [1, 1, 2, 1]
14 3 2 1
After:  [1, 3, 2, 1]

Before: [1, 0, 1, 1]
1 1 0 0
After:  [1, 0, 1, 1]

Before: [2, 1, 1, 3]
15 0 3 0
After:  [3, 1, 1, 3]

Before: [2, 1, 1, 3]
11 0 3 1
After:  [2, 0, 1, 3]

Before: [2, 3, 1, 3]
11 0 3 0
After:  [0, 3, 1, 3]

Before: [0, 3, 1, 1]
9 0 0 0
After:  [0, 3, 1, 1]

Before: [1, 0, 2, 1]
7 0 1 2
After:  [1, 0, 1, 1]

Before: [1, 0, 0, 0]
1 1 0 2
After:  [1, 0, 1, 0]

Before: [1, 0, 2, 1]
1 1 0 0
After:  [1, 0, 2, 1]

Before: [0, 3, 3, 0]
13 1 2 2
After:  [0, 3, 1, 0]

Before: [3, 2, 0, 2]
0 0 0 2
After:  [3, 2, 9, 2]

Before: [2, 1, 2, 3]
5 1 3 0
After:  [3, 1, 2, 3]

Before: [2, 2, 2, 2]
3 1 3 3
After:  [2, 2, 2, 1]

Before: [1, 3, 2, 3]
5 0 3 1
After:  [1, 3, 2, 3]

Before: [2, 1, 3, 2]
0 1 0 0
After:  [2, 1, 3, 2]

Before: [1, 0, 0, 1]
1 1 0 2
After:  [1, 0, 1, 1]

Before: [3, 3, 3, 0]
0 2 2 0
After:  [9, 3, 3, 0]

Before: [2, 2, 3, 3]
15 0 3 3
After:  [2, 2, 3, 3]

Before: [0, 3, 2, 3]
15 2 3 2
After:  [0, 3, 3, 3]

Before: [1, 2, 0, 2]
13 1 1 1
After:  [1, 1, 0, 2]

Before: [0, 3, 1, 1]
9 0 0 2
After:  [0, 3, 0, 1]

Before: [0, 1, 2, 2]
8 2 2 3
After:  [0, 1, 2, 4]

Before: [0, 0, 2, 1]
9 0 0 1
After:  [0, 0, 2, 1]

Before: [2, 3, 0, 3]
4 1 0 2
After:  [2, 3, 1, 3]

Before: [1, 0, 2, 2]
7 0 1 0
After:  [1, 0, 2, 2]

Before: [3, 1, 3, 3]
0 3 2 0
After:  [9, 1, 3, 3]

Before: [2, 3, 2, 3]
11 0 3 1
After:  [2, 0, 2, 3]

Before: [1, 0, 3, 0]
7 0 1 0
After:  [1, 0, 3, 0]

Before: [1, 3, 2, 3]
15 2 3 0
After:  [3, 3, 2, 3]

Before: [1, 3, 3, 0]
6 0 1 0
After:  [3, 3, 3, 0]

Before: [3, 0, 1, 2]
7 2 1 1
After:  [3, 1, 1, 2]

Before: [0, 3, 1, 1]
6 2 1 1
After:  [0, 3, 1, 1]

Before: [3, 3, 3, 0]
0 2 1 3
After:  [3, 3, 3, 9]

Before: [3, 2, 1, 1]
6 2 0 3
After:  [3, 2, 1, 3]

Before: [1, 0, 3, 0]
12 0 2 0
After:  [3, 0, 3, 0]

Before: [1, 2, 1, 0]
10 1 0 2
After:  [1, 2, 1, 0]

Before: [1, 3, 2, 1]
0 1 1 1
After:  [1, 9, 2, 1]

Before: [2, 3, 1, 3]
11 0 3 2
After:  [2, 3, 0, 3]

Before: [0, 1, 1, 0]
2 1 0 0
After:  [1, 1, 1, 0]

Before: [2, 2, 3, 3]
15 1 3 1
After:  [2, 3, 3, 3]

Before: [2, 3, 0, 1]
0 3 0 1
After:  [2, 2, 0, 1]

Before: [1, 0, 1, 0]
1 1 0 0
After:  [1, 0, 1, 0]

Before: [2, 2, 0, 0]
13 1 1 2
After:  [2, 2, 1, 0]

Before: [2, 0, 3, 1]
14 3 2 0
After:  [3, 0, 3, 1]

Before: [2, 1, 3, 1]
14 3 2 2
After:  [2, 1, 3, 1]

Before: [3, 3, 1, 3]
6 2 1 2
After:  [3, 3, 3, 3]

Before: [1, 2, 1, 2]
3 1 3 1
After:  [1, 1, 1, 2]

Before: [1, 2, 2, 3]
10 1 0 0
After:  [1, 2, 2, 3]

Before: [0, 1, 3, 3]
5 1 3 2
After:  [0, 1, 3, 3]

Before: [2, 1, 2, 1]
14 3 2 3
After:  [2, 1, 2, 3]

Before: [3, 3, 2, 3]
4 1 2 3
After:  [3, 3, 2, 1]

Before: [1, 3, 2, 3]
11 0 3 0
After:  [0, 3, 2, 3]

Before: [1, 2, 0, 2]
0 0 3 1
After:  [1, 2, 0, 2]

Before: [2, 2, 2, 2]
13 1 0 3
After:  [2, 2, 2, 1]

Before: [1, 3, 2, 1]
6 0 1 1
After:  [1, 3, 2, 1]

Before: [1, 3, 2, 1]
14 3 2 0
After:  [3, 3, 2, 1]

Before: [2, 3, 3, 1]
14 3 2 0
After:  [3, 3, 3, 1]

Before: [1, 2, 1, 3]
5 2 3 0
After:  [3, 2, 1, 3]

Before: [1, 1, 3, 3]
11 0 3 1
After:  [1, 0, 3, 3]

Before: [3, 1, 2, 1]
6 1 0 3
After:  [3, 1, 2, 3]

Before: [1, 0, 1, 0]
7 0 1 2
After:  [1, 0, 1, 0]

Before: [3, 0, 1, 2]
8 2 3 0
After:  [3, 0, 1, 2]

Before: [2, 3, 0, 2]
4 1 0 3
After:  [2, 3, 0, 1]

Before: [0, 2, 2, 2]
8 1 2 0
After:  [4, 2, 2, 2]

Before: [3, 2, 2, 0]
8 3 2 1
After:  [3, 2, 2, 0]

Before: [2, 1, 2, 0]
8 0 2 1
After:  [2, 4, 2, 0]

Before: [3, 1, 3, 1]
14 3 2 2
After:  [3, 1, 3, 1]

Before: [3, 3, 1, 2]
0 0 0 3
After:  [3, 3, 1, 9]

Before: [2, 1, 0, 2]
0 1 0 3
After:  [2, 1, 0, 2]

Before: [3, 3, 1, 3]
0 0 0 2
After:  [3, 3, 9, 3]

Before: [1, 0, 2, 0]
1 1 0 1
After:  [1, 1, 2, 0]

Before: [1, 2, 1, 1]
10 1 0 0
After:  [1, 2, 1, 1]

Before: [1, 0, 1, 2]
1 1 0 1
After:  [1, 1, 1, 2]



5 1 0 0
12 0 2 0
5 0 0 2
12 2 3 2
5 0 0 3
12 3 1 3
0 3 0 0
5 0 3 0
8 0 1 1
2 1 1 3
14 2 3 2
14 0 2 1
14 1 1 0
2 0 2 0
5 0 1 0
8 3 0 3
2 3 3 1
14 2 1 3
14 1 1 2
14 2 0 0
3 0 3 2
5 2 1 2
5 2 3 2
8 1 2 1
2 1 1 3
5 2 0 1
12 1 2 1
14 1 1 0
14 2 3 2
2 0 2 2
5 2 3 2
8 3 2 3
2 3 0 2
14 1 1 1
5 2 0 0
12 0 2 0
14 2 1 3
14 3 0 3
5 3 1 3
8 2 3 2
5 0 0 3
12 3 2 3
14 1 0 0
14 0 3 1
8 0 0 0
5 0 1 0
8 2 0 2
2 2 3 1
14 0 3 3
14 0 0 2
14 3 0 0
7 0 2 2
5 2 2 2
8 1 2 1
2 1 0 2
14 2 0 3
14 2 3 1
14 2 2 0
3 0 3 3
5 3 3 3
8 3 2 2
2 2 1 0
14 0 0 1
5 2 0 2
12 2 0 2
14 1 0 3
12 3 1 3
5 3 1 3
8 3 0 0
2 0 2 2
14 1 3 3
14 2 2 0
14 2 0 1
4 0 3 3
5 3 1 3
5 3 3 3
8 2 3 2
2 2 0 1
14 0 1 0
14 2 3 2
14 0 3 3
11 3 2 3
5 3 3 3
8 3 1 1
2 1 3 2
14 2 2 3
14 1 1 1
14 2 2 0
0 1 0 1
5 1 1 1
8 2 1 2
2 2 2 3
5 1 0 1
12 1 3 1
14 0 2 2
13 0 1 1
5 1 3 1
8 1 3 3
5 1 0 0
12 0 3 0
14 2 3 1
14 2 0 2
9 0 1 0
5 0 1 0
5 0 3 0
8 3 0 3
2 3 1 2
14 0 3 1
5 3 0 3
12 3 1 3
14 0 3 0
14 3 0 0
5 0 1 0
8 2 0 2
14 1 2 1
14 2 3 0
4 0 3 1
5 1 3 1
8 2 1 2
2 2 3 3
14 3 0 0
14 0 2 2
14 3 1 1
1 2 0 0
5 0 1 0
8 0 3 3
2 3 0 2
14 3 2 3
5 0 0 0
12 0 2 0
14 1 2 1
9 3 0 0
5 0 1 0
8 0 2 2
2 2 1 3
14 0 2 2
14 0 1 1
14 3 0 0
1 2 0 0
5 0 2 0
8 0 3 3
2 3 0 0
14 2 1 3
14 2 2 1
10 2 3 1
5 1 2 1
8 0 1 0
14 3 3 1
10 2 3 3
5 3 3 3
8 3 0 0
2 0 1 3
14 1 2 0
14 2 3 2
2 0 2 2
5 2 1 2
8 3 2 3
2 3 1 1
14 3 0 3
14 2 0 2
2 0 2 0
5 0 1 0
5 0 1 0
8 0 1 1
14 2 1 0
14 2 0 3
5 2 0 2
12 2 0 2
10 2 3 2
5 2 3 2
8 2 1 1
5 1 0 3
12 3 3 3
14 0 2 2
14 3 3 0
1 2 0 3
5 3 1 3
8 3 1 1
2 1 2 3
14 3 2 2
14 2 2 0
14 1 1 1
6 0 2 1
5 1 1 1
8 3 1 3
2 3 2 1
5 3 0 2
12 2 2 2
14 1 2 0
14 0 3 3
2 0 2 3
5 3 3 3
8 3 1 1
5 0 0 2
12 2 0 2
14 0 0 3
8 0 0 3
5 3 3 3
8 3 1 1
14 2 0 3
0 0 3 2
5 2 1 2
5 2 2 2
8 1 2 1
14 0 2 2
14 3 0 0
1 2 0 2
5 2 3 2
8 2 1 1
2 1 1 0
14 2 1 2
14 3 3 1
14 3 1 3
13 2 1 1
5 1 1 1
8 1 0 0
2 0 0 1
14 0 1 2
14 3 0 0
14 1 1 3
1 2 0 2
5 2 1 2
5 2 2 2
8 1 2 1
2 1 1 0
14 3 0 2
14 0 2 1
5 3 2 3
5 3 3 3
8 3 0 0
2 0 0 2
14 2 1 3
14 1 2 1
14 2 0 0
3 0 3 1
5 1 1 1
8 1 2 2
2 2 3 1
14 2 0 2
14 1 2 3
4 0 3 0
5 0 1 0
8 0 1 1
14 1 1 0
2 0 2 2
5 2 3 2
8 2 1 1
14 3 2 3
14 2 0 0
14 3 1 2
6 0 2 2
5 2 3 2
8 2 1 1
2 1 0 0
14 2 3 2
14 1 2 1
14 0 1 3
11 3 2 2
5 2 3 2
8 0 2 0
2 0 0 1
14 1 0 0
14 3 1 3
14 3 3 2
14 2 3 2
5 2 3 2
5 2 2 2
8 1 2 1
2 1 2 2
14 2 2 0
5 1 0 3
12 3 1 3
14 2 1 1
8 3 3 1
5 1 3 1
5 1 2 1
8 2 1 2
14 3 3 1
5 3 0 0
12 0 1 0
8 0 0 3
5 3 1 3
5 3 3 3
8 3 2 2
2 2 2 1
5 2 0 2
12 2 0 2
14 2 0 3
5 1 0 0
12 0 2 0
3 0 3 2
5 2 3 2
8 2 1 1
14 0 3 2
14 1 2 0
0 0 3 3
5 3 1 3
8 3 1 1
2 1 1 3
14 1 0 2
14 1 0 1
8 1 0 0
5 0 3 0
8 0 3 3
14 1 3 0
14 2 0 2
14 3 3 1
2 0 2 0
5 0 1 0
8 3 0 3
2 3 3 1
14 0 0 3
14 3 1 2
14 1 2 0
10 3 2 0
5 0 3 0
5 0 2 0
8 1 0 1
14 2 2 0
5 2 0 2
12 2 1 2
14 2 0 3
3 0 3 0
5 0 3 0
8 0 1 1
2 1 3 0
14 3 3 2
5 2 0 1
12 1 0 1
14 3 1 3
7 3 2 2
5 2 3 2
8 2 0 0
2 0 3 1
14 0 1 0
14 2 2 3
14 0 0 2
10 2 3 0
5 0 3 0
5 0 3 0
8 0 1 1
2 1 3 0
5 0 0 1
12 1 2 1
14 3 0 3
7 3 2 1
5 1 3 1
8 1 0 0
2 0 3 1
5 3 0 0
12 0 3 0
14 1 1 3
5 1 0 2
12 2 3 2
5 3 2 0
5 0 1 0
8 0 1 1
2 1 3 3
14 3 2 1
14 2 1 0
1 0 2 0
5 0 3 0
8 3 0 3
2 3 1 1
14 1 0 0
14 0 1 3
14 2 0 2
2 0 2 0
5 0 1 0
8 0 1 1
2 1 0 0
5 3 0 1
12 1 3 1
11 3 2 2
5 2 1 2
8 2 0 0
14 3 1 2
14 2 0 1
14 2 3 3
15 1 3 2
5 2 1 2
8 2 0 0
2 0 3 3
14 3 3 1
14 2 1 0
5 2 0 2
12 2 2 2
9 1 0 1
5 1 3 1
5 1 1 1
8 3 1 3
2 3 0 2
14 1 3 1
14 1 3 0
14 2 1 3
8 1 0 0
5 0 3 0
8 0 2 2
2 2 2 3
5 3 0 1
12 1 2 1
14 0 0 2
5 3 0 0
12 0 1 0
5 0 2 0
5 0 3 0
8 3 0 3
2 3 1 1
14 2 2 0
14 2 0 3
14 3 3 2
6 0 2 0
5 0 2 0
8 0 1 1
2 1 2 2
14 0 1 3
14 1 1 0
14 3 0 1
12 0 1 3
5 3 3 3
8 2 3 2
2 2 3 3
5 3 0 0
12 0 2 0
14 0 2 2
14 1 1 1
0 1 0 0
5 0 3 0
8 0 3 3
2 3 3 2
14 2 0 0
14 2 0 3
5 2 0 1
12 1 3 1
13 0 1 0
5 0 1 0
5 0 2 0
8 2 0 2
2 2 3 3
14 2 1 2
14 1 1 0
14 2 1 1
2 0 2 1
5 1 3 1
8 1 3 3
5 2 0 2
12 2 3 2
5 3 0 1
12 1 3 1
14 0 2 0
7 1 2 2
5 2 3 2
8 2 3 3
2 3 0 1
14 1 2 3
14 3 0 2
5 3 2 0
5 0 1 0
8 0 1 1
5 0 0 3
12 3 2 3
14 1 1 0
14 1 1 2
0 0 3 0
5 0 3 0
8 0 1 1
2 1 1 2
14 0 2 1
14 2 1 0
3 0 3 0
5 0 2 0
8 2 0 2
2 2 1 0
5 0 0 2
12 2 2 2
14 3 0 1
5 3 0 3
12 3 0 3
15 2 3 2
5 2 2 2
8 2 0 0
2 0 3 2
14 1 2 1
14 2 3 0
14 3 1 3
9 3 0 3
5 3 3 3
8 2 3 2
5 0 0 1
12 1 3 1
14 1 0 3
4 0 3 3
5 3 3 3
5 3 1 3
8 3 2 2
2 2 0 1
14 2 1 3
5 3 0 2
12 2 0 2
14 3 2 0
10 2 3 3
5 3 2 3
8 1 3 1
2 1 2 0
14 1 3 2
5 1 0 3
12 3 0 3
14 0 2 1
14 3 2 2
5 2 1 2
5 2 2 2
8 0 2 0
14 1 3 1
14 0 3 2
5 1 2 1
5 1 2 1
5 1 1 1
8 1 0 0
5 1 0 2
12 2 3 2
14 2 3 1
6 1 2 2
5 2 1 2
8 2 0 0
2 0 2 3
14 1 2 0
14 0 2 2
5 0 2 2
5 2 3 2
8 2 3 3
14 0 3 1
5 2 0 2
12 2 2 2
2 0 2 2
5 2 2 2
8 3 2 3
2 3 0 0
14 2 0 2
14 1 2 1
14 0 2 3
15 2 3 3
5 3 1 3
8 3 0 0
2 0 1 1
14 2 0 3
14 1 3 0
14 1 1 2
0 0 3 2
5 2 2 2
8 2 1 1
2 1 0 3
5 3 0 1
12 1 1 1
14 3 3 2
5 0 2 2
5 2 3 2
8 3 2 3
2 3 2 1
14 0 2 3
5 0 0 0
12 0 2 0
14 3 1 2
6 0 2 3
5 3 3 3
8 3 1 1
2 1 2 3
14 1 3 1
14 2 2 2
0 1 0 0
5 0 1 0
8 0 3 3
14 2 3 0
14 3 2 2
1 0 2 2
5 2 2 2
8 2 3 3
2 3 0 1
14 1 3 2
14 1 0 3
14 1 3 0
8 0 0 3
5 3 1 3
5 3 3 3
8 3 1 1
2 1 3 3
14 3 0 1
14 2 2 2
14 3 2 0
6 2 0 2
5 2 1 2
5 2 3 2
8 2 3 3
2 3 3 1
14 2 3 2
5 2 0 3
12 3 0 3
11 3 2 2
5 2 2 2
8 1 2 1
14 3 0 2
10 3 2 3
5 3 2 3
5 3 3 3
8 3 1 1
2 1 2 3
14 2 2 2
14 3 0 1
6 2 0 0
5 0 1 0
5 0 3 0
8 0 3 3
2 3 2 1
14 2 1 0
14 1 1 3
14 1 2 2
4 0 3 2
5 2 3 2
8 1 2 1
2 1 1 2
14 0 2 3
14 3 2 1
15 0 3 1
5 1 2 1
5 1 2 1
8 1 2 2
14 3 1 0
14 1 3 3
14 0 2 1
12 3 1 3
5 3 2 3
5 3 1 3
8 2 3 2
2 2 1 1
14 0 1 0
14 3 0 3
14 3 3 2
14 2 3 3
5 3 2 3
5 3 3 3
8 3 1 1
14 3 3 3
14 2 1 0
14 0 1 2
14 2 3 3
5 3 2 3
8 1 3 1
2 1 2 3
14 3 3 2
14 3 3 0
14 2 2 1
6 1 0 1
5 1 3 1
8 1 3 3
2 3 2 1
14 3 2 3
14 2 0 2
13 2 0 2
5 2 1 2
8 2 1 1
2 1 2 0
14 1 0 3
14 3 3 1
14 3 3 2
14 2 1 1
5 1 1 1
5 1 2 1
8 0 1 0
2 0 3 3
5 0 0 1
12 1 1 1
14 0 1 2
14 3 0 0
1 2 0 1
5 1 1 1
5 1 2 1
8 3 1 3
2 3 1 1
5 2 0 2
12 2 2 2
14 3 1 3
6 2 0 3
5 3 3 3
5 3 2 3
8 1 3 1
5 3 0 2
12 2 0 2
14 2 1 3
5 2 0 0
12 0 2 0
3 0 3 3
5 3 2 3
8 3 1 1
2 1 2 2
5 1 0 0
12 0 3 0
14 1 1 3
14 3 2 1
12 3 1 1
5 1 2 1
5 1 1 1
8 1 2 2
2 2 1 1
14 2 0 2
5 2 0 3
12 3 0 3
5 0 0 0
12 0 1 0
11 3 2 0
5 0 3 0
8 1 0 1
2 1 1 2
5 3 0 1
12 1 1 1
14 2 0 3
5 2 0 0
12 0 2 0
3 0 3 3
5 3 1 3
5 3 1 3
8 2 3 2
2 2 1 1
14 0 2 0
5 3 0 2
12 2 3 2
14 0 1 3
14 2 3 2
5 2 3 2
8 2 1 1
2 1 2 3
14 2 3 1
14 2 1 0
5 0 0 2
12 2 3 2
6 1 2 0
5 0 2 0
8 3 0 3
2 3 1 1
14 0 0 2
5 0 0 3
12 3 1 3
14 2 3 0
8 3 3 0
5 0 3 0
8 1 0 1
5 0 0 0
12 0 2 0
14 3 2 2
14 2 2 3
3 0 3 2
5 2 2 2
8 2 1 1
14 3 2 2
3 0 3 3
5 3 1 3
5 3 1 3
8 1 3 1
2 1 3 3
14 0 3 2
14 0 0 1
5 0 0 0
12 0 1 0
8 0 0 2
5 2 2 2
8 2 3 3
2 3 0 1
5 3 0 3
12 3 2 3
14 0 0 2
14 2 3 0
10 2 3 0
5 0 3 0
8 0 1 1
14 2 3 2
14 1 2 0
8 0 0 0
5 0 2 0
5 0 3 0
8 0 1 1
2 1 2 3
14 2 3 1
14 1 2 0
8 0 0 0
5 0 3 0
8 0 3 3
2 3 0 1
5 2 0 3
12 3 2 3
14 1 3 0
2 0 2 3
5 3 2 3
8 1 3 1
2 1 0 2
14 2 1 1
14 3 1 3
14 3 3 0
6 1 0 3
5 3 3 3
8 2 3 2
14 0 2 3
14 1 0 0
5 0 0 1
12 1 0 1
14 3 0 0
5 0 2 0
8 0 2 2
2 2 1 0
14 2 0 2
14 1 2 1
11 3 2 1
5 1 3 1
5 1 2 1
8 1 0 0
2 0 0 2
14 1 0 3
14 2 3 0
5 0 0 1
12 1 0 1
4 0 3 3
5 3 2 3
5 3 1 3
8 2 3 2
2 2 3 0
5 0 0 2
12 2 3 2
14 1 3 1
14 1 1 3
14 3 1 3
5 3 2 3
8 3 0 0
2 0 0 2
5 0 0 3
12 3 2 3
14 0 3 1
14 2 2 0
3 0 3 1
5 1 3 1
5 1 3 1
8 2 1 2
5 1 0 1
12 1 2 1
14 3 0 0
9 0 1 3
5 3 1 3
8 2 3 2
2 2 2 1
14 2 3 0
14 2 2 3
14 3 0 2
6 0 2 3
5 3 1 3
5 3 1 3
8 1 3 1
2 1 0 3
14 3 3 0
14 2 2 1
9 0 1 1
5 1 1 1
8 1 3 3
14 1 0 2
14 3 0 1
14 1 0 0
7 1 2 0
5 0 3 0
5 0 2 0
8 3 0 3
5 1 0 0
12 0 3 0
14 1 0 1
7 0 2 2
5 2 1 2
8 3 2 3
2 3 1 2
14 2 3 3
14 1 2 0
0 1 3 1
5 1 2 1
8 2 1 2
2 2 3 1
5 1 0 0
12 0 2 0
14 0 3 2
3 0 3 0
5 0 3 0
5 0 3 0
8 1 0 1
2 1 2 2
14 2 2 0
14 0 3 1
14 0 1 3
14 3 0 0
5 0 2 0
8 0 2 2
2 2 2 3
14 1 3 0
5 1 0 1
12 1 1 1
14 0 2 2
8 1 0 1
5 1 3 1
5 1 1 1
8 1 3 3
2 3 0 0`,
}
