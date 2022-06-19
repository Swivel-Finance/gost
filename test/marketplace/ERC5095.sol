// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./Interfaces.sol";
import "./ERC20Permit.sol";

contract ERC5095 is ERC20Permit, IERC5095 {
    /// @dev unix timestamp when the ERC5095 token can be redeemed
    uint256 public override immutable maturity;
    /// @dev address of the ERC20 token that is returned on ERC5095 redemption
    address public override immutable underlying;
    /// @dev address of the minting authority
    address public immutable lender;
    
    /////////////OPTIONAL///////////////// (Allows the calculation and distribution of yield post maturity)
    /// @dev address and interface for an external custody contract (necessary for some project's backwards compatability)
    address public immutable redeemer;

    error Maturity(uint256 timestamp);  

    constructor(address _underlying, uint256 _maturity, address _redeemer, address l, string memory name_, string memory symbol_, uint8 decimals_) ERC20Permit(name_, symbol_, decimals_) {
        underlying = _underlying;
        maturity = _maturity;
        redeemer = _redeemer;
        lender = l;
    }

    /// @notice Post maturity converts an amount of principal tokens to an amount of underlying that would be returned. Returns 0 pre-maturity.
    /// @param principalAmount The amount of principal tokens to convert
    /// @return underlyingAmount The amount of underlying tokens returned by the conversion
    function convertToUnderlying(uint256 principalAmount) external override view returns (uint256 underlyingAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (principalAmount);
    }

    /// @notice Post maturity converts a desired amount of underlying tokens returned to principal tokens needed. Returns 0 pre-maturity.
    /// @param underlyingAmount The amount of underlying tokens to convert
    /// @return principalAmount The amount of principal tokens returned by the conversion
    function convertToPrincipal(uint256 underlyingAmount) external override view returns (uint256 principalAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (underlyingAmount);
    }

    /// @notice Post maturity calculates the amount of principal tokens that `owner` can redeem. Returns 0 pre-maturity.
    /// @param owner The address of the owner for which redemption is calculated
    /// @return maxPrincipalAmount The maximum amount of principal tokens that `owner` can redeem.
    function maxRedeem(address owner) external override view returns (uint256 maxPrincipalAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (_balanceOf[owner]);
    }

    /// @notice Post maturity simulates the effects of redeemption at the current block. Returns 0 pre-maturity.
    /// @param principalAmount the amount of principal tokens redeemed in the simulation
    /// @return underlyingAmount The maximum amount of underlying returned by `principalAmount` of PT redemption
    function previewRedeem(uint256 principalAmount) external override view returns (uint256 underlyingAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (principalAmount);
    }

    /// @notice Post maturity calculates the amount of underlying tokens that `owner` can withdraw. Returns 0 pre-maturity.
    /// @param  owner The address of the owner for which withdrawal is calculated
    /// @return maxUnderlyingAmount The maximum amount of underlying tokens that `owner` can withdraw.
    function maxWithdraw(address owner) external override view returns (uint256 maxUnderlyingAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (_balanceOf[owner]);
    }

    /// @notice Post maturity simulates the effects of withdrawal at the current block. Returns 0 pre-maturity.
    /// @param underlyingAmount the amount of underlying tokens withdrawn in the simulation
    /// @return principalAmount The amount of principal tokens required for the withdrawal of `underlyingAmount`
    function previewWithdraw(uint256 underlyingAmount) external override view returns (uint256 principalAmount){
        if (block.timestamp < maturity) {
            return 0;
        }
        return (underlyingAmount);
    }

    /// @notice At or after maturity, Burns principalAmount from `owner` and sends exactly `underlyingAmount` of underlying tokens to `receiver`.
    /// @param underlyingAmount The amount of underlying tokens withdrawn
    /// @param receiver The receiver of the underlying tokens being withdrawn
    /// @return principalAmount The amount of principal tokens burnt by the withdrawal
    function withdraw(uint256 underlyingAmount, address receiver, address holder) external override returns (uint256 principalAmount){
        if (block.timestamp < maturity) {
            revert Maturity(maturity);
        }
        if (holder == msg.sender) {
            return IRedeemer(redeemer).authRedeem(underlying, maturity, msg.sender, receiver, underlyingAmount);
        }
        else {
            require(_allowance[holder][msg.sender] >= underlyingAmount, 'not enough approvals');
            return IRedeemer(redeemer).authRedeem(underlying, maturity, holder, receiver, underlyingAmount);     
        }
    }

    /// @notice At or after maturity, burns exactly `principalAmount` of Principal Tokens from `owner` and sends underlyingAmount of underlying tokens to `receiver`.
    /// @param receiver The receiver of the underlying tokens being withdrawn
    /// @return underlyingAmount The amount of underlying tokens distributed by the redemption
    function redeem(uint256 principalAmount, address receiver, address holder) external override returns (uint256 underlyingAmount){
        if (block.timestamp < maturity) {
            revert Maturity(maturity);
        }
        if (holder == msg.sender) {
            return IRedeemer(redeemer).authRedeem(underlying, maturity, msg.sender, receiver, principalAmount);
        }
        else {
            require(_allowance[holder][msg.sender] >= underlyingAmount, 'not enough approvals');
            return IRedeemer(redeemer).authRedeem(underlying, maturity, holder, receiver, principalAmount);     
        }
    }

    /// @param f Address to burn from
    /// @param a Amount to burn
    /// @return bool true if successful
    function burn(address f, uint256 a) external onlyAdmin(redeemer) returns (bool) {
        _burn(f, a);
        return true;
    }

    /// @param t Address recieving the minted amount
    /// @param a The amount to mint
    /// @return bool true if successful
    function mint(address t, uint256 a) external onlyAdmin(lender) returns (bool) {
        _mint(t, a);
        return true;
    }

    /// @param a Admin address
    modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
    }
}
