package cpmm

//From Uniswap
//amount sold to amount bought
func GetInputPrice(inputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputAmount * outputReserve
	denominator := inputReserve + inputAmount
	return numerator / denominator
}

// From Uniswap
//amount bought to amount sold
func GetOutputPrice(outputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputReserve * outputAmount 
	denominator := outputReserve - outputAmount
	return numerator/denominator
}
