package maker

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
	"github.com/DzananGanic/numericalgo/root"
)

type MarketMaker struct {
	profit        float32
	intermediates []markets.Contract
}

func NewMarketMaker() MarketMaker {
	mm := MarketMaker{0, make([]markets.Contract, 0)}
	return mm
}

func (mm MarketMaker) Make(cs *markets.ContractSet) {
	var totalPrice float64 = 0
	r := make([]float64, 0)
	c := make([]float64, 0)
	for i := 0; i < len(cs.Markets); i++ {
		r = append(r, float64(cs.Markets[i].P.Usd))
		c = append(c, float64(cs.Markets[i].P.Contract.Amount))
		totalPrice = totalPrice + r[i]/c[i]
	}

	if totalPrice > 1 {
		f := func(x float64) float64 {
			var eq float64 = -1
			for i := 0; i < len(cs.Markets); i++ {
				eq = eq + ((r[i]-(r[i]*x)/(c[i]+x))/(c[i] + x))
			}
			return eq
		}
		initialGuess := 1.5
		iter := 3

		result, bull := root.Newton(f, initialGuess, iter)
		fmt.Println(result, bull)
	}
}
