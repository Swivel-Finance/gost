# Lender


The lender contract executes loans on behalf of users. 
The contract holds the principal tokens for each market and mints an ERC-5095 position to users to represent their lent positions.


## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var | Type |
| --- | --- |
| HOLD | uint256 |
| admin | address |
| marketPlace | address |
| paused | mapping(address => mapping(uint256 => bool)) |
| swivelAddr | address |
| pendleAddr | address |
| tempusAddr | address |
| feenominator | uint256 |
| fees | mapping(address => uint256) |
| withdrawals | mapping(address => uint256) |


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

### unpaused
reverts on all markets where the paused mapping returns true



#### Declaration
```solidity
modifier unpaused(
address u,
uint256 m
)
```

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | address of the underlying
|`m` | uint256 | uint256 representing the maturity of the market


## Functions

### constructor
initializes the Lender contract



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
approves the redeemer contract to spend the principal tokens held by the lender contract.



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
bulk approves the usage of addresses at the given ERC20 addresses. 

> the lengths of the inputs must match because the arrays are paired by index


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
|`a` | address | address of a new admin

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### setFee
sets the feenominator to the given value



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
sets the address of the marketplace contract which contains the addresses of all the fixed rate markets



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
sets the feenominator to the given value



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
mint swaps the sender's principal tokens for illuminate's ERC5095 tokens in effect, this opens a new fixed rate position for the sender on illuminate



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
lend method signature for both illuminate and yield



#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
address y
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
lend method signature for swivel

> lends to yield pool. remaining balance is sent to the yield pool


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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
lend method signature for element



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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of principal tokens lent out

### lend
lend method signature for pendle



#### Declaration
```solidity
function lend(
uint8 p,
address u,
uint256 m,
uint256 a,
uint256 r,
uint256 d
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
lend method signature for tempus

> This method can be called before maturity to lend to Tempus while minting Illuminate tokens


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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
lend method signature for sense

> this method can be called before maturity to lend to Sense while minting Illuminate tokens
sense provides a [divider] contract that splits [target] assets (underlying) into PTs and YTs. Each [target] asset has a [series] of contracts, each identifiable by their [maturity].


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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
this method can be called before maturity to lend to APWine while minting Illuminate tokens



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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
) public unpaused returns
(uint256)
```

#### Modifiers:
| Modifier |
| --- |
| unpaused |

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
uint256 a,
address r
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
|`r` | address | the receiving address for PTs

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | the amount of tokens sent to the yield pool

### withdrawFee
withdraws accumulated lending fees of the underlying token



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
|`e` | address | address of the underlying token to withdraw

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### calculateFee
this method returns the fee based on the amount passed to it. If the feenominator is 0, then there is no fee.



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
|`a` | uint256 | amount of underlying tokens to calculate the fee for

#### Returns:
| Type | Description |
| --- | --- |
|`uint256` | The total for for the given amount

### scheduleWithdrawal
allows the admin to schedule the withdrawal of tokens



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
|`e` | address | address of (erc20) token to withdraw

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### blockWithdrawal
emergency function to block unplanned withdrawals



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
|`e` | address | address of token withdrawal to block

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### withdraw
allows the admin to withdraw the given token, provided the holding period has been observed



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

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful

### principalToken
retrieves the ERC5095 token for the given market



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
|`address` | of the ERC5095 token for the market

### pause
pauses a market and prevents execution of all lending for that market



#### Declaration
```solidity
function pause(
address u,
uint256 m,
bool b
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
|`u` | address | address of the underlying
|`m` | uint256 | uint256 representing the maturity of the market
|`b` | bool | bool representing whether to pause or unpause

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if successful



## Events

### Lend
emitted upon executed lend




### Mint
emitted upon minted ERC5095 to user




### ScheduleWithdrawal
emitted on token withdrawal scheduling




### BlockWithdrawal
emitted on token withdrawal blocking




