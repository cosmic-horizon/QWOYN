#!/bin/sh

rm -rf ~/.cohod

# Build genesis
cohod init --chain-id=qwoyn-1 test
cohod keys add validator --keyring-backend="test"
cohod add-genesis-account $(cohod keys show validator -a --keyring-backend="test") 100000000000000uqwoyn,100000000000000ucoho,100000000000000stake
cohod gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
cohod collect-gentxs

# Start node
cohod start --pruning=nothing --minimum-gas-prices="0stake"
