// Упражнение 1.6.
// Версия lissajous.go в которой сначала случайно добавляются цвета в палитру
// а зетем случайным образом вибирается цвет линий.
// Фон можно менять, если убрать color.Black из среза palette.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.Black} // определение среза
var a uint8 = 255

const (
	blackIndex = 0 // Перый цвет палитры, будет фоном
)

func lissajous(out io.Writer, cycles int) {
	const (
		//		cycles  = 5     // Количество полных колебаний
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano()) // настройка рандомизатора, иначе будет выдавать одинаковые значения

	for i := 0; i < 10; i++ { // генерация цветов и добавление в палитру
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
	}

	randIndex := uint8(rand.Intn(len(palette)))
	freq := rand.Float64() * 3.0 // Относительная частота колебаний Y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				randIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
