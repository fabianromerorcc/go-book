package lenconv

// FtToM converts Feet length to Meter
func FtToM(ft Feet) Meter {
	return Meter(ft * FeetInMeters)
}

func MToFt(m Meter) Feet {
	return Feet(m / FeetInMeters)
}
