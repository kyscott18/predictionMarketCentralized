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
	TokensSS  map[string]markets.PoolTokenSS
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
func (mp *MarketPlayer) AddLiquidity(cs *markets.ContractSet, m *markets.Market, numPoolTokens float32, v bool) {
	numReserve, numContracts := m.AddLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, numPoolTokens)

	//verbose statement
	if v {
		if numReserve != -1 {
			fmt.Println("User", mp.ID, "provided", numContracts, "contracts and", numReserve, "reserve in exchange for", numPoolTokens, "Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts or reserve to receive", numPoolTokens, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// RemoveLiquidity removes the amount of liquidity specified from the user perspective
func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, m *markets.Market, numPoolTokens float32, v bool) {
	numReserve, numContacts := m.RemoveLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, numPoolTokens)

	//verbose statement
	if v {
		if numReserve != -1 {
			fmt.Println("User", mp.ID, "exchanged", numPoolTokens, "Pool Tokens for", numContacts, "contracts and", numReserve, "reserve from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough Pool Tokens to exchange", numPoolTokens, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}

}

// AddLiquiditySS adds the amount of contracts specified and pairs them with reserve from backing to provide liquidity
func (mp *MarketPlayer) AddLiquiditySS(cs *markets.ContractSet, m *markets.Market, numContracts float32, v bool) {
	numPoolTokens := m.AddLiquiditySS(cs, &mp.contracts, &mp.TokensSS, numContracts)

	//verbose statement
	if v {
		if numContracts != -1 {
			fmt.Println("User", mp.ID, "provided", numContracts, "contracts in exchange for", numPoolTokens, "single sided Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough contracts to provide", numContracts, " for single sided Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

// RemoveLiquiditySS removes the amount of contracts specified and an equal amount of reserve
func (mp *MarketPlayer) RemoveLiquiditySS(cs *markets.ContractSet, m *markets.Market, numContracts float32, v bool) {
	numPoolTokens := m.RemoveLiquiditySS(cs, &mp.contracts, &mp.TokensSS, numContracts)

	//verbose statement
	if v {
		if numContracts != -1 {
			fmt.Println("User", mp.ID, "exchanged", numPoolTokens, "single sided Pool Tokens for", numContracts, "contracts from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.ID, "doesn't have enough single sided Pool Tokens to receive", numContracts, "contracts from the market with the condition", m.Condition)
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
	for _, element := range mp.TokensSS {
		fmt.Println("PoolTokenSS condition:", element.Condition, ", contracts:", element.OriginalNumContracts)
	}
	fmt.Printf("\n")
}

// NewMarketPlayer creates a new market player
func NewMarketPlayer(id int, startingBalance float32, v bool) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make(map[string]markets.Contract), make(map[string]markets.PoolToken), make(map[string]markets.PoolTokenSS)}
	if v {
		fmt.Println("New MarketPlayer")
		fmt.Println("id:", id)
		fmt.Println("startingBalance:", startingBalance)
		fmt.Println()
	}

	return mp
}
