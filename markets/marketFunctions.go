package markets

import (
	"fmt"

	"example.com/predictionMarketCentralized/cpmm"
)

func (m *Market) BuyContract(cs *ContractSet, balance *float32, contracts *[]Contract, amount float32) float32 {
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
	index := 0
	alreadyOwn := false
	for i := 0; i < len(*contracts); i++ {
		if (*contracts)[i].Condition == m.Condition {
			index = i
			alreadyOwn = true
			break
		}
	}
	if !alreadyOwn {
		index = len(*contracts)
		*contracts = append(*contracts, Contract{m.P.Contract.Condition, 0})
	}
	(*contracts)[index].Amount = (*contracts)[index].Amount + amount

	cs.Made = false

	return price
}

func (m *Market) SellContract(cs *ContractSet, balance *float32, contracts *[]Contract, amount float32) float32 {
	//input = contract, output = usd
	price := cpmm.GetInputPrice(amount, m.P.Contract.Amount, m.P.Usd)

	//check enough contracts to sell
	index := 0
	owned := false
	for i := 0; i < len(*contracts); i++ {
		if (*contracts)[i].Condition == m.Condition {
			if (*contracts)[i].Amount < amount {
				return -1
			}
			index = i
			owned = true
			break
		}
	}
	if !owned {
		return -1
	}

	//remove contracts from user
	(*contracts)[index].Amount = (*contracts)[index].Amount - amount

	//add contracts to pool
	m.P.Contract.Amount = m.P.Contract.Amount + amount

	//remove usd from pool
	m.P.Usd = m.P.Usd - price

	//add usd to user
	*balance = *balance + price

	cs.Made = false

	return price
}

func (m *Market) AddLiquidity(cs *ContractSet, balance *float32, contracts *[]Contract, tokens *[]PoolToken, amount float32) {
	price := amount * m.P.Usd / m.P.NumPoolTokens
	numContracts := amount * m.P.Contract.Amount / m.P.NumPoolTokens
	//check enough balance and contracts
	index := 0
	owned := false
	for i := 0; i < len(*contracts); i++ {
		if (*contracts)[i].Condition == m.Condition {
			if (*contracts)[i].Amount < amount {
				return
			}
			index = i
			owned = true
			break
		}
	}
	if !owned || *balance < price {
		return
	}

	//remove balance and contracts from user
	*balance = *balance - price
	(*contracts)[index].Amount = (*contracts)[index].Amount - numContracts

	// add balance and user to pool
	m.P.Usd = m.P.Usd + price
	m.P.Contract.Amount = m.P.Contract.Amount + numContracts

	//mint new poolTokens and add to user
	m.P.NumPoolTokens = m.P.NumPoolTokens + amount
	alreadyOwn := false
	for i := 0; i < len(*tokens); i++ {
		if (*tokens)[i].Condition == m.Condition {
			index = i
			alreadyOwn = true
			break
		}
	}
	if !alreadyOwn {
		index = len(*tokens)
		*tokens = append(*tokens, PoolToken{m.P.Contract.Condition, 0})
	}
	(*tokens)[index].Amount = (*tokens)[index].Amount + amount
}

func (m *Market) RemoveLiquidity(cs *ContractSet, balance *float32, contracts *[]Contract, tokens *[]PoolToken, amount float32) {
	price := amount * m.P.Usd / m.P.NumPoolTokens
	numContracts := amount * m.P.Contract.Amount / m.P.NumPoolTokens
	//check enough pool tokens
	tokenIndex := 0
	owned := false
	for i := 0; i < len(*tokens); i++ {
		if (*tokens)[i].Condition == m.Condition {
			if (*tokens)[i].Amount < amount {
				return
			}
			tokenIndex = i
			owned = true
			break
		}
	}
	if !owned {
		return
	}

	//remove balance and contacts from pool
	m.P.Usd = m.P.Usd - price
	m.P.Contract.Amount = m.P.Contract.Amount - numContracts

	//add balance and contracts to user
	index := 0
	*balance = *balance + price
	alreadyOwn := false
	for i := 0; i < len(*contracts); i++ {
		if (*contracts)[i].Condition == m.Condition {
			index = i
			alreadyOwn = true
			break
		}
	}
	if !alreadyOwn {
		index = len(*contracts)
		*contracts = append(*contracts, Contract{m.P.Contract.Condition, 0})
	}
	(*contracts)[index].Amount = (*contracts)[index].Amount + amount

	//remove pool tokens from user and burn them
	m.P.NumPoolTokens = m.P.NumPoolTokens - amount
	(*tokens)[tokenIndex].Amount = (*tokens)[tokenIndex].Amount - amount
}

func (cs *ContractSet) BuySet(balance *float32, contracts *[]Contract, amount float32) float32 {
	price := amount

	//check if enough usd
	if price > *balance {
		return -1
	}

	//remove usd from user
	*balance = *balance - price

	//add contracts to user
	for i := 0; i < len(cs.Markets); i++ {
		alreadyOwn := false
		for j := 0; j < len(*contracts); j++ {
			if (*contracts)[j].Condition == cs.Markets[i].P.Contract.Condition {
				(*contracts)[j].Amount = (*contracts)[j].Amount + amount
				alreadyOwn = true
				break
			}
		}
		if !alreadyOwn {
			(*contracts) = append((*contracts), Contract{cs.Markets[i].P.Contract.Condition, amount})
		}
	}

	//add the funds to the backing of the liquidity pool
	cs.Backing = cs.Backing + price

	return price
}

func (cs *ContractSet) SellSet(balance *float32, contracts *[]Contract, amount float32) float32 {
	price := amount

	//check if enough contracts
	for i := 0; i < len(cs.Markets); i++ {
		owned := false
		for j := 0; j < len(*contracts); j++ {
			if (*contracts)[j].Condition == cs.Markets[i].Condition {
				if (*contracts)[j].Amount < amount {
					fmt.Println("you don't have enough contracts")
					return -1
				}
				owned = true
			}
		}
		if !owned {
			fmt.Println("you don't have enough contracts")
			return -1
		}
	}

	//remove contracts from user
	for i := 0; i < len(cs.Markets); i++ {
		for j := 0; j < len(*contracts); j++ {
			if (*contracts)[j].Condition == cs.Markets[i].Condition {
				(*contracts)[j].Amount = (*contracts)[j].Amount - amount
			}
		}
	}

	//add usd to user
	*balance = *balance + price

	//remove backing from set
	cs.Backing = cs.Backing - price

	return price
}
