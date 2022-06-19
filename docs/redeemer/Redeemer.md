# Redeemer


The Redeemer contract is used to redeem the underlying lent capital of a loan.
Users may redeem their ERC-5095 tokens for the underlying asset represented by that token after maturity.


## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var | Type |
| --- | --- |
| admin | address |
| marketPlace | address |
| lender | address |
| swivelAddr | address |
| pendleAddr | address |
| tempusAddr | address |
| apwineAddr | address |


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
Initializes the Redeemer contract



#### Declaration
```solidity
function constructor(
address l,
address s,
address p,
address t,
address a
) public
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`l` | address | the lender contract
|`s` | address | the Swivel contract
|`p` | address | the Pendle contract
|`t` | address | the Tempus contract
|`a` | address | the APWine contract

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

### setLenderAddress
sets the address of the lender contract which contains the addresses of all the fixed rate markets



#### Declaration
```solidity
function setLenderAddress(
address l
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
|`l` | address | the address of the lender contract

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

### redeem
redeems underlying token for Illuminate, APWine and Tempus protocols

> Illuminate burns its tokens prior to redemption, unlike APWine and
Tempus, which redeem PTs to the redeemer, transferring the underlying to
this redeemer contract. Consequently, only Illuminate's redeem returns funds
to the user.


#### Declaration
```solidity
function redeem(
uint8 p,
address u,
uint256 m,
address o
) public returns
(bool)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | the underlying token being redeemed
|`m` | uint256 | the maturity of the market being redeemed
|`o` | address | address of the controller or contract that manages the underlying token

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the redemption was successful

### redeem
redeem method for Swivel, Yield, Element and Notional protocols



#### Declaration
```solidity
function redeem(
uint8 p,
address u,
uint256 m
) public returns
(bool)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | underlying token being redeemed
|`m` | uint256 | maturity of the market being redeemed

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the redemption was successful

### redeem
redeem method signature for Pendle



#### Declaration
```solidity
function redeem(
uint8 p,
address u,
uint256 m,
bytes32 i
) public returns
(bool)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | underlying token being redeemed
|`m` | uint256 | maturity of the market being redeemed
|`i` | bytes32 | forge id used by Pendle to redeem the underlying token

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the redemption was successful

### redeem
redeem method signature for Sense



#### Declaration
```solidity
function redeem(
uint8 p,
address u,
uint256 m,
address d,
address o
) public returns
(bool)
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`p` | uint8 | value of a specific principal according to the Illuminate Principals Enum
|`u` | address | underlying token being redeemed
|`m` | uint256 | maturity of the market being redeemed
|`d` | address | Sense contract that splits the loan's prinicpal and yield
|`o` | address | Sense contract that [d] calls into to adapt the underlying to Sense

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the redemption was successful

### authRedeem
implements the redeem method for the contract to fulfill the ERC-5095 interface



#### Declaration
```solidity
function authRedeem(
address u,
uint256 m,
address f,
address t,
uint256 a
) public authorized returns
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
|`m` | uint256 | maturity of the market
|`f` | address | address from where the underlying asset will be burned
|`t` | address | address to where the underlying asset will be transferred
|`a` | uint256 | amount of the underlying asset to be burned and sent to the to

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the underlying asset was burned successfully



## Events

### Redeem
emitted upon redemption of a loan




