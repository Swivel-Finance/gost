// SPDX-License-Identifier: MIT
pragma solidity >= 0.8.0;
import "./IERC20.sol";
import "./IERC2612.sol";
import "./IPT.sol";


interface IPool is IERC20, IERC2612 {
    function ts() external view returns(int128);
    function g1() external view returns(int128);
    function g2() external view returns(int128);
    function maturity() external view returns(uint32);
    function scaleFactor() external view returns(uint96);
    function getCache() external view returns (uint112, uint112, uint32);
    function underlying() external view returns(IERC20);
    function PT() external view returns(IPT);
    function getunderlyingBalance() external view returns(uint112);
    function getPTBalance() external view returns(uint112);
    function retrieveunderlying(address to) external returns(uint128 retrieved);
    function retrievePT(address to) external returns(uint128 retrieved);
    function sellunderlying(address to, uint128 min) external returns(uint128);
    function buyunderlying(address to, uint128 underlyingOut, uint128 max) external returns(uint128);
    function sellPT(address to, uint128 min) external returns(uint128);
    function buyPT(address to, uint128 PTOut, uint128 max) external returns(uint128);
    function sellunderlyingPreview(uint128 underlyingIn) external view returns(uint128);
    function buyunderlyingPreview(uint128 underlyingOut) external view returns(uint128);
    function sellPTPreview(uint128 PTIn) external view returns(uint128);
    function buyPTPreview(uint128 PTOut) external view returns(uint128);
    function mint(address to, address remainder, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256);
    function mintWithunderlying(address to, address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256);
    function burn(address underlyingTo, address PTTo, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256);
    function burnForunderlying(address to, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256);
    function cumulativeBalancesRatio() external view returns (uint256);
}