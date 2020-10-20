package simulatedPlayer

import (
	"example.com/predictionMarketCentralized/players"
	"example.com/predictionMarketCentralized/markets"
)

type SimulatedPlayer struct {
	mp players.MarketPlayer
}

func NewSimulatedPlayer(id int, balance float32) SimulatedPlayer {
	sp := SimulatedPlayer{players.NewMarketPlayer(id, balance)}
	return sp
}

func (sp *SimulatedPlayer) take(event string, m *markets.Market) {
	//get the ratio of the market
	//use a bernoulli distribution to predict if we will take or not
	//calculate an amount if you are going to take
	//buy the contracts if you have funds left
}

