package cpmm

// GetInputPrice translates the amount sold to the amount bought
func GetInputPrice(inputAmount float64, inputReserve float64, outputReserve float64) float64 {
	numerator := inputAmount * outputReserve
	denominator := inputReserve + inputAmount
	return numerator / denominator
}

// GetOutputPrice translates the amount bought to the amount sold
func GetOutputPrice(outputAmount float64, inputReserve float64, outputReserve float64) float64 {
	numerator := inputReserve * outputAmount
	denominator := outputReserve - outputAmount
	return numerator / denominator
}
