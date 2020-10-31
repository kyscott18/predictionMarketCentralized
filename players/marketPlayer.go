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

func (mp *MarketPlayer) BuyContract(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {

	price := m.BuyContract(cs, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "bought", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough funds to buy", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) SellContract(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {

	price := m.SellContract(cs, &mp.balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "sold", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough contracts to sell", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) BuySet(cs *markets.ContractSet, amount float32, v bool) {
	price := cs.BuySet(&mp.balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "bought", amount, "contracts sets from the event", cs.Event, "for $", price)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough funds to buy", amount, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) SellSet(cs *markets.ContractSet, amount float32, v bool) {
	price := cs.SellSet(&mp.balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "sold", amount, "contracts sets from the event", cs.Event, "for $", price)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough contracts to sell", amount, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) AddLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	price, numContacts := m.AddLiquidity(cs, &mp.balance, &mp.contracts, &mp.tokens, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "provided", numContacts, "contracts and", price, "usd in exchange for", amount, "Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough contracts or usd receive", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	price, numContacts := m.RemoveLiquidity(cs, &mp.balance, &mp.contracts, &mp.tokens, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "exchanged", amount, "Pool Tokens for", numContacts, "contracts and", price, "usd from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough Pool Tokens to exchange", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}

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

func NewMarketPlayer(id int, startingBalance float32, v bool) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make(map[string]markets.Contract), make(map[string]markets.PoolToken)}
	if v {
		fmt.Println("New MarketPlayer")
		fmt.Println("id:", id)
		fmt.Println("startingBalance:", startingBalance)
		fmt.Println("contracts: []")
		fmt.Println()
	}

	return mp
}
