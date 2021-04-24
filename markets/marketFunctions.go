package markets

import (
	"fmt"

	"example.com/predictionMarketCentralized/cpmm"
	"github.com/DzananGanic/numericalgo/root"
)

func (m *Market) buyContractRaw(cs *ContractSet, balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	//input = reserve, output = contracts
	numReserve := cpmm.GetOutputPrice(numContracts, m.P.Reserve, m.P.Contract.Amount)
	//check enough usd to buy
	if numReserve > *balance {
		return -1
	}

	//add usd to pool
	m.P.Reserve = m.P.Reserve + numReserve

	//remove contracts from pool
	m.P.Contract.Amount = m.P.Contract.Amount - numContracts

	//add contracts to user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + numContracts}

	//remove usd from user
	*balance = *balance - numReserve

	return numReserve
}

func (m *Market) sellContractRaw(cs *ContractSet, balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	//input = contract, output = usd
	numReserve := cpmm.GetInputPrice(numContracts, m.P.Contract.Amount, m.P.Reserve)

	//check enough contracts to sell
	_, ok := (*contracts)[m.Condition]
	if !ok {
		return -1
	} else if (*contracts)[m.Condition].Amount < numContracts {
		return -1
	}

	//remove contracts from user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - numContracts}

	//add contracts to pool
	m.P.Contract.Amount = m.P.Contract.Amount + numContracts

	//remove usd from pool
	m.P.Reserve = m.P.Reserve - numReserve

	//add usd to user
	*balance = *balance + numReserve

	return numReserve
}

