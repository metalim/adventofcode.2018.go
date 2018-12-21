package main

import (
	"fmt"
	"log"
	"math"
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
	if n < 0 {
		return -n
	}
	return n
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

func parse(in string) (task task) {
	ss := strings.Split(in, "\n")
	task.ip, _ = strconv.Atoi(strings.Split(ss[0], " ")[1])
	for _, s := range ss[1:] {
		c := strings.Split(s, " ")
		ns := sliceAtoi(c)
		var ok bool
		ns[0], ok = imap[c[0]]
		if !ok {
			panic("invalid instruction: " + s)
		}
		i := inst{}
		copy(i[:], ns)
		task.c = append(task.c, i)
	}
	return
}

func (reg *regs) exec(i inst) {
	code, a, b, c := i[0], i[1], i[2], i[3]
	switch code {
	case 0: // addr
		reg[c] = reg[a] + reg[b]
	case 1: // addi
		reg[c] = reg[a] + b
	case 2: // mulr
		reg[c] = reg[a] * reg[b]
	case 3: // muli
		reg[c] = reg[a] * b
	case 4: // banr
		reg[c] = reg[a] & reg[b]
	case 5: // bani
		reg[c] = reg[a] & b
	case 6: // borr
		reg[c] = reg[a] | reg[b]
	case 7: // bori
		reg[c] = reg[a] | b
	case 8: // setr
		reg[c] = reg[a]
	case 9: // seti
		reg[c] = a
	case 10: // gtir
		reg[c] = b2i(a > reg[b])
	case 11: // gtri
		reg[c] = b2i(reg[a] > b)
	case 12: // gtrr
		reg[c] = b2i(reg[a] > reg[b])
	case 13: // eqir
		reg[c] = b2i(a == reg[b])
	case 14: // eqri
		reg[c] = b2i(reg[a] == b)
	case 15: // eqrr
		reg[c] = b2i(reg[a] == reg[b])
	}
}

func (task *task) run(reg regs) regs {
	hot := make([]int, len(task.c))
	for 0 <= reg[task.ip] && reg[task.ip] < len(task.c) {
		ip := reg[task.ip]
		reg.exec(task.c[ip])
		if 0 <= ip && ip < len(hot) {
			hot[ip]++
		}
		reg[task.ip]++
	}
	// _log(hot) // show hot code
	return reg
}

func (task *task) part1() int {
	reg := task.run(regs{})
	return reg[0]
}

func (task *task) part2() int {
	task.c[26][1] = 999 // return part1 entry data
	task.c[35][1] = 999 // return part2 entry data
	i := task.c[24][3]  // index of entry register
	reg1 := task.run(regs{})
	_log("confirm part1:", Cyan(loop(reg1[i])))

	reg2 := task.run(regs{0: 1})
	return loop(reg2[i])
}

// corresponging to Assembly Go code. O(N**2), takes around 10 years for N~1e10
func loop0(n int) int {
	var sum, i, j int
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if j*i == n {
				sum += i
			}
		}
	}
	return sum
}

