This is going to be a proof of concept for my prediction market technology that I hope to implement in a decentralized setting in the future. 
Contracts are giving value 1 upon a condition. Contract sets are mutually exclusive and cover all outcomes of an event. Each contract has a separate market with an underlying liquidity pool. Each liquidity pool is guided with an automated constant product market maker. Each pool can be ineracted with independetly by buying or selling a contract. The structure also supports buying and selling complete sets of contracts for the derived value of 1. This enables second level market making to be done if the combined implied probability of all the markets in the set is not 1. Second level market making is done by either buying sets of contracts and selling to individual markets or buying from individual markets and selling sets of contracts. Liquidity can be provided to markets by providing contracts and reserve backing at the current ratio. Players are rewarded with a proportional amount of pool tokens which can be traded in at a later time for contracts and reserve backing at that current ratio. 
1. Allow for the creation of bets, creation of markets for bets and underlying liquidity pools, 
2. Create market players and swaps
3. Implement second level market making 
4. Simulated Participants using the ratio between contracts and usd in a pool and bernoulli distributions
5. Add balance pool to each ContractSet so that we can monitor value
6. Add pool tokens for each of the pools
7. Allow for adding contracts to pools and pulling funds from balance pool
8. Verification of contract outcome and redeeming, done by verifying outcome of event

PoolTokens:
    PoolTokens are proof that the marketPlayer added their contracts to the liquidity pool which can be redeemed at a later date.
    The creator of the market is rewarded pool tokens for providing the initial contracts and backing for those contracts in usd. There are two options for providing liquidity for the contracts. The first is very similar to Uniswap where the user provides an equal amount of both sides of the liquidity pool. When they redeem their token they are promised that the value will be the same but there may be a different ration between them. The other option is different from uniswaps liquidity tokens because the marketPlayer doesn't have to provide both assets in the pool just the contract and then the backing is used as the usd. This may cause some issues with redeeming the token.
    Two situations of liquidity providing:
        The user provides contracts and then the contracts are worth less than they were when they were entered. In this situation, the user does not get back the same amount of contracts as when they joined the liquidity pool.
        The other situation is that the user provides contracts and then the contracts are worth more than they were when they were entered. In this situation, the user may be limited to the amount that they can withdrawal if
