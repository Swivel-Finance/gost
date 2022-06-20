# MarketPlace


This contract is in charge of managing the avaialable principals for each loan market.
In addition, this contract routes swap orders between metaprincipal tokens and their respective underlying to YieldSpace pools.


## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var | Type |
| --- | --- |
| markets | mapping(address => mapping(uint256 => address[9])) |
| pools | mapping(address => mapping(uint256 => address)) |
| admin | address |
| redeemer | address |
| lender | address |


## Modifiers

### authorized
ensures that only a certain address can call the function



#### Declaration
```solidity
modifier authorized(
address a
)
```

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`a` | address | address that msg.sender must be to be authorized


## Functions

### constructor
intializes the MarketPlace contract



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
|`t` | address[8] | principal token addresses for this market minus the illuminate principal (which is added here)
|`n` | string | name for the illuminate token
|`s` | string | symbol for the illuminate token
|`d` | uint8 | decimals for the illuminate token

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### setPrincipal
allows the admin to create a market



#### Declaration
```solidity
function setPrincipal(
uint8 p,
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
|`p` | uint8 | enum value of the principal token
|`u` | address | underlying token address
|`m` | uint256 | maturity timestamp for the market
|`a` | address | address of the new market

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### setAdmin
sets the admin address



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
address u,
uint256 m,
uint128 a
) external returns
(uint128)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
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
address u,
uint256 m,
uint128 a
) external returns
(uint128)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
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
address u,
uint256 m,
uint128 a
) external returns
(uint128)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
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
address u,
uint256 m,
uint128 a
) external returns
(uint128)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of the underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of PT to swap

#### Returns:
| Type | Description |
| --- | --- |
|`uint128` | amount of underlying bought

### mint
mint liquidity tokens in exchange for adding underlying and PT

> amount of liquidity tokens to mint is calculated from the amount of unaccounted for PT in this contract.
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
|`minRatio` | uint256 | minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of liquidity tokens minted.

### mintWithUnderlying
Mint liquidity tokens in exchange for adding only underlying

> amount of liquidity tokens is calculated from the amount of PT to buy from the pool,
plus the amount of unaccounted for PT in this contract.


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
|`ptBought` | uint256 | amount of `PT` being bought in the Pool, from this we calculate how much underlying it will be taken in.
|`minRatio` | uint256 | minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of liquidity tokens minted

### burn
burn liquidity tokens in exchange for underlying and PT.



#### Declaration
```solidity
function burn(
address u,
uint256 m,
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
|`minRatio` | uint256 | minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | maximum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | amount of tokens burned and returned (tokensBurned, underlyings, PTs).

### burnForUnderlying
burn liquidity tokens in exchange for underlying.



#### Declaration
```solidity
function burnForUnderlying(
address u,
uint256 m,
uint256 minRatio,
uint256 maxRatio
) external returns
(uint256, uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | the address of the underlying token
|`m` | uint256 | the maturity of the principal token
|`minRatio` | uint256 | minimum ratio of underlying to PT in the pool.
|`maxRatio` | uint256 | minimum ratio of underlying to PT in the pool.

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | amount of underlying tokens returned
|`uint256` | amount of PT tokens sent to the pool



## Events

### CreateMarket
emitted upon the creation of a new market




