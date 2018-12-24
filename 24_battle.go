package main

import (
	"fmt"
	"image"
	"log"
	"os"
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
type pos = image.Point
type rect = image.Rectangle

type row []cell
type field []row

// map2d = field with bounds.
type map2d struct {
	f      field
	cap, b rect
	locked bool
}

func makeMap2d(w, h int) map2d {
	f := make(field, 0, h)
	for y := 0; y < h; y++ {
		f = append(f, make(row, w))
	}
	return map2d{f: f, cap: rect{pos{0, 0}, pos{w, h}}}
}

func (m *map2d) get(p pos) cell {
	return m.f[p.Y][p.X]
}

func (m *map2d) lock() {
	m.locked = true
}

func (m *map2d) set(p pos, c cell) {
	if m.locked {
		panic("map is locked")
	}
	m.f[p.Y][p.X] = c
	if !p.In(m.b) {
		m.b = m.b.Union(rect{p, p.Add(pos{1, 1})})
	}
}

// pStep returns position of step in specified direction. 0123 -> ESWN
func pStep(p pos, d int) pos {
	return pos{p.X + (1-d)%2, p.Y + (2-d)%2}
}

// pStep2 returns position of step in specified direction. 0123 -> NEWS
// [0,-1], [-1,0], [1,0], [0,1]
func pStep2(p pos, d int) pos {
	x := p.X + (1-d)%2
	y := p.Y + (2-d)%2
	panic("WIP")
	x = d % 3 * (1)
	y = 1
	return pos{x, y}
}

////////////////////////////////////////////////////////////////////////
// Solution
//

type attack = string
type group struct {
	n, hp  int
	im, we []attack
	pow    int
	at     attack
	init   int
}
type army []group
type task struct {
	imm army
	inf army
}

func parse(in string) (t task) {
	r := regexp.MustCompile("(\\d+) units each with (\\d+) hit points (\\(.*\\) )?with an attack that does (\\d+) (\\w+) damage at initiative (\\d+)")
	ri := regexp.MustCompile("immune to ([^;)]+)")
	rw := regexp.MustCompile("weak to ([^;)]+)")
	ins := strings.Split(in, "\n\n")

	parseArmy := func(sa string) (ar army) {
		ss := strings.Split(sa, "\n")
		for _, s := range ss[1:] {
			if len(s) == 0 {
				continue
			}
			m := r.FindStringSubmatch(s)
			m = m[1:]
			mn := sliceAtoi(m)
			mi := ri.FindStringSubmatch(m[2])
			mw := rw.FindStringSubmatch(m[2])
			var im, we []string
			if len(mi) > 0 {
				im = strings.Split(mi[1], ", ")
			}
			if len(mw) > 0 {
				we = strings.Split(mw[1], ", ")
			}
			g := group{n: mn[0], hp: mn[1], im: im, we: we, pow: mn[3], at: m[4], init: mn[5]}
			ar = append(ar, g)
		}
		return ar
	}

	t.imm = parseArmy(ins[0])
	t.inf = parseArmy(ins[1])
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
	test1(`Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4
`, 1)
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
	"github": `Immune System:
4647 units each with 7816 hit points with an attack that does 13 fire damage at initiative 1
301 units each with 3152 hit points (weak to fire) with an attack that does 104 fire damage at initiative 3
1508 units each with 8344 hit points with an attack that does 50 cold damage at initiative 9
2956 units each with 5021 hit points (weak to slashing; immune to bludgeoning) with an attack that does 13 slashing damage at initiative 15
898 units each with 11545 hit points with an attack that does 100 cold damage at initiative 2
207 units each with 6235 hit points (weak to cold) with an attack that does 242 slashing damage at initiative 17
7550 units each with 8773 hit points (immune to radiation; weak to fire, slashing) with an attack that does 11 radiation damage at initiative 11
1057 units each with 3791 hit points (immune to cold) with an attack that does 27 bludgeoning damage at initiative 5
5086 units each with 3281 hit points (weak to bludgeoning) with an attack that does 5 cold damage at initiative 19
330 units each with 4136 hit points with an attack that does 91 cold damage at initiative 6

Infection:
1755 units each with 6886 hit points (immune to slashing, radiation) with an attack that does 6 fire damage at initiative 4
2251 units each with 33109 hit points with an attack that does 29 cold damage at initiative 7
298 units each with 18689 hit points (weak to radiation, slashing) with an attack that does 123 slashing damage at initiative 13
312 units each with 15735 hit points (weak to bludgeoning, slashing) with an attack that does 99 cold damage at initiative 8
326 units each with 16400 hit points (weak to bludgeoning) with an attack that does 98 radiation damage at initiative 20
4365 units each with 54947 hit points with an attack that does 22 cold damage at initiative 14
1446 units each with 51571 hit points (weak to cold) with an attack that does 63 fire damage at initiative 18
8230 units each with 12331 hit points (weak to bludgeoning; immune to slashing) with an attack that does 2 fire damage at initiative 12
4111 units each with 17381 hit points with an attack that does 7 cold damage at initiative 10
366 units each with 28071 hit points (weak to cold, slashing) with an attack that does 150 fire damage at initiative 16
`,
	"google": `Immune System:
3400 units each with 1430 hit points (immune to fire, radiation, slashing) with an attack that does 4 radiation damage at initiative 4
138 units each with 8650 hit points (weak to bludgeoning; immune to slashing, cold, radiation) with an attack that does 576 slashing damage at initiative 16
255 units each with 9469 hit points (weak to radiation, fire) with an attack that does 351 bludgeoning damage at initiative 8
4145 units each with 2591 hit points (immune to cold; weak to slashing) with an attack that does 6 fire damage at initiative 12
3605 units each with 10989 hit points with an attack that does 26 fire damage at initiative 19
865 units each with 11201 hit points with an attack that does 102 slashing damage at initiative 10
633 units each with 10092 hit points (weak to slashing, radiation) with an attack that does 150 slashing damage at initiative 11
2347 units each with 3322 hit points with an attack that does 12 cold damage at initiative 2
7045 units each with 3877 hit points (weak to radiation) with an attack that does 5 bludgeoning damage at initiative 5
1086 units each with 8626 hit points (weak to radiation) with an attack that does 69 slashing damage at initiative 13

Infection:
2152 units each with 12657 hit points (weak to fire, cold) with an attack that does 11 fire damage at initiative 18
40 units each with 39458 hit points (immune to radiation, fire, slashing; weak to bludgeoning) with an attack that does 1519 slashing damage at initiative 7
59 units each with 35138 hit points (immune to radiation; weak to fire) with an attack that does 1105 fire damage at initiative 15
1569 units each with 51364 hit points (weak to radiation) with an attack that does 55 radiation damage at initiative 17
929 units each with 23887 hit points (weak to bludgeoning) with an attack that does 48 cold damage at initiative 14
5264 units each with 14842 hit points (immune to cold, fire; weak to slashing, bludgeoning) with an attack that does 4 bludgeoning damage at initiative 9
1570 units each with 30419 hit points (weak to radiation, cold; immune to fire) with an attack that does 35 slashing damage at initiative 1
1428 units each with 21393 hit points (weak to radiation) with an attack that does 29 cold damage at initiative 6
1014 units each with 25717 hit points (weak to fire) with an attack that does 47 fire damage at initiative 3
7933 units each with 29900 hit points (immune to bludgeoning, radiation, slashing) with an attack that does 5 slashing damage at initiative 20
`,
}
