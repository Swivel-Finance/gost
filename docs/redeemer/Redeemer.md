# Redeemer





## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

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
No description


#### Declaration
```solidity
modifier authorized
```



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
|`s` | address | the swivel contract
|`p` | address | the pendle contract
|`t` | address | the tempus contract
|`a` | address | the apwine contract

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
### setLenderAddress
No description


#### Declaration
```solidity
function setLenderAddress(
) external authorized returns
(bool)
```

#### Modifiers:
| Modifier |
| --- |
| authorized |



### redeem
Redeems underlying token for illuminate, apwine and tempus
protocols

> Illuminate burns its tokens prior to redemption, unlike APWine and
Tempus, which withdraw their tokens after transferring the underlying to
the redeem contract. As a result, only Illuminate's redeem returns funds
to the user.
We can avoid a require check on the principal at the start of this
redeem because there is no common business logic executed before the
protocol specific code is executed.


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
No description
> redeem method for swivel, yield, element and notional. This method redeems all


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
redeem method signature for pendle



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
|`i` | bytes32 | forge id used by pendle to redeem the underlying token

#### Returns:
| Type | Description |
| --- | --- |
|`bool` | true if the redemption was successful
### redeem
No description
> redeem method signature for sense


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
|`d` | address | sense contract that splits the loan's prinicpal and yield
|`o` | address | sense contract that [d] calls into to adapt the underlying to sense

### authRedeem
implements the redeem method for the contract to fulfill the
ERC-5095 interface



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
No description




