#!/bin/sh

rm -rf ~/.coho

# Build genesis
cohod init --chain-id=test test
cohod keys add validator --keyring-backend="test"
cohod add-genesis-account $(cohod keys show validator -a --keyring-backend="test") 100000000000000ucoho
cohod gentx validator 100000000ucoho --keyring-backend="test" --chain-id=test
cohod collect-gentxs

# Start node
cohod start --pruning=nothing --minimum-gas-prices="0ucoho"
