package simulatedPlayer

import (
	"example.com/predictionMarketCentralized/players"
	"example.com/predictionMarketCentralized/markets"
	"math/rand"
	"math"
	"time"
	"fmt"
)

type SimulatedPlayer struct {
	mp players.MarketPlayer
}

func NewSimulatedPlayer(id int, balance float32) SimulatedPlayer {
	sp := SimulatedPlayer{players.NewMarketPlayer(id, balance)}
	return sp
}

func (sp *SimulatedPlayer) Take(event string, m *markets.Market) {
	//get the ratio of the market
	ratio := m.P.Usd/m.P.Contract.Amount

	//use a bernoulli distribution to predict if we will take or not
	// bernoulli distribution with p = ratio
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	take := ratio > r1.Float32()

	//calculate an amount if you are going to take
	if take {
		amount := math.Round(r1.Float64()*10+1)
		fmt.Println("User", sp.mp.Id, "has chosen to buy", amount, "contracts from the event", event, "with the condition", m.P.Contract.Condition)
		//buy the contracts if you have funds left
		sp.mp.BuyContract(event, m, float32(amount))
	} else {
		fmt.Println("User", sp.mp.Id, "has chosen to buy 0 contracts from the event", event, "with the condition", m.P.Contract.Condition)
	}

}

