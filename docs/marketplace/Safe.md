# Safe


Safe ETH and ERC20 transfer library that gracefully handles missing return values.
  @author Modified from Gnosis (https://github.com/gnosis/gp-v2-contracts/blob/main/src/contracts/libraries/GPv2SafeERC20.sol)
  @dev Use with caution! Some functions in this library knowingly create dirty bits at the destination of the free memory pointer.


## Contents
<!-- START doctoc -->
<!-- END doctoc -->




## Functions

### approve
No description


#### Declaration
```solidity
function approve(
contract IERC20 e,
address t,
uint256 a
) internal
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`e` | contract IERC20 | Erc20 token to execute the call with
|`t` | address | To address
|`a` | uint256 | Amount being transferred

### transfer
No description


#### Declaration
```solidity
function transfer(
contract IERC20 e,
address t,
uint256 a
) internal
```

#### Modifiers:
No modifiers

#### Args:
| Arg | Type | Description |
| --- | --- | --- |
|`e` | contract IERC20 | Erc20 token to execute the call with
|`t` | address | To address
|`a` | uint256 | Amount being transferred



