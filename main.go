package main

import (
	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
)

func main() {
	cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 20)
	cs.PrintState()
	mp1 := players.NewMarketPlayer(1, 10)
	//mp1.BuySet(&cs, 5)
	mp1.PrintState()
	cs.PrintState()
	mp1.BuyContract(&cs, &cs.Markets[0], 2)
	mp1.PrintState()
	cs.PrintState()
	mm := maker.NewMarketMaker()
	mm.Make(&cs)
	mm.PrintState()
	cs.PrintState()
	// mm.Make(&cs)
	// mm.PrintState()
	// cs.PrintState()
	// mp1.SellContract(&cs, &cs.Markets[0], 5)
	// mp1.PrintState()
	// cs.PrintState()
	// mm.Make(&cs)
	// mm.PrintState()
	// cs.PrintState()

	//cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
	//mm := maker.NewMarketMaker()
	// bots := make([]simulatedPlayer.SimulatedPlayer, 0)
	// for i := 0; i < 100; i++ {
	// 	bots = append(bots, simulatedPlayer.NewSimulatedPlayer(i, 70))
	// }
	// for round := 0; round < 800; round++ {
	// 	for i := 0; i < 100; i++ {
	// 		for j := 0; j < len(cs.Markets); j++ {
	// 			bots[i].Take(&cs, &cs.Markets[j])
	// 			mm.Make(&cs)
	// 		}
	// 	}
	// }
	// mm.PrintState()

	// cs.PrintState()
}
