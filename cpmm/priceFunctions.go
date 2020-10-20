package cpmm

//From Uniswap
func GetInputPrice(inputAmount float32, inputReserve float32, outputReserve float32) float32 {
	inputAmountWithFee := inputAmount * 997
	numerator := inputAmountWithFee * outputReserve
	denominator := (inputReserve * 1000) + inputAmountWithFee
	return numerator / denominator
}

// From Uniswap
func GetOutputPrice(outputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputReserve * outputAmount * 1000
	denominator := (outputReserve - outputAmount) * 997
	return numerator/denominator + 1
}
