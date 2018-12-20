package field

type map1d map[int]Cell
type map2d map[int]map1d

// Map z
type Map struct {
	m map2d
	b Rect
}

// Bounds returns AABB bounds.
func (f *Map) Bounds() Rect {
	return f.b
}

// Get cell.
func (f *Map) Get(p Pos) Cell {
	if r, ok := f.m[p.Y]; ok {
		return r[p.X]
	}
	return 0
}

// Set cell.
func (f *Map) Set(p Pos, c Cell) {
	if f.m == nil {
		f.m = map2d{}
	}
	if _, ok := f.m[p.Y]; !ok {
		f.m[p.Y] = map1d{}
	}
	if !p.In(f.b) {
		f.b = f.b.Union(Rect{p, p.Add(Pos{1, 1})})
	}
	f.m[p.Y][p.X] = c
}