// refactor inside loop. O(N), <100ms for N~1e10
func loop1(n int) int {
	var sum, i int
	for i = 1; i <= n; i++ {
		// mod instead of iterating
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// refactor outside loop. Totally optional, but O(sqrt(N)), <100µs! for N~1e10
func loop(n int) int {
	var sum, i int
	sq := int(math.Sqrt(float64(n)))
	// iterate to sqrt(n) only and add both divisors
	for i = 1; i <= sq; i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}
	return sum
}

//
// tests
//

func verify(v, ex int) {
	if v != ex {
		log.Fatal(v, "!=", ex)
	}
}

func test() {
	task := parse(`#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`)
	verify(task.part1(), 7)
	fmt.Println("tests passed")
}

func main() {
	test()
	for i, in := range ins {
		fmt.Println(Brown(fmt.Sprint("=== for ", i, " ===")))

		t0 := time.Now()
		task := parse(in)
		t := time.Since(t0)
		fmt.Println(Gray("parse:"), Black(t).Bold())

		t0 = time.Now()
		v1 := task.part1()
		t = time.Since(t0)
		fmt.Println(Gray("part 1:"), Black(t).Bold(), Green(v1).Bold())

		t0 = time.Now()
		v2 := task.part2()
		t = time.Since(t0)
		fmt.Println(Gray("part 2:"), Black(t).Bold(), Green(v2).Bold())

		fmt.Println()
	}
}

// code analysis
/*
#ip 5
00 addi 5 16 5 ->jmp 17 [1 0 0 0 0 17]

L01:
01 seti 1 1 4 <- [0 10551398 0 10550400 0 1] -> [x x x x 1 2] [4]=1

L02:
02 seti 1 8 2 -> [x x 1 x 1 3] [2]=1

L03:
03 mulr 4 2 3 -> [x x 1 1 1 4] [3]=[4]*[2]=[4]
04 eqrr 3 1 3 -> [x x 1 ? 1 5] [3]=([3]==[1])=([4]==[1])
05 addr 3 5 5 -> skip [3]
06 addi 5 1 5 -> skip 1
07 addr 4 0 0 -> [0]+=[4]
08 addi 2 1 2 -> [2]++
09 gtrr 2 1 3 -> [3]=([2]>[1])
10 addr 5 3 5 -> skip [3]
11 seti 2 6 5 -> jmp L3

12 addi 4 1 4 -> [4]++
13 gtrr 4 1 3 -> [3]=([4]>[1])
14 addr 3 5 5 -> skip [3] [4]>[1]
15 seti 1 4 5 -> jmp L2

16 mulr 5 5 5 -> exit

17 addi 1 2 1 -> [1 2 0 0 0 18] [1]+=2
18 mulr 1 1 1 -> [1 4 0 0 0 19] [1]*=[1]
19 mulr 5 1 1 -> [1 76 0 0 0 20] [1]*=19
20 muli 1 11 1 ->[1 836 0 0 0 21] [1]*=11
21 addi 3 7 3 -> [1 836 0 7 0 22] [3]+=7
22 mulr 3 5 3 -> [1 836 0 154 0 23] [3]*=22
23 addi 3 8 3 -> [1 836 0 162 0 24] [3]+=8
24 addr 1 3 1 -> [1 998 0 162 0 25] [1]+=[3]
25 addr 5 0 5 -> [1 998 0 162 0 27]  <-- skip [part2]
26 seti 0 9 5 -> jmp L1

27 setr 5 8 3 -> [1 998 0 27 0 28] [3]=27
28 mulr 3 5 3 -> [1 998 0 756 0 29] [3]*=28
29 addr 5 3 3 -> [1 998 0 785 0 30] [3]+=29
30 mulr 5 3 3 -> [1 998 0 23550 0 31] [3]*=30
31 muli 3 14 3 ->[1 998 0 329700 0 32] [3]*=14
32 mulr 3 5 3 -> [1 998 0 10550400 0 33] [3]*=32
33 addr 1 3 1 -> [1 10551398 0 10550400 0 34] [1]+=[3] +=10550400
34 seti 0 4 0 -> [0 10551398 0 10550400 0 35] [0]=0
35 seti 0 3 5 -> jmp L1
*/
var ins = map[string]string{
	"github": `#ip 5
addi 5 16 5
seti 1 1 4
seti 1 8 2
mulr 4 2 3
eqrr 3 1 3
addr 3 5 5
addi 5 1 5
addr 4 0 0
addi 2 1 2
gtrr 2 1 3
addr 5 3 5
seti 2 6 5
addi 4 1 4
gtrr 4 1 3
addr 3 5 5
seti 1 4 5
mulr 5 5 5
addi 1 2 1
mulr 1 1 1
mulr 5 1 1
muli 1 11 1
addi 3 7 3
mulr 3 5 3
addi 3 8 3
addr 1 3 1
addr 5 0 5
seti 0 9 5
setr 5 8 3
mulr 3 5 3
addr 5 3 3
mulr 5 3 3
muli 3 14 3
mulr 3 5 3
addr 1 3 1
seti 0 4 0
seti 0 3 5`,
	"google": `#ip 4
addi 4 16 4
seti 1 9 5
seti 1 5 2
mulr 5 2 1
eqrr 1 3 1
addr 1 4 4
addi 4 1 4
addr 5 0 0
addi 2 1 2
gtrr 2 3 1
addr 4 1 4
seti 2 6 4
addi 5 1 5
gtrr 5 3 1
addr 1 4 4
seti 1 2 4
mulr 4 4 4
addi 3 2 3
mulr 3 3 3
mulr 4 3 3
muli 3 11 3
addi 1 5 1
mulr 1 4 1
addi 1 2 1
addr 3 1 3
addr 4 0 4
seti 0 2 4
setr 4 8 1
mulr 1 4 1
addr 4 1 1
mulr 4 1 1
muli 1 14 1
mulr 1 4 1
addr 3 1 3
seti 0 0 0
seti 0 2 4`,
}
