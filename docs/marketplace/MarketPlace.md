# 



## constructor(address)




**Params**
- `r`: address of the deployed redeemer contract

**Returns**

## createMarket(address,uint256,address[8],string,string,uint8)

creates a new market for the given underlying token and maturity
all seven principals should be provided in the order of their enum value



**Params**
- `u`: address of an underlying asset

- `m`: maturity (timestamp) of the market

- `t`: principal token addresses for this market minus the illuminate
principal (which is added here)

- `n`: name for the illuminate token

- `s`: symbol for the illuminate token

- `d`: decimals for the illuminate token


**Returns**
- `bool`: true if successful

## pause(uint8,bool)

allows the admin to pause a principal's pools



**Params**
- `p`: principal's enum value

- `s`: true if the pool should be paused, false otherwise


**Returns**
- `bool`: true if successful

## setPool(uint8,address,uint256,address)

sets the address for a pool



**Params**
- `p`: enum value of the principal token

- `u`: address of the underlying asset

- `m`: maturity (timestamp) of the market

- `a`: address of the pool


**Returns**
- `bool`: true if successful

## sellPrincipalToken(uint8,address,uint256,uint128)

sells the PT for the PT via the pool



**Params**
- `p`: enum value of the principal token

- `u`: address of the underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of PT to swap


**Returns**
- `uint128`: amount of PT bought

## buyPrincipalToken(uint8,address,uint256,uint128)

buys the underlying for the PT via the pool



**Params**
- `p`: enum value of the principal token

- `u`: address of the underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of underlying tokens to sell


**Returns**
- `uint128`: amount of PT received

## sellUnderlying(uint8,address,uint256,uint128)

sells the underlying for the PT via the pool



**Params**
- `p`: enum value of the principal token

- `u`: address of the underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of underlying to swap


**Returns**
- `uint128`: amount of underlying sold

## buyUnderlying(uint8,address,uint256,uint128)

buys the underlying for the PT via the pool



**Params**
- `p`: enum value of the principal token

- `u`: address of the underlying asset

- `m`: maturity (timestamp) of the market

- `a`: amount of PT to swap


**Returns**
- `uint128`: amount of underlying bought



## `event` CreateMarket(address,uint256)




**Params**

