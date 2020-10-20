package markets

import "fmt"

//Contract represents a contract on an event
type Contract struct {
	Condition string
	Amount    float32
}

//Pool is a liquidity pool containing contracts and usd
type Pool struct {
	Contract Contract
	Usd      float32
}

//Market is a market for the contract with the condition given
type Market struct {
	P         Pool
	Condition string
}

//ContractSet is the set of markets representing an event
type ContractSet struct {
	Markets []Market
	Event   string
}

func NewContractSet(event string, conditions []string, ratios []float32, numContracts float32) ContractSet {
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

func (cs ContractSet) PrintState() {
	fmt.Println("Event: ", cs.Event)
	for i := 0; i < len(cs.Markets); i++ {
		fmt.Println("Condition: ", cs.Markets[i].Condition)
		cs.Markets[i].P.printOdds()
		fmt.Println("Contracts in pool: ", cs.Markets[i].P.Contract.Amount)
		fmt.Println("USD in pool: ", cs.Markets[i].P.Usd)
		fmt.Printf("\n")
	}
}

func (p Pool) printOdds() {
	ratio := float32(p.Contract.Amount) / p.Usd
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