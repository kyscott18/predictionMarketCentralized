package main

import (
	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
)

func main() {
	cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
	//verbose statement
	cs.PrintState()
	mp1 := players.NewMarketPlayer(1, 20)
	mp1.BuyContract(cs.Event, &cs.Markets[0], 5)
	mp1.PrintState()
	cs.PrintState()
	mm := maker.NewMarketMaker()
	mm.Make(&cs)
	cs.PrintState()
}
