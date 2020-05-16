package weightconv

import "fmt"

// Pound - тип для фунтов
type Pound float64

// Kilogram - тип для килограммов
type Kilogram float64

// Константы для перевода
const (
	PoundRation   Pound    = 2.205
	KilogramRatio Kilogram = 2.205
)

func (p Pound) String() string    { return fmt.Sprintf("%.2flb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%.2fkg", k) }

// PToK - фунты в килограмы
func PToK(p Pound) Kilogram { return Kilogram(p / PoundRation) }

// KToP - килограмы в фунты
func KToP(k Kilogram) Pound { return Pound(k * KilogramRatio) }
