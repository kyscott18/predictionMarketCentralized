package markets

import "fmt"

type MarketCreator struct {
	Balance   float32
	Contracts map[string]Contract
	Tokens    map[string]PoolToken
}

var oracle MarketCreator

// Contract is the type that represents a contract for an event
type Contract struct {
	Condition string
	Amount    float32
}

// PoolToken is the type that represents stake in a liquidity pool
type PoolToken struct {
	Condition string
	Amount    float32
}

// PoolTokenSS is the type that represents stake in the contracts half of a liquidity pool
type PoolTokenSS struct {
	Condition            string
	Amount               float32
	OriginalNumContracts float32
}

// Pool is the type that represents a liquidity pool containing contracts and usd
type Pool struct {
	Contract      Contract
	Usd           float32
	NumPoolTokens float32
}

// Market is the type that represents a market for the contract with the condition given
type Market struct {
	P         Pool
	Condition string
}

// GetRatioFloat32 gets the price of the contract in the market in terms of reserve
func (m Market) GetRatioFloat32() float32 {
	return m.P.Usd / m.P.Contract.Amount
}

// GetRatioFloat64 get the price of the contract in the market in terms of reserve
func (m Market) GetRatioFloat64() float64 {
	return float64(m.P.Usd) / float64(m.P.Contract.Amount)
}

// ContractSet is the type that represents an event and corresponding set of markets
type ContractSet struct {
	Markets []Market
	Event   string
	Made    bool
	Backing float32
	Outcome string
}

// NewContractSet creates a new contract set
func NewContractSet(event string, conditions []string, ratios []float32, numContracts float32, v bool) ContractSet {
	markets := make([]Market, 0)
	oracle = MarketCreator{0, make(map[string]Contract), make(map[string]PoolToken)}
	for i := 0; i < len(conditions); i++ {
		contract := Contract{conditions[i], numContracts}
		usd := float32(numContracts) * ratios[i]
		p := Pool{contract, usd, numContracts}
		oracle.Tokens[conditions[i]] = PoolToken{conditions[i], numContracts}
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

// PrintState prints the current state of the ContractSet
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

// Validate determines the outcome for a ContractSet
func (cs *ContractSet) Validate(m Market, v bool) {
	cs.Outcome = m.Condition
	if v {
		fmt.Println("Outcome for the event", cs.Event, "has been determined to be", m.Condition)
	}
}

// printOdds prints the odds for a contract in several forms
func (p Pool) printOdds() {
	ratio := float32(p.Contract.Amount) / p.Usd
	fmt.Println("American odds: ", ratioToAmerican(ratio))
	fmt.Println("Implied probability: ", 1/ratio)
	fmt.Println("Decimal odds: ", ratio)
}

// ratioToAmerican converts the ratio of contracts and reserve to american odds
func ratioToAmerican(ratio float32) float32 {
	ratio = ratio - 1
	if ratio < .5 {
		return -100 / ratio
	}
	return 100 * ratio
}
