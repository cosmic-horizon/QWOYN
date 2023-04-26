# CLI examples for game

```sh
qwoynd query game params
qwoynd query game whitelisted-contracts
qwoynd query game all-deposit-balances
qwoynd query game deposit-balance $(qwoynd keys show -a validator --keyring-backend=test)
qwoynd query game all-unbondings
qwoynd query game user-unbondings $(qwoynd keys show -a validator --keyring-backend=test)

# query liquidity & swap rate
qwoynd query game liquidity
qwoynd query game estimated-swap-out 10000ucoho
qwoynd query game estimated-swap-out 10000uqwoyn
qwoynd query game swap-rate

# deposit tokens
qwoynd tx game deposit-token 1000000stake --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# withdraw tokens
qwoynd tx game withdraw-token 500000stake --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# stake ingame token
qwoynd tx game stake-ingame-token 500000stake --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# begin unstake of ingame token
qwoynd tx game begin-unstake-ingame-token 100000stake --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# claim staking reward
qwoynd tx game claim-ingame-staking-reward --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# whitelist contract
qwoynd tx game whitelist-contracts coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# remove whitelisted contract
qwoynd tx game remove-whitelisted-contracts coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# deposit nft from end wallet to game
qwoynd tx game deposit-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# query owner of nft
qwoynd query wasm contract-state smart coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"owner_of":{"token_id":"1"}}'

# game module address
# coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk

# add signer
qwoynd keys add signer --keyring-backend=test --home=$HOME/.qwoynd --recover
# pipe woman clutch absorb lonely cost credit math antique better thumb cook pave clarify hungry east garbage absent warfare song helmet anchor drift purity

qwoynd tx bank send validator $(qwoynd keys show -a signer --keyring-backend=test --home=$HOME/.qwoynd) 1000000stake --keyring-backend=test --home=$HOME/.qwoynd --broadcast-mode=block -y --chain-id=qwoyn-1

# upload wasm
qwoynd tx wasm store ./artifacts/ship_nft.wasm --from validator --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y

# instantiate ship nft with module account owned
qwoynd tx wasm instantiate 1 '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk","owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk"}' --from validator --label "Ship-NFT" --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ --no-admin -y

# transfer nft update ownership to module account - when owner is set to different account
qwoynd tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"transfer_ownership":{"owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk"}}' --from signer --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y


qwoynd tx game sign-withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/

qwoynd tx game withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' 42d6e9d3b62ffc9b0bc3f6a97cbc0857af1c7a7aa57549571d7bc72415a955d978a1790440ce53c8f9fbfa2ce70d967812eda6094d6f112d7e5736170e48e2a8 --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block --gas=400000

qwoynd tx game withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 2 '{"mint":{"token_id":"2","owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk","extension":{"ship_type":12,"owner":"300"}}}' 6a859feb353f0c49b9316cf871738f6d21a14320edb4817cdbfbcab7c6434cb10ab1b8debd61109c3726872d19b8090c7e3ae3c2cf87fe69de3533fe92127251 --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block --gas=300000

qwoynd tx game sign-withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 2 '{"mint":{"token_id":"2","owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk","extension":{"ship_type":12,"owner":"300"}}}' --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# send tokens to admin to add liquidity
qwoynd tx bank send validator $(qwoynd keys show -a signer --keyring-backend=test) 1000000ucoho,1000000uqwoyn --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block
# add liquidity
qwoynd tx game add-liquidity 1000000ucoho,1000000uqwoyn --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block
# remove liquidity
qwoynd tx game remove-liquidity 10000ucoho,10000uqwoyn --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

qwoynd tx game swap 10000ucoho --from=validator --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block

# transfer module ownership
qwoynd tx game transfer-module-ownership $(qwoynd keys show -a signer --keyring-backend=test) --from=signer --chain-id=qwoyn-1 --keyring-backend=test --home=$HOME/.qwoynd/ -y --broadcast-mode=block
```
