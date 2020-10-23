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
	if cs.Made == true {
		return
	}
	var totalPrice float64 = 0
	r := make([]float64, 0)
	c := make([]float64, 0)
	for i := 0; i < len(cs.Markets); i++ {
		r = append(r, float64(cs.Markets[i].P.Usd))
		c = append(c, float64(cs.Markets[i].P.Contract.Amount))
		totalPrice = totalPrice + cs.Markets[i].GetRatioFloat64()
	}

	if totalPrice > 1 {
		f := func(x float64) float64 {
			var eq float64 = -1
			for i := 0; i < len(cs.Markets); i++ {
				eq = eq + ((r[i] - (r[i]*x)/(c[i]+x)) / (c[i] + x))
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
		success := cs.BuySet(&mm.profit, &mm.intermediates, amount)
		mm.profit = mm.profit - amount
		if success != -1 {
			fmt.Println("MarketMaker bought", amount, "contracts sets from the event", cs.Event, "for $", amount)
		} else {
			fmt.Println("MarketMaker doesn't have enough funds to buy", amount, "contracts sets from the event", cs.Event)
			return
		}

		//sell contracts to individual markets
		for i := 0; i < len(cs.Markets); i++ {
			price := cs.Markets[i].SellContract(cs, &mm.profit, &mm.intermediates, amount)
			if price != -1 {
				fmt.Println("MarketMaker sold", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", price)
			} else {
				fmt.Println("Market Maker doesn't have enough contracts to sell", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
				return
			}
		}

		fmt.Printf("\n")

		cs.Made = true

	} else if totalPrice < 1 {
		f := func(x float64) float64 {
			var eq float64 = -1
			for i := 0; i < len(cs.Markets); i++ {
				eq = eq + ((r[i] + (r[i]*x)/(c[i]+x)) / (c[i] - x))
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
			price := cs.Markets[i].BuyContract(cs, &mm.profit, &mm.intermediates, amount)
			if price != -1 {
				fmt.Println("MarketMaker bought", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", price)
			} else {
				fmt.Println("MarketMaker doesn't have enough funds to buy", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
				return
			}
		}

		//Sell contracts as a set
		success := cs.SellSet(&mm.profit, &mm.intermediates, amount)
		//verbose statement
		if success != -1 {
			fmt.Println("MarketMaker sold", amount, "contracts sets from the event", cs.Event, "for $", amount)
		} else {
			fmt.Println("MarketMaker doesn't have enough contracts to sell", amount, "contracts sets from the event", cs.Event)
			return
		}

		fmt.Printf("\n")
		mm.profit = mm.profit - amount
		cs.Made = true
	}
}
