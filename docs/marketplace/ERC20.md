# ERC20



> Implementation of the {IERC20} interface.

This implementation is agnostic to the way tokens are created. This means
that a supply mechanism has to be added in a derived contract using {_mint}.

We have followed general OpenZeppelin guidelines: functions revert instead
of returning `false` on failure. This behavior is nonetheless conventional
and does not conflict with the expectations of ERC20 applications.

Additionally, an {Approval} event is emitted on calls to {transferFrom}.
This allows applications to reconstruct the allowance for all accounts just
by listening to said events. Other implementations of the EIP may not emit
these events, as it isn't required by the specification.

Calls to {transferFrom} do not check for allowance if the caller is the owner
of the funds. This allows to reduce the number of approvals that are necessary.

Finally, {transferFrom} does not decrease the allowance if it is set to
type(uint256).max. This reduces the gas costs without any likely impact.

## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

| Var | Type |
| --- | --- |
| _totalSupply | uint256 |
| _balanceOf | mapping(address => uint256) |
| _allowance | mapping(address => mapping(address => uint256)) |
| name | string |
| symbol | string |
| decimals | uint8 |



## Functions

### constructor
 @dev Sets the values for {name}, {symbol} and {decimals}.


#### Declaration
```solidity
function constructor(
) public
```

#### Modifiers:
No modifiers



### totalSupply
No description
> See {IERC20-totalSupply}.

#### Declaration
```solidity
function totalSupply(
) external returns
(uint256)
```

#### Modifiers:
No modifiers



### balanceOf
No description
> See {IERC20-balanceOf}.

#### Declaration
```solidity
function balanceOf(
) external returns
(uint256)
```

#### Modifiers:
No modifiers



### allowance
No description
> See {IERC20-allowance}.

#### Declaration
```solidity
function allowance(
) external returns
(uint256)
```

#### Modifiers:
No modifiers



### approve
No description
> See {IERC20-approve}.

#### Declaration
```solidity
function approve(
) external returns
(bool)
```

#### Modifiers:
No modifiers



### transfer
No description
> See {IERC20-transfer}.

Requirements:

- the caller must have a balance of at least `wad`.

#### Declaration
```solidity
function transfer(
) external returns
(bool)
```

#### Modifiers:
No modifiers



### transferFrom
if_succeeds {:msg "TransferFrom - decrease allowance"} msg.sender != src ==> old(_allowance[src][msg.sender]) >= wad;


#### Declaration
```solidity
function transferFrom(
) external returns
(bool)
```

#### Modifiers:
No modifiers



### _transfer
if_succeeds {:msg "Transfer - src decrease"} old(_balanceOf[src]) >= _balanceOf[src];
if_succeeds {:msg "Transfer - dst increase"} _balanceOf[dst] >= old(_balanceOf[dst]);
if_succeeds {:msg "Transfer - supply"} old(_balanceOf[src]) + old(_balanceOf[dst]) == _balanceOf[src] + _balanceOf[dst];


#### Declaration
```solidity
function _transfer(
) internal returns
(bool)
```

#### Modifiers:
No modifiers



### _setAllowance
No description
> Sets the allowance granted to `spender` by `owner`.

Emits an {Approval} event indicating the updated allowance.

#### Declaration
```solidity
function _setAllowance(
) internal returns
(bool)
```

#### Modifiers:
No modifiers



### _decreaseAllowance
if_succeeds {:msg "Decrease allowance - underflow"} old(_allowance[src][msg.sender]) <= _allowance[src][msg.sender];


#### Declaration
```solidity
function _decreaseAllowance(
) internal returns
(bool)
```

#### Modifiers:
No modifiers



### _mint
if_succeeds {:msg "Mint - balance overflow"} old(_balanceOf[dst]) >= _balanceOf[dst];
if_succeeds {:msg "Mint - supply overflow"} old(_totalSupply) >= _totalSupply;


#### Declaration
```solidity
function _mint(
) internal returns
(bool)
```

#### Modifiers:
No modifiers



### _burn
if_succeeds {:msg "Burn - balance underflow"} old(_balanceOf[src]) <= _balanceOf[src];
if_succeeds {:msg "Burn - supply underflow"} old(_totalSupply) <= _totalSupply;


#### Declaration
```solidity
function _burn(
) internal returns
(bool)
```

#### Modifiers:
No modifiers





