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
