package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
)

var palette color.Color

func main() {
	// const (
	// 	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	// 	width, height          = 1024, 1024
	// )

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png") // уточняю хедер
		xKey := r.URL.Query().Get("x")              // получаю параметры из url (?x=2&y=2&resolution=1024)
		yKey := r.URL.Query().Get("y")
		resolutionKey := r.URL.Query().Get("resolution")

		x, _ := strconv.Atoi(xKey)
		y, _ := strconv.Atoi(yKey)
		resolution, _ := strconv.Atoi(resolutionKey)
		xmin, ymin, xmax, ymax := -x, -y, x, y
		width, height := resolution, resolution

		red := rand.Intn(255)
		green := rand.Intn(255)
		blue := rand.Intn(255)

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/float64(height)*float64((ymax-ymin)) + float64(ymin)
			for px := 0; px < width; px++ {
				x := float64(px)/float64(width)*float64((xmax-xmin)) + float64(xmin)
				z := complex(x, y)
				// Точка (px, py) представляет комплексное число z.
				img.Set(px, py, mandelbrot(z, red, green, blue))
			}
		}
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128, r int, g int, b int) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			palette = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(contrast * n)}
			return palette
		}
	}
	//	palette = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	return color.Black
}
