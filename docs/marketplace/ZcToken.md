# ZcToken





## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

| Var | Type |
| --- | --- |
| approveCalled | mapping(address => uint256) |
| name | string |
| symbol | string |
| decimals | uint8 |



## Functions

### constructor
No description


#### Declaration
```solidity
function constructor(
address u,
uint256 m,
string n,
string s,
uint8 d
) public
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`u` | address | Underlying
|`m` | uint256 | Maturity
|`n` | string | Name
|`s` | string | Symbol
|`d` | uint8 | Decimals

### approve
No description


#### Declaration
```solidity
function approve(
) public returns
(bool)
```

#### Modifiers:
No modifiers



### approveReturns
No description


#### Declaration
```solidity
function approveReturns(
) public
```

#### Modifiers:
No modifiers



### balanceOf
No description


#### Declaration
```solidity
function balanceOf(
) public returns
(uint256)
```

#### Modifiers:
No modifiers



### balanceOfReturns
No description


#### Declaration
```solidity
function balanceOfReturns(
) public
```

#### Modifiers:
No modifiers





