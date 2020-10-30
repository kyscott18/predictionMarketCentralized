package main

import (
	"flag"
	"fmt"

	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
	"example.com/predictionMarketCentralized/simulatedPlayer"
)

func main() {
	typePtr := flag.String("type", "basic", "basic or simulated")
	verbosePtr := flag.Bool("v", false, "a bool")
	flag.Parse()

	if *verbosePtr {
		fmt.Println("Verbose")
	} else {
		fmt.Println("NotVerbose")
	}
	//TODO: add support for controlling verbose output

	if *typePtr == "basic" {
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
		mm := maker.NewMarketMaker()
		mp1 := players.NewMarketPlayer(1, 50)
		mp1.BuyContract(&cs, &cs.Markets[0], 5)
		mm.Make(&cs)
		mp1.AddLiquidity(&cs, &cs.Markets[0], 1.5)
		mp1.SellContract(&cs, &cs.Markets[0], 2)
		mm.Make(&cs)
		mp1.RemoveLiquidity(&cs, &cs.Markets[0], 1.5)
		mp1.PrintState()
		cs.PrintState()
	} else if *typePtr == "simulated" {
		//TODO: add simulation for adding and removing liquidity
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
		mm := maker.NewMarketMaker()
		bots := make([]simulatedPlayer.SimulatedPlayer, 0)
		for i := 0; i < 100; i++ {
			bots = append(bots, simulatedPlayer.NewSimulatedPlayer(i, 70))
		}
		for round := 0; round < 800; round++ {
			for i := 0; i < 100; i++ {
				for j := 0; j < len(cs.Markets); j++ {
					bots[i].Take(&cs, &cs.Markets[j])
					mm.Make(&cs)
				}
			}
		}
		mm.PrintState()
		cs.PrintState()

	}
}
