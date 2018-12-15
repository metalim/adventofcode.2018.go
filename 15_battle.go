package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	// fmt.Println(a...)
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
	s  string
	v  rune
	hp int
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
		_log(s)
		r := []rune(s)
		f = append(f, r)
		for x, v := range r {
			if v == 'E' || v == 'G' {
				us = append(us, unit{point{y, x}, string(v), v, 200})
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
	// _log("bfs", f, start, string(target))

	for step := 0; len(next) > 0 && len(out) == 0; step++ { //steps
		// time.Sleep(500 * time.Millisecond)
		// _log(step, len(next), next)
		cur, next = next, cur[:0]
		for _, p := range cur {
			for d := 0; d <= 3; d++ {
				x := p.x + (1-d)%2
				y := p.y + (2-d)%2
				t := point{y, x}
				if visited[t] {
					continue
				}
				// _log(t, string(f[y][x]))
				switch f[y][x] {
				case target:
					out[t] = true
				case '.':
					next = append(next, t)
					// _log("+", next)
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

func (f field) print() {
	// for _, r := range f {
	// 	_log(string(r))
	// }
}

var ds = [4][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func part1(p prep) int {
	f := p.f
	us := p.us
	for rounds := 0; ; rounds++ {
		// time.Sleep(time.Second)
		_log("\nafter", rounds, "rounds")
		f.print()
		// for _, u := range us {
		// 	if u.hp > 0 {
		// 		_log(u)
		// 	}
		// }
		_log("------")
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
			//   attack weakest, and by pos
			// else if has targets
			//   1. check reachable \
			//   2. find distances   } bfs
			//   3. find closest    /
			//   4. select by pos
			//   5. move by pos
			//   if can attack
			//     attack weakest, and by pos
			// else
			//   calc result

			if len(as) > 0 {
				for _, it := range as { // already sorted by pos
					a := us[it]
					if a.hp == hpMin {
						_log(".", u, "attacking", a)
						us[it].hp -= 3
						if us[it].hp <= 0 {
							f[a.y][a.x] = '.'
						}
						break
					}
				}
			} else if len(ts) > 0 {
				closest := f.bfs(u.point, us[ts[0]].v)
				for _, it := range ts {
					t := us[it]
					if closest[t.point] { // we have a target
						// find next step
						f.replace(u.point, '.', 'o')
						steps := f.bfs(t.point, 'o')
						_log(".", "steps", steps)
						for _, d := range ds { // select step by pos
							x := u.x + d[0]
							y := u.y + d[1]
							if steps[point{y, x}] {
								_log(" ", u, "moving to", y, x)
								//time.Sleep(500 * time.Millisecond)
								f.replace(u.point, 'o', '.')
								f[u.y][u.x] = '.'
								f[y][x] = u.v
								u.x = x
								u.y = y
								us[i] = u
								break //step sel
							}
						}
						break //target sel
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
					for _, it := range as { // already sorted by pos
						a := us[it]
						if a.hp == hpMin {
							_log(" and", u, "attacking", a)
							us[it].hp -= 3
							if us[it].hp <= 0 {
								f[a.y][a.x] = '.'
							}
							break
						}
					}
				}
			} else { //no more targets
				var sum int
				for _, u := range us {
					if u.hp > 0 {
						sum += u.hp
					}
				}
				_log("sum", sum, "rounds", rounds)
				return sum * rounds
			}
		}
	}
}

func part2(p prep) int {
	return 1
}

//
// tests
//

func verify(p prep, ex int) {
	v := part1(p)
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

var tests = map[string]int{
	`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`: 27730,
}

func test() {
	for t, v := range tests {
		p := prepare(t)
		verify(p, v)
	}
	fmt.Println("tests passed")
}

func main() {
	test()
	// delete(ins, "google")
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		p := prepare(in)
		fmt.Println("part 1:", part1(p))
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		// fmt.Println("part 2:", part2(p))
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
