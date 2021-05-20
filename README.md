This is going to be a proof of concept for my prediction market technology that I hope to implement in a decentralized setting in the future. 

The core priniciple is that contracts are tokens who represent a value if a condition is met and no value if it is not. It is easiest to determine this number to be 1. The easiest example is to think of a token having a value one if a coin is heads and zero if it is tails. 

Contracts sets can be formed where the set of contracts are mutually exclusive and cover the complete set of outcomes. In the example this would amount to a set of contracts where one represents heads and one represents tails. 

These contracts can be traded in a system that uses liquidity pools driven by an automated market maker. I chose to use the constant product market maker where x*y=k.

Each market offers a swap between a reserve currency and whatever contract that market supports. This means that you could swap a reserve currency and receive either a head or tails contract token or swap a head or tails contract token and receive the reserve currency. 

In order to be able to redeem all contracts, there must be enough reserve in the overall market. We must have a backing fund that is used for redeeming validated contracts. This means that once the event has taken place and the outcome has been determined, we need to guarantee that there will be enough reserve currency to redeem the valid contracts. 

Because of the nature of the contract sets the market also supports buying or selling complete contracts sets for a value of one. The funds are added or subtracted from the backing fund to allow for this. This means that a user could buy or sell both a heads and tails token and it would always be worth one unit of reserve currency. 

When a swap is made in one of the markets composing a set, this leaves the overall set unbalanced because the value of all the contracts in the set is not equal to one. Therefore, there is a number of contracts that can bought or sold and a set and sold or bought from individual markets, respectively. 

Impermanent loss leads to the loss of value from a market using the liquidity pool structure if left unchecked. The market must itself perform the second level market making that occurs when a swap is made in an individual market and use some of these profits to make sure that the set is sufficiently backed and return the rest to the participant. 

Liquidity can be provided to allow the market pools to grow which allows for lower price slippage. To do this a participant provides an equal ratio of the contract and the reserve. 

Because of impermanent loss, the liquidity provider opens themselves up to almost complete losses which is unexpected for many. Single sided liquidity provided allows for market participants to provide only the contracts that they have accumulated and the correct amount of reserve is taken from backing and added as liquidity. 