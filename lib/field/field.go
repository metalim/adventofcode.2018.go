package field

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

// Cell z
type Cell int

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
