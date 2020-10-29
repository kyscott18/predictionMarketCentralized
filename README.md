This is going to be a proof of concept for my prediction market technology that I hope to implement in a decentralized setting in the future. 
Contracts are giving value upon a condition. Contract sets are mutually exclusive and total. Each contract has a separate liquidity pool with a constant product automated market maker. Second level market making is done by either buying or selling a set of contracts. 
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
