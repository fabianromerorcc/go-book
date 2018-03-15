// Package lenconv performs feet and meters conversions
package lenconv

import "fmt"

type Feet float64
type Meter float64

const (
	FeetInMeters = 0.3048
)

func (ft Feet) String() string {
	return fmt.Sprintf("%g ft", ft)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}
