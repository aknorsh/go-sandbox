package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	if q, ok := queries["height"]; ok {
		height, _ = strconv.Atoi(q[0])
	} else {
		height = 320
	}
	if q, ok := queries["width"]; ok {
		width, _ = strconv.Atoi(q[0])
	} else {
		width = 600
	}
	if q, ok := queries["top"]; ok {
		switch q[0] {
		case "red":
			setPR(1)
		case "green":
			setPG(1)
		case "blue":
			setPB(1)
		}
	} else {
		setPR(1)
	}
	if q, ok := queries["bottom"]; ok {
		switch q[0] {
		case "red":
			setPR(-1)
		case "green":
			setPG(-1)
		case "blue":
			setPB(-1)
		}
	} else {
		setPB(-1)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	render(w)
	fmt.Fprintln(os.Stdout, "DEBUG: ", priorityR, priorityG, priorityB)
}

func setPR(val int8) {
	priorityR = val
	if priorityG == val {
		priorityG = 0
	}
	if priorityB == val {
		priorityB = 0
	}
}

func setPG(val int8) {
	priorityG = val
	if priorityG == val {
		priorityB = 0
	}
	if priorityB == val {
		priorityR = 0
	}
}

func setPB(val int8) {
	priorityB = val
	if priorityG == val {
		priorityR = 0
	}
	if priorityB == val {
		priorityG = 0
	}
}

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var width, height int = 600, 320
var priorityR, priorityG, priorityB int8 = 0, 1, -1
var xyscale float64 = float64(width) / 2.0 / xyrange
var zscale float64 = float64(height) * 0.2

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type grid struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
	z                              float64
}

var grids []grid

func render(out io.Writer) {

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
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
		rval := 0
		gval := 0
		bval := 0
		switch priorityR {
		case 1:
			rval = int(255 * col)
		case -1:
			rval = int(255 * (1.0 - col))
		}
		switch priorityG {
		case 1:
			gval = int(255 * col)
		case -1:
			gval = int(255 * (1.0 - col))
		}
		switch priorityB {
		case 1:
			bval = int(255 * col)
		case -1:
			bval = int(255 * (1.0 - col))
		}
		fmt.Fprintf(out, "<polygon fill='#%02x%02x%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			rval, gval, bval, val.ax, val.ay, val.bx, val.by, val.cx, val.cy, val.dx, val.dy)
	}
	fmt.Fprintln(out, "</svg>")

}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// z, ok := fSin(x, y)
	// z, ok := fMogul(x, y)
	z, ok := fEggcase(x, y)
	// z, ok := fSaddle(x, y)

	sx := float64(width)/2.0 + (x-y)*cos30*xyscale
	sy := float64(height)/2.0 + (x+y)*sin30*xyscale - z*zscale
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
