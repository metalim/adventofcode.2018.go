package field

import (
	"math/rand"
	"testing"
)

func runTestsOn(t *testing.T, f Field) {

	verify := func(v, ex interface{}) {
		if v != ex {
			t.Fatalf("got:%v expected:%v", v, ex)
		}
	}
	verifyc := func(v, ex Cell) {
		verify(v, ex)
	}
	p0 := Pos{0, 0}
	p1 := Pos{10, 10}
	p2 := Pos{-10, 3}
	p999 := Pos{-100, 99}

	// empty field returns default cells
	verifyc(f.Get(p0), 0)

	// can set cell in empty field, bounds are updated
	t.Log("set p1")
	f.Set(p1, 1)
	verify(f.Bounds().Min, p1)
	verifyc(f.Get(p0), 0)
	verifyc(f.Get(p1), 1)

	// getting point outside of bounds, does not change bounds
	t.Log("get p999")
	verifyc(f.Get(p999), 0)
	verify(f.Bounds().Min, p1)

	// extending works for negative coords
	t.Log("set p2")
	f.Set(p2, 2)
	verify(f.Bounds().Min, p2)
	verifyc(f.Get(p0), 0)
	verifyc(f.Get(p1), 1)
	verifyc(f.Get(p2), 2)

	// can change value of cell
	t.Log("set p1=3")
	f.Set(p1, 3)
	verifyc(f.Get(p0), 0)
	verifyc(f.Get(p1), 3)
	verifyc(f.Get(p2), 2)

	// can set value of distant cell
	t.Log("set p999=4")
	f.Set(p999, 4)
	verifyc(f.Get(p0), 0)
	verifyc(f.Get(p1), 3)
	verifyc(f.Get(p2), 2)
	verifyc(f.Get(p999), 4)
	//Print(f)
	t.Log("ok")
}

func TestSlice(t *testing.T) {
	var f Slice
	runTestsOn(t, &f)
}

func TestMap(t *testing.T) {
	var f Map
	runTestsOn(t, &f)
}

const (
	dim = 100
	d0  = -50
)

func runBenchOn(b *testing.B, f Field) {
	for i := 0; i < b.N; i++ {
		p := Pos{rand.Intn(dim) + d0, rand.Intn(dim) + d0}
		c := Cell(rand.Intn(100))
		f.Set(p, c)
		v := f.Get(p)
		if v != c {
			b.Fatalf("got:%v expected:%v", v, c)
		}
		g := Pos{rand.Intn(dim) + d0, rand.Intn(dim) + d0}
		f.Get(g)
	}
}

func BenchmarkSlice(b *testing.B) {
	var f Slice
	runBenchOn(b, &f)
}

func BenchmarkMap(b *testing.B) {
	var f Map
	runBenchOn(b, &f)
}
