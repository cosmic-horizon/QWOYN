# Aquifer module

The Qwoyn Aquifer is a revolutionary liquidity bootstrapping tool designed by Qwoyn. This decentralized method utilizes completely on-chain methods to create a liquidity pool for a new token. The Aquifer represents the source of liquidity for this new token, providing a steady flow of liquidity to support its growth. By utilizing The Qwoyn Aquifer, token creators can launch their projects with confidence, knowing that they have access to a reliable and sustainable source of liquidity. The Qwoyn Aquifer is a game-changer for the DeFi space, offering a secure and efficient way to launch new tokens and create a thriving ecosystem.
"Aquifer" - This name is a play on the word "aquarium," which represents a contained ecosystem. In this case, the ecosystem is the liquidity pool, which is created by your module.

## Params

- DepositToken: USDC
- AllocationToken: QWOYN
- VestingConfig: 1 year
- DepositEndTime: After 1 month
- InitialLiquidityPrice: 10 (10 USDC / QWOYN)
- LiquidityBootstrapPending: broadcasted ICA tx to bootstrap Osmosis LP
- LiquidityBootrapped: set as true once Osmosis liquidity created
- ICA account for the operation on Osmosis

## Endpoints

- PutAllocationTokenByAdmin (QWOYN)
- BuyAllocationToken (USDC -> QWOYN)
- SetDepositEndTime
- Initiate ICA account (auto?)
- Execute Add Liquidity (auto?)
