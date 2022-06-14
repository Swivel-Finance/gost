// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IPErc20 {
    function balanceOf(address) external view returns (uint256);

    function transfer(address, uint256) external returns (bool);

    function allowance(address, address) external view returns (uint256);

    function approve(address, uint256) external returns (bool);

    function transferFrom(
        address,
        address,
        uint256
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

interface IZcToken is IPErc20 {
    function burn(address, uint256) external returns (bool);

    function mint(address, uint256) external returns (bool);
}

interface IErc2612 {
    function permit(
        address,
        address,
        uint256,
        uint256,
        uint8,
        bytes32,
        bytes32
    ) external;

    function nonces(address) external view returns (uint256);
}
