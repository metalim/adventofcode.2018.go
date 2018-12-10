package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"
	"time"
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

type rect = image.Rectangle

func bounds(ps [][]int) rect {
	r := image.Rect(ps[0][0], ps[0][1], ps[0][0]+1, ps[0][1]+1)
	for _, p := range ps[1:] {
		if r.Min.X > p[0] {
			r.Min.X = p[0]
		}
		if r.Max.X < p[0]+1 {
			r.Max.X = p[0] + 1
		}
		if r.Min.Y > p[1] {
			r.Min.Y = p[1]
		}
		if r.Max.Y < p[1]+1 {
			r.Max.Y = p[1] + 1
		}
	}
	return r
}

type prep [][]int

func prepare(in string) prep {
	ss := strings.Split(in, "\n")
	ps := prep{}
	r := regexp.MustCompile("-?\\d+")
	for _, s := range ss {
		m := r.FindAllString(s, -1)
		p := sliceAtoi(m)
		ps = append(ps, p)
	}
	return ps
}

func _print(ps [][]int, b rect) {
	w := b.Dx()
	h := b.Dy()
	m := make([][]rune, h)
	for y := 0; y < h; y++ {
		m[y] = []rune(strings.Repeat(" ", w))
	}
	for _, p := range ps {
		m[p[1]-b.Min.Y][p[0]-b.Min.X] = '*'
	}
	for _, r := range m {
		fmt.Println(string(r))
	}
}

