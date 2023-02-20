#!/bin/sh

rm -rf ~/.qwoynd

# Build genesis
qwoynd init --chain-id=qwoyn-1 test
qwoynd keys add validator --keyring-backend="test"
qwoynd add-genesis-account $(qwoynd keys show validator -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
qwoynd collect-gentxs

# sed -i 's/stake/uqwoyn/g' $HOME/.qwoynd/config/genesis.json

# Start node
qwoynd start --pruning=nothing --minimum-gas-prices="0stake"
