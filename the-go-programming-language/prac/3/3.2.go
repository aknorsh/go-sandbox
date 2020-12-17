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
	zscale        = height * 0.2
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ok bool
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// z, ok := fSin(x, y)
	// z, ok := fMogul(x, y)
	// z, ok := fEggcase(x, y)
	// z, ok := fSaddle(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

func fSin(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	r = math.Sin(r)
	ok := false
	if r > 0 {
		ok = true
	}
	return r, ok
}

func fMogul(x, y float64) (float64, bool) {
	r := math.Sin(x) * math.Sin(y)
	r *= 0.2
	return r, true
}

func fEggcase(x, y float64) (float64, bool) {
	r := math.Sin(x) * math.Sin(y)
	r *= 0.4
	if r > 0 {
		r *= -1
	}
	ok := true
	if x < -math.Pi || 2*math.Pi < x || y < -3*math.Pi || 3*math.Pi < y {
		ok = false
	}
	return r, ok
}

func fSaddle(x, y float64) (float64, bool) {
	lx := 0.2 * x
	r := 1.0 - lx*lx
	ok := true
	if r < -2 {
		ok = false
	}
	return r, ok
}
