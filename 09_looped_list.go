package main

import (
	"fmt"
	"log"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

// flat version, O(nm^2)
func part1_flat(ne, nm int) int {
	elves := make([]int, ne)
	elf := 0
	pos := 0
	marbs := make([]int, 1, nm)
	max := 0
	for i := 1; i <= nm; i++ {
		if i%100000 == 0 {
			_log(i, len(marbs)) // "I'm not dead!"
		}
		elf = (elf + 1) % ne
		if i%23 == 0 {
			elves[elf] += i
			pos = (pos + len(marbs) - 7) % len(marbs)
			elves[elf] += marbs[pos]
			if max < elves[elf] {
				max = elves[elf]
			}
			marbs = append(marbs[:pos], marbs[pos+1:]...)
		} else {
			pos = (pos+1)%len(marbs) + 1
			marbs = append(marbs, 0)
			copy(marbs[pos+1:], marbs[pos:])
			marbs[pos] = i
		}
	}
	return max
}

// dl list
type marb struct {
	val        int
	prev, next *marb
}

// looped double-linked list version, O(nm)
func part1(ne, nm int) int {
	max := 0
	elves := make([]int, ne)
	elf := 0
	m := &marb{val: 0}
	m.next = m
	m.prev = m
	for i := 1; i <= nm; i++ {
		elf = (elf + 1) % ne
		if i%23 == 0 {
			elves[elf] += i
			m = m.prev.prev.prev.prev.prev.prev.prev //-7
			elves[elf] += m.val
			if max < elves[elf] {
				max = elves[elf]
			}
			m.prev.next = m.next
			m.next.prev = m.prev
			m = m.next
		} else {
			m = m.next.next //+2
			m2 := &marb{val: i, prev: m.prev, next: m}
			m.prev.next = m2
			m.prev = m2
			m = m2
		}
	}
	return max
}

func part2(ne, nm int) int {
	return part1(ne, nm*100)
}

func verify(ne, nm, ex int) {
	v := part1(ne, nm)
	if v != ex {
		log.Fatalln(v, "!=", ex)
	}
}

func test() {
	verify(9, 25, 32)
	verify(10, 1618, 8317)
	fmt.Println("tests passed")
}

func main() {
	test()
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		fmt.Println("part 1:", part1(in[0], in[1]))
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		fmt.Println("part 2:", part2(in[0], in[1]))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string][2]int{
	"github": [2]int{411, 71058},
	"google": [2]int{424, 71482}}
