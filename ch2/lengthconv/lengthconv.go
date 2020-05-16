package lengthconv

import "fmt"

// Foot - тип единицы измерения
type Foot float64

// Meter - тип единицы измерения
type Meter float64

// метр в футах и футы в мерах
const (
	FootRatio  Foot  = 3.281
	MeterRatio Meter = 3.281
)

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToM - Foot -> Meter
func FToM(f Foot) Meter { return Meter(f / FootRatio) }

// MToF - Meter -> Foot
func MToF(m Meter) Foot { return Foot(m * MeterRatio) }
