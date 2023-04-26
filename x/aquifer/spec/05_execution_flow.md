# Execution flow

1. Maintainer deposit QWOYN token - it's configured on params to put only uqwoyn
2. Maintainer set deposit end date (params) - can be set by gov as well
3. Users swap USDC into vested QWOYN token (vesting duration set on params)
4. Maintainer initiate ICA account
5. Maintainer initiate USDC token transfer to ICA account
6. Maintainer transfer QWOYN token to ICA account via IBC transfer
7. Maintainer transfer OSMO tokens to be used as fee to ICA account for pool creation e.g. 100 OSMO
8. Maintainer initiate USDC/QWOYN liquidity pool creation on Osmosis at predetermined price (USDC/QWOYN pool)

Here USDC should be the denom directly transferred from Osmosis network, not USDC received from other network e.g. Axelar.
