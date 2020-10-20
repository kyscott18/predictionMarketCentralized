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
	fmt.Println("New MarketMaker")
	fmt.Println("profit:", 0)
	fmt.Println("contracts: []")
	fmt.Println()
	return mm
}

func (mm MarketMaker) PrintState() {
	fmt.Println("State of MarketMaker")
	fmt.Println("profit:", mm.profit)
	fmt.Println("contracts:", mm.intermediates)
	fmt.Println()
}

func (mm *MarketMaker) Make(cs *markets.ContractSet) {
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
		//TODO: find a quick algorithm for initial guess
		initialGuess := 3.0
		iter := 7

		result, _ := root.Newton(f, initialGuess, iter)
		amount := float32(result)
		mm.profit = mm.profit + amount
		
		//buy contracts as a set
		cs.BuySet(&mm.profit, &mm.intermediates, amount)
		mm.profit = mm.profit - amount

		//sell contracts to individual markets
		for i := 0; i < len(cs.Markets); i++ {
			fmt.Println(cs.Markets[i].SellContract(cs.Event, &mm.profit, &mm.intermediates, amount), "sell")
		}

	} else if totalPrice < 1 {
		f := func(x float64) float64 {
			var eq float64 = -1
			for i := 0; i < len(cs.Markets); i++ {
				eq = eq + ((r[i]+(r[i]*x)/(c[i]+x))/(c[i] - x))
			}
			return eq
		}
		//TODO: find a quick algorithm for initial guess
		initialGuess := 3.0
		iter := 7

		result, _ := root.Newton(f, initialGuess, iter)
		amount := float32(result)
		mm.profit = mm.profit + amount

		//buy contracts from indiviual markets
		for i := 0; i < len(cs.Markets); i++ {
			fmt.Println(cs.Markets[i].BuyContract(cs.Event, &mm.profit, &mm.intermediates, amount), "buy")
		}

		//Sell contracts as a set
		cs.SellSet(&mm.profit, &mm.intermediates, amount)
		mm.profit = mm.profit - amount
	}
}
