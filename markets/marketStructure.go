package markets

import (
	"fmt"
	"math"
)

// MarketCreator is all the contracts, tokens, and reserve that the market creator holds
type MarketCreator struct {
	Balance   float64
	Contracts map[string]Contract
	Tokens    map[string]PoolToken
}

var oracle MarketCreator

// Contract is the type that represents a contract for an event
type Contract struct {
	Condition string
	Amount    float64
}

// PoolToken is the type that represents stake in a liquidity pool
type PoolToken struct {
	Condition string
	Amount    float64
}

// Pool is the type that represents a liquidity pool containing contracts and usd
type Pool struct {
	Contract  Contract
	Reserve   float64
	PoolToken PoolToken
}

// Market is the type that represents a market for the contract with the condition given
type Market struct {
	P         Pool
	Condition string
}

// GetRatioFloat64 get the price of the contract in the market in terms of reserve
func (m Market) GetRatioFloat64() float64 {
	return float64(m.P.Reserve) / float64(m.P.Contract.Amount)
}

// ContractSet is the type that represents an event and corresponding set of markets
type ContractSet struct {
	Markets []Market
	Event   string
	Backing float64
	Outcome string
	Mints   float64
}

// NewContractSet creates a new contract set
func NewContractSet(event string, conditions []string, ratios []float64, numContracts float64, v bool) ContractSet {
	markets := make([]Market, 0)
	oracle = MarketCreator{0, make(map[string]Contract), make(map[string]PoolToken)}
	Mints := numContracts
	for i := 0; i < len(conditions); i++ {
		contract := Contract{conditions[i], numContracts}
		usd := numContracts * ratios[i]
		poolToken := PoolToken{conditions[i], math.Sqrt(numContracts * usd)}
		p := Pool{contract, usd, poolToken}
		markets = append(markets, Market{p, conditions[i]})
	}

	contractSet := ContractSet{markets, event, 0, "none", Mints}
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
	fmt.Println("Backing:", cs.Backing)
	fmt.Println("Total Contracts Minted: ", cs.Mints)

	for i := 0; i < len(cs.Markets); i++ {
		fmt.Println("Condition: ", cs.Markets[i].Condition)
		cs.Markets[i].P.printOdds()
		fmt.Println("Contracts in pool: ", cs.Markets[i].P.Contract.Amount)
		fmt.Println("PoolTokens in pool: ", cs.Markets[i].P.PoolToken.Amount)
		fmt.Println("Reserve in pool: ", cs.Markets[i].P.Reserve)

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
	ratio := float64(p.Contract.Amount) / p.Reserve
	fmt.Println("American odds: ", ratioToAmerican(ratio))
	fmt.Println("Implied probability: ", 1/ratio)
	fmt.Println("Decimal odds: ", ratio)
}

// ratioToAmerican converts the ratio of contracts and reserve to american odds
func ratioToAmerican(ratio float64) float64 {
	ratio = ratio - 1
	if ratio < .5 {
		return -100 / ratio
	}
	return 100 * ratio
}
