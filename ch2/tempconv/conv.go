package tempconv

// CToF преобразует градус Цельсия в градус Фаренгейта
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK преобразует градус Цельсия в градус Кельвина
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC преобразует градус Фаренгейта в градус Цельсия
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK преобразует градус Фаренгейта в градус Кельвина
func FToK(f Fahrenheit) Kelvin { return Kelvin(273.15 + (f-32)*5/9) }

// KToC преобразует градус Кельвина в градус Цельсия
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF преобразует градус Кельвина в градус Фаренгейта
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }
