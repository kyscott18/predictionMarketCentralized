package players

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
)

type MarketPlayer struct {
	Id        int
	balance   float32
	contracts map[string]markets.Contract
	tokens    map[string]markets.PoolToken
}

func (mp *MarketPlayer) BuyContract(cs *markets.ContractSet, m *markets.Market, amount float32) {

	price := m.BuyContract(cs, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "bought", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough funds to buy", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
	}
	fmt.Printf("\n")
}

func (mp *MarketPlayer) SellContract(cs *markets.ContractSet, m *markets.Market, amount float32) {

	price := m.SellContract(cs, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if price != -1 {
		fmt.Println("User", mp.Id, "sold", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
	} else {
		fmt.Println("User", mp.Id, "doesn't have enough contracts to sell", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
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

func (mp *MarketPlayer) AddLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32) {
	m.AddLiquidity(cs, &mp.balance, &mp.contracts, &mp.tokens, amount)
}

func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32) {
	m.RemoveLiquidity(cs, &mp.balance, &mp.contracts, &mp.tokens, amount)
}

func (mp MarketPlayer) PrintState() {
	fmt.Println("State of MarketPlayer")
	fmt.Println("User", mp.Id, "has a balance of", mp.balance)
	fmt.Println("Contracts:")
	for _, element := range mp.contracts {
		fmt.Println("Contract condition:", element.Condition, ", amount:", element.Amount)
	}
	for _, element := range mp.tokens {
		fmt.Println("PoolToken condition:", element.Condition, ", amount:", element.Amount)
	}
	fmt.Printf("\n")
}

func NewMarketPlayer(id int, startingBalance float32) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make(map[string]markets.Contract), make(map[string]markets.PoolToken)}
	fmt.Println("New MarketPlayer")
	fmt.Println("id:", id)
	fmt.Println("startingBalance:", startingBalance)
	fmt.Println("contracts: []")
	fmt.Println()

	return mp
}
