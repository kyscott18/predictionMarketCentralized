package simulatedplayer

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
)

// SimulatedPlayer is the type that represents a simulated player
type SimulatedPlayer struct {
	mp players.MarketPlayer
}

// NewSimulatedPlayer creates a new simualted player
func NewSimulatedPlayer(id int, balance float64, v bool) SimulatedPlayer {
	sp := SimulatedPlayer{players.NewMarketPlayer(id, balance, v)}
	return sp
}

// BuyOrSell randomly chooses to buy or sell a contract for a simulated player
func (sp *SimulatedPlayer) BuyOrSell(cs *markets.ContractSet, m *markets.Market, v bool) {
	//get the ratio of the market
	ratio := m.GetRatioFloat64()

	// bernoulli distribution with p = ratio
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	take := ratio > r1.Float64()

	//calculate an amount if you are going to take
	if take {
		amount := math.Round(r1.Float64()*10 + 1)
		if v {
			fmt.Println("User", sp.mp.ID, "has chosen to buy", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		//buy the contracts if you have funds left
		sp.mp.BuyContract(cs, m, amount, v)
	} else {
		amount := math.Round(r1.Float64()*10 + 1)
		if v {
			fmt.Println("User", sp.mp.ID, "has chosen to sell", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		//sell the contracts if you have contracts left
		sp.mp.SellContract(cs, m, amount, v)
	}

}

// AddOrRemove randomly chooses to add or remove liquidity for a simulated player
func (sp *SimulatedPlayer) AddOrRemove(cs *markets.ContractSet, m *markets.Market, v bool) {
	//probabilty of adding to pool is unrelated to the current ratio????
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	take := math.Round(r1.Float64()*2) > 1

	if take {
		amount := math.Round(r1.Float64()*10 + 1)
		if v {
			fmt.Println("User", sp.mp.ID, "has chosen to pursue", amount, "pool tokens from the event", cs.Event, "with the condition", m.Condition)
		}
		sp.mp.AddLiquidity(cs, m, amount, v)
	} else {
		amount := math.Round(r1.Float64()*10 + 1)
		if v {
			fmt.Println("User", sp.mp.ID, "has chosen to redeem", amount, "pool tokens from the event", cs.Event, "with the condition", m.Condition)
		}
		sp.mp.RemoveLiquidity(cs, m, amount, v)
	}
}

// RemoveAll removes all of a simulated players liquidity
func (sp *SimulatedPlayer) RemoveAll(cs *markets.ContractSet, v bool) {
	for i := range cs.Markets {
		numPoolTokens := sp.mp.Tokens[cs.Markets[i].Condition].Amount
		sp.mp.RemoveLiquidity(cs, &cs.Markets[i], numPoolTokens, v)
		if v {
			fmt.Println("User", sp.mp.ID, "has chosen to redeem", numPoolTokens, "pool tokens from the market with the conditions", cs.Markets[i].Condition)
		}
	}
}

// Redeem redeems all a simualted players contracts for reserve
func (sp *SimulatedPlayer) Redeem(cs *markets.ContractSet, m *markets.Market, v bool) {
	sp.mp.Redeem(cs, m, v)
}

// SimulateValidation determines the outcome of an event
func SimulateValidation(cs *markets.ContractSet) {
	for _, m := range cs.Markets {
		if m.GetRatioFloat64() > .98 {
			cs.Validate(m, false)
			return
		}
	}
}

// SumPlayersBalance totals the balance for all the simulated players
func SumPlayersBalance(bots []SimulatedPlayer) float64 {
	var sum float64 = 0
	for _, b := range bots {
		sum = sum + b.mp.Balance
	}
	return sum
}
