```sh
qwoynd query game params

# query liquidity & swap rate
qwoynd query game liquidity
qwoynd query game estimated-swap-out 10000ucoho
qwoynd query game swap-rate

# deposit into outpost funding pool
qwoynd tx stimulus deposit-outpost-funding 1000ucoho --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

qwoynd keys add signer --keyring-backend=test --home=$HOME/.qwoynd --recover
# pipe woman clutch absorb lonely cost credit math antique better thumb cook pave clarify hungry east garbage absent warfare song helmet anchor drift purity

# send tokens to admin to add liquidity
qwoynd tx bank send validator $(qwoynd keys show -a signer --keyring-backend=test) 1000000ucoho,1000000uqwoyn,1000000stake --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# add liquidity
qwoynd tx game add-liquidity 1000000uqwoyn,1000000ucoho --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# withdraw from outpost funding pool
qwoynd tx stimulus withdraw-outpost-funding 500stake --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# check outpost funding pool balance
qwoynd query bank balances qwoyn19khjgsdu39ktw8k8ak5ajtruu8cdgmxsc0m2pc
```
