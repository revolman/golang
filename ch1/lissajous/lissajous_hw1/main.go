// Упражнение 1.5.
// Версия lissajous.go в которой измененены цвета с использованием инструкции color.RGBA{}.
// Зелёный на чёрном фоне.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x51, 0xA6, 0x00, 0xff}}
var a uint8 = 255

const (
	blackIndex = 0 // Перый цвет палитры, будет фоном
	greenIndex = 1 // Второй цвет палитры, будет линиями
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano()) // настройка рандомизатора, иначе будет выдавать одинаковые значения
	freq := rand.Float64() * 3.0           // Относительная частота колебаний Y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