// BuyContract swaps reserve for the amount of contracts specified
func (m *Market) BuyContract(cs *ContractSet, balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	//input = reserve, output = contracts
	numReserve := cpmm.GetOutputPrice(numContracts, m.P.Reserve, m.P.Contract.Amount)
	//check enough usd to buy
	if numReserve > *balance {
		return -1
	}

	//add usd to pool
	m.P.Reserve = m.P.Reserve + numReserve

	//remove contracts from pool
	m.P.Contract.Amount = m.P.Contract.Amount - numContracts

	//add contracts to user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + numContracts}

	// perform balancing
	f := func(x float64) float64 {
		var eq float64 = -1
		for _, m := range cs.Markets {
			r := float64(m.P.Reserve)
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
	initialGuess := float64(numReserve)
	iter := 20

	result, _ := root.Newton(f, initialGuess, iter)
	y := float32(result)

	//remove usd from user
	*balance = *balance - y

	oracle.Balance = oracle.Balance + y

	//buy contracts as a set
	success := cs.BuySet(&oracle.Balance, &oracle.Contracts, y)
	if success != -1 {
		fmt.Println("MarketMaker bought", y, "contracts sets from the event", cs.Event, "for $", y)
	} else {
		fmt.Println("MarketMaker doesn't have enough funds to buy", y, "contracts sets from the event", cs.Event)
		return -1
	}

	var deltaBacking float32 = -numReserve

	//sell contracts to individual markets
	for i := range cs.Markets {
		sellReserve := cs.Markets[i].sellContractRaw(cs, &oracle.Balance, &oracle.Contracts, y)
		if sellReserve != -1 {
			fmt.Println("MarketMaker sold", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", sellReserve)
			deltaBacking = deltaBacking + sellReserve
		} else {
			fmt.Println("Market Maker doesn't have enough contracts to sell", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
			return -1
		}
	}

	oracle.Balance = oracle.Balance - y

	//add backing to balance the pool
	cs.Backing = cs.Backing + deltaBacking

	return y
}

// SellContract swaps the amount of contracts specified for reserve
func (m *Market) SellContract(cs *ContractSet, balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	//input = contract, output = reserve
	numReserve := cpmm.GetInputPrice(numContracts, m.P.Contract.Amount, m.P.Reserve)

	//check enough contracts to sell
	_, ok := (*contracts)[m.Condition]
	if !ok {
		return -1
	} else if (*contracts)[m.Condition].Amount < numContracts {
		return -1
	}

	//remove contracts from user
	(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - numContracts}

	//add contracts to pool
	m.P.Contract.Amount = m.P.Contract.Amount + numContracts

	//remove usd from pool
	m.P.Reserve = m.P.Reserve - numReserve

	f := func(x float64) float64 {
		var eq float64 = -1
		for _, m := range cs.Markets {
			r := float64(m.P.Reserve)
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
	initialGuess := -float64(numReserve)
	iter := 20

	result, _ := root.Newton(f, initialGuess, iter)
	y := float32(result)

	y = y * -1

	//add usd to user
	*balance = *balance + y

	oracle.Balance = oracle.Balance + y

	deltaBacking := numReserve
	//buy contracts from indiviual markets
	for i := range cs.Markets {
		buyReserve := cs.Markets[i].buyContractRaw(cs, &oracle.Balance, &oracle.Contracts, y)
		if buyReserve != -1 {
			fmt.Println("MarketMaker bought", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition, "for $", buyReserve)
			deltaBacking = deltaBacking - buyReserve
		} else {
			fmt.Println("MarketMaker doesn't have enough funds to buy", y, "contracts from the event", cs.Event, "with the condition", cs.Markets[i].P.Contract.Condition)
			return -1
		}
	}

	//Sell contracts as a set
	success := cs.SellSet(&oracle.Balance, &oracle.Contracts, y)
	//verbose statement
	if success != -1 {
		fmt.Println("MarketMaker sold", y, "contracts sets from the event", cs.Event, "for $", y)
	} else {
		fmt.Println("MarketMaker doesn't have enough contracts to sell", y, "contracts sets from the event", cs.Event)
		return -1
	}

	oracle.Balance = oracle.Balance - y

	//add backing to balance the pool
	cs.Backing = cs.Backing + deltaBacking

	return y
}

// AddLiquidity provides the amount of liquidity tokens specified to the market
func (cs *ContractSet) AddLiquidity(balance *float32, tokens *float32, numPoolTokens float32) float32 {

	//check enough balance
	if *balance < numPoolTokens {
		return -1
	}

	//remove balance from user
	*balance = *balance - numPoolTokens

	//add pool tokens to user
	*tokens = *tokens + numPoolTokens

	// add liquidity to pool
	TotalContractsMinted = TotalContractsMinted + numPoolTokens
	for i := range cs.Markets {
		cs.Markets[i].P.Reserve = cs.Markets[i].P.Reserve + numPoolTokens*cs.Markets[i].GetRatioFloat32()
		cs.Markets[i].P.Contract.Amount = cs.Markets[i].P.Contract.Amount + numPoolTokens
	}

	return numPoolTokens
}

// RemoveLiquidity removes the amound of liquidity specified from the market
func (cs *ContractSet) RemoveLiquidity(balance *float32, tokens *float32, numPoolTokens float32) float32 {
	//check enough pool tokens
	if numPoolTokens > *tokens {
		return -1
	}

	//remove liquidity from pool
	for i := range cs.Markets {
		cs.Markets[i].P.Reserve = cs.Markets[i].P.Reserve - cs.Markets[i].P.Reserve*(numPoolTokens/TotalContractsMinted)
		cs.Markets[i].P.Contract.Amount = cs.Markets[i].P.Contract.Amount - cs.Markets[i].P.Contract.Amount*(numPoolTokens/TotalContractsMinted)
	}
	cs.Backing = cs.Backing - cs.Backing*(numPoolTokens/TotalContractsMinted)
	TotalContractsMinted = TotalContractsMinted - numPoolTokens

	//add balance and contracts to user
	*balance = *balance + numPoolTokens

	//remove pool tokens from user and burn them
	*tokens = *tokens - numPoolTokens

	return numPoolTokens
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
func (cs *ContractSet) BuySet(balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	numReserve := numContracts

	//check if enough usd
	if numReserve > *balance {
		return -1
	}

	//remove usd from user
	*balance = *balance - numReserve

	//add contracts to user
	for _, m := range cs.Markets {
		(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount + numContracts}
	}

	//add the funds to the backing of the liquidity pool
	cs.Backing = cs.Backing + numReserve

	TotalContractsMinted += numContracts

	return numReserve
}

// SellSet sells a set of contracts for a constant rate
func (cs *ContractSet) SellSet(balance *float32, contracts *map[string]Contract, numContracts float32) float32 {
	numReserve := numContracts

	//check if enough contracts
	for _, m := range cs.Markets {
		_, ok := (*contracts)[m.Condition]
		if !ok {
			return -1
		} else if (*contracts)[m.Condition].Amount < numContracts {
			return -1
		}
	}

	//remove contracts from user
	for _, m := range cs.Markets {
		(*contracts)[m.Condition] = Contract{m.Condition, (*contracts)[m.Condition].Amount - numContracts}
	}

	//add usd to user
	*balance = *balance + numReserve

	//remove backing from set
	cs.Backing = cs.Backing - numReserve

	TotalContractsMinted -= numContracts

	return numReserve
}
