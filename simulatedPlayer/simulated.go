package simulatedPlayer

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
)

type SimulatedPlayer struct {
	mp players.MarketPlayer
}

func NewSimulatedPlayer(id int, balance float32, v bool) SimulatedPlayer {
	sp := SimulatedPlayer{players.NewMarketPlayer(id, balance, v)}
	return sp
}

func (sp *SimulatedPlayer) BuyOrSell(cs *markets.ContractSet, m *markets.Market, v bool) {
	//get the ratio of the market
	ratio := m.GetRatioFloat32()

	// bernoulli distribution with p = ratio
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	take := ratio > r1.Float32()

	//calculate an amount if you are going to take
	if take {
		amount := math.Round(r1.Float64()*10 + 1)
		fmt.Println("User", sp.mp.Id, "has chosen to buy", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		//buy the contracts if you have funds left
		sp.mp.BuyContract(cs, m, float32(amount), v)
	} else {
		amount := math.Round(r1.Float64()*10 + 1)
		fmt.Println("User", sp.mp.Id, "has chosen to sell", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		//sell the contracts if you have contracts left
		sp.mp.SellContract(cs, m, float32(amount), v)
	}

}

func (sp *SimulatedPlayer) AddOrRemove(cs *markets.ContractSet, m *markets.Market, v bool) {
	//probabilty of adding to pool is unrelated to the current ratio????
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	take := math.Round(r1.Float64()*2) > 1

	if take {
		amount := math.Round(r1.Float64()*10 + 1)
		fmt.Println("User", sp.mp.Id, "has chosen to pursue", amount, "pool tokens from the event", cs.Event, "with the condition", m.Condition)
		sp.mp.AddLiquidity(cs, m, float32(amount), v)
	} else {
		amount := math.Round(r1.Float64()*10 + 1)
		fmt.Println("User", sp.mp.Id, "has chosen to redeem", amount, "pool tokens from the event", cs.Event, "with the condition", m.Condition)
		sp.mp.RemoveLiquidity(cs, m, float32(amount), v)
	}
}

func (sp *SimulatedPlayer) RemoveAll(cs *markets.ContractSet, m *markets.Market, v bool) {
	amount := sp.mp.Tokens[m.Condition].Amount
	fmt.Println("User", sp.mp.Id, "has chosen to redeem", amount, "pool tokens from the event", cs.Event, "with the condition", m.Condition)
	sp.mp.RemoveLiquidity(cs, m, float32(amount), v)
}
