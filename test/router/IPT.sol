// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "./IERC20.sol";
import "./IJoin.sol";

interface IPT is IERC20 {
    /// @dev Asset that is returned on redemption.
    function underlying() external view returns (address);

    /// @dev Source of redemption funds.
    function join() external view returns (IJoin);

    /// @dev Unix time at which redemption of PT for underlying are possible
    function maturity() external view returns (uint256);

    /// @dev Record price data at maturity
    function mature() external;

    /// @dev Mint PT providing an equal amount of underlying to the protocol
    function mintWithUnderlying(address to, uint256 amount) external;

    /// @dev Burn PT after maturity for an amount of underlying.
    function redeem(address to, uint256 amount) external returns (uint256);

    /// @dev Mint PT.
    /// This function can only be called by other Yield contracts, not users directly.
    /// @param to Wallet to mint the PT in.
    /// @param PTAmount Amount of PT to mint.
    function mint(address to, uint256 PTAmount) external;

    /// @dev Burn PT.
    /// This function can only be called by other Yield contracts, not users directly.
    /// @param from Wallet to burn the PT from.
    /// @param PTAmount Amount of PT to burn.
    function burn(address from, uint256 PTAmount) external;
}