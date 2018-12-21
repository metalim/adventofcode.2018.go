package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

func sliceAtoi(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i], _ = strconv.Atoi(v)
	}
	return out
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

//
// Solution
//

type regs [6]int
type inst [4]int
type task struct {
	ip int
	c  []inst
	sc []string
}

var imap = map[string]int{
	"addr": 0,
	"addi": 1,
	"mulr": 2,
	"muli": 3,
	"banr": 4,
	"bani": 5,
	"borr": 6,
	"bori": 7,
	"setr": 8,
	"seti": 9,
	"gtir": 10,
	"gtri": 11,
	"gtrr": 12,
	"eqir": 13,
	"eqri": 14,
	"eqrr": 15,
}

func parse(in string) (t task) {
	ss := strings.Split(in, "\n")
	t.ip, _ = strconv.Atoi(strings.Split(ss[0], " ")[1])
	for j, s := range ss[1:] {
		c := strings.Split(s, " ")
		ns := sliceAtoi(c)
		var ok bool
		ns[0], ok = imap[c[0]]
		if !ok {
			panic("invalid instruction: " + s)
		}
		i := inst{}
		copy(i[:], ns)
		t.c = append(t.c, i)
		t.sc = append(t.sc, "L"+strconv.Itoa(j)+": "+i.String())
	}
	return
}

var cmap = []string{
	"r%[3]d = r%[1]d + r%[2]d", // addr
	"r%[3]d = r%[1]d + %[2]d",  // addi
	"r%[3]d = r%[1]d * r%[2]d", // mulr
	"r%[3]d = r%[1]d * %[2]d",  // muli

	"r%[3]d = r%[1]d & r%[2]d", // banr
	"r%[3]d = r%[1]d & %[2]d",  // bani
	"r%[3]d = r%[1]d | r%[2]d", // borr
	"r%[3]d = r%[1]d | %[2]d",  // bori

	"r%[3]d = r%[1]d", // setr
	"r%[3]d = %[1]d",  // seti

	"r%[3]d = b2i(%[1]d > r%[2]d)",  // gtir
	"r%[3]d = b2i(r%[1]d > %[2]d)",  // gtri
	"r%[3]d = b2i(r%[1]d > r%[2]d)", // gtrr

	"r%[3]d = b2i(%[1]d == r%[2]d)",  // eqir
	"r%[3]d = b2i(r%[1]d == %[2]d)",  // eqri
	"r%[3]d = b2i(r%[1]d == r%[2]d)", // eqrr
}

func (i inst) String() string {
	code, a, b, c := i[0], i[1], i[2], i[3]
	return fmt.Sprintf(cmap[code], a, b, c)
}

func step8(init int) (r0 int) {
	var r1, r2, r4, r5 int

	r4 = 123
	for {
		r4 = r4 & 456
		r4 = b2i(r4 == 72)
		if r4 != 0 {
			break
		}
	}

	r4 = 0

	for {
		r1 = r4 | 65536
		r4 = init
		for {
			r5 = r1 & 255
			r4 = r4 + r5
			r4 = r4 & 16777215
			r4 = r4 * 65899
			r4 = r4 & 16777215
			r5 = b2i(256 > r1)
			if r5 != 0 {
				break
			}
			r5 = 0
			for {
				r2 = r5 + 1
				r2 = r2 * 256
				r2 = b2i(r2 > r1)
				if r2 != 0 {
					break
				}
				r5 = r5 + 1
			}
			r1 = r5
		}
		return r4
		// r5 = b2i(r4 == r0)
		// if r5 != 0 {
		// 	return
		// }
	}
}

func step13(init int) (r0 int) {
	var r1, r4, last int
	can := map[int]bool{}
	for {
		r1 = r4 | 65536
		for r4 = init; ; r1 /= 256 {
			r4 = ((r4 + r1&255) & 16777215 * 65899) & 16777215
			if r1 < 256 {
				break
			}
		}
		if can[r4] {
			_log(len(can))
			return last
		}
		can[r4] = true
		last = r4
	}
}

func (t *task) part1() int {
	return step8(t.c[7][1])
}

func (t *task) part2() int {
	return step13(t.c[7][1])
}

