package field

type row []Cell
type field []row

// Slice arbitrarily growing in any dimension, including negative.
type Slice struct {
	d, b Rect
	m    field
}

// Bounds returns AABB rect covering set cells.
func (f *Slice) Bounds() Rect {
	return f.b
}

// Get cell
func (f *Slice) Get(p Pos) Cell {
	if !p.In(f.d) {
		return 0
	}
	p = p.Sub(f.d.Min)
	return f.m[p.Y][p.X]
}

// Set cell
func (f *Slice) Set(p Pos, c Cell) {
	if !p.In(f.d) {
		d := Rect{p, p.Add(Pos{1, 1})}
		d = d.Inset(-(f.d.Dx()+f.d.Dy())/4 - 10).Union(f.d)
		f.growTo(d)
	}
	if !p.In(f.b) {
		f.b = f.b.Union(Rect{p, p.Add(Pos{1, 1})})
	}
	p = p.Sub(f.d.Min)
	f.m[p.Y][p.X] = c
}

func (f *Slice) growTo(d Rect) {
	w := d.Dx()
	h := d.Dy()
	dist := d.Min.Sub(f.d.Min)
	m := make(field, 0, h)
	for y := 0; y < h; y++ {
		r := make(row, w)
		fY := y + dist.Y
		if fY > 0 && fY < len(f.m) {
			copy(r[-dist.X:], f.m[fY])
		}
		m = append(m, r)
	}
	f.m = m
	f.d = d
}
