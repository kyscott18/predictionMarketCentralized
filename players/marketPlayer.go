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
	Tokens    float32
}

// BuyContract swaps reserve for the amount of contracts specified from the user perspective
func (mp *MarketPlayer) BuyContract(cs *markets.ContractSet, m *markets.Market, numContracts float32, v bool) {
	numReserve := m.BuyContract(cs, &mp.Balance, &mp.contracts, numContracts)

	//verbose statement
	if v {
		if numReserve != -1 {
			fmt.Println("User", mp.ID, "bought", numContracts, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", numReserve)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough funds to buy", numContracts, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// SellContract swaps the amount of contracts specified for reserve from the user perspective
func (mp *MarketPlayer) SellContract(cs *markets.ContractSet, m *markets.Market, numContracts float32, v bool) {

	numReserve := m.SellContract(cs, &mp.Balance, &mp.contracts, numContracts)

	//verbose statement
	if v {
		if numReserve != -1 {
			fmt.Println("User", mp.ID, "sold", numContracts, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition, "for $", numReserve)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts to sell", numContracts, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// BuySet purchases the amount of sets of contracts specified from the user prespective
func (mp *MarketPlayer) BuySet(cs *markets.ContractSet, numContracts float32, v bool) {
	numReserve := cs.BuySet(&mp.Balance, &mp.contracts, numContracts)

	//verbose statement
	if v {
		if numReserve != -1 {
			fmt.Println("User", mp.ID, "bought", numContracts, "contracts sets from the event", cs.Event, "for $", numReserve)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough funds to buy", numContracts, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// SellSet sells the amount of sets of contracts specified from the user perspective
func (mp *MarketPlayer) SellSet(cs *markets.ContractSet, numContracts float32, v bool) {
	numReserve := cs.SellSet(&mp.Balance, &mp.contracts, numContracts)

	//verbose statement
	if v {
		if numContracts != -1 {
			fmt.Println("User", mp.ID, "sold", numContracts, "contracts sets from the event", cs.Event, "for $", numReserve)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts to sell", numContracts, "contracts sets from the event", cs.Event)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// AddLiquidity adds the amount of liquidity specified to the market from the user perspective
func (mp *MarketPlayer) AddLiquidity(cs *markets.ContractSet, numPoolTokens float32, v bool) {
	tokens := cs.AddLiquidity(&mp.Balance, &mp.Tokens, numPoolTokens)

	//verbose statement
	if v {
		if tokens != -1 {
			fmt.Println("User", mp.ID, "provided", numPoolTokens, "reserve in exchange for", numPoolTokens, "Pool Tokens")
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough reserve to receive", numPoolTokens, "Pool Tokens")
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// RemoveLiquidity removes the amount of liquidity specified from the user perspective
func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, numPoolTokens float32, v bool) {
	tokens := cs.RemoveLiquidity(&mp.Balance, &mp.Tokens, numPoolTokens)

	//verbose statement
	if v {
		if tokens != -1 {
			fmt.Println("User", mp.ID, "exchanged", numPoolTokens, "Pool Tokens for", numPoolTokens, "reserve")
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough Pool Tokens to exchange", numPoolTokens, "Pool Tokens")
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
	fmt.Println("PoolTokens:", mp.Tokens)
	fmt.Printf("\n")
}

// NewMarketPlayer creates a new market player
func NewMarketPlayer(id int, startingBalance float32, v bool) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make(map[string]markets.Contract), 0}
	if v {
		fmt.Println("New MarketPlayer")
		fmt.Println("id:", id)
		fmt.Println("startingBalance:", startingBalance)
		fmt.Println()
	}

	return mp
}
