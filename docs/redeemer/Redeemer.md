# 



## constructor(address,address,address,address,address)

Initializes the Redeemer contract



**Params**
- `the`: lender contract

- `the`: swivel contract

- `the`: pendle contract

- `the`: tempus contract

- `the`: apwine contract

**Returns**

## setMarketPlaceAddress(address)

Sets the address of the marketplace contract which contains the
addresses of all the fixed rate markets



**Params**
- `m`: the address of the marketplace contract


**Returns**
- `bool`: true if the address was set, false otherwise

## setLenderAddress(address)




**Params**

**Returns**

## redeem(uint8,address,uint256,address)

Redeems underlying token for illuminate, apwine and tempus
protocols

Illuminate burns its tokens prior to redemption, unlike APWine and
Tempus, which withdraw their tokens after transferring the underlying to
the redeem contract. As a result, only Illuminate's redeem returns funds
to the user.
We can avoid a require check on the principal at the start of this
redeem because there is no common business logic executed before the
protocol specific code is executed.


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: the underlying token being redeemed

- `m`: the maturity of the market being redeemed

- `o`: address of the controller or contract that manages the underlying token


**Returns**
- `bool`: true if the redemption was successful

## redeem(uint8,address,uint256)


redeem method for swivel, yield, element and notional. This method redeems all


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: underlying token being redeemed

- `m`: maturity of the market being redeemed


**Returns**
- `bool`: true if the redemption was successful

## redeem(uint8,address,uint256,bytes32)

redeem method signature for pendle



**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: underlying token being redeemed

- `m`: maturity of the market being redeemed

- `i`: forge id used by pendle to redeem the underlying token


**Returns**
- `bool`: true if the redemption was successful

## redeem(uint8,address,uint256,address,address)


redeem method signature for sense


**Params**
- `p`: value of a specific principal according to the Illuminate Principals Enum

- `u`: underlying token being redeemed

- `m`: maturity of the market being redeemed

- `d`: sense contract that splits the loan's prinicpal and yield

- `o`: sense contract that [d] calls into to adapt the underlying to sense

**Returns**

## authRedeem(address,uint256,address,address,uint256)

implements the redeem method for the contract to fulfill the
ERC-5095 interface



**Params**
- `u`: address of the underlying asset

- `m`: maturity of the market

- `f`: address from where the underlying asset will be burned

- `t`: address to where the underlying asset will be transferred

- `a`: amount of the underlying asset to be burned and sent to the to


**Returns**
- `bool`: true if the underlying asset was burned successfully



## `event` Redeem(uint8,address,uint256,uint256)




**Params**

