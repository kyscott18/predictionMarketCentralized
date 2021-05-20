package cpmm

// GetAmountOut given an x amount of contracts to sell, how much reserve will you receive
func GetAmountOut(x float64, r float64, c float64) float64 {
	numerator := r * x * (1 - .003)
	denominator := c + (x * (1 - .003))
	return numerator / denominator
}

// GetAmountIn given an x amount of contracts to buy, how much reserve is required
func GetAmountIn(x float64, r float64, c float64) float64 {
	numerator := r * x
	denominator := (c - x) * (1 - .003)
	return numerator / denominator
}
