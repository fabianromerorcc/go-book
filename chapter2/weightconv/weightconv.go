// Package weightconv performs pound and kilogram conversions
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	PoundInKilo = 0.453592
)

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}
