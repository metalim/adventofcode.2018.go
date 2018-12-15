package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
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

//
// Solution
//

type point struct {
	y, x int
}

func (a point) dist(b point) int { return abs(a.x-b.x) + abs(a.y-b.y) }

type unit struct {
	point
	s   string
	v   rune
	hp  int
	pow int
}

type field [][]rune
type prep struct {
	f  field
	us []unit
}

type byPos []unit

func (a byPos) Len() int           { return len(a) }
func (a byPos) Less(i, j int) bool { return a[i].y < a[j].y || a[i].y == a[j].y && a[i].x < a[j].x }
func (a byPos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func prepare(in string) prep {
	ss := strings.Split(in, "\n")
	var f field
	var us []unit
	for y, s := range ss {
		r := []rune(s)
		f = append(f, r)
		for x, v := range r {
			if v == 'E' || v == 'G' {
				us = append(us, unit{point{y, x}, string(v), v, 200, 3})
			}
		}
	}
	return prep{f, us}
}

func (f field) bfs(start point, target rune) map[point]bool {
	out := map[point]bool{}
	visited := map[point]bool{start: true}
	next := []point{start}
	cur := []point{} // to avoid new slices - just swap them

	for step := 0; len(next) > 0 && len(out) == 0; step++ { //steps
		cur, next = next, cur[:0]
		for _, p := range cur {
			for d := 0; d <= 3; d++ {
				x := p.x + (1-d)%2
				y := p.y + (2-d)%2
				t := point{y, x}
				if visited[t] {
					continue
				}
				switch f[y][x] {
				case target:
					out[t] = true
				case '.':
					next = append(next, t)
				}
				visited[t] = true
			}
		}
	}
	return out
}

func (f field) replace(p point, from, to rune) {
	for d := 0; d <= 3; d++ {
		x := p.x + (1-d)%2
		y := p.y + (2-d)%2
		if f[y][x] == from {
			f[y][x] = to
		}
	}
}

var ds = [4][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func part1(p prep) int {
	f := p.f
	us := p.us
	for rounds := 0; ; rounds++ {
		sort.Sort(byPos(us))
		for i, u := range us {
			if u.hp <= 0 {
				continue
			}

			var ts []int
			var as []int
			hpMin := 999
			for it, t := range us {
				if u.v != t.v && t.hp > 0 {
					ts = append(ts, it)
					if u.point.dist(t.point) == 1 {
						as = append(as, it)
						if hpMin > t.hp {
							hpMin = t.hp
						}
					}
				}
			}
			// if can attack
			//   attack weakest, sorted by pos
			// else if has targets
			//   1. check reachable \
			//   2. find distances   } bfs
			//   3. find closest    /
			//   4. sort by pos
			//   5. move 1 step, sorted by pos
			//   if can attack
			//     attack weakest, sorted by pos
			// else
			//   calc result

			attack := func() {
				for _, it := range as { // already sorted by pos
					a := us[it]
					if a.hp == hpMin {
						us[it].hp -= u.pow
						if us[it].hp <= 0 {
							f[a.y][a.x] = '.'
						}
						break
					}
				}
			}

			if len(as) > 0 {
				attack()
			} else if len(ts) > 0 {
				closest := f.bfs(u.point, us[ts[0]].v)
				for _, it := range ts {
					t := us[it]
					if closest[t.point] { // we have a target
						// find next step
						f.replace(u.point, '.', 'o')
						steps := f.bfs(t.point, 'o')
						for _, d := range ds { // select step by pos
							x := u.x + d[0]
							y := u.y + d[1]
							if steps[point{y, x}] {
								f.replace(u.point, 'o', '.')
								f[u.y][u.x] = '.'
								f[y][x] = u.v
								u.x = x
								u.y = y
								us[i] = u
								break //sel step
							}
						}
						break //sel target
					}
				} // move completed
				for _, it := range ts {
					t := us[it]
					if u.point.dist(t.point) == 1 {
						as = append(as, it)
						if hpMin > t.hp {
							hpMin = t.hp
						}
					}
				}
				if len(as) > 0 { // can attack
					attack()
				}
			} else { //no more targets
				var sum int
				for _, u := range us {
					if u.hp > 0 {
						sum += u.hp
					}
				}
				return sum * rounds
			}
		}
	}
}

func part2(in string) [2]int {
	for pow := 4; ; pow++ {
		st := prepare(in)
		var elves, alive int
		for i := range st.us {
			if st.us[i].v == 'E' {
				elves++
				st.us[i].pow = pow
			}
		}
		score := part1(st)
		for i := range st.us {
			if st.us[i].v == 'E' && st.us[i].hp > 0 {
				alive++
			}
		}
		if elves == alive {
			return [2]int{score, pow}
		}
	}
}

//
// tests
//

func verify(v, ex int) {
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

var tests = map[string][2]int{
	`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`: {27730, 4988},
}

func test() {
	for t, v := range tests {
		p := prepare(t)
		verify(part1(p), v[0])
		verify(part2(t)[0], v[1])
	}
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
		fmt.Println("part 2:", part2(in))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `################################
###############G.G.#############
###############...G#############
##############.....#############
#############.G....#############
############.......#############
###########...##################
###########G.###################
#########....###################
##########..####################
##########G.###########...######
###########.G.G.......#....#####
###########...#####...#....#####
###########..#######..G.....##.#
###########.#########..........#
######..###.#########.G.#.....##
#####....#..#########....#######
#####.......#########..#.##..###
###..##.....#########..#.......#
###..........#######...........#
##..GG........#####G.......E...#
#...#....G...........G.......E.#
###.#...............E.EE.......#
###..............G#E......E...##
##.....G............#.....E..###
#.G........G..............E..###
#.#.....######.......E.......###
##...G...#####....#..#.#..######
#####..#...###....######..######
#####.......##..########..######
######....#.###########...######
################################`,
	"google": `################################
#######################.########
######################....######
#######################.....####
##################..##......####
###################.##.....#####
###################.....G..#####
##################.....G...#####
############.....GG.G...#..#####
##############...##....##.######
############...#..G............#
###########......E.............#
###########...#####..E........##
#...#######..#######.......#####
#..#..G....G#########.........##
#..#....G...#########..#....####
##.....G....#########.E......###
#####G.....G#########..E.....###
#####.......#########....#.....#
#####G#G....G#######.......#..E#
###.....G.....#####....#.#######
###......G.....G.G.......#######
###..................#..########
#####...................########
#####..............#...#########
####......G........#.E.#E..#####
####.###.........E...#E...######
####..##........#...##.....#####
########.#......######.....#####
########...E....#######....#####
#########...##..########...#####
################################`,
}