func main() {
	for i, in := range ins {
		fmt.Println(Brown(fmt.Sprint("=== for ", i, " ===")))
		var t0 time.Time
		var d time.Duration

		t0 = time.Now()
		t := parse(in)
		d = time.Since(t0)
		fmt.Println(Gray("parse:"), Black(d).Bold())

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

/*
// input comparison

L0: r4 = 123               <->  L0: r5 = 123
L1: r4 = r4 & 456          <->  L1: r5 = r5 & 456
L2: r4 = b2i(r4 == 72)     <->  L2: r5 = b2i(r5 == 72)
L3: r3 = r4 + r3           <->  L3: r1 = r5 + r1
L4: r3 = 0                 <->  L4: r1 = 0
L5: r4 = 0                 <->  L5: r5 = 0
L6: r1 = r4 | 65536        <->  L6: r4 = r5 | 65536
L7: r4 = 678134           <!!!> L7: r5 = 13431073
L8: r5 = r1 & 255          <->  L8: r3 = r4 & 255
L9: r4 = r4 + r5           <->  L9: r5 = r5 + r3
L10: r4 = r4 & 16777215    <->  L10: r5 = r5 & 16777215
L11: r4 = r4 * 65899       <->  L11: r5 = r5 * 65899
L12: r4 = r4 & 16777215    <->  L12: r5 = r5 & 16777215
L13: r5 = b2i(256 > r1)    <->  L13: r3 = b2i(256 > r4)
L14: r3 = r5 + r3          <->  L14: r1 = r3 + r1
L15: r3 = r3 + 1           <->  L15: r1 = r1 + 1
L16: r3 = 27               <->  L16: r1 = 27
L17: r5 = 0                <->  L17: r3 = 0
L18: r2 = r5 + 1           <->  L18: r2 = r3 + 1
L19: r2 = r2 * 256         <->  L19: r2 = r2 * 256
L20: r2 = b2i(r2 > r1)     <->  L20: r2 = b2i(r2 > r4)
L21: r3 = r2 + r3          <->  L21: r1 = r2 + r1
L22: r3 = r3 + 1           <->  L22: r1 = r1 + 1
L23: r3 = 25               <->  L23: r1 = 25
L24: r5 = r5 + 1           <->  L24: r3 = r3 + 1
L25: r3 = 17               <->  L25: r1 = 17
L26: r1 = r5               <->  L26: r4 = r3
L27: r3 = 7                <->  L27: r1 = 7
L28: r5 = b2i(r4 == r0)    <->  L28: r3 = b2i(r5 == r0)
L29: r3 = r5 + r3          <->  L29: r1 = r3 + r1
L30: r3 = 5                <->  L30: r1 = 5

*/
var ins = map[string]string{
	"github": `#ip 3
seti 123 0 4
bani 4 456 4
eqri 4 72 4
addr 4 3 3
seti 0 0 3
seti 0 6 4
bori 4 65536 1
seti 678134 1 4
bani 1 255 5
addr 4 5 4
bani 4 16777215 4
muli 4 65899 4
bani 4 16777215 4
gtir 256 1 5
addr 5 3 3
addi 3 1 3
seti 27 8 3
seti 0 1 5
addi 5 1 2
muli 2 256 2
gtrr 2 1 2
addr 2 3 3
addi 3 1 3
seti 25 7 3
addi 5 1 5
seti 17 1 3
setr 5 3 1
seti 7 8 3
eqrr 4 0 5
addr 5 3 3
seti 5 4 3`,
	"google": `#ip 1
seti 123 0 5
bani 5 456 5
eqri 5 72 5
addr 5 1 1
seti 0 0 1
seti 0 6 5
bori 5 65536 4
seti 13431073 4 5
bani 4 255 3
addr 5 3 5
bani 5 16777215 5
muli 5 65899 5
bani 5 16777215 5
gtir 256 4 3
addr 3 1 1
addi 1 1 1
seti 27 9 1
seti 0 1 3
addi 3 1 2
muli 2 256 2
gtrr 2 4 2
addr 2 1 1
addi 1 1 1
seti 25 4 1
addi 3 1 3
seti 17 8 1
setr 3 4 4
seti 7 7 1
eqrr 5 0 3
addr 3 1 1
seti 5 9 1`,
}
