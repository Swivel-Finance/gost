# ZcTokens

ERC-20 Tokens that represent a 1-1 claim for underlying upon maturity.

## Overview

Zero-coupon tokens, **zcTokens**, are ERC-20's which are redeemable 1:1 at maturity for an underlying token. 

When users make loans, their tokens deposit is sent to the lender contract and an equivalent number of zcTokens are minted to their account via `msg.sender`.

This leaves the protocol's principal token in the custody of the lender contract and the zcTokens in the possession of the user.

### Fixed-Yield Lending:
Alice has 1000 USDC.
Alice lends 1000 USDC via Illuminate at a rate of 5% over the term of the loan.
Alice receives 1050 zcUSDC while Illuminate lends out the 1050 USDC and receives 1050 ptUSDC.
Alice then has 1050 zcUSDC. At maturity Alice redeems her 1050 zcUSDC for 1050 USDC after maturity.

## Pricing

Given zcTokens represent a 1-1 redemption only upon maturity, they are discounted at a rate based on the predicted amount of interest that may have otherwise been generated until the maturity/redemption date (the nToken price).

Accepting that this potential yield decreases as time passes, ZcTokens appreciate towards par and reach par as maturity is reached.

Further, because an nToken represents the predicted future yield, the discount on a zcToken is inversely proportional to the cost of an nToken.

At maturity, zcTokens begin to appreciate above par, accruing the yield generated on a given money-market (e.g. Compound) until redeemed.
