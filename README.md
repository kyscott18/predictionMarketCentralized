This is going to be a proof of concept for my prediction market technology that I hope to implement in a decentralized setting in the future. 
Contracts are giving value upon a condition. Contract sets are mutually exclusive and total. Each contract has a separate liquidity pool with a constant product automated market maker. Second level market making is done by either buying or selling a set of contracts. 
1. Allow for the creation of bets, creation of markets for bets and underlying liquidity pools, 
2. Create market players and swaps
2. Implement second level market making 
3. Verification of bets and trading in, done by verifying outcome of event


class contract {
    condtion
    id
}

class pool {
    contracts[]
    usd
}

class market {
    pool
    condition
}

class ContractSet {
    markets[]
    buySet()
    sellSet()
}  

class marketPlayer {
    id
    balance
    contracts[]
    buyContract()
    sellContract()
}
