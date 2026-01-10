// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	s, e Point
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			pa := corner(i+1, j)
			pb := corner(i, j)
			pc := corner(i, j+1)
			pd := corner(i+1, j+1)
			if isCrossLine(Line{pa, pb}, Line{pc, pd}) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				pa.x, pa.y, pb.x, pb.y, pc.x, pc.y, pd.x, pd.y)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) Point {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return Point{0, 0}
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return Point{sx, sy}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if math.IsInf(r, 0) {
		return 1
	} else if math.IsNaN(r) {
		return 0
	} else if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}

func isCrossLine(u, v Line) bool {
	// 外界矩形快速排斥实验
	if math.Max(u.s.x, u.e.x) < math.Min(v.s.x, v.e.x) ||
		math.Max(v.s.x, v.e.x) < math.Min(u.s.x, u.e.x) ||
		math.Max(u.s.y, u.e.y) < math.Min(v.s.y, v.e.y) ||
		math.Max(v.s.y, v.e.y) < math.Min(u.s.y, u.e.y) {
		return false
	}
	// 正交跨立实验
	cross1 := crossNum(u.s, u.e, v.s)
	cross2 := crossNum(u.s, u.e, v.e)
	cross3 := crossNum(v.s, v.e, u.s)
	cross4 := crossNum(v.s, v.e, u.e)

	epsilon := 1e-9
	// 可能重叠
	return cross1*cross2 <= epsilon && cross3*cross4 <= epsilon
}

func crossNum(A, B, P Point) float64 {
	return (B.x-A.x)*(P.y-A.y) - (B.y-A.y)*(P.x-A.x)
}
