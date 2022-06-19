# MarketPlace





## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

| Var | Type |
| --- | --- |
| markets | mapping(address => mapping(uint256 => address[9])) |
| pools | mapping(address => mapping(uint256 => address)) |
| admin | address |
| redeemer | address |
| lender | address |
| paused | bool[9] |


## Modifiers

### authorized
No description


#### Declaration
```solidity
modifier authorized
```


### unpaused
No description


#### Declaration
```solidity
modifier unpaused
```



## Functions

### constructor
No description


#### Declaration
```solidity
function constructor(
address r,
address l
) public
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`r` | address | address of the deployed redeemer contract
|`l` | address | address of the deployed lender contract

### createMarket
creates a new market for the given underlying token and maturity
all seven principals should be provided in the order of their enum value



#### Declaration
```solidity
function createMarket(
address u,
uint256 m,
address[8] t,
string n,
string s,
uint8 d
) external authorized returns
(bool)
```

#### Modifiers:
| Modifier |
| --- |
| authorized |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`t` | address[8] | principal token addresses for this market minus the illuminate
principal (which is added here)
|`n` | string | name for the illuminate token
|`s` | string | symbol for the illuminate token
|`d` | uint8 | decimals for the illuminate token

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### pause
allows the admin to pause a principal's pools



#### Declaration
```solidity
function pause(
uint8 p,
bool s
) external authorized returns
(bool)
```

#### Modifiers:
| Modifier |
| --- |
| authorized |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | principal's enum value
|`s` | bool | true if the pool should be paused, false otherwise

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### setAdmin
No description


#### Declaration
```solidity
function setAdmin(
address a
) external authorized returns
(bool)
```

#### Modifiers:
| Modifier |
| --- |
| authorized |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`a` | address | Address of a new admin

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### setPool
sets the address for a pool



#### Declaration
```solidity
function setPool(
address u,
uint256 m,
address a
) external authorized returns
(bool)
```

#### Modifiers:
| Modifier |
| --- |
| authorized |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | address | address of the pool

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### sellPrincipalToken
sells the PT for the PT via the pool



#### Declaration
```solidity
function sellPrincipalToken(
uint8 p,
address u,
uint256 m,
uint128 a
) external unpaused returns
(uint128)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | enum value of the principal token
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of PT to swap

#### Returns:
| Type | Description |
| --- | --- |
|`uint128` | amount of PT bought
### buyPrincipalToken
buys the underlying for the PT via the pool



#### Declaration
```solidity
function buyPrincipalToken(
uint8 p,
address u,
uint256 m,
uint128 a
) external unpaused returns
(uint128)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | enum value of the principal token
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of underlying tokens to sell

#### Returns:
| Type | Description |
| --- | --- |
|`uint128` | amount of PT received
### sellUnderlying
sells the underlying for the PT via the pool



#### Declaration
```solidity
function sellUnderlying(
uint8 p,
address u,
uint256 m,
uint128 a
) external unpaused returns
(uint128)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | enum value of the principal token
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of underlying to swap

#### Returns:
| Type | Description |
| --- | --- |
|`uint128` | amount of underlying sold
### buyUnderlying
buys the underlying for the PT via the pool



#### Declaration
```solidity
function buyUnderlying(
uint8 p,
address u,
uint256 m,
uint128 a
) external unpaused returns
(uint128)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | enum value of the principal token
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of PT to swap

#### Returns:
| Type | Description |
| --- | --- |
|`uint128` | amount of underlying bought
### mint
No description
> Mint liquidity tokens in exchange for adding underlying and PT
The amount of liquidity tokens to mint is calculated from the amount of unaccounted for PT in this contract.
A proportional amount of underlying tokens need to be present in this contract, also unaccounted for.


#### Declaration
```solidity
function mint(
address u,
uint256 m,
uint256 uA,
uint256 ptA,
uint256 minRatio,
uint256 maxRatio
) external returns
(uint256, uint256, uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | the address of the underlying token
|`m` | uint256 | the maturity of the principal token
|`uA` | uint256 | the underlying amount being sent
|`ptA` | uint256 | the principal token amount being sent
|`minRatio` | uint256 | Minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | Maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`The` | amount of liquidity tokens minted.
### mintWithUnderlying
No description
> Mint liquidity tokens in exchange for adding only underlying
The amount of liquidity tokens is calculated from the amount of PT to buy from the pool,
plus the amount of unaccounted for PT in this contract.
The underlying tokens need to be present in this contract, unaccounted for.


#### Declaration
```solidity
function mintWithUnderlying(
address u,
uint256 m,
uint256 a,
uint256 ptBought,
uint256 minRatio,
uint256 maxRatio
) external returns
(uint256, uint256, uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | the address of the underlying token
|`m` | uint256 | the maturity of the principal token
|`a` | uint256 | the underlying amount being sent
|`ptBought` | uint256 | Amount of `PT` being bought in the Pool, from this we calculate how much underlying it will be taken in.
|`minRatio` | uint256 | Minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | Maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`The` | amount of liquidity tokens minted.
### burn
No description
> Burn liquidity tokens in exchange for underlying and PT.
The liquidity tokens need to be in this contract.


#### Declaration
```solidity
function burn(
address minRatio,
uint256 maxRatio
) external returns
(uint256, uint256, uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`minRatio` | address | Minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | Maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`The` | amount of tokens burned and returned (tokensBurned, underlyings, PTs).
### burnForUnderlying
No description
> Burn liquidity tokens in exchange for underlying.


#### Declaration
```solidity
function burnForUnderlying(
address minRatio,
uint256 maxRatio
) external returns
(uint256 tokensBurned, uint256 underlyingOut)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`minRatio` | address | Minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | Minimum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`tokensBurned` | The amount of lp tokens burned.
|`underlyingOut` | The amount of underlying tokens returned.


## Events

### CreateMarket
No description




