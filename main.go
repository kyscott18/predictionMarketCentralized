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

	if *typePtr == "basic" {
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200, *verbosePtr)
		mm := maker.NewMarketMaker(*verbosePtr)
		mp1 := players.NewMarketPlayer(1, 50, *verbosePtr)
		mp1.BuyContract(&cs, &cs.Markets[0], 5, *verbosePtr)
		mm.Make(&cs, *verbosePtr)
		if *verbosePtr {
			mm.PrintState()
			cs.PrintState()
		}
		mp1.AddLiquiditySS(&cs, &cs.Markets[0], 1.5, *verbosePtr)
		mp1.SellContract(&cs, &cs.Markets[0], 2, *verbosePtr)
		mm.Make(&cs, *verbosePtr)
		if *verbosePtr {
			mm.PrintState()
			cs.PrintState()
		}
		mp1.RemoveLiquiditySS(&cs, &cs.Markets[0], 1.5, *verbosePtr)
		cs.Validate(cs.Markets[0], *verbosePtr)
		mp1.Redeem(&cs, &cs.Markets[0], *verbosePtr)
		if !(*verbosePtr) {
			mp1.PrintState()
			cs.PrintState()
		}
	} else if *typePtr == "simulated" {
		//TODO: add support for controlling verbose output
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200, *verbosePtr)
		mm := maker.NewMarketMaker(*verbosePtr)
		bots := make([]simulatedPlayer.SimulatedPlayer, 0)
		for i := 0; i < 100; i++ {
			bots = append(bots, simulatedPlayer.NewSimulatedPlayer(i, 70, *verbosePtr))
		}
		for round := 0; round < 800; round++ {
			for i := range bots {
				for j := range cs.Markets {
					bots[i].BuyOrSell(&cs, &cs.Markets[j], *verbosePtr)
					mm.Make(&cs, *verbosePtr)
					bots[i].AddOrRemove(&cs, &cs.Markets[j], *verbosePtr)
				}
			}
		}
		//remove all liquidity
		for i := range bots {
			for j := range cs.Markets {
				bots[i].RemoveAll(&cs, &cs.Markets[j], *verbosePtr)
			}
		}

		//validate the outcome that is above .97 percent
		simulatedPlayer.SimulateValidation(&cs)
		//redeem all votes
		for i := range bots {
			for j := range cs.Markets {
				bots[i].Redeem(&cs, &cs.Markets[j], *verbosePtr)
			}
		}

		//print total player money
		fmt.Println("Total player money:", simulatedPlayer.SumPlayersBalance(bots))
		fmt.Println()

		mm.PrintState()
		cs.PrintState()

	}
}
