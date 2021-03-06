package main

import (
	"flag"
	"fmt"

	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
	simulatedplayer "example.com/predictionMarketCentralized/simulatedPlayer"
)

func main() {
	typePtr := flag.String("type", "basic", "basic or simulated")
	verbosePtr := flag.Bool("v", false, "a bool")
	flag.Parse()

	if *typePtr == "basic" {
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float64{.5, .5}, 20, *verbosePtr)
		mp1 := players.NewMarketPlayer(1, 10, *verbosePtr)
		if *verbosePtr {
			mp1.PrintState()
			cs.PrintState()
		}
		mp1.BuyContract(&cs, &cs.Markets[0], 2, *verbosePtr)
		mp1.SellContract(&cs, &cs.Markets[0], 2, *verbosePtr)

		mp1.BuySet(&cs, 2, *verbosePtr)
		mp1.SellSet(&cs, 2, *verbosePtr)

		mp1.AddLiquidity(&cs, &cs.Markets[0], 2, *verbosePtr)
		mp1.RemoveLiquidity(&cs, &cs.Markets[0], 2, *verbosePtr)

		cs.Validate(cs.Markets[0], *verbosePtr)
		mp1.Redeem(&cs, &cs.Markets[0], *verbosePtr)

		if !(*verbosePtr) {
			mp1.PrintState()
			cs.PrintState()
		}
	} else if *typePtr == "simulated" {
		cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float64{.5, .5}, 200, false)
		bots := make([]simulatedplayer.SimulatedPlayer, 0)
		if *verbosePtr {
			fmt.Println("Creating 100 simulated players")
			fmt.Println()
		}
		for i := 0; i < 100; i++ {
			bots = append(bots, simulatedplayer.NewSimulatedPlayer(i, 70, false))
		}
		if *verbosePtr {
			fmt.Println("Simulating trading")
			fmt.Println()
		}
		for round := 0; round < 800; round++ {
			for i := range bots {
				for j := range cs.Markets {
					bots[i].BuyOrSell(&cs, &cs.Markets[j], false)
					bots[i].AddOrRemove(&cs, &cs.Markets[j], false)
				}
			}
		}
		if *verbosePtr {
			fmt.Println("Ordering removal of all liquidity")
			fmt.Println()
		}
		//remove all liquidity
		for i := range bots {
			bots[i].RemoveAll(&cs, false)
		}
		if *verbosePtr {
			fmt.Println("Determining outcome and redeeming all contracts")
			fmt.Println()
		}

		//validate the outcome that is above .97 percent
		simulatedplayer.SimulateValidation(&cs)
		//redeem all votes
		for i := range bots {
			for j := range cs.Markets {
				bots[i].Redeem(&cs, &cs.Markets[j], false)
			}
		}

		//print total player money
		fmt.Println("Total player money:", simulatedplayer.SumPlayersBalance(bots))
		fmt.Println()

		//print total contracts minted
		fmt.Println("Total contracts minted:", cs.Mints)
		fmt.Println()

		cs.PrintState()

	}
}
