// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

interface IAavePool {
  /**
   * @notice Returns the normalized income normalized income of the reserve
   * @param asset The address of the underlying asset of the reserve
   * @return The reserve's normalized income
   */
  function getReserveNormalizedIncome(address asset) external view returns (uint256);

}
