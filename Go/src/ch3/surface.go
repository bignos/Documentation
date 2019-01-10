// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var color string

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='background-color: black; stroke: #00FF00; fill: black; stroke- width: 0.7' "+
		"width='%d' height='%d' >",
		width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {

			ax, ay, err := corner(i+1, j)
			if err {
				continue
			}

			bx, by, err := corner(i, j)
			if err {
				continue
			}

			cx, cy, err := corner(i, j+1)
			if err {
				continue
			}

			dx, dy, err := corner(i+1, j+1)
			if err {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) { // Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y) // Compute surface height z.
	color = fmt.Sprintf("#%06x", uint32(-1*z*100000))
	if math.IsInf(z, 0) {
		return .0, .0, true
	}
	sx := width/2 + (x-y)*cos30*xyscale // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, false
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	return math.Sin(r) / 10.
}
