package weightconv

// PToKg converts Pound weight to Kilogram
func PToKg(p Pound) Kilogram {
	return Kilogram(p * PoundInKilo)
}

// KgToP converts Kilogram weight to Pound
func KgToP(kg Kilogram) Pound {
	return Pound(kg / PoundInKilo)
}
