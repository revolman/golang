package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
)

var green = color.RGBA{0x51, 0xA6, 0x00, 0xff}

var r, g, b uint8
var a uint8 = 255
var palette = []color.Color{color.Black, color.RGBA{r, g, b, a}}

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(255)
		g := rand.Intn(255)
		b := rand.Intn(255)
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		fmt.Println(palette)
	}
	randColor := palette[rand.Intn(len(palette))]
	randIndex := rand.Intn(len(palette))
	fmt.Printf("Случайный цвет: %d\nСлучайный индекс: %d\n", randColor, randIndex)
}
