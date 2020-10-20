package players

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
)

type MarketPlayer struct {
	Id        int
	balance   float32
	contracts []markets.Contract
}

func (mp *MarketPlayer) BuyContract(event string, m *markets.Market, amount float32) {

	price := m.BuyContract(event, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "bought", amount, "contracts from the event", event, "with the condition", m.P.Contract.Condition, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough funds to buy", amount, "contracts from the event", event, "with the condition", m.P.Contract.Condition)
	}
	fmt.Printf("\n")
}

func (mp *MarketPlayer) SellContract(event string, m *markets.Market, amount float32) {

	price := m.SellContract(event, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "sold", amount, "contracts from the event", event, "with the condition", m.P.Contract.Condition, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough contracts to sell", amount, "contracts from the event", event, "with the condition", m.P.Contract.Condition)
	}
	fmt.Printf("\n")
}

func (mp *MarketPlayer) BuySet(cs *markets.ContractSet, amount float32) {
	price := cs.BuySet(&mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "bought", amount, "contracts sets from the event", cs.Event, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough funds to buy", amount, "contracts sets from the event", cs.Event)
	}
	fmt.Printf("\n")
}

func (mp *MarketPlayer) SellSet(cs *markets.ContractSet, amount float32) {
	price := cs.SellSet(&mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "sold", amount, "contracts sets from the event", cs.Event, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough contracts to sell", amount, "contracts sets from the event", cs.Event)
	}
	fmt.Printf("\n")
}

func (mp MarketPlayer) PrintState() {
	fmt.Println("State of MarketPlayer")
	fmt.Println("User", mp.Id, "has a balance of", mp.balance)
	fmt.Println("Contracts:")
	for i := 0; i < len(mp.contracts); i++ {
		fmt.Println("Condition:", mp.contracts[i].Condition, ", amount:", mp.contracts[i].Amount)
	}
	fmt.Printf("\n")
}

func NewMarketPlayer(id int, startingBalance float32) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make([]markets.Contract, 0)}
	fmt.Println("New MarketPlayer")
	fmt.Println("id:", id)
	fmt.Println("startingBalance:", startingBalance)
	fmt.Println("contracts: []")
	fmt.Println()

	return mp
}
