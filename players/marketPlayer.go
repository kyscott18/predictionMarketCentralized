package players

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
)

// MarketPlayer is the type that represents a participant in the market
type MarketPlayer struct {
	ID        int
	Balance   float32
	contracts map[string]markets.Contract
	Tokens    map[string]markets.PoolToken
}

// BuyContract swaps reserve for the amount of contracts specified from the user perspective
func (mp *MarketPlayer) BuyContract(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {

	price := m.BuyContract(cs, &mp.Balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "bought", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough funds to buy", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// SellContract swaps the amount of contracts specified for reserve from the user perspective
func (mp *MarketPlayer) SellContract(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {

	price := m.SellContract(cs, &mp.Balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "sold", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", price)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts to sell", amount, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// BuySet purchases the amount of sets of contracts specified from the user prespective
func (mp *MarketPlayer) BuySet(cs *markets.ContractSet, amount float32, v bool) {
	price := cs.BuySet(&mp.Balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "bought", amount, "contracts sets from the event", cs.Event, "for $", price)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough funds to buy", amount, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// SellSet sells the amount of sets of contracts specified from the user perspective
func (mp *MarketPlayer) SellSet(cs *markets.ContractSet, amount float32, v bool) {
	price := cs.SellSet(&mp.Balance, &mp.contracts, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "sold", amount, "contracts sets from the event", cs.Event, "for $", price)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts to sell", amount, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// AddLiquidity adds the amount of liquidity specified to the market from the user perspective
func (mp *MarketPlayer) AddLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	price, numContacts := m.AddLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "provided", numContacts, "contracts and", price, "usd in exchange for", amount, "Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts or usd to receive", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// RemoveLiquidity removes the amount of liquidity specified from the user perspective
func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	price, numContacts := m.RemoveLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.ID, "exchanged", amount, "Pool Tokens for", numContacts, "contracts and", price, "usd from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough Pool Tokens to exchange", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}

}

// Redeem swaps contracts for reserve from the user perspective if the event outcome has been determined
func (mp *MarketPlayer) Redeem(cs *markets.ContractSet, m *markets.Market, v bool) {
	price := m.Redeem(cs, &mp.Balance, &mp.contracts)

	if v {
		if price == -1 {
			fmt.Println("event was not decided yet")
		} else {
			fmt.Println("User", mp.ID, "redeemed", price, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// PrintState prints the current state of the market player
func (mp MarketPlayer) PrintState() {
	fmt.Println("State of MarketPlayer")
	fmt.Println("User", mp.ID, "has a balance of", mp.Balance)
	fmt.Println("Contracts:")
	for _, element := range mp.contracts {
		fmt.Println("Contract condition:", element.Condition, ", amount:", element.Amount)
	}
	for _, element := range mp.Tokens {
		fmt.Println("PoolToken condition:", element.Condition, ", amount:", element.Amount)
	}
	fmt.Printf("\n")
}

// NewMarketPlayer creates a new market player
func NewMarketPlayer(id int, startingBalance float32, v bool) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make(map[string]markets.Contract), make(map[string]markets.PoolToken)}
	if v {
		fmt.Println("New MarketPlayer")
		fmt.Println("id:", id)
		fmt.Println("startingBalance:", startingBalance)
		fmt.Println()
	}

	return mp
}
