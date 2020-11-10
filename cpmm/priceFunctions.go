package cpmm

// GetInputPrice translates the amount sold to the amount bought
func GetInputPrice(inputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputAmount * outputReserve
	denominator := inputReserve + inputAmount
	return numerator / denominator
}

// GetOutputPrice translates the amount bought to the amount sold
func GetOutputPrice(outputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputReserve * outputAmount
	denominator := outputReserve - outputAmount
	return numerator / denominator
}
