package main

import (
	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
	"example.com/predictionMarketCentralized/simulatedPlayers"

)

func main() {
	cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
	cs.PrintState()
	mp1 := players.NewMarketPlayer(1, 20)
	mp1.BuyContract(cs.Event, &cs.Markets[0], 5)
	mp1.PrintState()
	cs.PrintState()
	mm := maker.NewMarketMaker()
	mm.Make(&cs)
	mm.PrintState()
	cs.PrintState()
	mp1.SellContract(cs.Event, &cs.Markets[0], 5)
	mp1.PrintState()
	cs.PrintState()
	mm.Make(&cs)
	mm.PrintState()
	cs.PrintState()

	// bots := make([]SimulatedPlayer,0)
	// for i := 1; i < 100; i++ {
	// 	bots = append(bots, NewSimulatedPlayer(i, 50))
	// }
	// for round := 0; round < 20; round++ {
	// 	for i := 1; i < 100; i++ {
	// 		for j := 0; j < len(cs.Markets); j++ {
	// 			bots[i].take(cs.Event, &cs.Markets[j])
	// 		}
	// 	}
	// }
}
