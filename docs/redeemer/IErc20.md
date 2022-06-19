# IERC20





## Contents
<!-- START doctoc -->
<!-- END doctoc -->




## Functions

### totalSupply
No description
> Returns the amount of tokens in existence.

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
> Returns the amount of tokens owned by `account`.

#### Declaration
```solidity
function balanceOf(
) external returns
(uint256)
```

#### Modifiers:
No modifiers



### transfer
No description
> Moves `amount` tokens from the caller's account to `recipient`.

Returns a boolean value indicating whether the operation succeeded.

Emits a {Transfer} event.

#### Declaration
```solidity
function transfer(
) external returns
(bool)
```

#### Modifiers:
No modifiers



### allowance
No description
> Returns the remaining number of tokens that `spender` will be
allowed to spend on behalf of `owner` through {transferFrom}. This is
zero by default.

This value changes when {approve} or {transferFrom} are called.

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
> Sets `amount` as the allowance of `spender` over the caller's tokens.

Returns a boolean value indicating whether the operation succeeded.

IMPORTANT: Beware that changing an allowance with this method brings the risk
that someone may use both the old and the new allowance by unfortunate
transaction ordering. One possible solution to mitigate this race
condition is to first reduce the spender's allowance to 0 and set the
desired value afterwards:
https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729

Emits an {Approval} event.

#### Declaration
```solidity
function approve(
) external returns
(bool)
```

#### Modifiers:
No modifiers



### transferFrom
No description
> Moves `amount` tokens from `sender` to `recipient` using the
allowance mechanism. `amount` is then deducted from the caller's
allowance.

Returns a boolean value indicating whether the operation succeeded.

Emits a {Transfer} event.

#### Declaration
```solidity
function transferFrom(
) external returns
(bool)
```

#### Modifiers:
No modifiers





## Events

### Transfer
No description
> Emitted when `value` tokens are moved from one account (`from`) to
another (`to`).

Note that `value` may be zero.



### Approval
No description
> Emitted when the allowance of a `spender` for an `owner` is set by
a call to {approve}. `value` is the new allowance.



