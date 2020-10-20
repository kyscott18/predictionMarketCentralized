package main

import (
	"example.com/predictionMarketCentralized/maker"
	"example.com/predictionMarketCentralized/markets"
	"example.com/predictionMarketCentralized/players"
)

// //Contract represents a contract on an event
// type Contract struct {
// 	condition string
// 	amount    float32
// }

// //Pool is a liquidity pool containing contracts and usd
// type Pool struct {
// 	contract Contract
// 	usd      float32
// }

// //Market is a market for the contract with the condition given
// type Market struct {
// 	p         Pool
// 	condition string
// }

// func (m *Market) buyContract(event string, balance *float32, contracts *[]Contract, amount float32) float32 {
// 	//input = usd, output = contracts
// 	price := cpmm.GetOutputPrice(amount, m.p.usd, m.p.contract.amount)
// 	//check enough usd to buy
// 	if price > *balance {
// 		return -1
// 	}

// 	//remove usd from user
// 	*balance = *balance - price

// 	//add usd to pool
// 	m.p.usd = m.p.usd + price

// 	//remove contracts from pool
// 	m.p.contract.amount = m.p.contract.amount - amount

// 	//add contracts to user
// 	index := 0
// 	alreadyOwn := false
// 	for i := 0; i < len(*contracts); i++ {
// 		if (*contracts)[i].condition == m.condition {
// 			index = i
// 			alreadyOwn = true
// 			break
// 		}
// 	}
// 	if !alreadyOwn {
// 		index = len(*contracts)
// 		*contracts = append(*contracts, Contract{m.p.contract.condition, 0})
// 	}

// 	(*contracts)[index].amount = (*contracts)[index].amount + amount

// 	return price
// }

// func (m *Market) sellContract(event string, balance *float32, contracts *[]Contract, amount float32) float32 {
// 	//input = contract, output = usd
// 	price := cpmm.GetInputPrice(amount, m.p.contract.amount, m.p.usd)

// 	//check enough contracts to sell
// 	index := 0
// 	owned := false
// 	for i := 0; i < len(*contracts); i++ {
// 		if (*contracts)[i].condition == m.condition {
// 			if (*contracts)[i].amount < amount {
// 				return -1
// 			}
// 			index = i
// 			owned = true
// 			break
// 		}
// 	}

// 	if !owned {
// 		return -1
// 	}

// 	//remove contracts from user
// 	(*contracts)[index].amount = (*contracts)[index].amount - amount

// 	//add contracts to pool
// 	m.p.contract.amount = m.p.contract.amount + amount

// 	//remove usd from pool
// 	m.p.usd = m.p.usd - price

// 	//add usd to user
// 	*balance = *balance + price

// 	return price
// }

// // //ContractSet is the set of markets representing an event
// // type ContractSet struct {
// // 	markets []Market
// // 	event   string
// // }

// func (cs *ContractSet) buySet(balance *float32, contracts *[]Contract, amount float32) float32 {
// 	price := amount

// 	//check if enough usd
// 	if price > *balance {
// 		return -1
// 	}

// 	//remove usd from user
// 	*balance = *balance - price

// 	//add contracts to user
// 	for i := 0; i < len(cs.markets); i++ {
// 		alreadyOwn := false
// 		for j := 0; j < len(*contracts); j++ {
// 			if (*contracts)[j].condition == cs.markets[i].p.contract.condition {
// 				(*contracts)[j].amount = (*contracts)[j].amount + amount
// 				alreadyOwn = true
// 				break
// 			}
// 		}
// 		if !alreadyOwn {
// 			(*contracts) = append((*contracts), Contract{cs.markets[i].p.contract.condition, amount})
// 		}
// 	}

// 	// //verbose statement
// 	// fmt.Println("User", mp.id, "bought", amount, "contract sets from the event", cs.event)
// 	// fmt.Print("\n")
// 	return price

// }

// func (cs *ContractSet) sellSet(balance *float32, contracts *[]Contract, amount float32) float32 {
// 	price := amount

// 	//check if enough contracts
// 	for i := 0; i < len(cs.markets); i++ {
// 		owned := false
// 		for j := 0; j < len(*contracts); j++ {
// 			if (*contracts)[j].condition == cs.markets[i].p.contract.condition {
// 				if (*contracts)[j].amount < amount {
// 					fmt.Println("you don't have enough contracts")
// 					return -1
// 				}
// 				owned = true
// 			}
// 		}
// 		if !owned {
// 			fmt.Println("you don't have enough contracts")
// 			return -1
// 		}
// 	}

