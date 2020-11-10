package players

import (
	"fmt"

	"example.com/predictionMarketCentralized/markets"
)

type MarketPlayer struct {
	Id        int
	Balance   float32
	contracts map[string]markets.Contract
	Tokens    map[string]markets.PoolToken
	TokensSS  map[string]markets.PoolTokenSS
}

func (mp *MarketPlayer) BuyContract(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {

	price := m.BuyContract(cs, &mp.Balance, &mp.contracts, amount)

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

	price := m.SellContract(cs, &mp.Balance, &mp.contracts, amount)

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
	price := cs.BuySet(&mp.Balance, &mp.contracts, amount)

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
	price := cs.SellSet(&mp.Balance, &mp.contracts, amount)

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
	price, numContacts := m.AddLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, amount)

	//verbose statement
	if v {
		if price != -1 {
			fmt.Println("User", mp.Id, "provided", numContacts, "contracts and", price, "usd in exchange for", amount, "Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough contracts or usd to receive", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) RemoveLiquidity(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	price, numContacts := m.RemoveLiquidity(cs, &mp.Balance, &mp.contracts, &mp.Tokens, amount)

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

func (mp *MarketPlayer) AddLiquiditySS(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	numContracts := m.AddLiquiditySS(cs, &mp.contracts, &mp.TokensSS, amount)

	//verbose statement
	if v {
		if numContracts != -1 {
			fmt.Println("User", mp.Id, "provided", numContracts, "contracts in exchange for", amount, "single sided Pool Tokens from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough contracts to receive", amount, "single sided Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) RemoveLiquiditySS(cs *markets.ContractSet, m *markets.Market, amount float32, v bool) {
	numContracts := m.RemoveLiquiditySS(cs, &mp.contracts, &mp.TokensSS, amount)

	//verbose statement
	if v {
		if numContracts != -1 {
			fmt.Println("User", mp.Id, "exchanged", amount, "single sided Pool Tokens for", numContracts, "contracts from the market with the condition", m.Condition)
		} else {
			fmt.Println("User", mp.Id, "doesn't have enough single sided Pool Tokens to exchange", amount, "Pool Tokens from the market with the condition", m.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp *MarketPlayer) Redeem(cs *markets.ContractSet, m *markets.Market, v bool) {
	price := m.Redeem(cs, &mp.Balance, &mp.contracts)

	if v {
		if price == -1 {
			fmt.Println("event was not decided yet")
		} else {
			fmt.Println("User", mp.Id, "redeemed", price, "contracts from the event", cs.Event, "with the condition", m.P.Contract.Condition)
		}
		fmt.Printf("\n")
		mp.PrintState()
		cs.PrintState()
	}
}

func (mp MarketPlayer) PrintState() {
	fmt.Println("State of MarketPlayer")
	fmt.Println("User", mp.Id, "has a balance of", mp.Balance)
	fmt.Println("Contracts:")
	for _, element := range mp.contracts {
		fmt.Println("Contract condition:", element.Condition, ", amount:", element.Amount)
	}
	for _, element := range mp.Tokens {
		fmt.Println("PoolToken condition:", element.Condition, ", amount:", element.Amount)
	}
	for _, element := range mp.TokensSS {
		fmt.Println("PoolTokenSS condition:", element.Condition, ", amount:", element.Amount)
	}
	fmt.Printf("\n")
}

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
