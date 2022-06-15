## `Lender`





### `authorized(address a)`






### `constructor(address s, address p, address t)` (public)

Initializes the Lender contract


the ctor sets a default value for the feenominator


### `approve(address[] u, address[] a) → bool` (external)

Bulk approves the usage of addresses at the given ERC20 addresses
The lengths of the inputs must match because the arrays are paired by index




### `setFee(uint256 f) → bool` (external)

Sets the feenominator to the given value




### `setMarketPlaceAddress(address m) → bool` (external)

Sets the address of the marketplace contract which contains the
addresses of all the fixed rate markets




### `mint(uint8 p, address u, uint256 m, uint256 a) → bool` (public)

mint swaps the sender's principal tokens for illuminate's zc tokens
in effect, this opens a new fixed rate position for the sender on illuminate


mint is uniform across all principals, thus there is no need for a 'minter'


### `lend(uint8 p, address u, uint256 m, uint256 a, address y) → uint256` (public)



lend method signature for both illuminate and yield


### `lend(uint8 p, address u, uint256 m, uint256[] a, address y, struct Swivel.Order[] o, struct Swivel.Components[] s) → uint256` (public)

lends to yield pool. remaining balance is sent to the yield pool
TODO: this will change when we implement a check on the gas market


lend method signature for swivel


### `lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, uint256 d, address e, bytes32 i) → uint256` (public)



lend method signature for element


### `lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, uint256 d) → uint256` (public)



lend method signature for pendle


### `lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, uint256 d, address t, address x) → uint256` (public)

Can be called before maturity to lend to Tempus while minting Illuminate tokens


lend method signature for tempus


### `lend(uint8 p, address u, uint256 m, uint128 a, uint256 r, address x, address s) → uint256` (public)

Can be called before maturity to lend to Sense while minting Illuminate tokens


lend method signature for sense


### `lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address pool, address aave, uint256 i) → uint256` (public)

Can be called before maturity to lend to APWine while minting Illuminate tokens




### `lend(uint8 p, address u, uint256 m, uint256 a) → uint256` (public)



lend method signature for Notional


### `yield(address u, address y, uint256 a) → uint256` (internal)

transfers excess funds to yield pool after principal tokens have been lent out


this method is only used by the yield, illuminate and swivel protocols


### `withdraw(address e) → bool` (external)

Withdraws accumulated lending fees of the underlying token




### `calculateFee(uint256 a) → uint256` (internal)

This method returns the fee based on the amount passed to it. If
feenominator is 0, then there is no fee.




### `zcToken(address u, uint256 m) → address` (internal)

retrieves the zc token for the given market





### `Lend(uint8 principal, address underlying, uint256 maturity, uint256 returned)`





### `Mint(uint8 principal, address underlying, uint256 maturity, uint256 amount)`







