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

type grid struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
	z                              float64
}

var grids []grid

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ok bool
			ax, ay, az, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, bz, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, cz, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dz, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			grids = append(grids, grid{ax, ay, bx, by, cx, cy, dx, dy, (az + bz + cz + dz) / 4})
		}
	}
	max := math.Inf(-1)
	min := math.Inf(1)
	for _, val := range grids {
		if val.z < min {
			min = val.z
		}
		if val.z > max {
			max = val.z
		}
	}

	// render
	for _, val := range grids {
		col := (val.z - min) / (max - min)
		fmt.Printf("<polygon fill='#%02x00%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			int(255*col), int(255*(1.0-col)), val.ax, val.ay, val.bx, val.by, val.cx, val.cy, val.dx, val.dy)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// z, ok := fSin(x, y)
	z, ok := fMogul(x, y)
	// z, ok := fEggcase(x, y)
	// z, ok := fSaddle(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, ok
}

// -1 ~ 1
func fSin(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	r = math.Sin(r)
	return r, true
}

// -0.2 ~ +0.2
func fMogul(x, y float64) (float64, bool) {
	r := math.Sin(x) * math.Sin(y)
	r *= 0.2
	return r, true
}

// -0.4 ~ 0
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

// -2 ~ 1.0
func fSaddle(x, y float64) (float64, bool) {
	lx := 0.2 * x
	r := 1.0 - lx*lx
	ok := true
	if r < -2 {
		ok = false
	}
	return r, ok
}
