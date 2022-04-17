// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

interface ILidoETH {
    function getPooledEthByShares(uint256 _sharesAmount) external view returns (uint256);
    function submit(address _referral) external returns (uint256);
}