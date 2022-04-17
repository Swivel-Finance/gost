// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

interface IYearnVault {
    function token() external view returns (address);

    function underlying() external view returns (address);

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function decimals() external view returns (uint8);

    function controller() external view returns (address);

    function governance() external view returns (address);

    function pricePerShare() external view returns (uint256);

    function deposit(uint256) external;

    function depositAll() external;

    function withdraw(uint256) external;

    function withdrawAll() external;
}