package markets

import "fmt"

var MarketCreatorToken map[string]PoolToken

//Contract represents a contract on an event
type Contract struct {
	Condition string
	Amount    float32
}

//PoolToken represents stake in a liquidity pool
type PoolToken struct {
	Condition string
	Amount    float32
}

type PoolTokenSS struct {
	Condition string
	Amount    float32
}

//Pool is a liquidity pool containing contracts and usd
type Pool struct {
	Contract      Contract
	Usd           float32
	NumPoolTokens float32
}

//Market is a market for the contract with the condition given
type Market struct {
	P         Pool
	Condition string
}

//GetRatioFloat32 get the price of the contract in the market in terms of USD
func (m Market) GetRatioFloat32() float32 {
	return m.P.Usd / m.P.Contract.Amount
}

//GetRatioFloat64 get the price of the contract in the market in terms of USD
func (m Market) GetRatioFloat64() float64 {
	return float64(m.P.Usd) / float64(m.P.Contract.Amount)
}

//ContractSet is the set of markets representing an event
type ContractSet struct {
	Markets []Market
	Event   string
	Made    bool
	Backing float32
	Outcome string
}

//NewContractSet returns a newly created ContractSet
func NewContractSet(event string, conditions []string, ratios []float32, numContracts float32, v bool) ContractSet {
	markets := make([]Market, 0)
	MarketCreatorToken := make(map[string]PoolToken)
	for i := 0; i < len(conditions); i++ {
		contract := Contract{conditions[i], numContracts}
		usd := float32(numContracts) * ratios[i]
		p := Pool{contract, usd, numContracts}
		MarketCreatorToken[conditions[i]] = PoolToken{conditions[i], MarketCreatorToken[conditions[i]].Amount + numContracts}
		markets = append(markets, Market{p, conditions[i]})
	}

	contractSet := ContractSet{markets, event, true, 0, "none"}
	//verbose statement
	if v {
		fmt.Println("Newly created ContractSet")
		fmt.Println("Event:", event)
		fmt.Println("Conditions:", conditions)
		fmt.Println("Ratios:", ratios)
		fmt.Println("NumContracts", numContracts)
		fmt.Println()
	}

	return contractSet
}

//PrintState prints the state of the ContractSet
func (cs ContractSet) PrintState() {
	fmt.Println("State of ContractSet")
	fmt.Println("Event: ", cs.Event)
	fmt.Println("Make Status:", cs.Made)
	fmt.Println("Backing:", cs.Backing)

	for i := 0; i < len(cs.Markets); i++ {
		fmt.Println("Condition: ", cs.Markets[i].Condition)
		cs.Markets[i].P.printOdds()
		fmt.Println("Contracts in pool: ", cs.Markets[i].P.Contract.Amount)
		fmt.Println("USD in pool: ", cs.Markets[i].P.Usd)

		fmt.Printf("\n")
	}
}

func (cs *ContractSet) Validate(m Market, v bool) {
	cs.Outcome = m.Condition
	if v {
		fmt.Println("Outcome for the event", cs.Event, "has been determined to be", m.Condition)
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
