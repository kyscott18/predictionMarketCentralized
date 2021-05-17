package cpmm

// GetAmountOut translates the amount sold to the amount bought
func GetAmountOut(x float64, c float64, r float64) float64 {
	numerator := r * x * (1 - .003)
	denominator := c + (x * (1 - .003))
	return numerator / denominator
}

// GetAmountIn translates the amount bought to the amount sold
func GetAmountIn(x float64, r float64, c float64) float64 {
	numerator := r * x * (1 - .003)
	denominator := c - x
	return numerator / denominator
}
