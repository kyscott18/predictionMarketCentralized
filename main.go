package main

import (
	"flag"

	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
	"example.com/predictionMarketCentralized/simulatedPlayer"
)

func main() {
	typePtr := flag.String("type", "basic", "basic or simulated")
	verbosePtr := flag.Bool("v", false, "a bool")
	flag.Parse()

	if *typePtr == "basic" {
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200, *verbosePtr)
		mm := maker.NewMarketMaker(*verbosePtr)
		mp1 := players.NewMarketPlayer(1, 50, *verbosePtr)
		mp1.BuyContract(&cs, &cs.Markets[0], 5, *verbosePtr)
		mm.Make(&cs, *verbosePtr)
		if *verbosePtr {
			mm.PrintState()
		}
		mp1.AddLiquidity(&cs, &cs.Markets[0], 1.5, *verbosePtr)
		mp1.SellContract(&cs, &cs.Markets[0], 2, *verbosePtr)
		mm.Make(&cs, *verbosePtr)
		if *verbosePtr {
			mm.PrintState()
		}
		mp1.RemoveLiquidity(&cs, &cs.Markets[0], 1.5, *verbosePtr)
	} else if *typePtr == "simulated" {
		//TODO: add simulation for adding and removing liquidity
		//TODO: add support for controlling verbose output
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200, *verbosePtr)
		mm := maker.NewMarketMaker(*verbosePtr)
		bots := make([]simulatedPlayer.SimulatedPlayer, 0)
		for i := 0; i < 100; i++ {
			bots = append(bots, simulatedPlayer.NewSimulatedPlayer(i, 70, *verbosePtr))
		}
		for round := 0; round < 800; round++ {
			for i := 0; i < 100; i++ {
				for j := 0; j < len(cs.Markets); j++ {
					bots[i].Take(&cs, &cs.Markets[j], *verbosePtr)
					mm.Make(&cs, *verbosePtr)
				}
			}
		}
		mm.PrintState()
		cs.PrintState()

	}
}
