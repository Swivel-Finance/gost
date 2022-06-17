# Flow

This document describes the flow of assets in the process of lending on Illuminate. In this example, we will have the following assumptions:

- The underlying asset is USDC
- The loan term is for 1 year
- The protocol used by Illuminate is Notional
- The size of the principal of the loan is 1000 USDC
- The interest rate over the term of the loan is 5%
- The user is referred to as Alice

## Steps

The following describes the transactions that occur when a loan is made and redeemed on Illuminate.

1. Alice executes a one-year term loan of 1000 USDC. At the time she executes loan, the going interest rate for borrowing USDC is 5%.
2. Illuminate receives the `lend()` calls and determines that Notional is offering the best rate.
3. The underlying USDC is transferred to the lender contract
4. The lender deposits the USDC into Notional. At an interest rate of 5%, this should return 1050 as the amount redeemable at expiry.
5. Illuminate mints 1050 zcTokens to Alice's address (`msg.sender`).
6. One year passes.
7. Illuminate executes Notional's `redeem()` on the one-year, USDC market.
8. Illuminate transfers all Notional principal tokens from the lender contract to the redeemer contract.
9. Illuminate redeems all owed USDC tokens from Notional.
10. Alice now executes Illuminate's `redeem` with her mature zcTokens on the one-year USDC market.
11. Redeemer burns the 1050 zcTokens held by Alice.
12. Redeemer sends 1050 USDC to Alice.

**Enter visualization here**