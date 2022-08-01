```sh
cohod query game params
cohod query game whitelisted-contracts

# whitelist contract
cohod tx game whitelist-contracts coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc --from=signer --chain-id=test --keyring-backend=test --home=$HOME/.cohod/ -y --broadcast-mode=block

# deposit nft from end wallet to game
cohod tx game deposit-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 --from=signer --chain-id=test --keyring-backend=test --home=$HOME/.cohod/ -y --broadcast-mode=block

# query owner of nft
cohod query wasm contract-state smart coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"owner_of":{"token_id":"1"}}'

# game module address
# coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk

# instantiate ship nft with module account owned
cohod tx wasm instantiate 1 '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk","owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk"}' --from validator --label "Ship-NFT" --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/ --no-admin -y

# transfer nft update ownership to module account - when owner is set to different account
cohod tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"transfer_ownership":{"owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk"}}' --from signer --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/ -y


cohod tx game sign-withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' --from=signer --chain-id=test --keyring-backend=test --home=$HOME/.cohod/

cohod tx game withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' 42d6e9d3b62ffc9b0bc3f6a97cbc0857af1c7a7aa57549571d7bc72415a955d978a1790440ce53c8f9fbfa2ce70d967812eda6094d6f112d7e5736170e48e2a8 --from=signer --chain-id=test --keyring-backend=test --home=$HOME/.cohod/ -y --broadcast-mode=block --gas=400000

cohod tx game withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 2 '{"mint":{"token_id":"2","owner":"coho1djju4dm7wqk8s76vzjea80exht2rmfsxjx47wk","extension":{"ship_type":12,"owner":"300"}}}' --from=signer --chain-id=test --keyring-backend=test --home=$HOME/.cohod/ -y --broadcast-mode=block
```
