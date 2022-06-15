## `MarketPlace`





### `authorized(address a)`





### `unpaused(uint8 p)`






### `constructor(address r)` (public)





### `createMarket(address u, uint256 m, address[8] t, string n, string s, uint8 d) → bool` (external)

creates a new market for the given underlying token and maturity
all seven principals should be provided in the order of their enum value




### `pause(uint8 p, bool s) → bool` (external)

allows the admin to pause a principal's pools




### `setPool(uint8 p, address u, uint256 m, address a) → bool` (external)

sets the address for a pool




### `sellPrincipalToken(uint8 p, address u, uint256 m, uint128 a) → uint128` (external)

sells the PT for the PT via the pool




### `buyPrincipalToken(uint8 p, address u, uint256 m, uint128 a) → uint128` (external)

buys the underlying for the PT via the pool




### `sellUnderlying(uint8 p, address u, uint256 m, uint128 a) → uint128` (external)

sells the underlying for the PT via the pool




### `buyUnderlying(uint8 p, address u, uint256 m, uint128 a) → uint128` (external)

buys the underlying for the PT via the pool





### `CreateMarket(address underlying, uint256 maturity)`







### `Principals`





























