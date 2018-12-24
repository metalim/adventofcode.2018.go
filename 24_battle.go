package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

var debug = 1

func _log(a ...interface{}) {
	if debug > 0 {
		fmt.Println(a...)
	}
}

func _print(a ...interface{}) {
	if debug >= 0 {
		fmt.Println(a...)
	}
}

func sliceAtoi(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

////////////////////////////////////////////////////////////////////////
// Solution
//

type attack = string
type group struct {
	n, hp int
	def   map[attack]int
	pow   int
	at    attack
	init  int
	id    string
}
type army []*group
type task struct {
	in  string
	imm army
	inf army
}

func parse(in string) (t task) {
	r := regexp.MustCompile("(\\d+) units each with (\\d+) hit points (\\(.*\\) )?with an attack that does (\\d+) (\\w+) damage at initiative (\\d+)")
	ri := regexp.MustCompile("immune to ([^;)]+)")
	rw := regexp.MustCompile("weak to ([^;)]+)")
	ins := strings.Split(in, "\n\n")

	parseArmy := func(sa, name string) (ar army) {
		ss := strings.Split(sa, "\n")
		for i, s := range ss[1:] {
			if len(s) == 0 {
				continue
			}
			m := r.FindStringSubmatch(s)
			m = m[1:]
			mn := sliceAtoi(m)
			mi := ri.FindStringSubmatch(m[2])
			mw := rw.FindStringSubmatch(m[2])
			def := make(map[attack]int)
			if len(mi) > 0 {
				for _, k := range strings.Split(mi[1], ", ") {
					def[k] = 1
				}
			}
			if len(mw) > 0 {
				for _, k := range strings.Split(mw[1], ", ") {
					def[k] = -1
				}
			}
			g := group{n: mn[0], hp: mn[1], def: def, pow: mn[3], at: m[4], init: mn[5], id: name + strconv.Itoa(i+1)}
			ar = append(ar, &g)
		}
		return ar
	}

	t.in = in
	t.imm = parseArmy(ins[0], "Immune")
	t.inf = parseArmy(ins[1], "Infect")
	_log(len(t.imm), len(t.inf))
	return t
}

func (t *task) process() {
}

type fight struct {
	g, a *group
}

func (g *group) getDmgTo(g2 *group) int {
	dmg := g.n * g.pow * (1 - g2.def[g.at])
	return dmg
}

func (g *group) chooseFrom(ar army, taken map[*group]bool) (fight, bool) {
	var best *group
	bestDmg := -1

	for i := range ar {
		if taken[ar[i]] {
			continue
		}

		dmg := g.getDmgTo(ar[i])

		if bestDmg > dmg {
			continue
		}

		if bestDmg == dmg {
			if best.n*best.pow > ar[i].n*ar[i].pow {
				continue
			}

			if best.n*best.pow == ar[i].n*ar[i].pow {
				if best.init > ar[i].init {
					continue
				}
			}
		}

		// else -> choose new best.
		bestDmg = dmg
		best = ar[i]

	}
	if bestDmg <= 0 {
		return fight{g, nil}, false
	}
	taken[best] = true

	return fight{g, best}, bestDmg >= best.hp
}

type byPowInit army

func (a byPowInit) Len() int { return len(a) }
func (a byPowInit) Less(i, j int) bool {
	return a[i].n*a[i].pow > a[j].n*a[j].pow || a[i].n*a[i].pow == a[j].n*a[j].pow && a[i].init > a[j].init
}
func (a byPowInit) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (ar *army) chooseTargets(enemies army, queue map[int]fight) {
	taken := make(map[*group]bool)
	sort.Sort(byPowInit(*ar))
	for _, g := range *ar {
		if f, ok := g.chooseFrom(enemies, taken); ok {
			_log(g.id, g.n*g.pow, g.init, "will attack", f.a.id, "for", g.getDmgTo(f.a))
			queue[g.init] = f
			continue
		}
		_log(g.id, g.n*g.pow, g.init, "will skip")
	}
}

func (o fight) attack() {
	if o.g.n <= 0 {
		_log(o.g.id, "can't attack: it is dead")
		return
	}
	dmg := o.g.getDmgTo(o.a)
	kill := dmg / o.a.hp
	if kill > o.a.n {
		kill = o.a.n
	}
	// _log(o)
	// _log("attack", o.g.at, "defence", o.a.def)
	o.a.n -= kill
	_log(o.g.id, "deals", dmg, "dmg, and kills", kill, o.a.id, ", remaining:", o.a.n)
}

func (ar *army) removeDead() {
	var out int
	for _, g := range *ar {
		if g.n > 0 {
			(*ar)[out] = g
			out++
		}
	}
	*ar = (*ar)[:out]
}

func (ar *army) totalUnits() (sum int) {
	for _, g := range *ar {
		if g.n > 0 {
			sum += g.n
		}
	}
	return sum
}

////////////////////////////////////////////////////////////////////////
// part 1
//

func (t *task) part1() int {
	max := len(t.imm) + len(t.inf)

	for step := 1; len(t.imm) > 0 && len(t.inf) > 0; step++ {
		queue := make(map[int]fight)

		_log("\nstep", step)
		// _log("\nimm->")
		t.imm.chooseTargets(t.inf, queue)
		// _log("imm->inf", queue)

		// _log("\ninf->")
		t.inf.chooseTargets(t.imm, queue)
		// _log("inf->imm", queue)

		if len(queue) == 0 { // dead lock
			t.imm = army{}
			t.inf = army{}
			return 0
		}
		for i := max; i > 0; i-- {
			if o, ok := queue[i]; ok {
				o.attack()
			}
		}

		// remove dead groups
		t.imm.removeDead()
		t.inf.removeDead()
	}
	if len(t.imm) > 0 {
		return t.imm.totalUnits()
	}
	return t.inf.totalUnits()
}

////////////////////////////////////////////////////////////////////////
// part 2
//

func (t *task) part2() int {
	boost := 1
	d := 256

	for {
		*t = parse(t.in)
		for _, g := range t.imm {
			g.pow += boost
		}
		t.part1()
		if len(t.imm) <= 0 {
			boost += d
			continue
		}
		if d == 1 {
			_print(Black("boost is").Bold(), Cyan(boost))
			return t.imm.totalUnits()
		}
		d /= 2
		boost -= d
	}
	return 0
}

////////////////////////////////////////////////////////////////////////
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
	test2 := func(in string, ex int) {
		t := parse(in)
		t.process()
		verify(t.part2(), ex)
	}
	in := `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`
	test1(in, 5216)
	test2(in, 51)
	fmt.Println("tests passed", Black(time.Since(t0)).Bold())
}

////////////////////////////////////////////////////////////////////////
// main
//
func main() {
	debug = -1
	test()
	debug = 0

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
366 units each with 28071 hit points (weak to cold, slashing) with an attack that does 150 fire damage at initiative 16`,

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
7933 units each with 29900 hit points (immune to bludgeoning, radiation, slashing) with an attack that does 5 slashing damage at initiative 20`,
}
