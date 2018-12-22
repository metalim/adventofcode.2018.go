package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	. "github.com/logrusorgru/aurora"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
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

//
// Solution
//

// pStep returns position of step in specified direction. 0123 -> ESWN
func pStep(p pos, d int) pos {
	return pos{p.X + (1-d)%2, p.Y + (2-d)%2}
}

type task struct {
	d      int   // depth
	t      pos   // target
	e      map2d // erosion
	r      map2d // risk == type
	bounds rect
}

func parse(in [3]int) task {
	return task{d: in[0], t: pos{in[1], in[2]}}
}

////////////////////////////////////////////////////////////////////////
// build a map
//

func (t *task) process() {
	w := t.t.X + 1
	h := t.t.Y + 1
	// In theory it is possible to build a map, that needs 6xLargestDim extension from target position.
	w += h // extend map.
	h *= 2
	t.bounds = rect{pos{0, 0}, pos{w, h}}
	t.e = makeMap2d(w, h)
	t.r = makeMap2d(w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := pos{x, y}
			var g int
			switch {
			case p == t.t:
			case y == 0:
				g = x * 16807
			case x == 0:
				g = y * 48271
			default:
				g = t.e.get(pos{x - 1, y}) * t.e.get(pos{x, y - 1})
			}
			e := (g + t.d) % 20183
			t.e.set(p, e)
			r := e % 3
			t.r.set(p, r)
		}
	}
	t.e.lock()
	t.r.lock()
}

////////////////////////////////////////////////////////////////////////
// part 1
//

func (t *task) part1() (sum int) {
	for y := 0; y <= t.t.Y; y++ {
		for x := 0; x <= t.t.X; x++ {
			sum += t.r.get(pos{x, y})
		}
	}
	return sum
}

////////////////////////////////////////////////////////////////////////
// part 2
//

type hand int

const (
	nothing hand = iota
	torch
	gear
	change = 7
)

type step struct {
	p      pos
	h      hand
	future int
}

func (t *task) part2() int {
	ms := []map2d{ // for different tools
		makeMap2d(t.bounds.Max.X, t.bounds.Max.Y),
		makeMap2d(t.bounds.Max.X, t.bounds.Max.Y),
		makeMap2d(t.bounds.Max.X, t.bounds.Max.Y),
	}
	next := []step{{pos{0, 0}, torch, 0}}
	cur := []step{}
	tick := 1 // start from 1! because default field values are 0.

	var lock sync.Mutex
	done := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-done:
				return
			default:
				lock.Lock()
				t.print3Map(ms) // doesn't happen, unless bugs prevent getting the result in 1 second.
				_log(tick, len(cur), len(next))
				lock.Unlock()
			}
		}
	}()

	// BFS, but with 3 maps.
	for ; len(next) > 0; tick++ {

		lock.Lock()
		cur, next = next, cur[:0]

		found := ms[torch].get(t.t)
		if found > 0 && tick > found {
			done <- true
			return found - 1 // we started with 1, remember?
		}

		for _, s := range cur {

			v := ms[s.h].get(s.p)
			if s.future < 0 { // still changing hand? -> update map.
				if s.future != v { // it's not our counter? -> skip this step.
					continue
				}
				s.future++
				ms[s.h].set(s.p, s.future)
				next = append(next, s)
				continue
			}

			if v > 0 && v < tick { // arrived later with same hand? -> skip.
				continue
			}

			r := hand(t.r.get(s.p)) // forbidden hand for this place.

			for _, h := range [3]hand{nothing, torch, gear} { // set arrival ticks for each allowed hand.
				if h == r { // skip forbidden hand.
					continue
				}
				th := tick
				if h != s.h { // arrived with another hand? -> add 7 ticks to change it.
					th += change
				}
				tx := ms[h].get(s.p)
				if tx == 0 || tx > th { // if we can do faster
					ms[h].set(s.p, th)
				}
			}

			// now process next steps

			for d := 0; d < 4; d++ {
				p := pStep(s.p, d)

				if !p.In(t.bounds) {
					continue
				}

				// 0:   1,2
				// 1: 0,  2
				// 2: 0,1

				// 0+1 / 1+0 -> 2
				// 1+2 / 2+1 -> 0
				// 0+2 / 2+0 -> 1

				ns := step{p, s.h, 0}
				r2 := hand(t.r.get(p)) // second forbidden hand
				if ns.h == r2 {        // is current hand forbidden for next step? -> change hand.
					ns.h = 3 - r - r2
					ns.future = -change
				}

				visited := ms[ns.h].get(p)
				if visited > 0 && visited <= tick+1-ns.future { // is it already visited, and tick is not larger, than we can provide? -> skip.
					continue
				}

				if visited < 0 && visited >= ns.future { // not visited yet, but somebody arrives earlier? -> skip.
					continue
				}

				if ns.future == 0 {
					ms[ns.h].set(ns.p, tick+1) // next tick's value
				} else {
					ms[ns.h].set(ns.p, ns.future) // -7
				}
				next = append(next, ns)
			}
		}
		// time.Sleep(10 * time.Millisecond) // uncomment for fancy tracing.
		lock.Unlock()
	}
	log.Fatal("path not found")
	return 0
}

// print3Map prints 3 maps merged and colored. Optional.
func (t *task) print3Map(ms []map2d) {
	var b rect
	for _, m := range ms {
		b = b.Union(m.b)
	}
	b = b.Intersect(rect{pos{0, b.Max.Y - 15}, pos{20, b.Max.Y}})

	lim := strings.Repeat("-----", b.Dx())
	_log(lim)
	out := make([]string, 0, b.Dx())
	for y := b.Min.Y; y < b.Max.Y; y++ {
		out = out[:0]
		for x := b.Min.X; x < b.Max.X; x++ {
			im := 0
			vm := 0
			for i, m := range ms {
				v := m.get(pos{x, y})
				if v < 0 && (vm < v || vm >= 0) {
					vm = v
					im = i
				} else if v > 0 && (vm > v || vm == 0) {
					vm = v
					im = i
				}
			}
			if vm == 0 {
				out = append(out, "    "+Black([]string{".", "=", "|"}[t.r.get(pos{x, y})]).Bold().String())
				continue
			}

			col := Gray
			switch im {
			case 1:
				col = Red
			case 2:
				col = Green
			}

			sn := fmt.Sprintf(col("%5d").String(), vm)
			out = append(out, sn)
		}
		_log(strings.Join(out, ""))
	}
	_log(lim)
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
	test1 := func(in [3]int, ex int) {
		t := parse(in)
		t.process()
		verify(t.part1(), ex)
	}
	test1([3]int{510, 10, 10}, 114)
	fmt.Println("tests passed", Black(time.Since(t0)).Bold())
}

func main() {
	test()
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

var ins = map[string][3]int{
	"github": {11820, 7, 782},
	"google": {4845, 6, 770},
}
