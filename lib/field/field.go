package field

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

// Cell z
type Cell = int

// CellDefault z
const CellDefault = 0

// Pos z
type Pos = image.Point

// Rect z
type Rect = image.Rectangle

// Field two-dimensional.
type Field interface {
	Get(Pos) Cell
	Set(Pos, Cell)
	Bounds() Rect
}

// FillBFS z
func FillBFS(f Field, start Pos, fn func(Pos, Cell) Cell) {

}

// Print field
func Print(f Field) {
	bs := f.Bounds()
	r := make([]string, 0, bs.Dx())
	for y := bs.Min.Y; y < bs.Max.Y; y++ {
		r = r[:0]
		for x := bs.Min.X; x < bs.Max.X; x++ {
			c := f.Get(Pos{x, y})
			if c == 0 {
				r = append(r, ".")
				continue
			}
			r = append(r, strconv.Itoa(int(c)))
		}
		fmt.Println(strings.Join(r, " "))
	}
}

// Step in specified direction: 0123 -> ESWN
func Step(p Pos, d int) Pos {
	return Pos{p.X + (1-d)%2, p.Y + (2-d)%2}
}

// DStep in specified direction, including diagonals:
// ? var1: 01234567 -> E SE S SW W NW N NE
// ! var2: 0123:4567 -> E S W N : SE SW NW NE
func DStep(p Pos, d int) Pos {
	out := Step(p, d&3)
	if d > 3 { // diagonal? -> additional step in +1 direction.
		out = Step(out, (d+1)&3)
	}
	return out
}

func abs(n int) int {
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
}

// Manh distance.
func Manh(p1, p2 Pos) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}
