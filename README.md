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

#TODO:
Allow for single sided liquidity providing
Track all funds in the poll along with the gains and loses of all users
Integrate redeeming into basic and simulated main programs
