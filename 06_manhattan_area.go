package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type point struct {
	x, y int
}

const dim = 402

// Print map to confirm initial solution was correct, and AdventOfCode checker was wrong.
// Issue was fixed in 1:42 past unlock.
func _print(m [dim][dim]int, ps []point) {
	x0, x1, y0, y1 := 1000, 0, 1000, 0
	for _, p := range ps {
		if x0 > p.x {
			x0 = p.x
		}
		if x1 < p.x {
			x1 = p.x
		}
		if y0 > p.y {
			y0 = p.y
		}
		if y1 < p.y {
			y1 = p.y
		}
	}
	x0--
	x1++
	y0--
	y1++
	_log("bounds", x0, x1, y0, y1)
	ss := []rune("-123456789abcdefghijklmnopqrstuvwxyz/\\!%^&*()[]<>{}+:;")
	mr := make([][]rune, 0, y1-y0+1)
	for y := y0; y <= y1; y++ {
		sr := make([]rune, 0, x1-x0+1)
		for x := x0; x <= x1; x++ {
			sr = append(sr, ss[m[y][x]])
		}
		mr = append(mr, sr)
	}
	for _, p := range ps {
		mr[p.y-y0][p.x-x0] = '.'
	}
	for _, sr := range mr {
		_log(string(sr))
	}
}

func prepare(ss []string) []point {
	ps := make([]point, len(ss))
	for i, s := range ss {
		vs := strings.Split(s, ", ")
		x, _ := strconv.Atoi(vs[0])
		y, _ := strconv.Atoi(vs[1])
		ps[i] = point{x, y}
	}
	return ps
}

func part1(ps []point) int {
	m := [dim][dim]int{}
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			mins := []int{}
			mind := 1000
			for i, p := range ps {
				d := abs(x-p.x) + abs(y-p.y)
				switch {
				case d == mind:
					mins = append(mins, i+1)
				case d < mind:
					mind = d
					mins = []int{i + 1}
				}
			}
			if len(mins) == 1 {
				m[y][x] = mins[0]
			}
		}
	}
	//_print(m, ps)
	as := make([]int, len(ps)+1)
	for y := 1; y < dim-1; y++ {
		for x := 1; x < dim-1; x++ {
			as[m[y][x]]++
		}
	}
	for i := 0; i < dim; i++ {
		as[m[0][i]] = 0
		as[m[dim-1][i]] = 0
		as[m[i][0]] = 0
		as[m[i][dim-1]] = 0
	}
	sort.Ints(as)
	return as[len(as)-1]
}

func part2(ps []point) int {
	a := 0
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			sumd := 0
			for _, p := range ps {
				sumd += abs(y-p.x) + abs(x-p.y)
			}
			if sumd < 10000 {
				a++
			}
		}
	}
	return a
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		ss := strings.Split(in, "\n")
		p := prepare(ss)
		fmt.Println("part 1:", part1(p))
		fmt.Println("part 2:", part2(p))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `137, 282
229, 214
289, 292
249, 305
90, 289
259, 316
134, 103
96, 219
92, 308
269, 59
141, 132
71, 200
337, 350
40, 256
236, 105
314, 219
295, 332
114, 217
43, 202
160, 164
245, 303
339, 277
310, 316
164, 44
196, 335
228, 345
41, 49
84, 298
43, 51
158, 347
121, 51
176, 187
213, 120
174, 133
259, 263
210, 205
303, 233
265, 98
399, 332
186, 340
132, 99
174, 153
206, 142
341, 162
180, 166
152, 249
221, 118
95, 227
152, 186
72, 330`,
	"google": `177, 51
350, 132
276, 139
249, 189
225, 137
337, 354
270, 147
182, 329
118, 254
174, 280
42, 349
96, 341
236, 46
84, 253
292, 143
253, 92
224, 137
209, 325
243, 195
208, 337
197, 42
208, 87
45, 96
64, 295
266, 248
248, 298
194, 261
157, 74
52, 248
243, 201
242, 178
140, 319
69, 270
314, 302
209, 212
237, 217
86, 294
295, 144
248, 206
157, 118
155, 146
331, 40
247, 302
250, 95
193, 214
345, 89
183, 206
121, 169
79, 230
88, 155`}
