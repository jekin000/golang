//Package unitconv performs:
//Feets and Metres 1 ft = 0.3048m
//Pounds and Kilograms 1lb = 0.45359237 kg  1kg = 2.2046226lb
package unitconv

import "fmt"

type Feets float64
type Metres float64
type Pounds float64
type Kilograms float64

//Type API
func (f Feets) String() string     { return fmt.Sprintf("%g ft", f) }
func (m Metres) String() string    { return fmt.Sprintf("%g m", m) }
func (p Pounds) String() string    { return fmt.Sprintf("%g lb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g kg", k) }
