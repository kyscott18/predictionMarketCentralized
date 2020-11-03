package markets

import (
	"fmt"

	"example.com/predictionMarketCentralized/cpmm"
)

func (m *Market) BuyContract(cs *ContractSet, balance *float32, contracts *map[string]Contract, amount float32) float32 {
	//input = usd, output = contracts
	price := cpmm.GetOutputPrice(amount, m.P.Usd, m.P.Contract.Amount)
	//check enough usd to buy
	if price > *balance {
		return -1
	}

	//remove usd from user
	*balance = *balance - price

	//add usd to pool
	m.P.Usd = m.P.Usd + price

	//remove contracts from pool
	m.P.Contract.Amount = m.P.Contract.Amount - amount

	//add contracts to user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + amount}

	cs.Made = false

	return price
}

func (m *Market) SellContract(cs *ContractSet, balance *float32, contracts *map[string]Contract, amount float32) float32 {
	//input = contract, output = usd
	price := cpmm.GetInputPrice(amount, m.P.Contract.Amount, m.P.Usd)

	//check enough contracts to sell
	_, ok := (*contracts)[m.Condition]
	if !ok {
		return -1
	} else if (*contracts)[m.Condition].Amount < amount {
		return -1
	}

	//remove contracts from user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - amount}

	//add contracts to pool
	m.P.Contract.Amount = m.P.Contract.Amount + amount

	//remove usd from pool
	m.P.Usd = m.P.Usd - price

	//add usd to user
	*balance = *balance + price

	cs.Made = false

	return price
}

func (m *Market) AddLiquidity(cs *ContractSet, balance *float32, contracts *map[string]Contract, tokens *map[string]PoolToken, amount float32) (float32, float32) {
	price := amount * m.P.Usd / m.P.NumPoolTokens
	numContracts := amount * m.P.Contract.Amount / m.P.NumPoolTokens
	//check enough balance and contracts
	_, ok := (*contracts)[m.Condition]
	if !ok || *balance < price {
		return -1, -1
	} else if (*contracts)[m.Condition].Amount < amount {
		return -1, -1
	}

	//remove balance and contracts from user
	*balance = *balance - price
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - numContracts}

	// add balance and user to pool
	m.P.Usd = m.P.Usd + price
	m.P.Contract.Amount = m.P.Contract.Amount + numContracts

	//mint new poolTokens and add to user
	m.P.NumPoolTokens = m.P.NumPoolTokens + amount
	(*tokens)[m.Condition] = PoolToken{m.Condition, (*tokens)[m.Condition].Amount + amount}

	return price, numContracts
}

func (m *Market) RemoveLiquidity(cs *ContractSet, balance *float32, contracts *map[string]Contract, tokens *map[string]PoolToken, amount float32) (float32, float32) {
	price := amount * m.P.Usd / m.P.NumPoolTokens
	numContracts := amount * m.P.Contract.Amount / m.P.NumPoolTokens
	//check enough pool tokens
	_, ok := (*tokens)[m.Condition]
	if !ok {
		return -1, -1
	} else if (*tokens)[m.Condition].Amount < amount {
		return -1, -1
	}

	//remove balance and contacts from pool
	m.P.Usd = m.P.Usd - price
	m.P.Contract.Amount = m.P.Contract.Amount - numContracts

	//add balance and contracts to user
	*balance = *balance + price
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + amount}

	//remove pool tokens from user and burn them
	m.P.NumPoolTokens = m.P.NumPoolTokens - amount
	(*tokens)[m.Condition] = PoolToken{m.Condition, (*tokens)[m.Condition].Amount - amount}

	return price, numContracts
}

func (m *Market) Redeem(cs *ContractSet, balance *float32, contracts *map[string]Contract) float32 {
	//exchange contracts for value if the event has been decided
	if cs.Outcome == "none" {
		fmt.Println("Event", cs.Event, "has not been decided yet")
		return -1
	} else if cs.Outcome == m.Condition {
		_, ok := (*contracts)[m.Condition]
		if !ok {
			return 0
		}
		amount := (*contracts)[m.Condition].Amount
		*balance = *balance + (*contracts)[m.Condition].Amount
		(*contracts)[m.Condition] = Contract{m.Condition, 0}
		return amount
	} else {
		//burn contracts
		(*contracts)[m.Condition] = Contract{m.Condition, 0}
		return 0
	}
}

func (cs *ContractSet) BuySet(balance *float32, contracts *map[string]Contract, amount float32) float32 {
	price := amount

	//check if enough usd
	if price > *balance {
		return -1
	}

	//remove usd from user
	*balance = *balance - price

	//add contracts to user
	for _, m := range cs.Markets {
		(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + amount}
	}

	//add the funds to the backing of the liquidity pool
	cs.Backing = cs.Backing + price

	return price
}

func (cs *ContractSet) SellSet(balance *float32, contracts *map[string]Contract, amount float32) float32 {
	price := amount

	//check if enough contracts
	for _, m := range cs.Markets {
		_, ok := (*contracts)[m.Condition]
		if !ok {
			return -1
		} else if (*contracts)[m.Condition].Amount < amount {
			return -1
		}
	}

	//remove contracts from user
	for _, m := range cs.Markets {
		(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - amount}
	}

	//add usd to user
	*balance = *balance + price

	//remove backing from set
	cs.Backing = cs.Backing - price

	return price
}
