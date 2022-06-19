# Lender





## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

| Var | Type |
| --- | --- |
| HOLD | uint256 |
| admin | address |
| marketPlace | address |
| swivelAddr | address |
| pendleAddr | address |
| tempusAddr | address |
| feenominator | uint256 |
| fees | mapping(address => uint256) |
| withdrawals | mapping(address => uint256) |


## Modifiers

### authorized
No description


#### Declaration
```solidity
modifier authorized
```



## Functions

### constructor
Initializes the Lender contract

> the ctor sets a default value for the feenominator


#### Declaration
```solidity
function constructor(
address s,
address p,
address t
) public
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`s` | address | the swivel contract
|`p` | address | the pendle contract
|`t` | address | the tempus contract

### approve
Approves the redeemer contract to spend the principal tokens held by
the lender contract.



#### Declaration
```solidity
function approve(
address u,
uint256 m,
address r
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
|`u` | address | underlying token's address, used to define the market being approved
|`m` | uint256 | maturity of the underlying token, used to define the market being approved
|`r` | address | the address being approved, in this case the redeemer contract

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the approval was successful, false otherwise
### approve
Bulk approves the usage of addresses at the given ERC20 addresses
The lengths of the inputs must match because the arrays are paired by index



#### Declaration
```solidity
function approve(
address[] u,
address[] a
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
|`u` | address[] | array of ERC20 token addresses that will be approved on
|`a` | address[] | array of addresses that will be approved

#### Returns:
| Type | Description |
| --- | --- |
|`true` | if successful
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

### setFee
Sets the feenominator to the given value



#### Declaration
```solidity
function setFee(
uint256 f
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
|`f` | uint256 | the new value of the feenominator, fees are not collected when the feenominator is 0

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### setMarketPlaceAddress
Sets the address of the marketplace contract which contains the
addresses of all the fixed rate markets



#### Declaration
```solidity
function setMarketPlaceAddress(
address m
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
|`m` | address | the address of the marketplace contract

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the address was set, false otherwise
### setSwivel
Sets the feenominator to the given value



#### Declaration
```solidity
function setSwivel(
address s
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
|`s` | address | the address of the Swivel.sol Router

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### mint
mint swaps the sender's principal tokens for illuminate's zc tokens
in effect, this opens a new fixed rate position for the sender on illuminate

> mint is uniform across all principals, thus there is no need for a 'minter'


#### Declaration
```solidity
function mint(
uint8 p,
address u,
uint256 m,
uint256 a
) public returns
(bool)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the MarketPlace Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount being minted

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the mint was successful, false otherwise
### lend
No description
> lend method signature for both illuminate and yield


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
address y
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the MarketPlace Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount of underlying tokens to lend
|`y` | address | yieldspace pool that will execute the swap for the principal token

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
lends to yield pool. remaining balance is sent to the yield pool
TODO: this will change when we implement a check on the gas market

> lend method signature for swivel


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256[] a,
address y,
struct Swivel.Order[] o,
struct Swivel.Components[] s
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256[] | array of amounts of underlying tokens lent to each order in the orders array
|`y` | address | yield pool
|`o` | struct Swivel.Order[] | array of swivel orders being filled
|`s` | struct Swivel.Components[] | array of signatures for each order in the orders array

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
No description
> lend method signature for element


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
uint256 r,
uint256 d,
address e,
bytes32 i
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount of principal tokens to lend
|`r` | uint256 | minimum amount to return, this puts a cap on allowed slippage
|`d` | uint256 | deadline is a timestamp by which the swap must be executed deadline is a timestamp by which the swap must be executed
|`e` | address | element pool that is lent to
|`i` | bytes32 | the id of the pool

### lend
No description
> lend method signature for pendle


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
uint256 r,
uint256 d
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount of principal tokens to lend
|`r` | uint256 | minimum amount to return, this puts a cap on allowed slippage
|`d` | uint256 | deadline is a timestamp by which the swap must be executed

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
Can be called before maturity to lend to Tempus while minting Illuminate tokens

> lend method signature for tempus


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
uint256 r,
uint256 d,
address t,
address x
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount of principal tokens to lend
|`r` | uint256 | minimum amount to return when executing the swap (sets a limit to slippage)
|`d` | uint256 | deadline is a timestamp by which the swap must be executed
|`t` | address | tempus pool that houses the underlying principal tokens
|`x` | address | tempus amm that executes the swap

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
Can be called before maturity to lend to Sense while minting Illuminate tokens

> lend method signature for sense
sense provides a [divider] contract that splits [target] assets (underlying)
into PTs and YTs. Each [target] asset has a [series] of contracts, each
identifiable by their [maturity].


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint128 a,
uint256 r,
address x,
address s
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint128 | amount of underlying tokens to lend
|`r` | uint256 | minimum number of tokens to lend (sets a limit on the order)
|`x` | address | amm that is used to conduct the swap
|`s` | address | contract that holds the principal token for this market

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
Can be called before maturity to lend to APWine while minting Illuminate tokens



#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
uint256 r,
address pool,
address i
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | the amount of underlying tokens to lend
|`r` | uint256 | the minimum amount of zero-coupon tokens to return accounting for slippage
|`pool` | address | the address of a given APWine pool
|`i` | address | the id of the pool

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### lend
No description
> lend method signature for Notional


#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a
) public returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | address of an underlying asset
|`m` | uint256 | maturity (timestamp) of the market
|`a` | uint256 | amount of principal tokens to lend

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out
### yield
transfers excess funds to yield pool after principal tokens have been lent out

> this method is only used by the yield, illuminate and swivel protocols


#### Declaration
```solidity
function yield(
address u,
address y,
uint256 a
) internal returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of an underlying asset
|`y` | address | the yield pool to lend to
|`a` | uint256 | the amount of underlying tokens to lend

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of tokens sent to the yield pool
### withdrawFee
Withdraws accumulated lending fees of the underlying token



#### Declaration
```solidity
function withdrawFee(
address e
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
|`e` | address | Address of the underlying token to withdraw

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful
### calculateFee
This method returns the fee based on the amount passed to it. If
feenominator is 0, then there is no fee.



#### Declaration
```solidity
function calculateFee(
uint256 a
) internal returns
(uint256)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`a` | uint256 | Amount of underlying tokens to calculate the fee for

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | The total for for the given amount
### scheduleWithdrawal
Allows the admin to schedule the withdrawal of tokens



#### Declaration
```solidity
function scheduleWithdrawal(
address e
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
|`e` | address | Address of (erc20) token to withdraw

### blockWithdrawal
Emergency function to block unplanned withdrawals



#### Declaration
```solidity
function blockWithdrawal(
address e
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
|`e` | address | Address of token withdrawal to block

### withdraw
Allows the admin to withdraw the given token, provided the holding 
period has been observed



#### Declaration
```solidity
function withdraw(
address e
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
|`e` | address | Address of token to withdraw

### principalToken
retrieves the zc token for the given market



#### Declaration
```solidity
function principalToken(
address u,
uint256 m
) internal returns
(address)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of the underlying
|`m` | uint256 | uint256 representing the maturity of the market

#### Returns:
| Type | Description |
| --- | --- |
|`address` | of the zc token for the market


## Events

### Lend
Emitted upon executed lend




### Mint
Emitted upon minted ERC5095 to user




### ScheduleWithdrawal
Emitted on token withdrawal scheduling




### BlockWithdrawal
Emitted on token withdrawal blocking




