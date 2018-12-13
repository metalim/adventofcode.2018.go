package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

func _file(name string) string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

//
// Solution
//

type point struct {
	y, x int
}
type cart struct {
	point
	d int
	t int
}
type prep struct {
	f  [][]rune
	cs []cart
}
type byPos []cart

func (a byPos) Len() int           { return len(a) }
func (a byPos) Less(i, j int) bool { return a[i].y < a[j].y || a[i].y == a[j].y && a[i].x < a[j].x }
func (a byPos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func prepare(in string) prep {
	ss := strings.Split(in, "\n")
	f := make([][]rune, len(ss))
	carts := []cart{}
	for y, r := range ss {
		f[y] = make([]rune, len(r))
		for x, v := range r {
			d := strings.IndexRune(">v<^", v)
			if d < 0 {
				f[y][x] = v
			} else {
				carts = append(carts, cart{point{y, x}, d, 0})
				f[y][x] = rune("-|-|"[d])
			}
		}
	}
	return prep{f, carts}
}

func step(c cart, f [][]rune) cart {
	c.x += (1 - c.d) % 2
	c.y += (2 - c.d) % 2
	r := f[c.y][c.x]
	switch r {
	case '+':
		c.d = (c.d + c.t + 3) % 4
		c.t = (c.t + 1) % 3
	case '\\':
		c.d ^= 1
	case '/':
		c.d ^= 3
	}
	return c
}

func part1(st prep) string {
	cs := make([]cart, len(st.cs))
	copy(cs, st.cs)
	for {
		sort.Sort(byPos(cs))
		for i, c := range cs {
			c = step(c, st.f)
			for j, c2 := range cs {
				if i != j && c.point == c2.point {
					return fmt.Sprintf("%d,%d", c.x, c.y)
				}
			}
			cs[i] = c
		}
	}
}

func part2(st prep) string {
	cs := make([]cart, len(st.cs))
	copy(cs, st.cs)
	for {
		sort.Sort(byPos(cs))
		for i, c := range cs {
			if c.x < 0 {
				continue
			}
			c = step(c, st.f)
			for j, c2 := range cs {
				if i != j && c.point == c2.point {
					c.x = -1
					cs[j].x = -1
					break
				}
			}
			cs[i] = c
		}
		cs2 := make([]cart, 0, len(cs))
		for _, c := range cs {
			if c.x >= 0 {
				cs2 = append(cs2, c)
			}
		}
		cs = cs2
		if len(cs) < 2 {
			return fmt.Sprintf("%d,%d", cs[0].x, cs[0].y)
		}
	}
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		f := _file(in)
		p := prepare(f)
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
	"github": `13_github.txt`,
	"google": `13_google.txt`}
