// Package tempconv выполняет преобразование температур
package tempconv

import "fmt"

// Celsius - Цельсий
type Celsius float64

// Fahrenheit - Фаренгейт
type Fahrenheit float64

// Kelvin - Кельвин
type Kelvin float64

// Эти константы остались из старого примера. В новом они не используются
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Определение методов типа
func (c Celsius) String() string    { return fmt.Sprintf("%.1f°С", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }
