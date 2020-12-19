package markets

import (
	"fmt"

	"example.com/predictionMarketCentralized/cpmm"
	"github.com/DzananGanic/numericalgo/root"
)

// BuyContract swaps reserve for the amount of contracts specified
func (m *Market) BuyContract(cs *ContractSet, balance *float32, contracts *map[string]Contract, amount float32) float32 {
	//input = usd, output = contracts
	price := cpmm.GetOutputPrice(amount, m.P.Usd, m.P.Contract.Amount)
	//check enough usd to buy
	if price > *balance {
		return -1
	}

	//add usd to pool
	m.P.Usd = m.P.Usd + price

	//remove contracts from pool
	m.P.Contract.Amount = m.P.Contract.Amount - amount

	//add contracts to user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + amount}

	// mm := maker.NewMarketMaker(*verbosePtr)
	// mm.make(&cs, *verbosePtr)
	var profit float32 = 0
	var intermediates map[string]Contract = make(map[string]Contract)
	// perform balancing
	f := func(x float64) float64 {
		var eq float64 = -1
		for _, m := range cs.Markets {
			r := float64(m.P.Usd)
			c := float64(m.P.Contract.Amount)
			eq = eq + ((r - (r*x)/(c+x)) / (c + x))
		}
		return eq
	}

	var totalPrice float64 = 0
	for _, m := range cs.Markets {
		totalPrice = totalPrice + m.GetRatioFloat64()
	}

	//TODO: find a quick algorithm for initial guess
	initialGuess := totalPrice
	iter := 7

	result, _ := root.Newton(f, initialGuess, iter)
	y := float32(result)

	//remove usd from user
	*balance = *balance - y

	profit = profit + y
	//buy contracts as a set
	success := cs.BuySet(&profit, &intermediates, y)
	if success != -1 {
		fmt.Println("MarketMaker bought", y, "contracts sets from the event", cs.Event, "for $", y)
	} else {
		fmt.Println("MarketMaker doesn't have enough funds to buy", y, "contracts sets from the event", cs.Event)
		return -1
	}

	var deltaBacking float32 = -price

	//sell contracts to individual markets
	for i := range cs.Markets {
		price := cs.Markets[i].SellContract(cs, &profit, &intermediates, y)
		if price != -1 {
			fmt.Println("MarketMaker sold", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", price)
			deltaBacking = deltaBacking + price
		} else {
			fmt.Println("Market Maker doesn't have enough contracts to sell", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
			return -1
		}
	}

	//add backing to balance the pool
	println("@@@delta backing", deltaBacking)
	cs.Backing = cs.Backing + deltaBacking

	return price
}

// SellContract swaps the amount of contracts specified for reserve
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

// AddLiquidity provides the amount of liquidity tokens specified to the market
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

// RemoveLiquidity removes the amound of liquidity specified from the market
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

// Redeem swaps contracts for reserve if the outcome of the event has been determined
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

// BuySet purchases a set of contracts for a constant rate
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

// SellSet sells a set of contracts for a constant rate
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
