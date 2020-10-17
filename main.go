package main

import (
	"fmt"
)

//Contract represents a contract on an event
type Contract struct {
	condition string
	amount    float32
}

//Pool is a liquidity pool containing contracts and usd
type Pool struct {
	contract Contract
	usd      float32
}

//Market is a market for the contract with the condition given
type Market struct {
	p         Pool
	condition string
}

func (m *Market) buyContract(event string, mp *MarketPlayer, amount float32) {
	//input = usd, output = contracts
	price := getOutputPrice(amount, m.p.usd, m.p.contract.amount)
	//check enough usd to buy
	if price > mp.balance {
		fmt.Println("You don't have enough usd")
		return
	}

	//remove usd from user
	mp.balance = mp.balance - price

	//add usd to pool
	m.p.usd = m.p.usd + price

	//remove contracts from pool
	m.p.contract.amount = m.p.contract.amount - amount

	//add contracts to user
	alreadyOwn := false
	for i := 0; i < len(mp.contracts); i++ {
		if mp.contracts[i].condition == m.p.contract.condition {
			mp.contracts[i].amount = mp.contracts[i].amount + amount
			alreadyOwn = true
			break
		}
	}
	if !alreadyOwn {
		mp.contracts = append(mp.contracts, Contract{m.p.contract.condition, amount})
	}

	//verbose statement
	fmt.Println("User", mp.id, "bought", amount, "contracts from the event", event, "with the condition", m.p.contract.condition, "for $", price)
	fmt.Printf("\n")
}

func (m *Market) sellContract(event string, mp *MarketPlayer, amount float32) {
	//input = contract, output = usd
	price := getInputPrice(amount, m.p.contract.amount, m.p.usd)

	//check enough contracts to sell
	//remove contracts from user
	owned := false
	for i := 0; i < len(mp.contracts); i++ {
		if mp.contracts[i].condition == m.p.contract.condition {
			if mp.contracts[i].amount < amount {
				fmt.Println("you don't have enough contracts")
				return
			}
			mp.contracts[i].amount = mp.contracts[i].amount - amount
			owned = true
			break
		}
	}
	if !owned {
		fmt.Println("you don't have enough contracts")
		return
	}

	//add contracts to pool
	m.p.contract.amount = m.p.contract.amount + amount

	//remove usd from pool
	m.p.usd = m.p.usd - price

	//add usd to user
	mp.balance = mp.balance + price

	//verbose statement
	fmt.Println("User", mp.id, "sold", amount, "contracts from the event", event, "with the condition", m.p.contract.condition, "for $", price)
	fmt.Printf("\n")
}

//ContractSet is the set of markets representing an event
type ContractSet struct {
	markets []Market
	event   string
}

func (cs *ContractSet) buySet(mp *MarketPlayer, amount float32) {
	price := amount

	if price > mp.balance {
		fmt.Println("You don't have enough usd")
		return
	}

	mp.balance = mp.balance - price

	//add contracts to user
	for i := 0; i < len(cs.markets); i++ {
		alreadyOwn := false
		for j := 0; j < len(mp.contracts); j++ {
			if mp.contracts[j].condition == cs.markets[i].p.contract.condition {
				mp.contracts[j].amount = mp.contracts[j].amount + amount
				alreadyOwn = true
				break
			}
		}
		if !alreadyOwn {
			mp.contracts = append(mp.contracts, Contract{cs.markets[i].p.contract.condition, amount})
		}
	}

	//verbose statement
	fmt.Println("User", mp.id, "bought", amount, "contract sets from the event", cs.event)

}

func newContractSet(event string, conditions []string, ratios []float32, numContracts float32) ContractSet {
	markets := make([]Market, 0)
	for i := 0; i < len(conditions); i++ {

		contract := Contract{conditions[i], numContracts}
		usd := float32(numContracts) * ratios[i]
		p := Pool{contract, usd}
		markets = append(markets, Market{p, conditions[i]})
	}
	contractSet := ContractSet{markets, event}
	return contractSet
}

func (cs ContractSet) printState() {
	fmt.Println("Event: ", cs.event)
	for i := 0; i < len(cs.markets); i++ {
		fmt.Println("Condition: ", cs.markets[i].condition)
		cs.markets[i].p.printOdds()
		fmt.Println("Contracts in pool: ", cs.markets[i].p.contract.amount)
		fmt.Println("USD in pool: ", cs.markets[i].p.usd)
		fmt.Printf("\n")
	}
}

func (p Pool) printOdds() {
	ratio := float32(p.contract.amount) / p.usd
	fmt.Println("American odds: ", ratioToAmerican(ratio))
	fmt.Println("Implied probability: ", 1/ratio)
	fmt.Println("Decimal odds: ", ratio)
}

func ratioToAmerican(ratio float32) float32 {
	ratio = ratio - 1
	if ratio < .5 {
		return -100 / ratio
	}
	return 100 * ratio
}

// From Uniswap
func getInputPrice(inputAmount float32, inputReserve float32, outputReserve float32) float32 {
	inputAmountWithFee := inputAmount * 997
	numerator := inputAmountWithFee * outputReserve
	denominator := (inputReserve * 1000) + inputAmountWithFee
	return numerator / denominator
}

// From Uniswap
func getOutputPrice(outputAmount float32, inputReserve float32, outputReserve float32) float32 {
	numerator := inputReserve * outputAmount * 1000
	denominator := (outputReserve - outputAmount) * 997
	return numerator/denominator + 1
}

//MarketPlayer represents market participants
type MarketPlayer struct {
	id        int
	balance   float32
	contracts []Contract
}

func (mp *MarketPlayer) buyContract(event string, m *Market, amount float32) {
	m.buyContract(event, mp, amount)
}

func (mp *MarketPlayer) sellContract(event string, m *Market, amount float32) {
	m.sellContract(event, mp, amount)
}

func (mp *MarketPlayer) buySet(cs *ContractSet, amount float32) {
	cs.buySet(mp, amount)
}

func newMarketPlayer(id int, startingBalance float32) MarketPlayer {
	mp := MarketPlayer{id, startingBalance, make([]Contract, 0)}
	return mp
}

func main() {
	cs := newContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
	//verbose statement
	cs.printState()
	//mp1 := newMarketPlayer(1, 20)
}
