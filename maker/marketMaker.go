package maker

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
	"github.com/DzananGanic/numericalgo/root"
)

// MarketMaker is a type that represents a designated market maker
type MarketMaker struct {
	profit        float32
	intermediates map[string]markets.Contract
}

// NewMarketMaker creates a market maker
func NewMarketMaker(v bool) MarketMaker {
	mm := MarketMaker{0, make(map[string]markets.Contract)}
	if v {
		fmt.Println("New MarketMaker")
		fmt.Println("profit:", 0)
		fmt.Println("contracts: []")
		fmt.Println()
	}
	return mm
}

// PrintState prints the current state of the market maker
func (mm MarketMaker) PrintState() {
	fmt.Println("State of MarketMaker")
	fmt.Println("profit:", mm.profit)
	fmt.Println("contracts:", mm.intermediates)
	fmt.Println()
}

// Make is the market maker function of balancing the market
func (mm *MarketMaker) Make(cs *markets.ContractSet, v bool) {
	if cs.Made == true {
		return
	}

	f := func(x float64) float64 {
		var eq float64 = -1
		for _, m := range cs.Markets {
			r := float64(m.P.Usd)
			c := float64(m.P.Contract.Amount)
			eq = eq + ((r - (r*x)/(c+x)) / (c + x))
		}
		return eq
	}

	var totalPrice float64 = 0
	for _, m := range cs.Markets {
		totalPrice = totalPrice + m.GetRatioFloat64()
	}

	//TODO: find a quick algorithm for initial guess
	initialGuess := totalPrice
	iter := 7

	result, _ := root.Newton(f, initialGuess, iter)
	amount := float32(result)

	if amount > 0 {
		mm.profit = mm.profit + amount
		//buy contracts as a set
		success := cs.BuySet(&mm.profit, &mm.intermediates, amount)
		if v {
			if success != -1 {
				fmt.Println("MarketMaker bought", amount, "contracts sets from the event", cs.Event, "for $", amount)
			} else {
				fmt.Println("MarketMaker doesn't have enough funds to buy", amount, "contracts sets from the event", cs.Event)
				return
			}
		}

		//sell contracts to individual markets
		for i := range cs.Markets {
			price := cs.Markets[i].SellContract(cs, &mm.profit, &mm.intermediates, amount)
			if v {
				if price != -1 {
					fmt.Println("MarketMaker sold", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", price)
				} else {
					fmt.Println("Market Maker doesn't have enough contracts to sell", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
					return
				}
			}
		}

	} else if amount < 0 {
		amount = amount * -1
		mm.profit = mm.profit + amount
		//buy contracts from indiviual markets
		for i := range cs.Markets {
			// fmt.Println(cs.Markets[i].Condition)
			// fmt.Println(mm.profit)
			price := cs.Markets[i].BuyContract(cs, &mm.profit, &mm.intermediates, amount)
			if v {
				if price != -1 {
					fmt.Println("MarketMaker bought", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", price)
				} else {
					fmt.Println("MarketMaker doesn't have enough funds to buy", amount, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
					return
				}
			}
		}

		//Sell contracts as a set
		success := cs.SellSet(&mm.profit, &mm.intermediates, amount)
		//verbose statement
		if v {
			if success != -1 {
				fmt.Println("MarketMaker sold", amount, "contracts sets from the event", cs.Event, "for $", amount)
			} else {
				fmt.Println("MarketMaker doesn't have enough contracts to sell", amount, "contracts sets from the event", cs.Event)
				return
			}
		}
	}
	if v {
		fmt.Println()
	}
	mm.profit = mm.profit - amount
	cs.Made = true
}
