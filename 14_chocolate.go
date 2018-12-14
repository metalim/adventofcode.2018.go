package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

func _log(a ...interface{}) {
	fmt.Println(a...)
}

//
// Solution
//

func part1(in int) string {
	rs := make([]int, 2, in+20)
	rs[0] = 3
	rs[1] = 7
	i, j := 0, 1
	for len(rs) < in+10 {
		sum := rs[i] + rs[j]
		a, b := sum/10, sum%10
		if a > 0 {
			rs = append(rs, a)
		}
		rs = append(rs, b)
		i = (i + 1 + rs[i]) % len(rs)
		j = (j + 1 + rs[j]) % len(rs)
	}
	return strings.Trim(strings.Replace(fmt.Sprint(rs[in:in+10]), " ", "", -1), "[]")
}

func part2(in int) int {
	rs := []int{3, 7}
	ln := len(fmt.Sprint(in))
	mn := int(math.Pow10(ln - 1))
	i, j := 0, 1
	t := 37
	for {
		sum := rs[i] + rs[j]
		a, b := sum/10, sum%10
		if a > 0 {
			rs = append(rs, a)
			t = t%mn*10 + a
			if t == in {
				return len(rs) - ln
			}
		}
		rs = append(rs, b)
		t = t%mn*10 + b
		if t == in {
			return len(rs) - ln
		}
		i = (i + 1 + rs[i]) % len(rs)
		j = (j + 1 + rs[j]) % len(rs)
	}
}

//
// tests
//

func verify1(in int, ex string) {
	v := part1(in)
	if v != ex {
		log.Fatal(v, " != ", ex)
	}
}

func verify2(in int, ex int) {
	v := part2(in)
	if v != ex {
		log.Fatal(v, " != ", ex)
	}
}

func test() {
	verify1(5, "0124515891")
	verify1(9, "5158916779")
	verify1(18, "9251071085")
	verify1(2018, "5941429882")

	verify2(1245, 5+1) // skip leading 0, and don't care about string representation, as our inputs don't start with '0' anyway
	verify2(51589, 9)
	verify2(92510, 18)
	verify2(59414, 2018)

	fmt.Println("\ntests passed\n")
}

func main() {
	test()
	// delete(ins, "github")
	// delete(ins, "google")
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		fmt.Println("part 1:", part1(in))
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		fmt.Println("part 2:", part2(in))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		fmt.Println()
	}
}

var ins = map[string]int{
	"github": 637061,
	"google": 306281,
}
