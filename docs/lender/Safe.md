# 
Safe ETH and ERC20 transfer library that gracefully handles missing return values.
  @author Modified from Gnosis (https://github.com/gnosis/gp-v2-contracts/blob/main/src/contracts/libraries/GPv2SafeERC20.sol)
  @dev Use with caution! Some functions in this library knowingly create dirty bits at the destination of the free memory pointer.


## transfer(contract IErc20,address,uint256)




**Params**
- `e`: Erc20 token to execute the call with

- `t`: To address

- `a`: Amount being transferred

**Returns**

## transferFrom(contract IErc20,address,address,uint256)




**Params**
- `e`: Erc20 token to execute the call with

- `f`: From address

- `t`: To address

- `a`: Amount being transferred

**Returns**

## approve(contract IErc20,address,uint256)




**Params**

**Returns**



