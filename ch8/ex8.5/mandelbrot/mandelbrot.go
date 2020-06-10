// алгоритм взял тут: https://github.com/renatofq/gopl/blob/master/e8_5/newton.go
// значение proc оказывает влияние, только при высоких значениях

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

// вычисленные координаты
type operand struct {
	px, py int
	z      complex128
}

// точка и её цвет
type result struct {
	px, py int
	c      color.Color
}

func main() {
	start := time.Now()
	proc := runtime.NumCPU()
	var wg sync.WaitGroup

	in := make(chan operand, proc)
	out := make(chan result, proc)
	// defer close(out)

	// вычисление и отправка значений
	go func() {
		defer close(in)

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {

				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				in <- operand{px, py, z}
			}
		}
	}()

	// получение и многопоточная обработка значений
	for i := 0; i < proc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range in {
				out <- result{
					px: item.px,
					py: item.py,
					c:  newton(item.z),
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// отрисовка результата
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for res := range out {
		img.Set(res.px, res.py, res.c)
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors

	log.Println(time.Since(start))

}

// func mandelbrot(z complex128) color.Color {
// 	const iterations = 200
// 	const contrast = 15

// 	var v complex128
// 	for n := uint8(0); n < iterations; n++ {
// 		v = v*v + z
// 		if cmplx.Abs(v) > 2 {
// 			return color.Gray{255 - contrast*n}
// 		}
// 	}
// 	return color.Black
// }

func newton(z complex128) color.Color {
	const iterations = 100
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
