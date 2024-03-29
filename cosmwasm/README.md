# Cosmic Horizon Network - Smart Contracts

## Install Docker version 4.10.0

Check Docker [Docs](https://docs.docker.com/desktop/mac/install/) page for the Docker installation guide.

## Install Rust and Cargo

Here's [installation guide](https://doc.rust-lang.org/cargo/getting-started/installation.html) for Rust and Cargo.

## Add contract owner account

```sh
qwoynd keys add signer --keyring-backend=test --home=$HOME/.qwoynd --recover
# pipe woman clutch absorb lonely cost credit math antique better thumb cook pave clarify hungry east garbage absent warfare song helmet anchor drift purity

qwoynd tx bank send validator $(qwoynd keys show -a signer --keyring-backend=test --home=$HOME/.qwoynd) 1000000stake --keyring-backend=test --home=$HOME/.qwoynd --broadcast-mode=block -y --chain-id=qwoyn-1
```

## Build Contracts

Before you run below command you need to run Docker on your local machine.

This command will compile and build CosmWasm smart contracts.

```
# go to cosmwasm dir
cd cosmwasm

# execute build script
sh build_optimised_release.sh
```

You can find compiled `.wasm` files under `artifacts` directory.

## Upload WASM binary

```
qwoynd tx wasm store $WASM_FILE --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.qwoynd/ -y

# $WASM_FILE - Path to wasm binary file (Example: ./artifacts/ship_nft.wasm)
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to upload Ship NFT wasm binary to local test node

```
qwoynd tx wasm store ./artifacts/ship_nft.wasm --from validator --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y
```

## Instantiate Contract

```
qwoynd tx wasm instantiate $CODE_ID $INSTANTIATE_MSG --from $FROM --label $LABEL --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.qwoynd/ --no-admin -y

# $CODE_ID - Uploaded Code ID (Example: 1)
# $INSTANTIATE_MSG - JSON encoded init args (Example: '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf"}')
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to init Ship NFT contract on local test node

```
qwoynd tx wasm instantiate 1 '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v","owner":"coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v"}' --from validator --label "Ship-NFT" --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ --no-admin -y
```

## Get Contract Address

```
qwoynd query wasm list-contract-by-code $CODE_ID

# $CODE_ID - Uploaded Code ID (Example: 1)
```

Example command to get Ship NFT contract address on local test node

```
qwoynd query wasm list-contract-by-code 1
```

## Mint NFT

```
qwoynd tx wasm execute $CONTRACT_ADDRESS $EXECUTE_MSG --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.qwoynd/ -y

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# $EXECUTE_MSG - JSON encoded send args (Example: '{"mint":{"token_id":"1","owner":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf","extension":{"ship_type":10,"owner":"100"}}}')
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to mint Ship NFT with token_id "1"

```
qwoynd tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"mint":{"token_id":"1","owner":"coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v","extension":{"ship_type":10,"owner":"100"}}}' --from signer --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y
```

## Transfer NFT

Example command to transfer minted Ship NFT with token_id "1"

```
qwoynd tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"transfer_nft":{"token_id":"1","recipient":"coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v"}}' --from signer --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y
```

## Get NFT Info

```
qwoynd query wasm contract-state smart $CONTRACT_ADDRESS $QUERY_MSG

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# QUERY_MSG - JSON encoded read args (Example: '{"nft_info":{"token_id":"1"}}')
```

Example command to get Ship NFT contract info

```
qwoynd query wasm contract-state smart coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"contract_info":{}}'
```

Example command to get Ship NFT info with token_id "1"

```
qwoynd query wasm contract-state smart coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"nft_info":{"token_id":"1"}}'
```

## Update NFT

```
qwoynd tx wasm execute $CONTRACT_ADDRESS $EXECUTE_MSG --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.qwoynd/ -y

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# $EXECUTE_MSG - JSON encoded send args (Example: '{"update_nft":{"token_id":"1","extension":{"ship_type":20,"owner":"200"}}}')
# $FROM - Name or address of signer account (Example: signer)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to update Ship NFT with token_id "1"

```
qwoynd tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"update_nft":{"token_id":"1","extension":{"ship_type":20,"owner":"200"}}}' --from signer --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y
```

## Transfer Ownership

```
qwoynd tx wasm execute $CONTRACT_ADDRESS $EXECUTE_MSG --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.qwoynd/ -y

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# $EXECUTE_MSG - JSON encoded send args (Example: '{"transfer_ownership":{"owner":"coho1f0yhatqxayku7y00k448n43qwch320v4da2plx"}}')
# $FROM - Name or address of signer account (Example: signer)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to transfer ownership to `coho1f0yhatqxayku7y00k448n43qwch320v4da2plx`

```
qwoynd tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"transfer_ownership":{"owner":"coho1f0yhatqxayku7y00k448n43qwch320v4da2plx"}}' --from signer --chain-id qwoyn-1 --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.qwoynd/ -y
```
