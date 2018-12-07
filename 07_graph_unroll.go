package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

type rset map[rune]bool
type rmap map[rune]rset
type prep struct {
	before, after rmap
}

func prepare(ss []string) prep {
	before := rmap{}
	after := rmap{}
	r := regexp.MustCompile("Step (\\w) must be finished before step (\\w) can begin.")
	for r := 'A'; r <= 'Z'; r++ {
		before[r] = rset{}
		after[r] = rset{}
	}
	for _, s := range ss {
		m := r.FindStringSubmatch(s)
		r1 := []rune(m[1])[0]
		r2 := []rune(m[2])[0]
		before[r1][r2] = true
		after[r2][r1] = true
	}
	return prep{before, after}
}

func part1(p prep) string {
	out := []rune{}
	for len(p.after) > 0 {
		for r := 'A'; r <= 'Z'; r++ {
			if m, ok := p.after[r]; ok && len(m) == 0 {
				out = append(out, r)
				for k := range p.before[r] {
					delete(p.after[k], r)
				}
				delete(p.after, r)
				break
			}
		}
	}
	return string(out)
}

type worker struct {
	r rune
	t int
	//s string // for debug printing
}

func part2(p prep) int {
	t := 0
	ws := make([]worker, 0, 5)
	ws2 := make([]worker, 0, 5)
	taken := rset{}
	for len(p.after) > 0 {
		for r := 'A'; r <= 'Z'; r++ {
			if m, ok := p.after[r]; ok && len(m) == 0 && !taken[r] {
				ws = append(ws, worker{r, int(r-'A') + 61 /*, string(r)*/})
				taken[r] = true
			}
			if len(ws) == 5 {
				break
			}
		}

		mint := ws[0].t
		for _, w := range ws[1:] {
			if mint > w.t {
				mint = w.t
			}
		}
		t += mint
		for _, w := range ws {
			w.t -= mint
			if w.t > 0 {
				ws2 = append(ws2, w)
			} else {
				for k := range p.before[w.r] {
					delete(p.after[k], w.r)
				}
				delete(p.after, w.r)
			}
		}
		ws, ws2 = ws2, ws // just to avoid new slice allocation
		ws2 = ws2[:0]
	}
	return t
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		ss := strings.Split(in, "\n")
		t0 := time.Now()
		p := prepare(ss)
		fmt.Println("part 1:", part1(p))
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		p = prepare(ss)
		fmt.Println("part 2:", part2(p))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `Step P must be finished before step F can begin.
Step F must be finished before step M can begin.
Step Q must be finished before step S can begin.
Step K must be finished before step G can begin.
Step W must be finished before step X can begin.
Step V must be finished before step I can begin.
Step S must be finished before step Y can begin.
Step U must be finished before step D can begin.
Step J must be finished before step B can begin.
Step Z must be finished before step C can begin.
Step Y must be finished before step D can begin.
Step X must be finished before step A can begin.
Step E must be finished before step N can begin.
Step M must be finished before step B can begin.
Step N must be finished before step I can begin.
Step I must be finished before step T can begin.
Step H must be finished before step A can begin.
Step A must be finished before step B can begin.
Step O must be finished before step L can begin.
Step T must be finished before step L can begin.
Step D must be finished before step R can begin.
Step G must be finished before step L can begin.
Step C must be finished before step R can begin.
Step R must be finished before step L can begin.
Step L must be finished before step B can begin.
Step O must be finished before step R can begin.
Step Q must be finished before step I can begin.
Step M must be finished before step L can begin.
Step R must be finished before step B can begin.
Step J must be finished before step O can begin.
Step O must be finished before step B can begin.
Step Y must be finished before step L can begin.
Step G must be finished before step R can begin.
Step P must be finished before step Z can begin.
Step K must be finished before step Y can begin.
Step X must be finished before step I can begin.
Step E must be finished before step H can begin.
Step I must be finished before step H can begin.
Step P must be finished before step K can begin.
Step G must be finished before step B can begin.
Step H must be finished before step L can begin.
Step X must be finished before step C can begin.
Step P must be finished before step X can begin.
Step X must be finished before step M can begin.
Step Q must be finished before step H can begin.
Step S must be finished before step Z can begin.
Step C must be finished before step B can begin.
Step N must be finished before step A can begin.
Step M must be finished before step R can begin.
Step X must be finished before step E can begin.
Step P must be finished before step L can begin.
Step H must be finished before step G can begin.
Step E must be finished before step D can begin.
Step D must be finished before step L can begin.
Step W must be finished before step A can begin.
Step S must be finished before step X can begin.
Step V must be finished before step O can begin.
Step H must be finished before step B can begin.
Step T must be finished before step B can begin.
Step Y must be finished before step C can begin.
Step A must be finished before step R can begin.
Step N must be finished before step L can begin.
Step V must be finished before step Z can begin.
Step W must be finished before step V can begin.
Step S must be finished before step M can begin.
Step Z must be finished before step A can begin.
Step W must be finished before step S can begin.
Step Q must be finished before step R can begin.
Step N must be finished before step G can begin.
Step Z must be finished before step L can begin.
Step K must be finished before step O can begin.
Step X must be finished before step R can begin.
Step V must be finished before step H can begin.
Step P must be finished before step R can begin.
Step M must be finished before step A can begin.
Step K must be finished before step L can begin.
Step P must be finished before step M can begin.
Step F must be finished before step N can begin.
Step W must be finished before step H can begin.
Step K must be finished before step B can begin.
Step H must be finished before step C can begin.
Step X must be finished before step H can begin.
Step V must be finished before step U can begin.
Step S must be finished before step H can begin.
Step J must be finished before step X can begin.
Step S must be finished before step N can begin.
Step V must be finished before step A can begin.
Step H must be finished before step O can begin.
Step Y must be finished before step O can begin.
Step H must be finished before step R can begin.
Step X must be finished before step T can begin.
Step J must be finished before step H can begin.
Step G must be finished before step C can begin.
Step E must be finished before step R can begin.
Step W must be finished before step J can begin.
Step F must be finished before step E can begin.
Step P must be finished before step I can begin.
Step F must be finished before step T can begin.
Step J must be finished before step L can begin.
Step U must be finished before step Z can begin.
Step Q must be finished before step D can begin.`,
	"google": `Step B must be finished before step X can begin.
Step H must be finished before step P can begin.
Step Y must be finished before step J can begin.
Step Z must be finished before step I can begin.
Step T must be finished before step U can begin.
Step R must be finished before step C can begin.
Step S must be finished before step J can begin.
Step W must be finished before step J can begin.
Step C must be finished before step L can begin.
Step L must be finished before step F can begin.
Step E must be finished before step G can begin.
Step A must be finished before step G can begin.
Step V must be finished before step X can begin.
Step U must be finished before step O can begin.
Step P must be finished before step F can begin.
Step O must be finished before step I can begin.
Step I must be finished before step F can begin.
Step K must be finished before step F can begin.
Step J must be finished before step F can begin.
Step G must be finished before step X can begin.
Step M must be finished before step X can begin.
Step F must be finished before step Q can begin.
Step Q must be finished before step N can begin.
Step D must be finished before step N can begin.
Step X must be finished before step N can begin.
Step I must be finished before step Q can begin.
Step U must be finished before step I can begin.
Step D must be finished before step X can begin.
Step B must be finished before step W can begin.
Step L must be finished before step N can begin.
Step U must be finished before step X can begin.
Step U must be finished before step J can begin.
Step C must be finished before step V can begin.
Step G must be finished before step N can begin.
Step S must be finished before step K can begin.
Step Q must be finished before step D can begin.
Step J must be finished before step X can begin.
Step V must be finished before step K can begin.
Step Z must be finished before step A can begin.
Step L must be finished before step M can begin.
Step H must be finished before step D can begin.
Step V must be finished before step Q can begin.
Step L must be finished before step V can begin.
Step S must be finished before step D can begin.
Step C must be finished before step Q can begin.
Step S must be finished before step L can begin.
Step E must be finished before step V can begin.
Step E must be finished before step P can begin.
Step C must be finished before step I can begin.
Step O must be finished before step K can begin.
Step H must be finished before step V can begin.
Step M must be finished before step F can begin.
Step K must be finished before step N can begin.
Step C must be finished before step X can begin.
Step G must be finished before step D can begin.
Step E must be finished before step U can begin.
Step R must be finished before step L can begin.
Step K must be finished before step G can begin.
Step W must be finished before step C can begin.
Step B must be finished before step L can begin.
Step L must be finished before step J can begin.
Step U must be finished before step D can begin.
Step I must be finished before step G can begin.
Step Q must be finished before step X can begin.
Step B must be finished before step M can begin.
Step T must be finished before step P can begin.
Step G must be finished before step Q can begin.
Step Y must be finished before step U can begin.
Step M must be finished before step D can begin.
Step P must be finished before step I can begin.
Step I must be finished before step K can begin.
Step O must be finished before step M can begin.
Step H must be finished before step Z can begin.
Step V must be finished before step M can begin.
Step P must be finished before step J can begin.
Step B must be finished before step U can begin.
Step E must be finished before step X can begin.
Step M must be finished before step Q can begin.
Step W must be finished before step L can begin.
Step O must be finished before step J can begin.
Step I must be finished before step X can begin.
Step J must be finished before step N can begin.
Step Y must be finished before step S can begin.
Step E must be finished before step D can begin.
Step M must be finished before step N can begin.
Step E must be finished before step O can begin.
Step I must be finished before step D can begin.
Step V must be finished before step N can begin.
Step R must be finished before step X can begin.
Step Z must be finished before step O can begin.
Step O must be finished before step X can begin.
Step I must be finished before step J can begin.
Step S must be finished before step E can begin.
Step E must be finished before step Q can begin.
Step J must be finished before step Q can begin.
Step H must be finished before step Y can begin.
Step T must be finished before step G can begin.
Step S must be finished before step A can begin.
Step P must be finished before step K can begin.
Step A must be finished before step D can begin.
Step B must be finished before step P can begin.`}

//*/
