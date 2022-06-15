## `Redeemer`





### `authorized(address a)`






### `constructor(address l, address s, address p, address t, address a)` (public)

Initializes the Redeemer contract




### `setMarketPlaceAddress(address m) → bool` (external)

Sets the address of the marketplace contract which contains the
addresses of all the fixed rate markets




### `setLenderAddress(address l) → bool` (external)





### `redeem(uint8 p, address u, uint256 m, address o) → bool` (public)

Redeems underlying token for illuminate, apwine and tempus
protocols


Illuminate burns its tokens prior to redemption, unlike APWine and
Tempus, which withdraw their tokens after transferring the underlying to
the redeem contract
We can avoid a require check on the principal at the start of this
redeem because there is no common business logic executed before the
protocol specific code is executed.


### `redeem(uint8 p, address u, uint256 m) → bool` (public)



redeem method for swivel, yield, element and notional. This method redeems all


### `redeem(uint8 p, address u, uint256 m, bytes32 i) → bool` (public)

redeem method signature for pendle




### `redeem(uint8 p, address u, uint256 m, address d, address o) → bool` (public)



redeem method signature for sense


### `authRedeem(address u, uint256 m, address f, address t, uint256 a) → bool` (public)

implements the redeem method for the contract to fulfill the
ERC-5095 interface





### `Redeem(uint8 principal, address underlying, uint256 maturity, uint256 amount)`