// 	//remove contracts from user
// 	for i := 0; i < len(cs.markets); i++ {
// 		for j := 0; j < len(*contracts); j++ {
// 			if (*contracts)[j].condition == cs.markets[i].condition {
// 				(*contracts)[j].amount = (*contracts)[j].amount - amount
// 			}
// 		}
// 	}

// 	//add usd to user
// 	*balance = *balance + price

// 	// //verbose statement
// 	// fmt.Println("User", mp.id, "sold", amount, "contract sets from the event", cs.event)
// 	// fmt.Print("\n")
// 	return price

// }

// func newContractSet(event string, conditions []string, ratios []float32, numContracts float32) ContractSet {
// 	markets := make([]Market, 0)
// 	for i := 0; i < len(conditions); i++ {

// 		contract := Contract{conditions[i], numContracts}
// 		usd := float32(numContracts) * ratios[i]
// 		p := Pool{contract, usd}
// 		markets = append(markets, Market{p, conditions[i]})
// 	}
// 	contractSet := ContractSet{markets, event}
// 	return contractSet
// }

// func (cs ContractSet) printState() {
// 	fmt.Println("Event: ", cs.event)
// 	for i := 0; i < len(cs.markets); i++ {
// 		fmt.Println("Condition: ", cs.markets[i].condition)
// 		cs.markets[i].p.printOdds()
// 		fmt.Println("Contracts in pool: ", cs.markets[i].p.contract.amount)
// 		fmt.Println("USD in pool: ", cs.markets[i].p.usd)
// 		fmt.Printf("\n")
// 	}
// }

// func (p Pool) printOdds() {
// 	ratio := float32(p.contract.amount) / p.usd
// 	fmt.Println("American odds: ", ratioToAmerican(ratio))
// 	fmt.Println("Implied probability: ", 1/ratio)
// 	fmt.Println("Decimal odds: ", ratio)
// }

// func ratioToAmerican(ratio float32) float32 {
// 	ratio = ratio - 1
// 	if ratio < .5 {
// 		return -100 / ratio
// 	}
// 	return 100 * ratio
// }

// //MarketPlayer represents market participants
// type MarketPlayer struct {
// 	id        int
// 	balance   float32
// 	contracts []Contract
// }

// func (mp *MarketPlayer) buyContract(event string, m *Market, amount float32) {

// 	price := m.buyContract(event, &mp.balance, &mp.contracts, amount)

// 	//verbose statement
// 	if price != -1 {
// 		fmt.Println("User", mp.id, "bought", amount, "contracts from the event", event, "with the condition", m.p.contract.condition, "for $", price)
// 	} else {
// 		fmt.Println("User", mp.id, "doesn't have enough funds to buy", amount, "contracts from the event", event, "with the condition", m.p.contract.condition)
// 	}
// 	fmt.Printf("\n")

// }

// func (mp *MarketPlayer) sellContract(event string, m *Market, amount float32) {

// 	price := m.sellContract(event, &mp.balance, &mp.contracts, amount)

// 	//verbose statement
// 	if price != -1 {
// 		fmt.Println("User", mp.id, "sold", amount, "contracts from the event", event, "with the condition", m.p.contract.condition, "for $", price)
// 	} else {
// 		fmt.Println("User", mp.id, "doesn't have enough contracts to sell", amount, "contracts from the event", event, "with the condition", m.p.contract.condition)
// 	}
// 	fmt.Printf("\n")
// }

// func (mp *MarketPlayer) buySet(cs *ContractSet, amount float32) {
// 	cs.buySet(&mp.balance, &mp.contracts, amount)
// }

// func (mp *MarketPlayer) sellSet(cs *ContractSet, amount float32) {
// 	cs.sellSet(&mp.balance, &mp.contracts, amount)
// }

// func (mp MarketPlayer) printState() {
// 	fmt.Println("User", mp.id, "has a balance of", mp.balance)
// 	fmt.Println("Contracts:")
// 	for i := 0; i < len(mp.contracts); i++ {
// 		fmt.Println("Condition:", mp.contracts[i].condition, ", amount:", mp.contracts[i].amount)
// 	}
// 	fmt.Printf("\n")
// }

// func newMarketPlayer(id int, startingBalance float32) MarketPlayer {
// 	mp := MarketPlayer{id, startingBalance, make([]Contract, 0)}
// 	return mp
// }

func main() {
	cs := markets.NewContractSet("coin flip", []string{"heads", "tails"}, []float32{.5, .5}, 200)
	//verbose statement
	cs.PrintState()
	mp1 := players.NewMarketPlayer(1, 20)
	mp1.BuyContract(cs.Event, &cs.Markets[0], 5)
	mp1.PrintState()
	cs.PrintState()
	mm := maker.NewMarketMaker()
	mm.Make(&cs)
}
