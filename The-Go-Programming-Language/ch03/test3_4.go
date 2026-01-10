// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", myhandler)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprint(w, getSVG())
}

func getSVG() string {
	svg := ""
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			pa, za := corner(i+1, j)
			pb, zb := corner(i, j)
			pc, zc := corner(i, j+1)
			pd, zd := corner(i+1, j+1)
			color := getColorFromZ(math.Max(math.Max(za, zb), math.Max(zc, zd)))
			if isCrossLine(Line{pa, pb}, Line{pc, pd}) {
				continue
			}
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:%s; fill:%s'/>\n",
				pa.x, pa.y, pb.x, pb.y, pc.x, pc.y, pd.x, pd.y, color, color)
		}
	}
	svg += fmt.Sprintln("</svg>")
	return svg
}

func getColorFromZ(z float64) string {
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return "#ffffff"
	}
	z = (z + 1) / 2
	r := int(255 * z)
	g := 0
	b := int(255 * (1 - z))

	color := (r << 16) | (g << 8) | b
	// color := int(z*(0xff0000-0xff)) + 0xff
	return fmt.Sprintf("#%06x", color)
}

func corner(i, j int) (Point, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return Point{0, 0}, z
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return Point{sx, sy}, z
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

func eggBox(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y)
}

func moguls(x, y float64) float64 {
	const (
		A1 = 0.5
		w1 = 2
		A2 = 0.3
		w2 = 5
	)
	return A1*math.Sin(w1*x)*math.Sin(w1*y) + A2*math.Sin(w2*x)*math.Sin(w2*y)
}

func saddle(x, y float64) float64 {
	return x*x - y*y
}

func gaussian(x, y float64) float64 {
	ro := 1.0
	return math.Exp((-(x*x + y*y)) / 2 * ro * ro)
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
