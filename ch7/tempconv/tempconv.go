package main

import (
	"flag"
	"fmt"
)

// Celsius ...
type Celsius float64

// Fahrenheit - Фаренгейт
type Fahrenheit float64

// Kelvin - Кельвин
type Kelvin float64

// Для удовлетворения интерфейса нужны два метода: String() string и Set(string) error.
// Тип Celsius уже имеет метод String() string, а этого уже достаточно.
// Остаётся написать Set(string) error
type celsiusFlag struct{ Celsius }

func (c Celsius) String() string    { return fmt.Sprintf("%1.f°С", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }

// Set получит строковое значение при установке флага
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	// разбирает строку s на части, значения которых присваивает переменным 36,6°С
	// value = 36,6		unit = °С
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°С":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("неверная температура %q", s)
}

// CelsiusFlag определяет флаг Celsius с указанным именем, значением по умолчанию
// и строкой-инструкцией по примененю и возвращает адрес переменной-флага.
// Аргумент должен содержать числовое значение и единицу измерения, например "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value} // value будет обрабатываться методом Set(string)
	// Var присваивает аргумент *celsiusFlag параметру flag.Value, заставляя компилятор
	// выполнить проверку того, что тип *celsiusFlag имеет необходимые методы.
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
