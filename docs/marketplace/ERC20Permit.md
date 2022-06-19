# ERC20Permit



> Extension of {ERC20} that allows token holders to use their tokens
without sending any transactions by setting {IERC20-allowance} with a
signature using the {permit} method, and then spend them via
{IERC20-transferFrom}.

The {permit} signature mechanism conforms to the {IERC2612} interface.

## Contents
<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this contains internal vars as well due to a bug in the docgen procedure

| Var | Type |
| --- | --- |
| nonces | mapping(address => uint256) |
| PERMIT_TYPEHASH | bytes32 |
| deploymentChainId | uint256 |



## Functions

### constructor
No description


#### Declaration
```solidity
function constructor(
) internal ERC20
```

#### Modifiers:
| Modifier |
| --- |
| ERC20 |



### DOMAIN_SEPARATOR
No description
> Return the DOMAIN_SEPARATOR.

#### Declaration
```solidity
function DOMAIN_SEPARATOR(
) external returns
(bytes32)
```

#### Modifiers:
No modifiers



### version
No description
> Setting the version as a function so that it can be overriden

#### Declaration
```solidity
function version(
) public returns
(string)
```

#### Modifiers:
No modifiers



### permit
No description
> See {IERC2612-permit}.

In cases where the free option is not a concern, deadline can simply be
set to uint(-1), so it should be seen as an optional parameter

#### Declaration
```solidity
function permit(
) external
```

#### Modifiers:
No modifiers





