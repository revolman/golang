package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 120
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Print("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			_, _, z := xyz(i, j)
			points := [8]float64{ax, ay, bx, by, cx, cy, dx, dy}
			valid := true
			for _, val := range points {
				if math.IsNaN(val) {
					valid = false
				}
			}
			if !valid {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='fill: #%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(z))
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64) {
	x, y, z := xyz(i, j)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx, sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	return math.Sin(r) / r
}

// этой функцией выковыривается z, а заодно x и y
func xyz(i, j int) (x, y, z float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z.
	z = f(x, y)
	return x, y, z
}

func color(z float64) uint64 {
	redFactor := (z + 1) * 0.5
	blueFactor := (-z + 1) * 0.5

	redByte := int(255.0 * redFactor)
	blueByte := int(255.0 * blueFactor)

	// сдвиг значения в "зону красного цвета"
	var redWord uint64 = uint64(redByte) << 16
	var blueWord uint64 = uint64(blueByte)

	// Логическое сложение двух полученых слов
	return blueWord | redWord
}
