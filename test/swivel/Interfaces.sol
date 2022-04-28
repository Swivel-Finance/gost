// SPDX-License-Identifier: UNLICENSED

pragma solidity >= 0.8.4;

interface Erc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IERC20 is Erc20 {
  function totalSupply() external view returns (uint256);
}

interface CErc20 {
	function mint(uint256) external returns (uint256);
	function redeemUnderlying(uint256) external returns (uint256);
}

interface MarketPlace {

  struct Market {
    address cToken;
    address zcToken;
    address vault;
    address adapter;
    uint256 maturityRate;
  }

  function markets (uint8, address, uint256) external view returns (Market memory);

  // adds notional and mints zctokens
  function mintZcTokenAddingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  // removes notional and burns zctokens
  function burnZcTokenRemovingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  // returns the amount of underlying interest to send
  function redeemVaultInterest(uint8, address, uint256, address) external returns (uint256);
  // returns the amount of underlying principal to send
  function redeemZcToken(uint8, address, uint256, address, uint256) external returns (uint256);
  // returns the cToken address for a given market
  function cTokenAddress(uint8, address, uint256) external returns (address);
  // EVFZE FF EZFVE call this which would then burn zctoken and remove notional
  function custodialExit(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IZFVI call this which would then mint zctoken and add notional
  function custodialInitiate(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IZFZE && EZFZI call this, tranferring zctoken from one party to another
  function p2pZcTokenExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFVE && EVFVI call this, removing notional from one party and adding to the other
  function p2pVaultExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IVFVE call this which then transfers notional from msg.sender (taker) to swivel
  function transferVaultNotionalFee(uint8, address, uint256, address, uint256) external returns (bool);
}

interface IYearnVault {
    function token() external view returns (address);

    function underlying() external view returns (address);

    function pricePerShare() external view returns (uint256);

    function deposit(uint256) external;

    function withdraw(uint256) external;
}

interface IAavePool {
  /**
   * @notice Returns the normalized income normalized income of the reserve
   * @param asset The address of the underlying asset of the reserve
   * @return The reserve's normalized income
   */
  function getReserveNormalizedIncome(address asset) external view returns (uint256);
  /**
   * @dev Emitted on deposit()
   * @param asset The address of the underlying asset of the reserve
   * @param amount The amount deposited
   * @param onBehalfOf The beneficiary of the deposit, receiving the aTokens
   * @param referralCode The referral code used
   **/
  function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) external;

  function withdraw(address asset, uint256 amount, address to) external;
}

interface IEulerToken {    
    function deposit(uint subAccountId, uint amount) external;
    function withdraw(uint subAccountId, uint amount) external;
}

interface IERC4626 is IERC20 {

    /// @notice Mints `shares` IN ASSETS Vault shares to `receiver` by
    /// depositing exactly `assets` of underlying tokens.
    function deposit(uint256 assets, address receiver) external returns (uint256 shares);

    /// @notice Mints exactly `shares` IN SHARES Vault shares to `receiver`
    /// by depositing `assets` of underlying tokens.
    function mint(uint256 shares, address receiver) external returns (uint256 assets);

    /// @notice Redeems `shares` IN ASSETS from `owner` and sends `assets`
    /// of underlying tokens to `receiver`.
    function withdraw(uint256 assets, address receiver, address owner) external returns (uint256 shares);

    /// @notice Redeems `shares` IN SHARES from `owner` and sends `assets`
    /// of underlying tokens to `receiver`.
    function redeem(uint256 shares, address receiver, address owner) external returns (uint256 assets);

    /*////////////////////////////////////////////////////////
                      Vault Accounting Logic
    ////////////////////////////////////////////////////////*/

    /// @notice The amount of shares that the vault would
    /// exchange for the amount of assets provided, in an
    /// ideal scenario where all the conditions are met.
    function convertToShares(uint256 assets) external view returns (uint256 shares);

    /// @notice The amount of assets that the vault would
    /// exchange for the amount of shares provided, in an
    /// ideal scenario where all the conditions are met.
    function convertToAssets(uint256 shares) external view returns (uint256 assets);
}
