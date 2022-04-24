// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

interface Erc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface CErc20 {
	function mint(uint256) external returns (uint256);
	function redeemUnderlying(uint256) external returns (uint256);
}

interface MarketPlace {
  // adds notional and mints zctokens
  function mintZcTokenAddingNotional(address, uint256, address, uint256) external returns (bool);
  // removes notional and burns zctokens
  function burnZcTokenRemovingNotional(address, uint256, address, uint256) external returns (bool);
  // returns the amount of underlying principal to send
  function redeemZcToken(address, uint256, address, uint256) external returns (uint256);
  // returns the amount of underlying interest to send
  function redeemVaultInterest(address, uint256, address) external returns (uint256);
  // returns the cToken address for a given market
  function cTokenAddress(address, uint256) external returns (address);
  // EVFZE FF EZFVE call this which would then burn zctoken and remove notional
  function custodialExit(address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IZFVI call this which would then mint zctoken and add notional
  function custodialInitiate(address, uint256, address, address, uint256) external returns (bool);
  // IZFZE && EZFZI call this, tranferring zctoken from one party to another
  function p2pZcTokenExchange(address, uint256, address, address, uint256) external returns (bool);
  // IVFVE && EVFVI call this, removing notional from one party and adding to the other
  function p2pVaultExchange(address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IVFVE call this which then transfers notional from msg.sender (taker) to swivel
  function transferVaultNotionalFee(address, uint256, address, uint256) external returns (bool);
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