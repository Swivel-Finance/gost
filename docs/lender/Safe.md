# Safe


Safe ETH and ERC20 transfer library that gracefully handles missing return values.
  @author Modified from Gnosis (https://github.com/gnosis/gp-v2-contracts/blob/main/src/contracts/libraries/GPv2SafeERC20.sol)
  @dev Use with caution! Some functions in this library knowingly create dirty bits at the destination of the free memory pointer.


## Contents
<!-- START doctoc -->
<!-- END doctoc -->




## Functions

### transfer
No description


#### Declaration
```solidity
function transfer(
contract IErc20 e,
address t,
uint256 a
) internal
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`e` | contract IErc20 | Erc20 token to execute the call with
|`t` | address | To address
|`a` | uint256 | Amount being transferred

### transferFrom
No description


#### Declaration
```solidity
function transferFrom(
contract IErc20 e,
address f,
address t,
uint256 a
) internal
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`e` | contract IErc20 | Erc20 token to execute the call with
|`f` | address | From address
|`t` | address | To address
|`a` | uint256 | Amount being transferred

### approve
No description


#### Declaration
```solidity
function approve(
) internal
```

#### Modifiers:
No modifiers





