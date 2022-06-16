# 



## constructor(address,address,address)

Initializes the Lender contract

the ctor sets a default value for the feenominator


**Params**
- `the`: swivel contract

- `the`: pendle contract

- `the`: tempus contract

**Returns**

## approve(address,uint256,address)

Approves the redeemer contract to spend the principal tokens held by
the lender contract.



**Params**
- `underlying`: token's address, used to define the market being approved

- `maturity`: of the underlying token, used to define the market being approved

- `the`: address being approved, in this case the redeemer contract


**Returns**
- `bool`: true if the approval was successful, false otherwise

## approve(address[],address[])

Bulk approves the usage of addresses at the given ERC20 addresses
The lengths of the inputs must match because the arrays are paired by index



**Params**
- `u`: array of ERC20 token addresses that will be approved on

- `a`: array of addresses that will be approved


**Returns**
- `true`: if successful

## setFee(uint256)

Sets the feenominator to the given value



**Params**
- `the`: new value of the feenominator, fees are not collected when the feenominator is 0


**Returns**
- `bool`: true if successful

## setMarketPlaceAddress(address)

Sets the address of the marketplace contract which contains the
addresses of all the fixed rate markets



**Params**
- `the`: address of the marketplace contract


**Returns**
- `bool`: true if the address was set, false otherwise

## mint(uint8,address,uint256,uint256)

mint swaps the sender's principal tokens for illuminate's zc tokens
in effect, this opens a new fixed rate position for the sender on illuminate

mint is uniform across all principals, thus there is no need for a 'minter'


**Params**
- `p`: value of a specific principal according to the MarketPlace Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount being minted


**Returns**
- `bool`: true if the mint was successful, false otherwise

## lend(uint8,address,uint256,uint256,address)


lend method signature for both illuminate and yield


**Params**
- `p`: value of a specific principal according to the MarketPlace Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of underlying tokens to lend

- `y`: yield pool that will execute the swap for the principal token


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint256[],address,struct Swivel.Order[],struct Swivel.Components[])

lends to yield pool. remaining balance is sent to the yield pool
TODO: this will change when we implement a check on the gas market

lend method signature for swivel


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: array of amounts of underlying tokens lent to each order in the orders array

- `y`: yield pool

- `o`: array of swivel orders being filled

- `s`: array of signatures for each order in the orders array


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint256,uint256,uint256,address,bytes32)


lend method signature for element


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of principal tokens to lend

- `r`: minimum amount to return, this puts a cap on allowed slippage

- `d`: deadline is a timestamp by which the swap must be executed deadline is a timestamp by which the swap must be executed

- `e`: element pool that is lent to

- `i`: the id of the pool

**Returns**

## lend(uint8,address,uint256,uint256,uint256,uint256)


lend method signature for pendle


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of principal tokens to lend

- `r`: minimum amount to return, this puts a cap on allowed slippage

- `d`: deadline is a timestamp by which the swap must be executed


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint256,uint256,uint256,address,address)

Can be called before maturity to lend to Tempus while minting Illuminate tokens

lend method signature for tempus


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of principal tokens to lend

- `r`: minimum amount to return when executing the swap (sets a limit to slippage)

- `d`: deadline is a timestamp by which the swap must be executed

- `t`: tempus pool that houses the underlying principal tokens

- `x`: tempus amm that executes the swap


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint128,uint256,address,address)

Can be called before maturity to lend to Sense while minting Illuminate tokens

lend method signature for sense
sense provides a [divider] contract that splits [target] assets (underlying)
into PTs and YTs. Each [target] asset has a [series] of contracts, each
identifiable by their [maturity].


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of underlying tokens to lend

- `r`: minimum number of tokens to lend (sets a limit on the order)

- `x`: amm that is used to conduct the swap

- `s`: contract that holds the principal token for this market


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint256,uint256,address,address,uint256)

Can be called before maturity to lend to APWine while minting Illuminate tokens



**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: the amount of underlying tokens to lend

- `r`: the minimum amount of zero-coupon tokens to return accounting for slippage

- `pool`: the address of a given APWine pool

- `i`: the id of the pool


**Returns**
- `uint256`: the amount of principal tokens lent out

## lend(uint8,address,uint256,uint256)


lend method signature for Notional


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of principal tokens to lend


**Returns**
- `uint256`: the amount of principal tokens lent out

## yield(address,address,uint256)

transfers excess funds to yield pool after principal tokens have been lent out

this method is only used by the yield, illuminate and swivel protocols


**Params**
- `u`: address of an underlying asset

- `the`: yield pool to lend to

- `the`: amount of underlying tokens to lend


**Returns**
- `uint256`: the amount of tokens sent to the yield pool

## withdraw(address)

Withdraws accumulated lending fees of the underlying token



**Params**
- `e`: Address of the underlying token to withdraw


**Returns**
- `bool`: true if successful

## calculateFee(uint256)

This method returns the fee based on the amount passed to it. If
feenominator is 0, then there is no fee.



**Params**
- `a`: Amount of underlying tokens to calculate the fee for


**Returns**
- `uint256`: The total for for the given amount

## zcToken(address,uint256)

retrieves the zc token for the given market



**Params**
- `u`: address of the underlying

- `m`: uint256 representing the maturity of the market


**Returns**
- `address`: of the zc token for the market



## `event` Lend(uint8,address,uint256,uint256)




**Params**

## `event` Mint(uint8,address,uint256,uint256)




**Params**

