# Cosmic Horizon Network - Smart Contracts

## Install Docker version 4.10.0

Check Docker [Docs](https://docs.docker.com/desktop/mac/install/) page for the Docker installation guide.

## Install Rust and Cargo

Here's [installation guide](https://doc.rust-lang.org/cargo/getting-started/installation.html) for Rust and Cargo.

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
cohod tx wasm store $WASM_FILE --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.cohod/

# $WASM_FILE - Path to wasm binary file (Example: ./artifacts/ship_nft.wasm)
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to upload Ship NFT wasm binary to local test node

```
cohod tx wasm store ./artifacts/ship_nft.wasm --from coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/
```

## Instantiate Contract

```
cohod tx wasm instantiate $CODE_ID $INSTANTIATE_MSG --amount 50000ucoho --from $FROM --label $LABEL --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.cohod/ --no-admin

# $CODE_ID - Uploaded Code ID (Example: 1)
# $INSTANTIATE_MSG - JSON encoded init args (Example: '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf"}')
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to init Ship NFT contract on local test node

```
cohod tx wasm instantiate 1 '{"name":"Ship NFT","symbol":"SHIP","minter":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf"}' --amount 50000ucoho --from coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf --label "Ship-NFT" --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/ --no-admin
```

## Get Contract Address

```
cohod query wasm list-contract-by-code $CODE_ID

# $CODE_ID - Uploaded Code ID (Example: 1)
```

Example command to get Ship NFT contract address on local test node

```
cohod query wasm list-contract-by-code 1
```

## Mint NFT

```
cohod tx wasm execute $CONTRACT_ADDRESS $EXECUTE_MSG --amount 50000ucoho --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.cohod/

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# $EXECUTE_MSG - JSON encoded send args (Example: '{"mint":{"token_id":"1","owner":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf","extension":{"ship_type":10,"owner":"100"}}}')
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to mint Ship NFT with token_id "1"

```
cohod tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"mint":{"token_id":"1","owner":"coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf","extension":{"ship_type":10,"owner":"100"}}}' --amount 50000ucoho --from validator --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/
```

## Get NFT Info

```
cohod query wasm contract-state smart $CONTRACT_ADDRESS $QUERY_MSG

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# QUERY_MSG - JSON encoded read args (Example: '{"nft_info":{"token_id":"1"}}')
```

Example command to get Ship NFT info with token_id "1"

```
cohod query wasm contract-state smart coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"nft_info":{"token_id":"1"}}'
```

## Update NFT

```
cohod tx wasm execute $CONTRACT_ADDRESS $EXECUTE_MSG --amount 50000ucoho --from $FROM --chain-id $CHAIN_ID --gas auto --gas-adjustment 1.3 -b block --keyring-backend=$KEYRING_BACKEND --home=$HOME/.cohod/

# $CONTRACT_ADDRESS - Contract address (Example: coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc)
# $EXECUTE_MSG - JSON encoded send args (Example: '{"update_nft":{"token_id":"1","extension":{"ship_type":20,"owner":"200"}}}')
# $FROM - Name or address of signer account (Example: coho1m6auqrjwertsnccvkk9tts3lzw0hcz0jn2v3lf)
# $LABEL - Label string
# $CHAIN_ID - Chain ID (Example: test)
# $KEYRING_BACKEND - Can be os|file|kwallet|pass|test|memory (Example: test)
```

Example command to update Ship NFT with token_id "1"

```
cohod tx wasm execute coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc '{"update_nft":{"token_id":"1","extension":{"ship_type":20,"owner":"200"}}}' --amount 50000ucoho --from validator --chain-id test --gas auto --gas-adjustment 1.3 -b block --keyring-backend=test --home=$HOME/.cohod/
```
