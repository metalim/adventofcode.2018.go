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

type cart struct {
	y, x int
	d    int
	t    int
}
type prep struct {
	f  [][]rune
	cs []cart
}
type byPos []cart

func (a byPos) Len() int           { return len(a) }
func (a byPos) Less(i, j int) bool { return a[i].y*1000+a[i].x < a[j].y*1000+a[j].x }
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
				carts = append(carts, cart{y, x, d, 0})
				f[y][x] = rune("-|-|"[d])
			}
		}
	}
	return prep{f, carts}
}

func part1(st prep) string {
	cs := make([]cart, len(st.cs))
	copy(cs, st.cs)
	for {
		sort.Sort(byPos(cs))
		for i, c := range cs {
			c.x += (1 - c.d) % 2
			c.y += (2 - c.d) % 2
			r := st.f[c.y][c.x]
			switch r {
			case '+':
				c.d = (c.d + c.t + 3) % 4
				c.t = (c.t + 1) % 3
			case '\\':
				c.d ^= 1
			case '/':
				c.d ^= 3
			}
			for _, c2 := range cs {
				if c.x == c2.x && c.y == c2.y {
					return fmt.Sprintf("%d,%d", c.x, c.y)
				}
			}
			cs[i] = c
		}
	}
}

func part2(st prep) string {
	for {
		sort.Sort(byPos(st.cs))
		for i, c := range st.cs {
			if c.x < 0 {
				continue
			}
			c.x += (1 - c.d) % 2
			c.y += (2 - c.d) % 2
			r := st.f[c.y][c.x]
			switch r {
			case '+':
				c.d = (c.d + c.t + 3) % 4
				c.t = (c.t + 1) % 3
			case '\\':
				c.d ^= 1
			case '/':
				c.d ^= 3
			}
			for j, c2 := range st.cs {
				if c.x == c2.x && c.y == c2.y {
					c.x = -1
					st.cs[j].x = -1
					break
				}
			}
			st.cs[i] = c
		}
		cs := make([]cart, 0, len(st.cs))
		for _, c := range st.cs {
			if c.x >= 0 {
				cs = append(cs, c)
			}
		}
		if len(cs) < 2 {
			return fmt.Sprintf("%d,%d", cs[0].x, cs[0].y)
		}
		st.cs = cs
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
		p = prepare(f)
		fmt.Println("part 2:", part2(p))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `13_github.txt`,
	"google": `13_google.txt`}