func fly(ps prep) {
	for i := 1; ; i++ {
		for _, p := range ps {
			p[0] += p[2]
			p[1] += p[3]
		}
		b := bounds(ps)
		if b.Dy() == 10 {
			fmt.Println("part 1:")
			_print(ps, b)
			fmt.Println("part 2:", i)
			return
		}
	}
}

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		t0 := time.Now()
		ps := prepare(in)
		fly(ps)
		t1 := time.Now()
		fmt.Println(t1.Sub(t0))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `position=< 20247,  40241> velocity=<-2, -4>
position=< 10184, -29948> velocity=<-1,  3>
position=< 50313, -39966> velocity=<-5,  4>
position=<-19870, -19921> velocity=< 2,  2>
position=< 10224, -49995> velocity=<-1,  5>
position=<-19904,  20191> velocity=< 2, -2>
position=< 50300,  -9887> velocity=<-5,  1>
position=<-29924, -29942> velocity=< 3,  3>
position=<-29903, -29939> velocity=< 3,  3>
position=<-29905,  50273> velocity=< 3, -5>
position=< 20222, -29944> velocity=<-2,  3>
position=< 40293,  50268> velocity=<-4, -5>
position=< 40249,  40243> velocity=<-4, -4>
position=<-19869,  50274> velocity=< 2, -5>
position=<-29887,  50277> velocity=< 3, -5>
position=<-49943, -49998> velocity=< 5,  5>
position=<-49992,  10164> velocity=< 5, -1>
position=< 50316,  -9888> velocity=<-5,  1>
position=< 10208, -29942> velocity=<-1,  3>
position=< 40275, -49998> velocity=<-4,  5>
position=< 20254,  40246> velocity=<-2, -4>
position=< 30231,  10164> velocity=<-3, -1>
position=<-19884,  50271> velocity=< 2, -5>
position=<-39935,  -9890> velocity=< 4,  1>
position=< 30241,  30214> velocity=<-3, -3>
position=< 20216, -49993> velocity=<-2,  5>
position=<-49962, -29943> velocity=< 5,  3>
position=<-29884,  10164> velocity=< 3, -1>
position=<-29936,  30214> velocity=< 3, -3>
position=< 30274, -29948> velocity=<-3,  3>
position=< 30246, -29946> velocity=<-3,  3>
position=<-29895,  -9893> velocity=< 3,  1>
position=<-29879,  10165> velocity=< 3, -1>
position=<-19895, -49998> velocity=< 2,  5>
position=<-39930,  40250> velocity=< 4, -4>
position=<-29888, -29945> velocity=< 3,  3>
position=< 40270,  20189> velocity=<-4, -2>
position=< 20212, -49998> velocity=<-2,  5>
position=< 40265,  30222> velocity=<-4, -3>
position=< -9876, -50002> velocity=< 1,  5>
position=<-39951,  40247> velocity=< 4, -4>
position=<-19865,  20194> velocity=< 2, -2>
position=<-29924,  40242> velocity=< 3, -4>
position=< 50313,  30216> velocity=<-5, -3>
position=< 10216, -49994> velocity=<-1,  5>
position=< -9846, -39968> velocity=< 1,  4>
position=< 40265,  20194> velocity=<-4, -2>
position=< -9870, -29939> velocity=< 1,  3>
position=<-29903, -19918> velocity=< 3,  2>
position=< 30255, -49998> velocity=<-3,  5>
position=<-29879,  20192> velocity=< 3, -2>
position=< 50277,  40241> velocity=<-5, -4>
position=< 40269,  40245> velocity=<-4, -4>
position=< 40265,  50277> velocity=<-4, -5>
position=<-29940, -50000> velocity=< 3,  5>
position=<-39935,  40247> velocity=< 4, -4>
position=<-29914,  40250> velocity=< 3, -4>
position=< -9829,  40242> velocity=< 1, -4>
position=<-29891,  20196> velocity=< 3, -2>
position=<-39954, -19919> velocity=< 4,  2>
position=< 10202, -39971> velocity=<-1,  4>
position=< 40258,  30214> velocity=<-4, -3>
position=< 30243, -29940> velocity=<-3,  3>
position=<-19893,  30220> velocity=< 2, -3>
position=<-29932, -39966> velocity=< 3,  4>
position=< 10176, -29941> velocity=<-1,  3>
position=<-39942,  40245> velocity=< 4, -4>
position=< 30242,  10164> velocity=<-3, -1>
position=< -9868, -39975> velocity=< 1,  4>
position=< -9825,  30218> velocity=< 1, -3>
position=< 40305, -19915> velocity=<-4,  2>
position=<-49954, -39967> velocity=< 5,  4>
position=< 20215,  20193> velocity=<-2, -2>
position=<-49994,  50275> velocity=< 5, -5>
position=<-19856,  40246> velocity=< 2, -4>
position=< 40273, -39972> velocity=<-4,  4>
position=< 50327, -19912> velocity=<-5,  2>
position=< 20256,  20193> velocity=<-2, -2>
position=<-19868,  10163> velocity=< 2, -1>
position=<-29930, -50002> velocity=< 3,  5>
position=<-19897,  50275> velocity=< 2, -5>
position=<-29911,  20195> velocity=< 3, -2>
position=< 20238,  20191> velocity=<-2, -2>
position=<-39959, -19918> velocity=< 4,  2>
position=< 50284,  40249> velocity=<-5, -4>
position=< -9857, -50001> velocity=< 1,  5>
position=< 50284, -39972> velocity=<-5,  4>
position=<-29938,  10164> velocity=< 3, -1>
position=< 50333, -49997> velocity=<-5,  5>
position=<-19878, -19916> velocity=< 2,  2>
position=< 10178,  30214> velocity=<-1, -3>
position=<-39919, -50002> velocity=< 4,  5>
position=< 10176, -49995> velocity=<-1,  5>
position=< 20251,  -9890> velocity=<-2,  1>
position=< 20252,  30219> velocity=<-2, -3>
position=< 30281,  40246> velocity=<-3, -4>
position=< 40289,  10160> velocity=<-4, -1>
position=< 30222,  20192> velocity=<-3, -2>
position=<-29899, -39971> velocity=< 3,  4>
position=< 30278,  10163> velocity=<-3, -1>
position=< -9857,  -9889> velocity=< 1,  1>
position=<-39906, -39968> velocity=< 4,  4>
position=<-39959, -50002> velocity=< 4,  5>
position=<-19895,  20187> velocity=< 2, -2>
position=< 20253,  30219> velocity=<-2, -3>
position=< 20211,  -9885> velocity=<-2,  1>
position=< 50313,  10162> velocity=<-5, -1>
position=< 30250,  50268> velocity=<-3, -5>
position=<-39951, -29941> velocity=< 4,  3>
position=< 10229, -49995> velocity=<-1,  5>
position=<-19868,  20196> velocity=< 2, -2>
position=< 30246, -39971> velocity=<-3,  4>
position=<-19864, -39975> velocity=< 2,  4>
position=< 20248,  40243> velocity=<-2, -4>
position=< 50287,  -9890> velocity=<-5,  1>
position=<-19873,  30214> velocity=< 2, -3>
position=<-39942, -39975> velocity=< 4,  4>
position=< 30262, -19914> velocity=<-3,  2>
position=<-39918, -29948> velocity=< 4,  3>
position=< -9844,  30214> velocity=< 1, -3>
position=< -9825,  20191> velocity=< 1, -2>
position=< 20240,  -9885> velocity=<-2,  1>
position=< 40258,  30214> velocity=<-4, -3>
position=<-29896,  20194> velocity=< 3, -2>
position=<-19871, -29948> velocity=< 2,  3>
position=< 40249, -49993> velocity=<-4,  5>
position=< 40281,  -9885> velocity=<-4,  1>
position=< 50284,  50273> velocity=<-5, -5>
position=<-29892, -29941> velocity=< 3,  3>
position=< 10189,  40243> velocity=<-1, -4>
position=< 40299,  40241> velocity=<-4, -4>
position=< -9854,  -9885> velocity=< 1,  1>
position=<-49953,  50268> velocity=< 5, -5>
position=<-19868,  20195> velocity=< 2, -2>
position=< 40301, -39972> velocity=<-4,  4>
position=< 40301, -19912> velocity=<-4,  2>
position=<-19901, -29948> velocity=< 2,  3>
position=<-19912,  40245> velocity=< 2, -4>
position=< 20222, -39975> velocity=<-2,  4>
position=< 40270,  10161> velocity=<-4, -1>
position=< 50305,  40249> velocity=<-5, -4>
position=<-39947, -19915> velocity=< 4,  2>
position=< 40301,  30214> velocity=<-4, -3>
position=< 20255,  50273> velocity=<-2, -5>
position=<-39955, -49998> velocity=< 4,  5>
position=< 30278, -50000> velocity=<-3,  5>
position=< 10212,  10167> velocity=<-1, -1>
position=< 30258,  30221> velocity=<-3, -3>
position=<-49985, -49998> velocity=< 5,  5>
position=<-39940,  50268> velocity=< 4, -5>
position=<-39914,  30214> velocity=< 4, -3>
position=<-39947,  50275> velocity=< 4, -5>
position=< 40273,  40249> velocity=<-4, -4>
position=< 30254,  20195> velocity=<-3, -2>
position=< 30283,  50276> velocity=<-3, -5>
position=<-39914,  20189> velocity=< 4, -2>
position=< -9853, -49998> velocity=< 1,  5>
position=< -9854, -39972> velocity=< 1,  4>
position=<-39935,  50270> velocity=< 4, -5>
position=<-29898,  40241> velocity=< 3, -4>
position=< 10193, -49993> velocity=<-1,  5>
position=<-29880,  50273> velocity=< 3, -5>
position=< 40289, -19917> velocity=<-4,  2>
position=< -9866,  30214> velocity=< 1, -3>
position=< 50321,  50270> velocity=<-5, -5>
position=<-49951,  20187> velocity=< 5, -2>
position=<-19900, -29945> velocity=< 2,  3>
position=< 50336, -29947> velocity=<-5,  3>
position=< -9838, -39975> velocity=< 1,  4>
position=< 10196,  10169> velocity=<-1, -1>
position=<-49965,  20194> velocity=< 5, -2>
position=< 20231,  40241> velocity=<-2, -4>
position=< 30273,  50272> velocity=<-3, -5>
position=< 40261, -49998> velocity=<-4,  5>
position=<-39964, -39975> velocity=< 4,  4>
position=<-39930,  20188> velocity=< 4, -2>
position=<-29916,  30214> velocity=< 3, -3>
position=< 10229,  40245> velocity=<-1, -4>
position=<-49994, -29942> velocity=< 5,  3>
position=< -9858,  20191> velocity=< 1, -2>
position=< 40306,  20188> velocity=<-4, -2>
position=< 10176,  30216> velocity=<-1, -3>
position=< 30230, -19916> velocity=<-3,  2>
position=<-49975,  40246> velocity=< 5, -4>
position=< 10168, -50002> velocity=<-1,  5>
position=<-49962, -29942> velocity=< 5,  3>
position=<-49970,  40244> velocity=< 5, -4>
position=<-19868, -39974> velocity=< 2,  4>
position=< 30273,  40241> velocity=<-3, -4>
position=<-29879,  40243> velocity=< 3, -4>
position=<-19887, -39966> velocity=< 2,  4>
position=<-49943, -50002> velocity=< 5,  5>
position=<-39951, -39967> velocity=< 4,  4>
position=<-49993,  40245> velocity=< 5, -4>
position=< 20246, -39975> velocity=<-2,  4>
position=< 40273,  40250> velocity=<-4, -4>
position=<-39939,  -9890> velocity=< 4,  1>
position=< -9886,  40245> velocity=< 1, -4>
position=< 40275, -39975> velocity=<-4,  4>
position=< 40274,  -9885> velocity=<-4,  1>
position=< 30256,  50272> velocity=<-3, -5>
position=< 40301,  10169> velocity=<-4, -1>
position=<-29936,  10160> velocity=< 3, -1>
position=< -9870, -39974> velocity=< 1,  4>
position=< 50321, -19919> velocity=<-5,  2>
position=< -9870,  10164> velocity=< 1, -1>
position=< 30278,  -9885> velocity=<-3,  1>
position=<-49966, -29948> velocity=< 5,  3>
position=< 20245,  10160> velocity=<-2, -1>
position=<-39967, -49993> velocity=< 4,  5>
position=< 20211,  50270> velocity=<-2, -5>
position=<-49960,  50268> velocity=< 5, -5>
position=< 30278,  40246> velocity=<-3, -4>
position=< 30246,  50275> velocity=<-3, -5>
position=< 10168,  -9890> velocity=<-1,  1>
position=< -9846,  -9894> velocity=< 1,  1>
position=< 50300, -29942> velocity=<-5,  3>
position=<-29911, -29946> velocity=< 3,  3>
position=<-49991,  10160> velocity=< 5, -1>
position=< 40286,  40244> velocity=<-4, -4>
position=<-29879,  20193> velocity=< 3, -2>
position=< 20256,  50273> velocity=<-2, -5>
position=< -9869, -19921> velocity=< 1,  2>
position=<-19863,  30223> velocity=< 2, -3>
position=<-29896, -19914> velocity=< 3,  2>
position=<-49938,  20189> velocity=< 5, -2>
position=<-19905, -50000> velocity=< 2,  5>
position=<-29913, -49998> velocity=< 3,  5>
position=<-19887, -39971> velocity=< 2,  4>
position=< 10218,  40246> velocity=<-1, -4>
position=< 10168,  10161> velocity=<-1, -1>
position=< 40275,  40245> velocity=<-4, -4>
position=< 10217, -50002> velocity=<-1,  5>
position=< 20219, -50002> velocity=<-2,  5>
position=< -9867, -49997> velocity=< 1,  5>
position=<-49957,  50276> velocity=< 5, -5>
position=<-39938, -50001> velocity=< 4,  5>
position=<-19878,  30214> velocity=< 2, -3>
position=< 30246,  40242> velocity=<-3, -4>
position=<-19881,  -9889> velocity=< 2,  1>
position=< 10168, -29940> velocity=<-1,  3>
position=<-49960,  30214> velocity=< 5, -3>
position=<-29940, -39972> velocity=< 3,  4>
position=<-29900,  30215> velocity=< 3, -3>
position=< 20240,  20195> velocity=<-2, -2>
position=< 30263, -39971> velocity=<-3,  4>
position=< 50287, -19921> velocity=<-5,  2>
position=< -9886,  10160> velocity=< 1, -1>
position=< 30230,  20188> velocity=<-3, -2>
position=< 50281, -19921> velocity=<-5,  2>
position=< 20223, -39966> velocity=<-2,  4>
position=<-19863, -50002> velocity=< 2,  5>
position=< 50328,  20190> velocity=<-5, -2>
position=<-29921, -19917> velocity=< 3,  2>
position=< 40253, -29944> velocity=<-4,  3>
position=<-49986, -29943> velocity=< 5,  3>
position=< 50296, -19917> velocity=<-5,  2>
position=< 50316,  30216> velocity=<-5, -3>
position=< 30279,  20188> velocity=<-3, -2>
position=<-19854,  30214> velocity=< 2, -3>
position=<-19903, -39971> velocity=< 2,  4>
position=<-19897,  30219> velocity=< 2, -3>
position=< 10197, -50000> velocity=<-1,  5>
position=<-49978,  -9891> velocity=< 5,  1>
position=< 40265, -19916> velocity=<-4,  2>
position=<-49954, -39973> velocity=< 5,  4>
position=< 10170,  20187> velocity=<-1, -2>
position=< 30243, -29946> velocity=<-3,  3>
position=< -9875,  20191> velocity=< 1, -2>
position=<-19881, -39968> velocity=< 2,  4>
position=<-49976,  10160> velocity=< 5, -1>
position=<-39927,  30223> velocity=< 4, -3>
position=< 40282, -19921> velocity=<-4,  2>
position=< -9830,  20192> velocity=< 1, -2>
position=<-39959,  50274> velocity=< 4, -5>
position=< 30254, -19920> velocity=<-3,  2>
position=< 50335,  40241> velocity=<-5, -4>
position=< -9833,  40242> velocity=< 1, -4>
position=<-19881,  -9891> velocity=< 2,  1>
position=<-19894, -50002> velocity=< 2,  5>
position=< 30278, -19917> velocity=<-3,  2>
position=<-29932,  50272> velocity=< 3, -5>
position=< 10224,  10167> velocity=<-1, -1>
position=<-49950,  20193> velocity=< 5, -2>
position=<-29919, -39972> velocity=< 3,  4>
position=< -9851,  40245> velocity=< 1, -4>
position=< -9861,  20191> velocity=< 1, -2>
position=< 50276, -29940> velocity=<-5,  3>
position=< -9862,  -9887> velocity=< 1,  1>
position=<-49936,  10160> velocity=< 5, -1>
position=<-39919,  30223> velocity=< 4, -3>
position=<-29919,  -9885> velocity=< 3,  1>
position=< -9850, -29944> velocity=< 1,  3>
position=<-39938, -19915> velocity=< 4,  2>
position=< 10192,  20193> velocity=<-1, -2>
position=<-19889,  40246> velocity=< 2, -4>
position=< -9830, -49999> velocity=< 1,  5>
position=<-39926,  20191> velocity=< 4, -2>
position=<-39906, -19913> velocity=< 4,  2>
position=< 40273,  30216> velocity=<-4, -3>
position=<-19897, -50001> velocity=< 2,  5>
position=< -9878,  30216> velocity=< 1, -3>
position=< 20212,  40241> velocity=<-2, -4>
position=<-39951,  30220> velocity=< 4, -3>
position=< -9830,  30222> velocity=< 1, -3>
position=< -9859,  50272> velocity=< 1, -5>
position=< 20251, -19913> velocity=<-2,  2>
position=<-49961,  30218> velocity=< 5, -3>
position=< 50321,  -9885> velocity=<-5,  1>
position=< 20196,  -9894> velocity=<-2,  1>
position=< 20227,  10160> velocity=<-2, -1>
position=<-29892,  -9886> velocity=< 3,  1>
position=< 30262, -49997> velocity=<-3,  5>
position=< 20230, -29943> velocity=<-2,  3>
position=<-39964,  40245> velocity=< 4, -4>
position=< 50316, -49999> velocity=<-5,  5>
position=<-39918, -29942> velocity=< 4,  3>
position=< 30222,  50268> velocity=<-3, -5>
position=< 10184,  10165> velocity=<-1, -1>
position=< 50311,  50272> velocity=<-5, -5>
position=< 30258,  50274> velocity=<-3, -5>
position=<-39950,  30218> velocity=< 4, -3>
position=<-49941,  30214> velocity=< 5, -3>
position=<-19913,  50275> velocity=< 2, -5>
position=< -9870,  30218> velocity=< 1, -3>
position=< 30226,  30218> velocity=<-3, -3>
position=<-29919, -49999> velocity=< 3,  5>
position=< 20219, -39967> velocity=<-2,  4>
position=< -9882, -39975> velocity=< 1,  4>
position=< 10171, -39975> velocity=<-1,  4>
position=<-29908,  10165> velocity=< 3, -1>
position=<-39923, -49998> velocity=< 4,  5>
position=<-39924,  -9889> velocity=< 4,  1>
position=<-29908,  30216> velocity=< 3, -3>
position=< 30273, -29944> velocity=<-3,  3>
position=<-29903, -39972> velocity=< 3,  4>
position=<-29927,  20188> velocity=< 3, -2>
position=< 50319, -29943> velocity=<-5,  3>
position=< 30274, -29939> velocity=<-3,  3>
position=<-39956, -50002> velocity=< 4,  5>
position=< 30283, -49995> velocity=<-3,  5>
position=< 30275,  50270> velocity=<-3, -5>
position=<-49952, -39971> velocity=< 5,  4>
position=< 40307,  20187> velocity=<-4, -2>
position=<-49986,  10166> velocity=< 5, -1>
position=<-29924,  20187> velocity=< 3, -2>
position=< -9846,  -9890> velocity=< 1,  1>
position=< 40257,  30221> velocity=<-4, -3>
position=< 30271, -49993> velocity=<-3,  5>
position=<-19889,  50274> velocity=< 2, -5>
position=<-29913, -29939> velocity=< 3,  3>
position=< 30238,  30218> velocity=<-3, -3>
position=< 20227, -19918> velocity=<-2,  2>
position=<-49933,  -9888> velocity=< 5,  1>
position=< 40268, -39971> velocity=<-4,  4>
position=<-49965,  -9887> velocity=< 5,  1>
position=<-29932,  10169> velocity=< 3, -1>
position=<-49933, -49999> velocity=< 5,  5>
position=<-19880,  50268> velocity=< 2, -5>
position=< 50276,  -9887> velocity=<-5,  1>
position=< 10197,  -9892> velocity=<-1,  1>
position=< 30266,  -9894> velocity=<-3,  1>
position=<-49954,  40247> velocity=< 5, -4>
position=< 40305, -29939> velocity=<-4,  3>
position=< 30250, -39971> velocity=<-3,  4>
position=< 10181,  50271> velocity=<-1, -5>
position=< 20232, -29946> velocity=<-2,  3>
position=< 30272,  50273> velocity=<-3, -5>
position=< 10229, -29939> velocity=<-1,  3>
position=<-19873, -19917> velocity=< 2,  2>
position=<-19886,  -9894> velocity=< 2,  1>
position=<-49934, -29943> velocity=< 5,  3>
position=< -9828,  30219> velocity=< 1, -3>
position=<-19893,  50268> velocity=< 2, -5>
position=<-39965,  50268> velocity=< 4, -5>
position=<-19881,  20196> velocity=< 2, -2>
position=< -9862, -39972> velocity=< 1,  4>
position=< -9849,  10168> velocity=< 1, -1>
position=< 50316, -29946> velocity=<-5,  3>
position=< -9838,  30214> velocity=< 1, -3>
position=< 50305,  50269> velocity=<-5, -5>
position=<-49933,  -9891> velocity=< 5,  1>
position=< 50321,  30215> velocity=<-5, -3>`,
	"google": `position=<-40181, -50237> velocity=< 4,  5>
position=<-40122,  30405> velocity=< 4, -3>
position=<-40158, -50246> velocity=< 4,  5>
position=<-50220,  50573> velocity=< 5, -5>
position=< 20317, -30082> velocity=<-2,  3>
position=<-50207,  40490> velocity=< 5, -4>
position=< 10260,  -9913> velocity=<-1,  1>
position=<-50212, -20003> velocity=< 5,  2>
position=< 40498, -40158> velocity=<-4,  4>
position=< 30446,  20330> velocity=<-3, -2>
position=< 10255, -50241> velocity=<-1,  5>
position=< 10263, -40158> velocity=<-1,  4>
position=<-40178,  40492> velocity=< 4, -4>
position=< 40487,  30411> velocity=<-4, -3>
position=<-19978,  30411> velocity=< 2, -3>
position=<-50215, -40162> velocity=< 5,  4>
position=< 40474, -30078> velocity=<-4,  3>
position=< 20309,  10241> velocity=<-2, -1>
position=<-50235,  40484> velocity=< 5, -4>
position=< 10259, -50246> velocity=<-1,  5>
position=< 10228,  10248> velocity=<-1, -1>
position=<-19970,  40487> velocity=< 2, -4>
position=<-30093, -50246> velocity=< 3,  5>
position=< -9929,  -9918> velocity=< 1,  1>
position=< 20307, -50246> velocity=<-2,  5>
position=< 30417, -19994> velocity=<-3,  2>
position=< 10279,  40491> velocity=<-1, -4>
position=<-30089, -19997> velocity=< 3,  2>
position=<-40132,  -9922> velocity=< 4,  1>
position=<-40132,  10244> velocity=< 4, -1>
position=< 10234, -50246> velocity=<-1,  5>
position=< 10247,  50569> velocity=<-1, -5>
position=<-20012,  50572> velocity=< 2, -5>
position=< 20348, -50243> velocity=<-2,  5>
position=< 40506,  -9922> velocity=<-4,  1>
position=<-30067,  30411> velocity=< 3, -3>
position=< -9928,  30402> velocity=< 1, -3>
position=< 40490,  40492> velocity=<-4, -4>
position=<-50247, -40156> velocity=< 5,  4>
position=<-40182, -40158> velocity=< 4,  4>
position=< 10282,  40492> velocity=<-1, -4>
position=<-40169, -50245> velocity=< 4,  5>
position=< 30436,  10244> velocity=<-3, -1>
position=<-30042,  30406> velocity=< 3, -3>
position=< 10236, -50244> velocity=<-1,  5>
position=< 40523, -50246> velocity=<-4,  5>
position=<-30101,  -9916> velocity=< 3,  1>
position=< 50571, -50238> velocity=<-5,  5>
position=<-19980, -40157> velocity=< 2,  4>
position=<-30092,  30406> velocity=< 3, -3>
position=<-20009,  10244> velocity=< 2, -1>
position=< 30433,  30411> velocity=<-3, -3>
position=< 40514, -50242> velocity=<-4,  5>
position=< 40479, -30075> velocity=<-4,  3>
position=<-20020,  50566> velocity=< 2, -5>
position=< 20309, -30076> velocity=<-2,  3>
position=<-50237,  20324> velocity=< 5, -2>
position=< 20341, -40165> velocity=<-2,  4>
position=< -9919,  30409> velocity=< 1, -3>
position=< 50555,  30405> velocity=<-5, -3>
position=<-50205,  30402> velocity=< 5, -3>
position=< 20365,  40485> velocity=<-2, -4>
position=< -9939,  30407> velocity=< 1, -3>
position=< 30406,  30410> velocity=<-3, -3>
position=< 40527,  30411> velocity=<-4, -3>
position=< 20339, -40156> velocity=<-2,  4>
position=<-20012,  -9922> velocity=< 2,  1>
position=<-30049,  30406> velocity=< 3, -3>
position=<-20020,  30407> velocity=< 2, -3>
position=< 20330, -50240> velocity=<-2,  5>
position=< 20338,  40487> velocity=<-2, -4>
position=< 40526,  50567> velocity=<-4, -5>
position=<-30040, -50244> velocity=< 3,  5>
position=<-40148, -19999> velocity=< 4,  2>
position=< 20325,  -9922> velocity=<-2,  1>
position=< 20305,  30411> velocity=<-2, -3>
position=<-20020,  30409> velocity=< 2, -3>
position=< -9931,  50566> velocity=< 1, -5>
position=< 50582,  50564> velocity=<-5, -5>
position=<-50259, -30084> velocity=< 5,  3>
position=<-19985,  50564> velocity=< 2, -5>
position=< 10256, -19999> velocity=<-1,  2>
position=< -9918, -19995> velocity=< 1,  2>
position=< 20349,  20330> velocity=<-2, -2>
position=< 50592, -50246> velocity=<-5,  5>
position=<-30093,  40485> velocity=< 3, -4>
position=< -9894,  50565> velocity=< 1, -5>
position=< 30418, -19999> velocity=<-3,  2>
position=<-30065,  10249> velocity=< 3, -1>
position=<-40123,  40492> velocity=< 4, -4>
position=<-40150,  40489> velocity=< 4, -4>
position=<-40182,  20329> velocity=< 4, -2>
position=<-19977,  -9918> velocity=< 2,  1>
position=< 30388,  30411> velocity=<-3, -3>
position=< 30393,  10245> velocity=<-3, -1>
position=< 40471,  -9914> velocity=<-4,  1>
position=<-40122,  10240> velocity=< 4, -1>
position=< 40524, -40156> velocity=<-4,  4>
position=< -9923,  40484> velocity=< 1, -4>
position=<-30073,  40491> velocity=< 3, -4>
position=< 50590,  40483> velocity=<-5, -4>
position=< 10255, -40159> velocity=<-1,  4>
position=< 10241, -30079> velocity=<-1,  3>
position=< 20364, -40156> velocity=<-2,  4>
position=<-40133,  -9922> velocity=< 4,  1>
position=< -9897,  40483> velocity=< 1, -4>
position=<-50247,  30411> velocity=< 5, -3>
position=< 10227,  -9922> velocity=<-1,  1>
position=< 30421, -30080> velocity=<-3,  3>
position=<-50239,  10244> velocity=< 5, -1>
position=<-30077, -19998> velocity=< 3,  2>
position=<-19994, -19997> velocity=< 2,  2>
position=<-30091,  20321> velocity=< 3, -2>
position=<-50239,  -9915> velocity=< 5,  1>
position=<-50239, -50238> velocity=< 5,  5>
position=<-50254,  50568> velocity=< 5, -5>
position=< 20312,  30410> velocity=<-2, -3>
position=< 10258,  30411> velocity=<-1, -3>
position=<-30061,  50573> velocity=< 3, -5>
position=< -9928,  10245> velocity=< 1, -1>
position=< 10271, -30080> velocity=<-1,  3>
position=<-19972,  10247> velocity=< 2, -1>
position=< 40493,  50566> velocity=<-4, -5>
position=<-30096, -30083> velocity=< 3,  3>
position=< 20332,  -9921> velocity=<-2,  1>
position=< 20357, -30081> velocity=<-2,  3>
position=<-50263,  40486> velocity=< 5, -4>
position=< 10267, -50237> velocity=<-1,  5>
position=<-50239,  10242> velocity=< 5, -1>
position=< 50583,  20321> velocity=<-5, -2>
position=< 30425, -30076> velocity=<-3,  3>
position=<-50227,  30411> velocity=< 5, -3>
position=<-40125,  20330> velocity=< 4, -2>
position=<-20020, -40163> velocity=< 2,  4>
position=< 10268, -40163> velocity=<-1,  4>
position=<-40181,  20321> velocity=< 4, -2>
position=< 10231,  20330> velocity=<-1, -2>
position=<-30101,  20322> velocity=< 3, -2>
position=<-40157,  20326> velocity=< 4, -2>
position=< 30394,  10244> velocity=<-3, -1>
position=<-40125,  50564> velocity=< 4, -5>
position=< 40498,  20323> velocity=<-4, -2>
position=< 50563,  10248> velocity=<-5, -1>
position=<-30080,  50573> velocity=< 3, -5>
position=< -9896,  20321> velocity=< 1, -2>
position=<-50236,  50566> velocity=< 5, -5>
position=<-40142, -19994> velocity=< 4,  2>
position=< -9939,  -9916> velocity=< 1,  1>
position=<-40155,  -9915> velocity=< 4,  1>
position=< 20362,  20326> velocity=<-2, -2>
position=< 20347,  20325> velocity=<-2, -2>
position=< 30397, -19999> velocity=<-3,  2>
position=< 10235,  20327> velocity=<-1, -2>
position=<-50207, -30075> velocity=< 5,  3>
position=< 20333, -20003> velocity=<-2,  2>
position=< 40499,  40483> velocity=<-4, -4>
position=< 30385,  -9921> velocity=<-3,  1>
position=<-50210,  10242> velocity=< 5, -1>
position=<-40148,  50564> velocity=< 4, -5>
position=< 50582,  40483> velocity=<-5, -4>
position=<-30060, -30078> velocity=< 3,  3>
position=< 10271,  30406> velocity=<-1, -3>
position=< 50599, -40165> velocity=<-5,  4>
position=< 30401,  50565> velocity=<-3, -5>
position=< 20333, -50237> velocity=<-2,  5>
position=<-30053,  10249> velocity=< 3, -1>
position=< 20308, -50237> velocity=<-2,  5>
position=<-40163, -40161> velocity=< 4,  4>
position=< 40474, -50238> velocity=<-4,  5>
position=<-30053,  20321> velocity=< 3, -2>
position=<-30082, -50241> velocity=< 3,  5>
position=< 30429, -50246> velocity=<-3,  5>
position=< 30426,  20321> velocity=<-3, -2>
position=<-50250, -40157> velocity=< 5,  4>
position=< 40483, -50244> velocity=<-4,  5>
position=< 20320, -20003> velocity=<-2,  2>
position=< 40479, -40156> velocity=<-4,  4>
position=< 50568, -50238> velocity=<-5,  5>
position=< 50591, -30081> velocity=<-5,  3>
position=< 40526,  10240> velocity=<-4, -1>
position=< 50595,  20329> velocity=<-5, -2>
position=< 50592, -30075> velocity=<-5,  3>
position=<-30089,  50571> velocity=< 3, -5>
position=<-40169,  20330> velocity=< 4, -2>
position=<-20011,  30402> velocity=< 2, -3>
position=< 10265,  20330> velocity=<-1, -2>
position=<-19988, -20001> velocity=< 2,  2>
position=< 30397, -30077> velocity=<-3,  3>
position=< 50595,  30407> velocity=<-5, -3>
position=<-19972,  -9920> velocity=< 2,  1>
position=<-19978, -30084> velocity=< 2,  3>
position=< 30401,  10240> velocity=<-3, -1>
position=<-30042,  20325> velocity=< 3, -2>
position=< 10267,  40483> velocity=<-1, -4>
position=< -9939, -20001> velocity=< 1,  2>
position=< -9922, -30081> velocity=< 1,  3>
position=< 30393,  20321> velocity=<-3, -2>
position=< 10235, -30077> velocity=<-1,  3>
position=< 50607,  50564> velocity=<-5, -5>
position=<-40162, -20001> velocity=< 4,  2>
position=< 10271,  20326> velocity=<-1, -2>
position=< -9881, -40160> velocity=< 1,  4>
position=< 10247,  10246> velocity=<-1, -1>
position=<-20020, -30079> velocity=< 2,  3>
position=<-40123, -19994> velocity=< 4,  2>
position=< 30442, -40159> velocity=<-3,  4>
position=< 40483,  10246> velocity=<-4, -1>
position=< -9914, -30080> velocity=< 1,  3>
position=< 50572, -30079> velocity=<-5,  3>
position=< 30435, -30084> velocity=<-3,  3>
position=<-19995, -19999> velocity=< 2,  2>
position=<-30077,  -9916> velocity=< 3,  1>
position=<-19988,  40484> velocity=< 2, -4>
position=< 50555,  30403> velocity=<-5, -3>
position=< 50595,  30402> velocity=<-5, -3>
position=< -9878,  50564> velocity=< 1, -5>
position=< -9922, -50244> velocity=< 1,  5>
position=<-40141,  30402> velocity=< 4, -3>
position=< 30398, -50244> velocity=<-3,  5>
position=< -9904, -19999> velocity=< 1,  2>
position=<-50243,  40490> velocity=< 5, -4>
position=< 40490,  -9918> velocity=<-4,  1>
position=< -9907,  -9914> velocity=< 1,  1>
position=<-50242,  10249> velocity=< 5, -1>
position=<-40171,  20326> velocity=< 4, -2>
position=< 10247, -40164> velocity=<-1,  4>
position=<-30098, -30084> velocity=< 3,  3>
position=<-40155,  10242> velocity=< 4, -1>
position=< -9895,  10249> velocity=< 1, -1>
position=< 30442,  50573> velocity=<-3, -5>
position=<-19967,  50565> velocity=< 2, -5>
position=< 50559,  20321> velocity=<-5, -2>
position=< -9919, -30081> velocity=< 1,  3>
position=< -9905,  30406> velocity=< 1, -3>
position=< 50572,  -9917> velocity=<-5,  1>
position=< 10255,  10248> velocity=<-1, -1>
position=< -9938,  -9922> velocity=< 1,  1>
position=< 40483,  40486> velocity=<-4, -4>
position=<-50242,  40484> velocity=< 5, -4>
position=< 20352, -30083> velocity=<-2,  3>
position=< 30446,  40492> velocity=<-3, -4>
position=< -9936, -50237> velocity=< 1,  5>
position=< 50568, -20002> velocity=<-5,  2>
position=<-20016,  20330> velocity=< 2, -2>
position=< 50596, -20003> velocity=<-5,  2>
position=< 30425,  10240> velocity=<-3, -1>
position=< 50567, -30077> velocity=<-5,  3>
position=< 50606,  20325> velocity=<-5, -2>
position=< 50571, -20000> velocity=<-5,  2>
position=< 20337,  40492> velocity=<-2, -4>
position=< 40478,  -9918> velocity=<-4,  1>
position=<-30069,  20321> velocity=< 3, -2>
position=< 30401, -19994> velocity=<-3,  2>
position=< 50583, -50237> velocity=<-5,  5>
position=< 20317, -20002> velocity=<-2,  2>
position=< 10231,  -9915> velocity=<-1,  1>
position=< 20328,  30409> velocity=<-2, -3>
position=< 10233,  10244> velocity=<-1, -1>
position=<-19972,  50565> velocity=< 2, -5>
position=<-40134,  40491> velocity=< 4, -4>
position=< 40466,  50568> velocity=<-4, -5>
position=<-50215, -30079> velocity=< 5,  3>
position=< 10259,  20325> velocity=<-1, -2>
position=< -9939,  40484> velocity=< 1, -4>
position=<-50231,  50568> velocity=< 5, -5>
position=< 40467,  40492> velocity=<-4, -4>
position=< 10255, -30083> velocity=<-1,  3>
position=<-50222,  -9916> velocity=< 5,  1>
position=< -9927,  30408> velocity=< 1, -3>
position=< 20304,  20329> velocity=<-2, -2>
position=<-30057,  10249> velocity=< 3, -1>
position=< 20312,  20323> velocity=<-2, -2>
position=<-40182,  30405> velocity=< 4, -3>
position=<-30050, -40161> velocity=< 3,  4>
position=< 20317,  10243> velocity=<-2, -1>
position=<-50250,  -9919> velocity=< 5,  1>
position=<-30093,  50568> velocity=< 3, -5>
position=< 50555, -30080> velocity=<-5,  3>
position=< 50571,  -9915> velocity=<-5,  1>
position=< 40468,  30411> velocity=<-4, -3>
position=<-40180, -50237> velocity=< 4,  5>
position=< 10264,  -9913> velocity=<-1,  1>
position=< 50560, -50243> velocity=<-5,  5>
position=<-40170, -50242> velocity=< 4,  5>
position=< 20347,  10240> velocity=<-2, -1>
position=<-20018, -20003> velocity=< 2,  2>
position=< -9878,  20322> velocity=< 1, -2>
position=<-30069,  40486> velocity=< 3, -4>
position=<-50214,  10244> velocity=< 5, -1>
position=< 30427,  30407> velocity=<-3, -3>
position=< -9918,  -9922> velocity=< 1,  1>
position=< 40469,  10249> velocity=<-4, -1>
position=< 50575,  50572> velocity=<-5, -5>
position=< 50567,  50570> velocity=<-5, -5>
position=< 30427,  -9917> velocity=<-3,  1>
position=<-40162,  40485> velocity=< 4, -4>
position=<-19988,  40488> velocity=< 2, -4>
position=< -9890,  50568> velocity=< 1, -5>
position=<-30098,  20321> velocity=< 3, -2>
position=<-50207,  40483> velocity=< 5, -4>
position=< 40517,  50568> velocity=<-4, -5>
position=< 10275,  30402> velocity=<-1, -3>
position=< 50588,  10246> velocity=<-5, -1>
position=< -9883,  -9914> velocity=< 1,  1>
position=<-40166, -50238> velocity=< 4,  5>
position=< 10255, -40156> velocity=<-1,  4>
position=< 40524,  -9913> velocity=<-4,  1>
position=< 50551,  30402> velocity=<-5, -3>
position=<-30077, -20003> velocity=< 3,  2>
position=<-40182, -50243> velocity=< 4,  5>
position=<-30053, -19994> velocity=< 3,  2>
position=< 40506, -30084> velocity=<-4,  3>
position=<-50244,  -9918> velocity=< 5,  1>
position=< 30395,  20321> velocity=<-3, -2>
position=< 20338, -40156> velocity=<-2,  4>
position=<-30057,  50567> velocity=< 3, -5>
position=<-30100, -30084> velocity=< 3,  3>
position=<-20003,  50571> velocity=< 2, -5>
position=< 20328, -30075> velocity=<-2,  3>
position=< 20317,  -9921> velocity=<-2,  1>
position=< 20336, -40157> velocity=<-2,  4>
position=<-19988,  10246> velocity=< 2, -1>
position=<-40158,  50565> velocity=< 4, -5>
position=< 40506,  30410> velocity=<-4, -3>
position=<-19971,  -9918> velocity=< 2,  1>
position=< 50571, -19995> velocity=<-5,  2>
position=<-40131,  10240> velocity=< 4, -1>
position=<-30099, -19994> velocity=< 3,  2>
position=< 10223,  30409> velocity=<-1, -3>
position=< 20364, -50243> velocity=<-2,  5>
position=< 30393, -40162> velocity=<-3,  4>
position=< 40526,  40492> velocity=<-4, -4>
position=<-50239,  -9917> velocity=< 5,  1>
position=< -9907, -30077> velocity=< 1,  3>
position=< -9939,  40491> velocity=< 1, -4>
position=<-19964,  30410> velocity=< 2, -3>
position=< 50580, -19994> velocity=<-5,  2>
position=<-40164,  50568> velocity=< 4, -5>
position=<-30053, -40165> velocity=< 3,  4>
position=< 10240, -40158> velocity=<-1,  4>
position=< -9923,  50564> velocity=< 1, -5>
position=< 50605,  50564> velocity=<-5, -5>
position=< 30393, -40159> velocity=<-3,  4>
position=< -9939,  50568> velocity=< 1, -5>
position=<-40150,  50573> velocity=< 4, -5>
position=< 50566, -40161> velocity=<-5,  4>
position=<-19975,  -9913> velocity=< 2,  1>
position=<-40162,  40489> velocity=< 4, -4>
position=< -9913, -30081> velocity=< 1,  3>
position=< 10271, -50240> velocity=<-1,  5>
position=<-40164, -30079> velocity=< 4,  3>
position=< -9883, -40158> velocity=< 1,  4>
position=< 20355, -30084> velocity=<-2,  3>
position=< 40510,  -9922> velocity=<-4,  1>
position=<-30090,  20325> velocity=< 3, -2>
position=<-20003, -19996> velocity=< 2,  2>
position=<-30085,  20322> velocity=< 3, -2>
position=<-30099, -20003> velocity=< 3,  2>
position=< -9880,  -9922> velocity=< 1,  1>
position=< 50607, -19994> velocity=<-5,  2>
position=<-50207,  30402> velocity=< 5, -3>`}
