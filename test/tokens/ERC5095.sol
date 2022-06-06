// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import "./PErc20.sol";
import "./IAdapter.sol"; 
import "./IERC5095.sol"; 
import "./IRedeemer.sol";

abstract contract ERC5095 is PErc20, IERC5095 {

    uint256 public override immutable maturity;
    address public override immutable underlying;
    uint256 public override maturityRate;
    IRedeemer public immutable redeemer;
    IAdapter public adapter;
    address public cToken;

    error Maturity(uint256 maturity);  

    constructor(address u, uint256 m, address a, address mp) {
        underlying = u;
        maturity = m;
        adapter = IAdapter(a);
        redeemer = IRedeemer(mp);
    }

    function convertToUnderlying(uint256 principalAmount) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (principalAmount * (adapter.exchangeRateCurrent(cToken) / maturityRate));
    }

    function convertToPrincipal(uint256 underlyingAmount) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return(underlyingAmount * maturityRate / adapter.exchangeRateCurrent(cToken));
    }

    function maxRedeem(address owner) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return(_balanceOf[owner]);
    }

    function previewRedeem(uint256 principalAmount) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return(principalAmount * (adapter.exchangeRateCurrent(cToken) / maturityRate));
    }

    function maxWithdraw(address owner) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return(_balanceOf[owner] * (adapter.exchangeRateCurrent(cToken) / maturityRate));
    }

    function previewWithdraw(uint256 underlyingAmount) external view returns (uint256){
        if (block.timestamp < maturity) {
            return 0;
        }
        return(underlyingAmount * maturityRate / adapter.exchangeRateCurrent(cToken));
    }

    function withdraw(uint256 underlyingAmount, address owner, address receiver) external returns (uint256){
        if (maturityRate == 0) {
            if (block.timestamp < maturity) {
                revert Maturity(maturity);
            }
            maturityRate = adapter.exchangeRateCurrent(cToken);
            return redeemer.adminRedeem(underlying, maturity, owner, receiver, underlyingAmount);
        }
        return redeemer.adminRedeem(underlying, maturity, owner, receiver, (underlyingAmount * maturityRate / adapter.exchangeRateCurrent(cToken)));
    }

    function redeem(uint256 principalAmount, address owner, address receiver) external returns (uint256){
        if (maturityRate == 0) {
            if (block.timestamp < maturity) {
                revert Maturity(maturity);
            }
            maturityRate = adapter.exchangeRateCurrent(cToken);
        }
        return redeemer.adminRedeem(underlying, maturity, owner, receiver, principalAmount);
    }
}